// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	context "context"

	job "github.com/smartcontractkit/chainlink/v2/core/services/job"
	mock "github.com/stretchr/testify/mock"
)

// ORM is an autogenerated mock type for the ORM type
type ORM struct {
	mock.Mock
}

type ORM_Expecter struct {
	mock *mock.Mock
}

func (_m *ORM) EXPECT() *ORM_Expecter {
	return &ORM_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, secretsURL, hash, contents
func (_m *ORM) Create(ctx context.Context, secretsURL string, hash string, contents string) (int64, error) {
	ret := _m.Called(ctx, secretsURL, hash, contents)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) (int64, error)); ok {
		return rf(ctx, secretsURL, hash, contents)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) int64); ok {
		r0 = rf(ctx, secretsURL, hash, contents)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, secretsURL, hash, contents)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ORM_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type ORM_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - secretsURL string
//   - hash string
//   - contents string
func (_e *ORM_Expecter) Create(ctx interface{}, secretsURL interface{}, hash interface{}, contents interface{}) *ORM_Create_Call {
	return &ORM_Create_Call{Call: _e.mock.On("Create", ctx, secretsURL, hash, contents)}
}

func (_c *ORM_Create_Call) Run(run func(ctx context.Context, secretsURL string, hash string, contents string)) *ORM_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *ORM_Create_Call) Return(_a0 int64, _a1 error) *ORM_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ORM_Create_Call) RunAndReturn(run func(context.Context, string, string, string) (int64, error)) *ORM_Create_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteWorkflowSpec provides a mock function with given fields: ctx, owner, name
func (_m *ORM) DeleteWorkflowSpec(ctx context.Context, owner string, name string) error {
	ret := _m.Called(ctx, owner, name)

	if len(ret) == 0 {
		panic("no return value specified for DeleteWorkflowSpec")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, owner, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ORM_DeleteWorkflowSpec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteWorkflowSpec'
type ORM_DeleteWorkflowSpec_Call struct {
	*mock.Call
}

// DeleteWorkflowSpec is a helper method to define mock.On call
//   - ctx context.Context
//   - owner string
//   - name string
func (_e *ORM_Expecter) DeleteWorkflowSpec(ctx interface{}, owner interface{}, name interface{}) *ORM_DeleteWorkflowSpec_Call {
	return &ORM_DeleteWorkflowSpec_Call{Call: _e.mock.On("DeleteWorkflowSpec", ctx, owner, name)}
}

func (_c *ORM_DeleteWorkflowSpec_Call) Run(run func(ctx context.Context, owner string, name string)) *ORM_DeleteWorkflowSpec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ORM_DeleteWorkflowSpec_Call) Return(_a0 error) *ORM_DeleteWorkflowSpec_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ORM_DeleteWorkflowSpec_Call) RunAndReturn(run func(context.Context, string, string) error) *ORM_DeleteWorkflowSpec_Call {
	_c.Call.Return(run)
	return _c
}

// GetContents provides a mock function with given fields: ctx, url
func (_m *ORM) GetContents(ctx context.Context, url string) (string, error) {
	ret := _m.Called(ctx, url)

	if len(ret) == 0 {
		panic("no return value specified for GetContents")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, url)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, url)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ORM_GetContents_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetContents'
type ORM_GetContents_Call struct {
	*mock.Call
}

// GetContents is a helper method to define mock.On call
//   - ctx context.Context
//   - url string
func (_e *ORM_Expecter) GetContents(ctx interface{}, url interface{}) *ORM_GetContents_Call {
	return &ORM_GetContents_Call{Call: _e.mock.On("GetContents", ctx, url)}
}

func (_c *ORM_GetContents_Call) Run(run func(ctx context.Context, url string)) *ORM_GetContents_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ORM_GetContents_Call) Return(_a0 string, _a1 error) *ORM_GetContents_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ORM_GetContents_Call) RunAndReturn(run func(context.Context, string) (string, error)) *ORM_GetContents_Call {
	_c.Call.Return(run)
	return _c
}

