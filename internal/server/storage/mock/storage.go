// Code generated by MockGen. DO NOT EDIT.
// Source: ./interface.go

// Package storagemock is a generated GoMock package.
package storagemock

import (
	context "context"
	model "gophkeeper/internal/server/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockUserRepository) Create(ctx context.Context, m *model.User) (*model.User, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, m)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserRepositoryMockRecorder) Create(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRepository)(nil).Create), ctx, m)
}

// Read mocks base method.
func (m *MockUserRepository) Read(ctx context.Context, id uuid.UUID) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, id)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockUserRepositoryMockRecorder) Read(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockUserRepository)(nil).Read), ctx, id)
}

// ReadByEmailAndPassword mocks base method.
func (m *MockUserRepository) ReadByEmailAndPassword(ctx context.Context, name, password string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadByEmailAndPassword", ctx, name, password)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadByEmailAndPassword indicates an expected call of ReadByEmailAndPassword.
func (mr *MockUserRepositoryMockRecorder) ReadByEmailAndPassword(ctx, name, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadByEmailAndPassword", reflect.TypeOf((*MockUserRepository)(nil).ReadByEmailAndPassword), ctx, name, password)
}

// MockSecretRepository is a mock of SecretRepository interface.
type MockSecretRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSecretRepositoryMockRecorder
}

// MockSecretRepositoryMockRecorder is the mock recorder for MockSecretRepository.
type MockSecretRepositoryMockRecorder struct {
	mock *MockSecretRepository
}

// NewMockSecretRepository creates a new mock instance.
func NewMockSecretRepository(ctrl *gomock.Controller) *MockSecretRepository {
	mock := &MockSecretRepository{ctrl: ctrl}
	mock.recorder = &MockSecretRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretRepository) EXPECT() *MockSecretRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockSecretRepository) Create(ctx context.Context, uid uuid.UUID, m *model.Secret) (*model.Secret, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, uid, m)
	ret0, _ := ret[0].(*model.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockSecretRepositoryMockRecorder) Create(ctx, uid, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSecretRepository)(nil).Create), ctx, uid, m)
}

// DeleteByName mocks base method.
func (m *MockSecretRepository) DeleteByName(ctx context.Context, uid uuid.UUID, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByName", ctx, uid, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByName indicates an expected call of DeleteByName.
func (mr *MockSecretRepositoryMockRecorder) DeleteByName(ctx, uid, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByName", reflect.TypeOf((*MockSecretRepository)(nil).DeleteByName), ctx, uid, name)
}

// List mocks base method.
func (m *MockSecretRepository) List(ctx context.Context, uid uuid.UUID) ([]*model.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, uid)
	ret0, _ := ret[0].([]*model.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockSecretRepositoryMockRecorder) List(ctx, uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSecretRepository)(nil).List), ctx, uid)
}

// ReadByName mocks base method.
func (m *MockSecretRepository) ReadByName(ctx context.Context, uid uuid.UUID, name string) (*model.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadByName", ctx, uid, name)
	ret0, _ := ret[0].(*model.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadByName indicates an expected call of ReadByName.
func (mr *MockSecretRepositoryMockRecorder) ReadByName(ctx, uid, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadByName", reflect.TypeOf((*MockSecretRepository)(nil).ReadByName), ctx, uid, name)
}
