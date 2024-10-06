// Code generated by mockery. DO NOT EDIT.

package eventstore

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	egopb "github.com/tochemey/ego/v3/egopb"
)

// EventsStore is an autogenerated mock type for the EventsStore type
type EventsStore struct {
	mock.Mock
}

type EventsStore_Expecter struct {
	mock *mock.Mock
}

func (_m *EventsStore) EXPECT() *EventsStore_Expecter {
	return &EventsStore_Expecter{mock: &_m.Mock}
}

// Connect provides a mock function with given fields: ctx
func (_m *EventsStore) Connect(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EventsStore_Connect_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Connect'
type EventsStore_Connect_Call struct {
	*mock.Call
}

// Connect is a helper method to define mock.On call
//   - ctx context.Context
func (_e *EventsStore_Expecter) Connect(ctx interface{}) *EventsStore_Connect_Call {
	return &EventsStore_Connect_Call{Call: _e.mock.On("Connect", ctx)}
}

func (_c *EventsStore_Connect_Call) Run(run func(ctx context.Context)) *EventsStore_Connect_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *EventsStore_Connect_Call) Return(_a0 error) *EventsStore_Connect_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EventsStore_Connect_Call) RunAndReturn(run func(context.Context) error) *EventsStore_Connect_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteEvents provides a mock function with given fields: ctx, persistenceID, toSequenceNumber
func (_m *EventsStore) DeleteEvents(ctx context.Context, persistenceID string, toSequenceNumber uint64) error {
	ret := _m.Called(ctx, persistenceID, toSequenceNumber)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, uint64) error); ok {
		r0 = rf(ctx, persistenceID, toSequenceNumber)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EventsStore_DeleteEvents_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteEvents'
type EventsStore_DeleteEvents_Call struct {
	*mock.Call
}

// DeleteEvents is a helper method to define mock.On call
//   - ctx context.Context
//   - persistenceID string
//   - toSequenceNumber uint64
func (_e *EventsStore_Expecter) DeleteEvents(ctx interface{}, persistenceID interface{}, toSequenceNumber interface{}) *EventsStore_DeleteEvents_Call {
	return &EventsStore_DeleteEvents_Call{Call: _e.mock.On("DeleteEvents", ctx, persistenceID, toSequenceNumber)}
}

func (_c *EventsStore_DeleteEvents_Call) Run(run func(ctx context.Context, persistenceID string, toSequenceNumber uint64)) *EventsStore_DeleteEvents_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(uint64))
	})
	return _c
}

func (_c *EventsStore_DeleteEvents_Call) Return(_a0 error) *EventsStore_DeleteEvents_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EventsStore_DeleteEvents_Call) RunAndReturn(run func(context.Context, string, uint64) error) *EventsStore_DeleteEvents_Call {
	_c.Call.Return(run)
	return _c
}

// Disconnect provides a mock function with given fields: ctx
func (_m *EventsStore) Disconnect(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EventsStore_Disconnect_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Disconnect'
type EventsStore_Disconnect_Call struct {
	*mock.Call
}

// Disconnect is a helper method to define mock.On call
//   - ctx context.Context
func (_e *EventsStore_Expecter) Disconnect(ctx interface{}) *EventsStore_Disconnect_Call {
	return &EventsStore_Disconnect_Call{Call: _e.mock.On("Disconnect", ctx)}
}

func (_c *EventsStore_Disconnect_Call) Run(run func(ctx context.Context)) *EventsStore_Disconnect_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *EventsStore_Disconnect_Call) Return(_a0 error) *EventsStore_Disconnect_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EventsStore_Disconnect_Call) RunAndReturn(run func(context.Context) error) *EventsStore_Disconnect_Call {
	_c.Call.Return(run)
	return _c
}

