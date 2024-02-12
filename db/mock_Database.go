// Code generated by mockery v2.40.3. DO NOT EDIT.

package db

import mock "github.com/stretchr/testify/mock"

// MockDatabase is an autogenerated mock type for the Database type
type MockDatabase struct {
	mock.Mock
}

type MockDatabase_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDatabase) EXPECT() *MockDatabase_Expecter {
	return &MockDatabase_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *MockDatabase) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDatabase_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockDatabase_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockDatabase_Expecter) Close() *MockDatabase_Close_Call {
	return &MockDatabase_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockDatabase_Close_Call) Run(run func()) *MockDatabase_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDatabase_Close_Call) Return(_a0 error) *MockDatabase_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDatabase_Close_Call) RunAndReturn(run func() error) *MockDatabase_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Connect provides a mock function with given fields:
func (_m *MockDatabase) Connect() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Connect")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDatabase_Connect_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Connect'
type MockDatabase_Connect_Call struct {
	*mock.Call
}

// Connect is a helper method to define mock.On call
func (_e *MockDatabase_Expecter) Connect() *MockDatabase_Connect_Call {
	return &MockDatabase_Connect_Call{Call: _e.mock.On("Connect")}
}

func (_c *MockDatabase_Connect_Call) Run(run func()) *MockDatabase_Connect_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDatabase_Connect_Call) Return(_a0 error) *MockDatabase_Connect_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDatabase_Connect_Call) RunAndReturn(run func() error) *MockDatabase_Connect_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: key
func (_m *MockDatabase) Get(key string) (string, error) {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDatabase_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockDatabase_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - key string
func (_e *MockDatabase_Expecter) Get(key interface{}) *MockDatabase_Get_Call {
	return &MockDatabase_Get_Call{Call: _e.mock.On("Get", key)}
}

func (_c *MockDatabase_Get_Call) Run(run func(key string)) *MockDatabase_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockDatabase_Get_Call) Return(value string, err error) *MockDatabase_Get_Call {
	_c.Call.Return(value, err)
	return _c
}

func (_c *MockDatabase_Get_Call) RunAndReturn(run func(string) (string, error)) *MockDatabase_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function with given fields: key, value
func (_m *MockDatabase) Set(key string, value string) error {
	ret := _m.Called(key, value)

	if len(ret) == 0 {
		panic("no return value specified for Set")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDatabase_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type MockDatabase_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - key string
//   - value string
func (_e *MockDatabase_Expecter) Set(key interface{}, value interface{}) *MockDatabase_Set_Call {
	return &MockDatabase_Set_Call{Call: _e.mock.On("Set", key, value)}
}

func (_c *MockDatabase_Set_Call) Run(run func(key string, value string)) *MockDatabase_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockDatabase_Set_Call) Return(err error) *MockDatabase_Set_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockDatabase_Set_Call) RunAndReturn(run func(string, string) error) *MockDatabase_Set_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDatabase creates a new instance of MockDatabase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDatabase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDatabase {
	mock := &MockDatabase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
