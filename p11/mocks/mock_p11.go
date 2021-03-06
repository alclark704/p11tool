// Code generated by MockGen. DO NOT EDIT.
// Source: ./mocks/../p11.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	pkcs11 "github.com/miekg/pkcs11"
	reflect "reflect"
)

// MockTokenCtx is a mock of TokenCtx interface
type MockTokenCtx struct {
	ctrl     *gomock.Controller
	recorder *MockTokenCtxMockRecorder
}

// MockTokenCtxMockRecorder is the mock recorder for MockTokenCtx
type MockTokenCtxMockRecorder struct {
	mock *MockTokenCtx
}

// NewMockTokenCtx creates a new mock instance
func NewMockTokenCtx(ctrl *gomock.Controller) *MockTokenCtx {
	mock := &MockTokenCtx{ctrl: ctrl}
	mock.recorder = &MockTokenCtxMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTokenCtx) EXPECT() *MockTokenCtxMockRecorder {
	return m.recorder
}

// CloseSession mocks base method
func (m *MockTokenCtx) CloseSession(sh pkcs11.SessionHandle) error {
	ret := m.ctrl.Call(m, "CloseSession", sh)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSession indicates an expected call of CloseSession
func (mr *MockTokenCtxMockRecorder) CloseSession(sh interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSession", reflect.TypeOf((*MockTokenCtx)(nil).CloseSession), sh)
}

// CreateObject mocks base method
func (m *MockTokenCtx) CreateObject(sh pkcs11.SessionHandle, temp []*pkcs11.Attribute) (pkcs11.ObjectHandle, error) {
	ret := m.ctrl.Call(m, "CreateObject", sh, temp)
	ret0, _ := ret[0].(pkcs11.ObjectHandle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateObject indicates an expected call of CreateObject
func (mr *MockTokenCtxMockRecorder) CreateObject(sh, temp interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateObject", reflect.TypeOf((*MockTokenCtx)(nil).CreateObject), sh, temp)
}

// Destroy mocks base method
func (m *MockTokenCtx) Destroy() {
	m.ctrl.Call(m, "Destroy")
}

// Destroy indicates an expected call of Destroy
func (mr *MockTokenCtxMockRecorder) Destroy() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockTokenCtx)(nil).Destroy))
}

// DestroyObject mocks base method
func (m *MockTokenCtx) DestroyObject(sh pkcs11.SessionHandle, oh pkcs11.ObjectHandle) error {
	ret := m.ctrl.Call(m, "DestroyObject", sh, oh)
	ret0, _ := ret[0].(error)
	return ret0
}

// DestroyObject indicates an expected call of DestroyObject
func (mr *MockTokenCtxMockRecorder) DestroyObject(sh, oh interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyObject", reflect.TypeOf((*MockTokenCtx)(nil).DestroyObject), sh, oh)
}

// Encrypt mocks base method
func (m *MockTokenCtx) Encrypt(sh pkcs11.SessionHandle, message []byte) ([]byte, error) {
	ret := m.ctrl.Call(m, "Encrypt", sh, message)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encrypt indicates an expected call of Encrypt
func (mr *MockTokenCtxMockRecorder) Encrypt(sh, message interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encrypt", reflect.TypeOf((*MockTokenCtx)(nil).Encrypt), sh, message)
}

// EncryptInit mocks base method
func (m_2 *MockTokenCtx) EncryptInit(sh pkcs11.SessionHandle, m []*pkcs11.Mechanism, o pkcs11.ObjectHandle) error {
	ret := m_2.ctrl.Call(m_2, "EncryptInit", sh, m, o)
	ret0, _ := ret[0].(error)
	return ret0
}

// EncryptInit indicates an expected call of EncryptInit
func (mr *MockTokenCtxMockRecorder) EncryptInit(sh, m, o interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EncryptInit", reflect.TypeOf((*MockTokenCtx)(nil).EncryptInit), sh, m, o)
}

// Finalize mocks base method
func (m *MockTokenCtx) Finalize() error {
	ret := m.ctrl.Call(m, "Finalize")
	ret0, _ := ret[0].(error)
	return ret0
}

// Finalize indicates an expected call of Finalize
func (mr *MockTokenCtxMockRecorder) Finalize() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finalize", reflect.TypeOf((*MockTokenCtx)(nil).Finalize))
}