// GetLatestEvent provides a mock function with given fields: ctx, persistenceID
func (_m *EventsStore) GetLatestEvent(ctx context.Context, persistenceID string) (*egopb.Event, error) {
	ret := _m.Called(ctx, persistenceID)

	var r0 *egopb.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*egopb.Event, error)); ok {
		return rf(ctx, persistenceID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *egopb.Event); ok {
		r0 = rf(ctx, persistenceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*egopb.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, persistenceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventsStore_GetLatestEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLatestEvent'
type EventsStore_GetLatestEvent_Call struct {
	*mock.Call
}

// GetLatestEvent is a helper method to define mock.On call
//   - ctx context.Context
//   - persistenceID string
func (_e *EventsStore_Expecter) GetLatestEvent(ctx interface{}, persistenceID interface{}) *EventsStore_GetLatestEvent_Call {
	return &EventsStore_GetLatestEvent_Call{Call: _e.mock.On("GetLatestEvent", ctx, persistenceID)}
}

func (_c *EventsStore_GetLatestEvent_Call) Run(run func(ctx context.Context, persistenceID string)) *EventsStore_GetLatestEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *EventsStore_GetLatestEvent_Call) Return(_a0 *egopb.Event, _a1 error) *EventsStore_GetLatestEvent_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EventsStore_GetLatestEvent_Call) RunAndReturn(run func(context.Context, string) (*egopb.Event, error)) *EventsStore_GetLatestEvent_Call {
	_c.Call.Return(run)
	return _c
}

// GetShardEvents provides a mock function with given fields: ctx, shardNumber, offset, max
func (_m *EventsStore) GetShardEvents(ctx context.Context, shardNumber uint64, offset int64, max uint64) ([]*egopb.Event, int64, error) {
	ret := _m.Called(ctx, shardNumber, offset, max)

	var r0 []*egopb.Event
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, int64, uint64) ([]*egopb.Event, int64, error)); ok {
		return rf(ctx, shardNumber, offset, max)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, int64, uint64) []*egopb.Event); ok {
		r0 = rf(ctx, shardNumber, offset, max)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*egopb.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, int64, uint64) int64); ok {
		r1 = rf(ctx, shardNumber, offset, max)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, uint64, int64, uint64) error); ok {
		r2 = rf(ctx, shardNumber, offset, max)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EventsStore_GetShardEvents_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetShardEvents'
type EventsStore_GetShardEvents_Call struct {
	*mock.Call
}

// GetShardEvents is a helper method to define mock.On call
//   - ctx context.Context
//   - shardNumber uint64
//   - offset int64
//   - max uint64
func (_e *EventsStore_Expecter) GetShardEvents(ctx interface{}, shardNumber interface{}, offset interface{}, max interface{}) *EventsStore_GetShardEvents_Call {
	return &EventsStore_GetShardEvents_Call{Call: _e.mock.On("GetShardEvents", ctx, shardNumber, offset, max)}
}

func (_c *EventsStore_GetShardEvents_Call) Run(run func(ctx context.Context, shardNumber uint64, offset int64, max uint64)) *EventsStore_GetShardEvents_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(int64), args[3].(uint64))
	})
	return _c
}

func (_c *EventsStore_GetShardEvents_Call) Return(_a0 []*egopb.Event, _a1 int64, _a2 error) *EventsStore_GetShardEvents_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *EventsStore_GetShardEvents_Call) RunAndReturn(run func(context.Context, uint64, int64, uint64) ([]*egopb.Event, int64, error)) *EventsStore_GetShardEvents_Call {
	_c.Call.Return(run)
	return _c
}

// PersistenceIDs provides a mock function with given fields: ctx, pageSize, pageToken
func (_m *EventsStore) PersistenceIDs(ctx context.Context, pageSize uint64, pageToken string) ([]string, string, error) {
	ret := _m.Called(ctx, pageSize, pageToken)

	var r0 []string
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, string) ([]string, string, error)); ok {
		return rf(ctx, pageSize, pageToken)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, string) []string); ok {
		r0 = rf(ctx, pageSize, pageToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, string) string); ok {
		r1 = rf(ctx, pageSize, pageToken)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, uint64, string) error); ok {
		r2 = rf(ctx, pageSize, pageToken)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EventsStore_PersistenceIDs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PersistenceIDs'
type EventsStore_PersistenceIDs_Call struct {
	*mock.Call
}

// PersistenceIDs is a helper method to define mock.On call
//   - ctx context.Context
//   - pageSize uint64
//   - pageToken string
func (_e *EventsStore_Expecter) PersistenceIDs(ctx interface{}, pageSize interface{}, pageToken interface{}) *EventsStore_PersistenceIDs_Call {
	return &EventsStore_PersistenceIDs_Call{Call: _e.mock.On("PersistenceIDs", ctx, pageSize, pageToken)}
}

func (_c *EventsStore_PersistenceIDs_Call) Run(run func(ctx context.Context, pageSize uint64, pageToken string)) *EventsStore_PersistenceIDs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(string))
	})
	return _c
}

func (_c *EventsStore_PersistenceIDs_Call) Return(persistenceIDs []string, nextPageToken string, err error) *EventsStore_PersistenceIDs_Call {
	_c.Call.Return(persistenceIDs, nextPageToken, err)
	return _c
}

func (_c *EventsStore_PersistenceIDs_Call) RunAndReturn(run func(context.Context, uint64, string) ([]string, string, error)) *EventsStore_PersistenceIDs_Call {
	_c.Call.Return(run)
	return _c
}

// Ping provides a mock function with given fields: ctx
func (_m *EventsStore) Ping(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EventsStore_Ping_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ping'
type EventsStore_Ping_Call struct {
	*mock.Call
}

// Ping is a helper method to define mock.On call
//   - ctx context.Context
func (_e *EventsStore_Expecter) Ping(ctx interface{}) *EventsStore_Ping_Call {
	return &EventsStore_Ping_Call{Call: _e.mock.On("Ping", ctx)}
}

func (_c *EventsStore_Ping_Call) Run(run func(ctx context.Context)) *EventsStore_Ping_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *EventsStore_Ping_Call) Return(_a0 error) *EventsStore_Ping_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EventsStore_Ping_Call) RunAndReturn(run func(context.Context) error) *EventsStore_Ping_Call {
	_c.Call.Return(run)
	return _c
}