// GetContentsByHash provides a mock function with given fields: ctx, hash
func (_m *ORM) GetContentsByHash(ctx context.Context, hash string) (string, error) {
	ret := _m.Called(ctx, hash)

	if len(ret) == 0 {
		panic("no return value specified for GetContentsByHash")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, hash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, hash)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ORM_GetContentsByHash_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetContentsByHash'
type ORM_GetContentsByHash_Call struct {
	*mock.Call
}

// GetContentsByHash is a helper method to define mock.On call
//   - ctx context.Context
//   - hash string
func (_e *ORM_Expecter) GetContentsByHash(ctx interface{}, hash interface{}) *ORM_GetContentsByHash_Call {
	return &ORM_GetContentsByHash_Call{Call: _e.mock.On("GetContentsByHash", ctx, hash)}
}

func (_c *ORM_GetContentsByHash_Call) Run(run func(ctx context.Context, hash string)) *ORM_GetContentsByHash_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ORM_GetContentsByHash_Call) Return(_a0 string, _a1 error) *ORM_GetContentsByHash_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ORM_GetContentsByHash_Call) RunAndReturn(run func(context.Context, string) (string, error)) *ORM_GetContentsByHash_Call {
	_c.Call.Return(run)
	return _c
}

// GetContentsByWorkflowID provides a mock function with given fields: ctx, workflowID
func (_m *ORM) GetContentsByWorkflowID(ctx context.Context, workflowID string) (string, string, error) {
	ret := _m.Called(ctx, workflowID)

	if len(ret) == 0 {
		panic("no return value specified for GetContentsByWorkflowID")
	}

	var r0 string
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, string, error)); ok {
		return rf(ctx, workflowID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, workflowID)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) string); ok {
		r1 = rf(ctx, workflowID)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, workflowID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ORM_GetContentsByWorkflowID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetContentsByWorkflowID'
type ORM_GetContentsByWorkflowID_Call struct {
	*mock.Call
}

// GetContentsByWorkflowID is a helper method to define mock.On call
//   - ctx context.Context
//   - workflowID string
func (_e *ORM_Expecter) GetContentsByWorkflowID(ctx interface{}, workflowID interface{}) *ORM_GetContentsByWorkflowID_Call {
	return &ORM_GetContentsByWorkflowID_Call{Call: _e.mock.On("GetContentsByWorkflowID", ctx, workflowID)}
}

func (_c *ORM_GetContentsByWorkflowID_Call) Run(run func(ctx context.Context, workflowID string)) *ORM_GetContentsByWorkflowID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ORM_GetContentsByWorkflowID_Call) Return(_a0 string, _a1 string, _a2 error) *ORM_GetContentsByWorkflowID_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ORM_GetContentsByWorkflowID_Call) RunAndReturn(run func(context.Context, string) (string, string, error)) *ORM_GetContentsByWorkflowID_Call {
	_c.Call.Return(run)
	return _c
}

// GetSecretsURLByHash provides a mock function with given fields: ctx, hash
func (_m *ORM) GetSecretsURLByHash(ctx context.Context, hash string) (string, error) {
	ret := _m.Called(ctx, hash)

	if len(ret) == 0 {
		panic("no return value specified for GetSecretsURLByHash")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, hash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, hash)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ORM_GetSecretsURLByHash_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSecretsURLByHash'
type ORM_GetSecretsURLByHash_Call struct {
	*mock.Call
}

// GetSecretsURLByHash is a helper method to define mock.On call
//   - ctx context.Context
//   - hash string
func (_e *ORM_Expecter) GetSecretsURLByHash(ctx interface{}, hash interface{}) *ORM_GetSecretsURLByHash_Call {
	return &ORM_GetSecretsURLByHash_Call{Call: _e.mock.On("GetSecretsURLByHash", ctx, hash)}
}

func (_c *ORM_GetSecretsURLByHash_Call) Run(run func(ctx context.Context, hash string)) *ORM_GetSecretsURLByHash_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ORM_GetSecretsURLByHash_Call) Return(_a0 string, _a1 error) *ORM_GetSecretsURLByHash_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ORM_GetSecretsURLByHash_Call) RunAndReturn(run func(context.Context, string) (string, error)) *ORM_GetSecretsURLByHash_Call {
	_c.Call.Return(run)
	return _c
}