// FindObjects mocks base method
func (m *MockTokenCtx) FindObjects(sh pkcs11.SessionHandle, max int) ([]pkcs11.ObjectHandle, bool, error) {
	ret := m.ctrl.Call(m, "FindObjects", sh, max)
	ret0, _ := ret[0].([]pkcs11.ObjectHandle)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindObjects indicates an expected call of FindObjects
func (mr *MockTokenCtxMockRecorder) FindObjects(sh, max interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindObjects", reflect.TypeOf((*MockTokenCtx)(nil).FindObjects), sh, max)
}

// FindObjectsFinal mocks base method
func (m *MockTokenCtx) FindObjectsFinal(sh pkcs11.SessionHandle) error {
	ret := m.ctrl.Call(m, "FindObjectsFinal", sh)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindObjectsFinal indicates an expected call of FindObjectsFinal
func (mr *MockTokenCtxMockRecorder) FindObjectsFinal(sh interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindObjectsFinal", reflect.TypeOf((*MockTokenCtx)(nil).FindObjectsFinal), sh)
}

// FindObjectsInit mocks base method
func (m *MockTokenCtx) FindObjectsInit(sh pkcs11.SessionHandle, temp []*pkcs11.Attribute) error {
	ret := m.ctrl.Call(m, "FindObjectsInit", sh, temp)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindObjectsInit indicates an expected call of FindObjectsInit
func (mr *MockTokenCtxMockRecorder) FindObjectsInit(sh, temp interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindObjectsInit", reflect.TypeOf((*MockTokenCtx)(nil).FindObjectsInit), sh, temp)
}

// GenerateKey mocks base method
func (m *MockTokenCtx) GenerateKey(sh pkcs11.SessionHandle, mech []*pkcs11.Mechanism, temp []*pkcs11.Attribute) (pkcs11.ObjectHandle, error) {
	ret := m.ctrl.Call(m, "GenerateKey", sh, mech, temp)
	ret0, _ := ret[0].(pkcs11.ObjectHandle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateKey indicates an expected call of GetAttributeValue
func (mr *MockTokenCtxMockRecorder) GenerateKey(sh, mech, temp interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateKey", reflect.TypeOf((*MockTokenCtx)(nil).GenerateKey), sh, mech, temp)
}

// GenerateKeyPair mocks base method
func (m *MockTokenCtx) GenerateKeyPair(sh pkcs11.SessionHandle, mech []*pkcs11.Mechanism, public, private []*pkcs11.Attribute) (pkcs11.ObjectHandle, pkcs11.ObjectHandle, error) {
	ret := m.ctrl.Call(m, "GenerateKeyPair", sh, mech, public, private)
	ret0, _ := ret[0].(pkcs11.ObjectHandle)
	ret1, _ := ret[1].(pkcs11.ObjectHandle)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GenerateKeyPair indicates an expected call of GetAttributeValue
func (mr *MockTokenCtxMockRecorder) GenerateKeyPair(sh, mech, public, private interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateKeyPair", reflect.TypeOf((*MockTokenCtx)(nil).GenerateKeyPair), sh, mech, public, private)
}

// GetAttributeValue mocks base method
func (m *MockTokenCtx) GetAttributeValue(sh pkcs11.SessionHandle, o pkcs11.ObjectHandle, a []*pkcs11.Attribute) ([]*pkcs11.Attribute, error) {
	ret := m.ctrl.Call(m, "GetAttributeValue", sh, o, a)
	ret0, _ := ret[0].([]*pkcs11.Attribute)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAttributeValue indicates an expected call of GetAttributeValue
func (mr *MockTokenCtxMockRecorder) GetAttributeValue(sh, o, a interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttributeValue", reflect.TypeOf((*MockTokenCtx)(nil).GetAttributeValue), sh, o, a)
}

// GetSlotList mocks base method
func (m *MockTokenCtx) GetSlotList(tokenPresent bool) ([]uint, error) {
	ret := m.ctrl.Call(m, "GetSlotList", tokenPresent)
	ret0, _ := ret[0].([]uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSlotList indicates an expected call of GetSlotList
func (mr *MockTokenCtxMockRecorder) GetSlotList(tokenPresent interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSlotList", reflect.TypeOf((*MockTokenCtx)(nil).GetSlotList), tokenPresent)
}

// GetTokenInfo mocks base method
func (m *MockTokenCtx) GetTokenInfo(slotID uint) (pkcs11.TokenInfo, error) {
	ret := m.ctrl.Call(m, "GetTokenInfo", slotID)
	ret0, _ := ret[0].(pkcs11.TokenInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTokenInfo indicates an expected call of GetTokenInfo
func (mr *MockTokenCtxMockRecorder) GetTokenInfo(slotID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTokenInfo", reflect.TypeOf((*MockTokenCtx)(nil).GetTokenInfo), slotID)
}

// Initialize mocks base method
func (m *MockTokenCtx) Initialize() error {
	ret := m.ctrl.Call(m, "Initialize")
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize
func (mr *MockTokenCtxMockRecorder) Initialize() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockTokenCtx)(nil).Initialize))
}

// Login mocks base method
func (m *MockTokenCtx) Login(sh pkcs11.SessionHandle, userType uint, pin string) error {
	ret := m.ctrl.Call(m, "Login", sh, userType, pin)
	ret0, _ := ret[0].(error)
	return ret0
}

// Login indicates an expected call of Login
func (mr *MockTokenCtxMockRecorder) Login(sh, userType, pin interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockTokenCtx)(nil).Login), sh, userType, pin)
}

// OpenSession mocks base method
func (m *MockTokenCtx) OpenSession(slotID, flags uint) (pkcs11.SessionHandle, error) {
	ret := m.ctrl.Call(m, "OpenSession", slotID, flags)
	ret0, _ := ret[0].(pkcs11.SessionHandle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenSession indicates an expected call of OpenSession
func (mr *MockTokenCtxMockRecorder) OpenSession(slotID, flags interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenSession", reflect.TypeOf((*MockTokenCtx)(nil).OpenSession), slotID, flags)
}

// MockToken is a mock of Token interface
type MockToken struct {
	ctrl     *gomock.Controller
	recorder *MockTokenMockRecorder
}

// MockTokenMockRecorder is the mock recorder for MockToken
type MockTokenMockRecorder struct {
	mock *MockToken
}

// NewMockToken creates a new mock instance
func NewMockToken(ctrl *gomock.Controller) *MockToken {
	mock := &MockToken{ctrl: ctrl}
	mock.recorder = &MockTokenMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockToken) EXPECT() *MockTokenMockRecorder {
	return m.recorder
}

// Checksum mocks base method
func (m *MockToken) Checksum(keyLabel string) ([]byte, error) {
	ret := m.ctrl.Call(m, "Checksum", keyLabel)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Checksum indicates an expected call of Checksum
func (mr *MockTokenMockRecorder) Checksum(keyLabel interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Checksum", reflect.TypeOf((*MockToken)(nil).Checksum), keyLabel)
}

// ImportKey mocks base method
func (m *MockToken) ImportKey(keyBytes []byte, label string) error {
	ret := m.ctrl.Call(m, "ImportKey", keyBytes, label)
	ret0, _ := ret[0].(error)
	return ret0
}

// ImportKey indicates an expected call of ImportKey
func (mr *MockTokenMockRecorder) ImportKey(keyBytes, label interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportKey", reflect.TypeOf((*MockToken)(nil).ImportKey), keyBytes, label)
}

// DeleteAllExcept mocks base method
func (m *MockToken) DeleteAllExcept(keyLabels []string) error {
	ret := m.ctrl.Call(m, "DeleteAllExcept", keyLabels)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllExcept indicates an expected call of DeleteAllExcept
func (mr *MockTokenMockRecorder) DeleteAllExcept(keyLabels interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllExcept", reflect.TypeOf((*MockToken)(nil).DeleteAllExcept), keyLabels)
}

// PrintObjects mocks base method
func (m *MockToken) PrintObjects(label *string) error {
	ret := m.ctrl.Call(m, "PrintObjects", label)
	ret0, _ := ret[0].(error)
	return ret0
}

// PrintObjects indicates an expected call of PrintObjects
func (mr *MockTokenMockRecorder) PrintObjects(label interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrintObjects", reflect.TypeOf((*MockToken)(nil).PrintObjects), label)
}

// Finalise mocks base method
func (m *MockToken) Finalise() error {
	ret := m.ctrl.Call(m, "Finalise")
	ret0, _ := ret[0].(error)
	return ret0
}

// Finalise indicates an expected call of Finalise
func (mr *MockTokenMockRecorder) Finalise() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finalise", reflect.TypeOf((*MockToken)(nil).Finalise))
}