// ReplayEvents provides a mock function with given fields: ctx, persistenceID, fromSequenceNumber, toSequenceNumber, max
func (_m *EventsStore) ReplayEvents(ctx context.Context, persistenceID string, fromSequenceNumber uint64, toSequenceNumber uint64, max uint64) ([]*egopb.Event, error) {
	ret := _m.Called(ctx, persistenceID, fromSequenceNumber, toSequenceNumber, max)

	var r0 []*egopb.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, uint64, uint64, uint64) ([]*egopb.Event, error)); ok {
		return rf(ctx, persistenceID, fromSequenceNumber, toSequenceNumber, max)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, uint64, uint64, uint64) []*egopb.Event); ok {
		r0 = rf(ctx, persistenceID, fromSequenceNumber, toSequenceNumber, max)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*egopb.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, uint64, uint64, uint64) error); ok {
		r1 = rf(ctx, persistenceID, fromSequenceNumber, toSequenceNumber, max)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventsStore_ReplayEvents_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReplayEvents'
type EventsStore_ReplayEvents_Call struct {
	*mock.Call
}

// ReplayEvents is a helper method to define mock.On call
//   - ctx context.Context
//   - persistenceID string
//   - fromSequenceNumber uint64
//   - toSequenceNumber uint64
//   - max uint64
func (_e *EventsStore_Expecter) ReplayEvents(ctx interface{}, persistenceID interface{}, fromSequenceNumber interface{}, toSequenceNumber interface{}, max interface{}) *EventsStore_ReplayEvents_Call {
	return &EventsStore_ReplayEvents_Call{Call: _e.mock.On("ReplayEvents", ctx, persistenceID, fromSequenceNumber, toSequenceNumber, max)}
}

func (_c *EventsStore_ReplayEvents_Call) Run(run func(ctx context.Context, persistenceID string, fromSequenceNumber uint64, toSequenceNumber uint64, max uint64)) *EventsStore_ReplayEvents_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(uint64), args[3].(uint64), args[4].(uint64))
	})
	return _c
}

func (_c *EventsStore_ReplayEvents_Call) Return(_a0 []*egopb.Event, _a1 error) *EventsStore_ReplayEvents_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EventsStore_ReplayEvents_Call) RunAndReturn(run func(context.Context, string, uint64, uint64, uint64) ([]*egopb.Event, error)) *EventsStore_ReplayEvents_Call {
	_c.Call.Return(run)
	return _c
}

// ShardNumbers provides a mock function with given fields: ctx
func (_m *EventsStore) ShardNumbers(ctx context.Context) ([]uint64, error) {
	ret := _m.Called(ctx)

	var r0 []uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]uint64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []uint64); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventsStore_ShardNumbers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ShardNumbers'
type EventsStore_ShardNumbers_Call struct {
	*mock.Call
}

// ShardNumbers is a helper method to define mock.On call
//   - ctx context.Context
func (_e *EventsStore_Expecter) ShardNumbers(ctx interface{}) *EventsStore_ShardNumbers_Call {
	return &EventsStore_ShardNumbers_Call{Call: _e.mock.On("ShardNumbers", ctx)}
}

func (_c *EventsStore_ShardNumbers_Call) Run(run func(ctx context.Context)) *EventsStore_ShardNumbers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *EventsStore_ShardNumbers_Call) Return(_a0 []uint64, _a1 error) *EventsStore_ShardNumbers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EventsStore_ShardNumbers_Call) RunAndReturn(run func(context.Context) ([]uint64, error)) *EventsStore_ShardNumbers_Call {
	_c.Call.Return(run)
	return _c
}

// WriteEvents provides a mock function with given fields: ctx, events
func (_m *EventsStore) WriteEvents(ctx context.Context, events []*egopb.Event) error {
	ret := _m.Called(ctx, events)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*egopb.Event) error); ok {
		r0 = rf(ctx, events)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EventsStore_WriteEvents_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteEvents'
type EventsStore_WriteEvents_Call struct {
	*mock.Call
}

// WriteEvents is a helper method to define mock.On call
//   - ctx context.Context
//   - events []*egopb.Event
func (_e *EventsStore_Expecter) WriteEvents(ctx interface{}, events interface{}) *EventsStore_WriteEvents_Call {
	return &EventsStore_WriteEvents_Call{Call: _e.mock.On("WriteEvents", ctx, events)}
}

func (_c *EventsStore_WriteEvents_Call) Run(run func(ctx context.Context, events []*egopb.Event)) *EventsStore_WriteEvents_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]*egopb.Event))
	})
	return _c
}

func (_c *EventsStore_WriteEvents_Call) Return(_a0 error) *EventsStore_WriteEvents_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EventsStore_WriteEvents_Call) RunAndReturn(run func(context.Context, []*egopb.Event) error) *EventsStore_WriteEvents_Call {
	_c.Call.Return(run)
	return _c
}

// NewEventsStore creates a new instance of EventsStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEventsStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *EventsStore {
	mock := &EventsStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