// GetSecretsURLByID provides a mock function with given fields: ctx, id
func (_m *ORM) GetSecretsURLByID(ctx context.Context, id int64) (string, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetSecretsURLByID")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (string, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) string); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ORM_GetSecretsURLByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSecretsURLByID'
type ORM_GetSecretsURLByID_Call struct {
	*mock.Call
}

// GetSecretsURLByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *ORM_Expecter) GetSecretsURLByID(ctx interface{}, id interface{}) *ORM_GetSecretsURLByID_Call {
	return &ORM_GetSecretsURLByID_Call{Call: _e.mock.On("GetSecretsURLByID", ctx, id)}
}

func (_c *ORM_GetSecretsURLByID_Call) Run(run func(ctx context.Context, id int64)) *ORM_GetSecretsURLByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *ORM_GetSecretsURLByID_Call) Return(_a0 string, _a1 error) *ORM_GetSecretsURLByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ORM_GetSecretsURLByID_Call) RunAndReturn(run func(context.Context, int64) (string, error)) *ORM_GetSecretsURLByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetSecretsURLHash provides a mock function with given fields: owner, secretsURL
func (_m *ORM) GetSecretsURLHash(owner []byte, secretsURL []byte) ([]byte, error) {
	ret := _m.Called(owner, secretsURL)

	if len(ret) == 0 {
		panic("no return value specified for GetSecretsURLHash")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte, []byte) ([]byte, error)); ok {
		return rf(owner, secretsURL)
	}
	if rf, ok := ret.Get(0).(func([]byte, []byte) []byte); ok {
		r0 = rf(owner, secretsURL)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte, []byte) error); ok {
		r1 = rf(owner, secretsURL)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ORM_GetSecretsURLHash_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSecretsURLHash'
type ORM_GetSecretsURLHash_Call struct {
	*mock.Call
}

// GetSecretsURLHash is a helper method to define mock.On call
//   - owner []byte
//   - secretsURL []byte
func (_e *ORM_Expecter) GetSecretsURLHash(owner interface{}, secretsURL interface{}) *ORM_GetSecretsURLHash_Call {
	return &ORM_GetSecretsURLHash_Call{Call: _e.mock.On("GetSecretsURLHash", owner, secretsURL)}
}

func (_c *ORM_GetSecretsURLHash_Call) Run(run func(owner []byte, secretsURL []byte)) *ORM_GetSecretsURLHash_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte), args[1].([]byte))
	})
	return _c
}

func (_c *ORM_GetSecretsURLHash_Call) Return(_a0 []byte, _a1 error) *ORM_GetSecretsURLHash_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ORM_GetSecretsURLHash_Call) RunAndReturn(run func([]byte, []byte) ([]byte, error)) *ORM_GetSecretsURLHash_Call {
	_c.Call.Return(run)
	return _c
}

