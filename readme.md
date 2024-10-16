## eGo

[![build](https://img.shields.io/github/actions/workflow/status/Tochemey/ego/build.yml?branch=main)](https://github.com/Tochemey/ego/actions/workflows/build.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/tochemey/ego.svg)](https://pkg.go.dev/github.com/tochemey/ego)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/tochemey/ego)](https://go.dev/doc/install)
![GitHub Release](https://img.shields.io/github/v/release/Tochemey/ego)
[![codecov](https://codecov.io/gh/Tochemey/ego/branch/main/graph/badge.svg?token=Z5b9gM6Mnt)](https://codecov.io/gh/Tochemey/ego)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2FTochemey%2Fego.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2FTochemey%2Fego?ref=badge_shield)

eGo is a minimal library that help build event-sourcing and CQRS application through a simple interface, and it allows developers to describe their **_commands_**, **_events_** and **_states_** **_are defined using google protocol buffers_**.
Under the hood, ego leverages [Go-Akt](https://github.com/Tochemey/goakt) to scale out and guarantee performant, reliable persistence.

### Features

#### Domain Entity/Aggregate Root

The aggregate root is crucial for maintaining data consistency, especially in distributed systems. It defines how to handle the various commands (requests to perform actions) that are always directed at the aggregate root.
In eGo commands sent the aggregate root are processed in order. When a command is processed, it may result in the generation of events, which are then stored in an event store. Every event persisted has a revision number
and timestamp that can help track it. The aggregate root in eGo is responsible for defining how to handle events that are the result of command handlers. The end result of events handling is to build the new state of the aggregate root.
When running in cluster mode, aggregate root are sharded.

- `Commands handler`: The command handlers define how to handle each incoming command,
  which validations must be applied, and finally, which events will be persisted if any. When there is no event to be persisted a nil can
  be returned as a no-op. Command handlers are the meat of the event sourced actor.
  They encode the business rules of your event sourced actor and act as a guardian of the Aggregate consistency.
  The command handler must first validate that the incoming command can be applied to the current model state.
  Any decision should be solely based on the data passed in the commands and the state of the Behavior.
  In case of successful validation, one or more events expressing the mutations are persisted. Once the events are persisted, they are applied to the state producing a new valid state.
- `Events handler`: The event handlers are used to mutate the state of the Aggregate by applying the events to it.
  Event handlers must be pure functions as they will be used when instantiating the Aggregate and replaying the event store.

#### Howto

To define an Aggregate Root, one needs to:
1. the state of the aggregate root using google protocol buffers message
2. the various commands that will be handled by the aggregate root
3. the various events that are result of the command handlers and that will be handled by the aggregate root to return the new state of the aggregate root
2. implements the [`EntityBehavior`](./behavior.go) interface.

#### Events Stream

Every event handled by Aggregate Root are pushed to an events stream. That enables real-time processing of events without having to interact with the events store

#### Projection

One can add a projection to the eGo engine to help build a read model. Projections in eGo rely on an offset store to track how far they have consumed events
persisted by the write model. The offset used in eGo is a timestamp-based offset. One can also:
- remove a given projection: this will stop the projection and remove it from the system
- check the status of a given projection

#### Events Store

One can implement a custom events store. See [EventsStore](./eventstore/iface.go). eGo comes packaged with two events store:
- [Postgres](./eventstore/postgres/postgres.go): Schema can be found [here](./resources/eventstore_postgres.sql)
- [Memory](./eventstore/memory/memory.go) (for testing purpose only)

#### Offsets Store

One can implement a custom offsets store. See [OffsetStore](./offsetstore/iface.go). eGo comes packaged with two offset store:
- [Postgres](./offsetstore/postgres/postgres.go): Schema can be found [here](./resources/offsetstore_postgres.sql)
- [Memory](./offsetstore/memory/memory.go) (for testing purpose only)

#### Cluster

The cluster mode heavily relies on [Go-Akt](https://github.com/Tochemey/goakt#clustering) clustering.

#### Mocks

eGo ships in some [mocks](./mocks)

### Examples

Check the [examples](./example)

### Installation

```bash
go get github.com/tochemey/ego
```

### Sample

```go
package main

import (
  "context"
  "errors"
  "log"
  "os"
  "os/signal"
  "syscall"
  "time"

  "github.com/google/uuid"
  "google.golang.org/protobuf/proto"

  "github.com/tochemey/ego/v3"
  "github.com/tochemey/ego/v3/eventstore/memory"
  samplepb "github.com/tochemey/ego/v3/example/pbs/sample/pb/v1"
)

func main() {
  // create the go context
  ctx := context.Background()
  // create the event store
  eventStore := memory.NewEventsStore()
  // connect the event store
  _ = eventStore.Connect(ctx)
  // create the ego engine
  engine := ego.NewEngine("Sample", eventStore)
  // start ego engine
  _ = engine.Start(ctx)
  // create a persistence id
  entityID := uuid.NewString()
  // create an entity behavior with a given id
  behavior := NewAccountBehavior(entityID)
  // create an entity
  _ = engine.Entity(ctx, behavior)

  // send some commands to the pid
  var command proto.Message
  // create an account
  command = &samplepb.CreateAccount{
    AccountId:      entityID,
    AccountBalance: 500.00,
  }
  // send the command to the actor. Please don't ignore the error in production grid code
  reply, _, _ := engine.SendCommand(ctx, entityID, command, time.Minute)
  account := reply.(*samplepb.Account)
  log.Printf("current balance: %v", account.GetAccountBalance())

  // send another command to credit the balance
  command = &samplepb.CreditAccount{
    AccountId: entityID,
    Balance:   250,
  }

  reply, _, _ = engine.SendCommand(ctx, entityID, command, time.Minute)
  account = reply.(*samplepb.Account)
  log.Printf("current balance: %v", account.GetAccountBalance())

  // capture ctrl+c
  interruptSignal := make(chan os.Signal, 1)
  signal.Notify(interruptSignal, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
  <-interruptSignal

  // disconnect the event store
  _ = eventStore.Disconnect(ctx)
  // stop the actor system
  _ = engine.Stop(ctx)
  os.Exit(0)
}

// AccountBehavior implements EntityBehavior
type AccountBehavior struct {
  id string
}

// make sure that AccountBehavior is a true persistence behavior
var _ ego.EntityBehavior = (*AccountBehavior)(nil)

// NewAccountBehavior creates an instance of AccountBehavior
func NewAccountBehavior(id string) *AccountBehavior {
  return &AccountBehavior{id: id}
}

// ID returns the id
func (a *AccountBehavior) ID() string {
  return a.id
}

// InitialState returns the initial state
func (a *AccountBehavior) InitialState() ego.State {
  return ego.State(new(samplepb.Account))
}

// HandleCommand handles every command that is sent to the persistent behavior
func (a *AccountBehavior) HandleCommand(_ context.Context, command ego.Command, _ ego.State) (events []ego.Event, err error) {
  switch cmd := command.(type) {
  case *samplepb.CreateAccount:
    // TODO in production grid app validate the command using the prior state
    return []ego.Event{
      &samplepb.AccountCreated{
        AccountId:      cmd.GetAccountId(),
        AccountBalance: cmd.GetAccountBalance(),
      },
    }, nil

  case *samplepb.CreditAccount:
    // TODO in production grid app validate the command using the prior state
    return []ego.Event{
      &samplepb.AccountCredited{
        AccountId:      cmd.GetAccountId(),
        AccountBalance: cmd.GetBalance(),
      },
    }, nil

  default:
    return nil, errors.New("unhandled command")
  }
}

// HandleEvent handles every event emitted
func (a *AccountBehavior) HandleEvent(_ context.Context, event ego.Event, priorState ego.State) (state ego.State, err error) {
  switch evt := event.(type) {
  case *samplepb.AccountCreated:
    return &samplepb.Account{
      AccountId:      evt.GetAccountId(),
      AccountBalance: evt.GetAccountBalance(),
    }, nil

  case *samplepb.AccountCredited:
    account := priorState.(*samplepb.Account)
    bal := account.GetAccountBalance() + evt.GetAccountBalance()
    return &samplepb.Account{
      AccountId:      evt.GetAccountId(),
      AccountBalance: bal,
    }, nil

  default:
    return nil, errors.New("unhandled event")
  }
}

```

## Versioning

The version system adopted in eGo deviates a bit from the standard semantic versioning system.
The version format is as follows:

- The `MAJOR` part of the version will stay at `v3` for the meantime.
- The `MINOR` part of the version will cater for any new _features_, _breaking changes_  with a note on the breaking changes.
- The `PATCH` part of the version will cater for dependencies upgrades, bug fixes, security patches and co.

The versioning will remain like `v3.x.x` until further notice.

### Contribution

Contributions are welcome!
The project adheres to [Semantic Versioning](https://semver.org) and [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/).
This repo uses [Earthly](https://earthly.dev/get-earthly).

There are two ways you can become a contributor:

1. Request to become a collaborator, and then you can just open pull requests against the repository without forking it.
2. Follow these steps
  - Fork the repository
  - Create a feature branch
  - Set your docker credentials on your fork using the following secret names: `DOCKER_USER` and `DOCKER_PASS`
  - Submit a [pull request](https://help.github.com/articles/using-pull-requests)

## Test & Linter

Prior to submitting a [pull request](https://help.github.com/articles/using-pull-requests), please run:

```bash
earthly +test
```


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2FTochemey%2Fego.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2FTochemey%2Fego?ref=badge_large)