// GetWorkflowSpec provides a mock function with given fields: ctx, owner, name
func (_m *ORM) GetWorkflowSpec(ctx context.Context, owner string, name string) (*job.WorkflowSpec, error) {
	ret := _m.Called(ctx, owner, name)

	if len(ret) == 0 {
		panic("no return value specified for GetWorkflowSpec")
	}

	var r0 *job.WorkflowSpec
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*job.WorkflowSpec, error)); ok {
		return rf(ctx, owner, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *job.WorkflowSpec); ok {
		r0 = rf(ctx, owner, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*job.WorkflowSpec)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, owner, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ORM_GetWorkflowSpec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetWorkflowSpec'
type ORM_GetWorkflowSpec_Call struct {
	*mock.Call
}

// GetWorkflowSpec is a helper method to define mock.On call
//   - ctx context.Context
//   - owner string
//   - name string
func (_e *ORM_Expecter) GetWorkflowSpec(ctx interface{}, owner interface{}, name interface{}) *ORM_GetWorkflowSpec_Call {
	return &ORM_GetWorkflowSpec_Call{Call: _e.mock.On("GetWorkflowSpec", ctx, owner, name)}
}

func (_c *ORM_GetWorkflowSpec_Call) Run(run func(ctx context.Context, owner string, name string)) *ORM_GetWorkflowSpec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ORM_GetWorkflowSpec_Call) Return(_a0 *job.WorkflowSpec, _a1 error) *ORM_GetWorkflowSpec_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ORM_GetWorkflowSpec_Call) RunAndReturn(run func(context.Context, string, string) (*job.WorkflowSpec, error)) *ORM_GetWorkflowSpec_Call {
	_c.Call.Return(run)
	return _c
}

// GetWorkflowSpecByID provides a mock function with given fields: ctx, id
func (_m *ORM) GetWorkflowSpecByID(ctx context.Context, id string) (*job.WorkflowSpec, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetWorkflowSpecByID")
	}

	var r0 *job.WorkflowSpec
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*job.WorkflowSpec, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *job.WorkflowSpec); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*job.WorkflowSpec)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ORM_GetWorkflowSpecByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetWorkflowSpecByID'
type ORM_GetWorkflowSpecByID_Call struct {
	*mock.Call
}

// GetWorkflowSpecByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *ORM_Expecter) GetWorkflowSpecByID(ctx interface{}, id interface{}) *ORM_GetWorkflowSpecByID_Call {
	return &ORM_GetWorkflowSpecByID_Call{Call: _e.mock.On("GetWorkflowSpecByID", ctx, id)}
}

func (_c *ORM_GetWorkflowSpecByID_Call) Run(run func(ctx context.Context, id string)) *ORM_GetWorkflowSpecByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ORM_GetWorkflowSpecByID_Call) Return(_a0 *job.WorkflowSpec, _a1 error) *ORM_GetWorkflowSpecByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ORM_GetWorkflowSpecByID_Call) RunAndReturn(run func(context.Context, string) (*job.WorkflowSpec, error)) *ORM_GetWorkflowSpecByID_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, secretsURL, contents
func (_m *ORM) Update(ctx context.Context, secretsURL string, contents string) (int64, error) {
	ret := _m.Called(ctx, secretsURL, contents)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (int64, error)); ok {
		return rf(ctx, secretsURL, contents)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) int64); ok {
		r0 = rf(ctx, secretsURL, contents)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, secretsURL, contents)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ORM_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type ORM_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - secretsURL string
//   - contents string
func (_e *ORM_Expecter) Update(ctx interface{}, secretsURL interface{}, contents interface{}) *ORM_Update_Call {
	return &ORM_Update_Call{Call: _e.mock.On("Update", ctx, secretsURL, contents)}
}

func (_c *ORM_Update_Call) Run(run func(ctx context.Context, secretsURL string, contents string)) *ORM_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ORM_Update_Call) Return(_a0 int64, _a1 error) *ORM_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ORM_Update_Call) RunAndReturn(run func(context.Context, string, string) (int64, error)) *ORM_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UpsertWorkflowSpec provides a mock function with given fields: ctx, spec
func (_m *ORM) UpsertWorkflowSpec(ctx context.Context, spec *job.WorkflowSpec) (int64, error) {
	ret := _m.Called(ctx, spec)

	if len(ret) == 0 {
		panic("no return value specified for UpsertWorkflowSpec")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *job.WorkflowSpec) (int64, error)); ok {
		return rf(ctx, spec)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *job.WorkflowSpec) int64); ok {
		r0 = rf(ctx, spec)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *job.WorkflowSpec) error); ok {
		r1 = rf(ctx, spec)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ORM_UpsertWorkflowSpec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpsertWorkflowSpec'
type ORM_UpsertWorkflowSpec_Call struct {
	*mock.Call
}

// UpsertWorkflowSpec is a helper method to define mock.On call
//   - ctx context.Context
//   - spec *job.WorkflowSpec
func (_e *ORM_Expecter) UpsertWorkflowSpec(ctx interface{}, spec interface{}) *ORM_UpsertWorkflowSpec_Call {
	return &ORM_UpsertWorkflowSpec_Call{Call: _e.mock.On("UpsertWorkflowSpec", ctx, spec)}
}

func (_c *ORM_UpsertWorkflowSpec_Call) Run(run func(ctx context.Context, spec *job.WorkflowSpec)) *ORM_UpsertWorkflowSpec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*job.WorkflowSpec))
	})
	return _c
}

func (_c *ORM_UpsertWorkflowSpec_Call) Return(_a0 int64, _a1 error) *ORM_UpsertWorkflowSpec_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ORM_UpsertWorkflowSpec_Call) RunAndReturn(run func(context.Context, *job.WorkflowSpec) (int64, error)) *ORM_UpsertWorkflowSpec_Call {
	_c.Call.Return(run)
	return _c
}

// UpsertWorkflowSpecWithSecrets provides a mock function with given fields: ctx, spec, url, hash, contents
func (_m *ORM) UpsertWorkflowSpecWithSecrets(ctx context.Context, spec *job.WorkflowSpec, url string, hash string, contents string) (int64, error) {
	ret := _m.Called(ctx, spec, url, hash, contents)

	if len(ret) == 0 {
		panic("no return value specified for UpsertWorkflowSpecWithSecrets")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *job.WorkflowSpec, string, string, string) (int64, error)); ok {
		return rf(ctx, spec, url, hash, contents)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *job.WorkflowSpec, string, string, string) int64); ok {
		r0 = rf(ctx, spec, url, hash, contents)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *job.WorkflowSpec, string, string, string) error); ok {
		r1 = rf(ctx, spec, url, hash, contents)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ORM_UpsertWorkflowSpecWithSecrets_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpsertWorkflowSpecWithSecrets'
type ORM_UpsertWorkflowSpecWithSecrets_Call struct {
	*mock.Call
}

// UpsertWorkflowSpecWithSecrets is a helper method to define mock.On call
//   - ctx context.Context
//   - spec *job.WorkflowSpec
//   - url string
//   - hash string
//   - contents string
func (_e *ORM_Expecter) UpsertWorkflowSpecWithSecrets(ctx interface{}, spec interface{}, url interface{}, hash interface{}, contents interface{}) *ORM_UpsertWorkflowSpecWithSecrets_Call {
	return &ORM_UpsertWorkflowSpecWithSecrets_Call{Call: _e.mock.On("UpsertWorkflowSpecWithSecrets", ctx, spec, url, hash, contents)}
}

func (_c *ORM_UpsertWorkflowSpecWithSecrets_Call) Run(run func(ctx context.Context, spec *job.WorkflowSpec, url string, hash string, contents string)) *ORM_UpsertWorkflowSpecWithSecrets_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*job.WorkflowSpec), args[2].(string), args[3].(string), args[4].(string))
	})
	return _c
}

func (_c *ORM_UpsertWorkflowSpecWithSecrets_Call) Return(_a0 int64, _a1 error) *ORM_UpsertWorkflowSpecWithSecrets_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ORM_UpsertWorkflowSpecWithSecrets_Call) RunAndReturn(run func(context.Context, *job.WorkflowSpec, string, string, string) (int64, error)) *ORM_UpsertWorkflowSpecWithSecrets_Call {
	_c.Call.Return(run)
	return _c
}

// NewORM creates a new instance of ORM. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewORM(t interface {
	mock.TestingT
	Cleanup(func())
}) *ORM {
	mock := &ORM{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
