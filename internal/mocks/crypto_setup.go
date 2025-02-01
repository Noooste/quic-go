// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Noooste/quic-go/internal/handshake (interfaces: CryptoSetup)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package mocks -destination crypto_setup_tmp.go github.com/Noooste/quic-go/internal/handshake CryptoSetup
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	handshake "github.com/Noooste/quic-go/internal/handshake"
	protocol "github.com/Noooste/quic-go/internal/protocol"
	gomock "go.uber.org/mock/gomock"
)

// MockCryptoSetup is a mock of CryptoSetup interface.
type MockCryptoSetup struct {
	ctrl     *gomock.Controller
	recorder *MockCryptoSetupMockRecorder
	isgomock struct{}
}

// MockCryptoSetupMockRecorder is the mock recorder for MockCryptoSetup.
type MockCryptoSetupMockRecorder struct {
	mock *MockCryptoSetup
}

// NewMockCryptoSetup creates a new mock instance.
func NewMockCryptoSetup(ctrl *gomock.Controller) *MockCryptoSetup {
	mock := &MockCryptoSetup{ctrl: ctrl}
	mock.recorder = &MockCryptoSetupMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCryptoSetup) EXPECT() *MockCryptoSetupMockRecorder {
	return m.recorder
}

// ChangeConnectionID mocks base method.
func (m *MockCryptoSetup) ChangeConnectionID(arg0 protocol.ConnectionID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ChangeConnectionID", arg0)
}

// ChangeConnectionID indicates an expected call of ChangeConnectionID.
func (mr *MockCryptoSetupMockRecorder) ChangeConnectionID(arg0 any) *MockCryptoSetupChangeConnectionIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeConnectionID", reflect.TypeOf((*MockCryptoSetup)(nil).ChangeConnectionID), arg0)
	return &MockCryptoSetupChangeConnectionIDCall{Call: call}
}

// MockCryptoSetupChangeConnectionIDCall wrap *gomock.Call
type MockCryptoSetupChangeConnectionIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCryptoSetupChangeConnectionIDCall) Return() *MockCryptoSetupChangeConnectionIDCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCryptoSetupChangeConnectionIDCall) Do(f func(protocol.ConnectionID)) *MockCryptoSetupChangeConnectionIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCryptoSetupChangeConnectionIDCall) DoAndReturn(f func(protocol.ConnectionID)) *MockCryptoSetupChangeConnectionIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Close mocks base method.
func (m *MockCryptoSetup) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockCryptoSetupMockRecorder) Close() *MockCryptoSetupCloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCryptoSetup)(nil).Close))
	return &MockCryptoSetupCloseCall{Call: call}
}

// MockCryptoSetupCloseCall wrap *gomock.Call
type MockCryptoSetupCloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCryptoSetupCloseCall) Return(arg0 error) *MockCryptoSetupCloseCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCryptoSetupCloseCall) Do(f func() error) *MockCryptoSetupCloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCryptoSetupCloseCall) DoAndReturn(f func() error) *MockCryptoSetupCloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ConnectionState mocks base method.
func (m *MockCryptoSetup) ConnectionState() handshake.ConnectionState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectionState")
	ret0, _ := ret[0].(handshake.ConnectionState)
	return ret0
}

// ConnectionState indicates an expected call of ConnectionState.
func (mr *MockCryptoSetupMockRecorder) ConnectionState() *MockCryptoSetupConnectionStateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectionState", reflect.TypeOf((*MockCryptoSetup)(nil).ConnectionState))
	return &MockCryptoSetupConnectionStateCall{Call: call}
}

// MockCryptoSetupConnectionStateCall wrap *gomock.Call
type MockCryptoSetupConnectionStateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCryptoSetupConnectionStateCall) Return(arg0 handshake.ConnectionState) *MockCryptoSetupConnectionStateCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCryptoSetupConnectionStateCall) Do(f func() handshake.ConnectionState) *MockCryptoSetupConnectionStateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCryptoSetupConnectionStateCall) DoAndReturn(f func() handshake.ConnectionState) *MockCryptoSetupConnectionStateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DiscardInitialKeys mocks base method.
func (m *MockCryptoSetup) DiscardInitialKeys() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DiscardInitialKeys")
}

// DiscardInitialKeys indicates an expected call of DiscardInitialKeys.
func (mr *MockCryptoSetupMockRecorder) DiscardInitialKeys() *MockCryptoSetupDiscardInitialKeysCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscardInitialKeys", reflect.TypeOf((*MockCryptoSetup)(nil).DiscardInitialKeys))
	return &MockCryptoSetupDiscardInitialKeysCall{Call: call}
}

// MockCryptoSetupDiscardInitialKeysCall wrap *gomock.Call
type MockCryptoSetupDiscardInitialKeysCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCryptoSetupDiscardInitialKeysCall) Return() *MockCryptoSetupDiscardInitialKeysCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCryptoSetupDiscardInitialKeysCall) Do(f func()) *MockCryptoSetupDiscardInitialKeysCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCryptoSetupDiscardInitialKeysCall) DoAndReturn(f func()) *MockCryptoSetupDiscardInitialKeysCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Get0RTTOpener mocks base method.
func (m *MockCryptoSetup) Get0RTTOpener() (handshake.LongHeaderOpener, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get0RTTOpener")
	ret0, _ := ret[0].(handshake.LongHeaderOpener)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get0RTTOpener indicates an expected call of Get0RTTOpener.
func (mr *MockCryptoSetupMockRecorder) Get0RTTOpener() *MockCryptoSetupGet0RTTOpenerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get0RTTOpener", reflect.TypeOf((*MockCryptoSetup)(nil).Get0RTTOpener))
	return &MockCryptoSetupGet0RTTOpenerCall{Call: call}
}

// MockCryptoSetupGet0RTTOpenerCall wrap *gomock.Call
type MockCryptoSetupGet0RTTOpenerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCryptoSetupGet0RTTOpenerCall) Return(arg0 handshake.LongHeaderOpener, arg1 error) *MockCryptoSetupGet0RTTOpenerCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCryptoSetupGet0RTTOpenerCall) Do(f func() (handshake.LongHeaderOpener, error)) *MockCryptoSetupGet0RTTOpenerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCryptoSetupGet0RTTOpenerCall) DoAndReturn(f func() (handshake.LongHeaderOpener, error)) *MockCryptoSetupGet0RTTOpenerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Get0RTTSealer mocks base method.
func (m *MockCryptoSetup) Get0RTTSealer() (handshake.LongHeaderSealer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get0RTTSealer")
	ret0, _ := ret[0].(handshake.LongHeaderSealer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get0RTTSealer indicates an expected call of Get0RTTSealer.
func (mr *MockCryptoSetupMockRecorder) Get0RTTSealer() *MockCryptoSetupGet0RTTSealerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get0RTTSealer", reflect.TypeOf((*MockCryptoSetup)(nil).Get0RTTSealer))
	return &MockCryptoSetupGet0RTTSealerCall{Call: call}
}

// MockCryptoSetupGet0RTTSealerCall wrap *gomock.Call
type MockCryptoSetupGet0RTTSealerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCryptoSetupGet0RTTSealerCall) Return(arg0 handshake.LongHeaderSealer, arg1 error) *MockCryptoSetupGet0RTTSealerCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCryptoSetupGet0RTTSealerCall) Do(f func() (handshake.LongHeaderSealer, error)) *MockCryptoSetupGet0RTTSealerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCryptoSetupGet0RTTSealerCall) DoAndReturn(f func() (handshake.LongHeaderSealer, error)) *MockCryptoSetupGet0RTTSealerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Get1RTTOpener mocks base method.
func (m *MockCryptoSetup) Get1RTTOpener() (handshake.ShortHeaderOpener, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get1RTTOpener")
	ret0, _ := ret[0].(handshake.ShortHeaderOpener)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get1RTTOpener indicates an expected call of Get1RTTOpener.
func (mr *MockCryptoSetupMockRecorder) Get1RTTOpener() *MockCryptoSetupGet1RTTOpenerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get1RTTOpener", reflect.TypeOf((*MockCryptoSetup)(nil).Get1RTTOpener))
	return &MockCryptoSetupGet1RTTOpenerCall{Call: call}
}

// MockCryptoSetupGet1RTTOpenerCall wrap *gomock.Call
type MockCryptoSetupGet1RTTOpenerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCryptoSetupGet1RTTOpenerCall) Return(arg0 handshake.ShortHeaderOpener, arg1 error) *MockCryptoSetupGet1RTTOpenerCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCryptoSetupGet1RTTOpenerCall) Do(f func() (handshake.ShortHeaderOpener, error)) *MockCryptoSetupGet1RTTOpenerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCryptoSetupGet1RTTOpenerCall) DoAndReturn(f func() (handshake.ShortHeaderOpener, error)) *MockCryptoSetupGet1RTTOpenerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Get1RTTSealer mocks base method.
func (m *MockCryptoSetup) Get1RTTSealer() (handshake.ShortHeaderSealer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get1RTTSealer")
	ret0, _ := ret[0].(handshake.ShortHeaderSealer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get1RTTSealer indicates an expected call of Get1RTTSealer.
func (mr *MockCryptoSetupMockRecorder) Get1RTTSealer() *MockCryptoSetupGet1RTTSealerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get1RTTSealer", reflect.TypeOf((*MockCryptoSetup)(nil).Get1RTTSealer))
	return &MockCryptoSetupGet1RTTSealerCall{Call: call}
}

// MockCryptoSetupGet1RTTSealerCall wrap *gomock.Call
type MockCryptoSetupGet1RTTSealerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCryptoSetupGet1RTTSealerCall) Return(arg0 handshake.ShortHeaderSealer, arg1 error) *MockCryptoSetupGet1RTTSealerCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCryptoSetupGet1RTTSealerCall) Do(f func() (handshake.ShortHeaderSealer, error)) *MockCryptoSetupGet1RTTSealerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCryptoSetupGet1RTTSealerCall) DoAndReturn(f func() (handshake.ShortHeaderSealer, error)) *MockCryptoSetupGet1RTTSealerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetHandshakeOpener mocks base method.
func (m *MockCryptoSetup) GetHandshakeOpener() (handshake.LongHeaderOpener, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHandshakeOpener")
	ret0, _ := ret[0].(handshake.LongHeaderOpener)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHandshakeOpener indicates an expected call of GetHandshakeOpener.
func (mr *MockCryptoSetupMockRecorder) GetHandshakeOpener() *MockCryptoSetupGetHandshakeOpenerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHandshakeOpener", reflect.TypeOf((*MockCryptoSetup)(nil).GetHandshakeOpener))
	return &MockCryptoSetupGetHandshakeOpenerCall{Call: call}
}

// MockCryptoSetupGetHandshakeOpenerCall wrap *gomock.Call
type MockCryptoSetupGetHandshakeOpenerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCryptoSetupGetHandshakeOpenerCall) Return(arg0 handshake.LongHeaderOpener, arg1 error) *MockCryptoSetupGetHandshakeOpenerCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCryptoSetupGetHandshakeOpenerCall) Do(f func() (handshake.LongHeaderOpener, error)) *MockCryptoSetupGetHandshakeOpenerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCryptoSetupGetHandshakeOpenerCall) DoAndReturn(f func() (handshake.LongHeaderOpener, error)) *MockCryptoSetupGetHandshakeOpenerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetHandshakeSealer mocks base method.
func (m *MockCryptoSetup) GetHandshakeSealer() (handshake.LongHeaderSealer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHandshakeSealer")
	ret0, _ := ret[0].(handshake.LongHeaderSealer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHandshakeSealer indicates an expected call of GetHandshakeSealer.
func (mr *MockCryptoSetupMockRecorder) GetHandshakeSealer() *MockCryptoSetupGetHandshakeSealerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHandshakeSealer", reflect.TypeOf((*MockCryptoSetup)(nil).GetHandshakeSealer))
	return &MockCryptoSetupGetHandshakeSealerCall{Call: call}
}

// MockCryptoSetupGetHandshakeSealerCall wrap *gomock.Call
type MockCryptoSetupGetHandshakeSealerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCryptoSetupGetHandshakeSealerCall) Return(arg0 handshake.LongHeaderSealer, arg1 error) *MockCryptoSetupGetHandshakeSealerCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCryptoSetupGetHandshakeSealerCall) Do(f func() (handshake.LongHeaderSealer, error)) *MockCryptoSetupGetHandshakeSealerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCryptoSetupGetHandshakeSealerCall) DoAndReturn(f func() (handshake.LongHeaderSealer, error)) *MockCryptoSetupGetHandshakeSealerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetInitialOpener mocks base method.
func (m *MockCryptoSetup) GetInitialOpener() (handshake.LongHeaderOpener, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInitialOpener")
	ret0, _ := ret[0].(handshake.LongHeaderOpener)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInitialOpener indicates an expected call of GetInitialOpener.
func (mr *MockCryptoSetupMockRecorder) GetInitialOpener() *MockCryptoSetupGetInitialOpenerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInitialOpener", reflect.TypeOf((*MockCryptoSetup)(nil).GetInitialOpener))
	return &MockCryptoSetupGetInitialOpenerCall{Call: call}
}

// MockCryptoSetupGetInitialOpenerCall wrap *gomock.Call
type MockCryptoSetupGetInitialOpenerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCryptoSetupGetInitialOpenerCall) Return(arg0 handshake.LongHeaderOpener, arg1 error) *MockCryptoSetupGetInitialOpenerCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCryptoSetupGetInitialOpenerCall) Do(f func() (handshake.LongHeaderOpener, error)) *MockCryptoSetupGetInitialOpenerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCryptoSetupGetInitialOpenerCall) DoAndReturn(f func() (handshake.LongHeaderOpener, error)) *MockCryptoSetupGetInitialOpenerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetInitialSealer mocks base method.
func (m *MockCryptoSetup) GetInitialSealer() (handshake.LongHeaderSealer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInitialSealer")
	ret0, _ := ret[0].(handshake.LongHeaderSealer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInitialSealer indicates an expected call of GetInitialSealer.
func (mr *MockCryptoSetupMockRecorder) GetInitialSealer() *MockCryptoSetupGetInitialSealerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInitialSealer", reflect.TypeOf((*MockCryptoSetup)(nil).GetInitialSealer))
	return &MockCryptoSetupGetInitialSealerCall{Call: call}
}

// MockCryptoSetupGetInitialSealerCall wrap *gomock.Call
type MockCryptoSetupGetInitialSealerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCryptoSetupGetInitialSealerCall) Return(arg0 handshake.LongHeaderSealer, arg1 error) *MockCryptoSetupGetInitialSealerCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCryptoSetupGetInitialSealerCall) Do(f func() (handshake.LongHeaderSealer, error)) *MockCryptoSetupGetInitialSealerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCryptoSetupGetInitialSealerCall) DoAndReturn(f func() (handshake.LongHeaderSealer, error)) *MockCryptoSetupGetInitialSealerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetSessionTicket mocks base method.
func (m *MockCryptoSetup) GetSessionTicket() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionTicket")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionTicket indicates an expected call of GetSessionTicket.
func (mr *MockCryptoSetupMockRecorder) GetSessionTicket() *MockCryptoSetupGetSessionTicketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionTicket", reflect.TypeOf((*MockCryptoSetup)(nil).GetSessionTicket))
	return &MockCryptoSetupGetSessionTicketCall{Call: call}
}

// MockCryptoSetupGetSessionTicketCall wrap *gomock.Call
type MockCryptoSetupGetSessionTicketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCryptoSetupGetSessionTicketCall) Return(arg0 []byte, arg1 error) *MockCryptoSetupGetSessionTicketCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCryptoSetupGetSessionTicketCall) Do(f func() ([]byte, error)) *MockCryptoSetupGetSessionTicketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCryptoSetupGetSessionTicketCall) DoAndReturn(f func() ([]byte, error)) *MockCryptoSetupGetSessionTicketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// HandleMessage mocks base method.
func (m *MockCryptoSetup) HandleMessage(arg0 []byte, arg1 protocol.EncryptionLevel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleMessage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleMessage indicates an expected call of HandleMessage.
func (mr *MockCryptoSetupMockRecorder) HandleMessage(arg0, arg1 any) *MockCryptoSetupHandleMessageCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleMessage", reflect.TypeOf((*MockCryptoSetup)(nil).HandleMessage), arg0, arg1)
	return &MockCryptoSetupHandleMessageCall{Call: call}
}

// MockCryptoSetupHandleMessageCall wrap *gomock.Call
type MockCryptoSetupHandleMessageCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCryptoSetupHandleMessageCall) Return(arg0 error) *MockCryptoSetupHandleMessageCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCryptoSetupHandleMessageCall) Do(f func([]byte, protocol.EncryptionLevel) error) *MockCryptoSetupHandleMessageCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCryptoSetupHandleMessageCall) DoAndReturn(f func([]byte, protocol.EncryptionLevel) error) *MockCryptoSetupHandleMessageCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// NextEvent mocks base method.
func (m *MockCryptoSetup) NextEvent() handshake.Event {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextEvent")
	ret0, _ := ret[0].(handshake.Event)
	return ret0
}

// NextEvent indicates an expected call of NextEvent.
func (mr *MockCryptoSetupMockRecorder) NextEvent() *MockCryptoSetupNextEventCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextEvent", reflect.TypeOf((*MockCryptoSetup)(nil).NextEvent))
	return &MockCryptoSetupNextEventCall{Call: call}
}

// MockCryptoSetupNextEventCall wrap *gomock.Call
type MockCryptoSetupNextEventCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCryptoSetupNextEventCall) Return(arg0 handshake.Event) *MockCryptoSetupNextEventCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCryptoSetupNextEventCall) Do(f func() handshake.Event) *MockCryptoSetupNextEventCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCryptoSetupNextEventCall) DoAndReturn(f func() handshake.Event) *MockCryptoSetupNextEventCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetHandshakeConfirmed mocks base method.
func (m *MockCryptoSetup) SetHandshakeConfirmed() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetHandshakeConfirmed")
}

// SetHandshakeConfirmed indicates an expected call of SetHandshakeConfirmed.
func (mr *MockCryptoSetupMockRecorder) SetHandshakeConfirmed() *MockCryptoSetupSetHandshakeConfirmedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHandshakeConfirmed", reflect.TypeOf((*MockCryptoSetup)(nil).SetHandshakeConfirmed))
	return &MockCryptoSetupSetHandshakeConfirmedCall{Call: call}
}

// MockCryptoSetupSetHandshakeConfirmedCall wrap *gomock.Call
type MockCryptoSetupSetHandshakeConfirmedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCryptoSetupSetHandshakeConfirmedCall) Return() *MockCryptoSetupSetHandshakeConfirmedCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCryptoSetupSetHandshakeConfirmedCall) Do(f func()) *MockCryptoSetupSetHandshakeConfirmedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCryptoSetupSetHandshakeConfirmedCall) DoAndReturn(f func()) *MockCryptoSetupSetHandshakeConfirmedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetLargest1RTTAcked mocks base method.
func (m *MockCryptoSetup) SetLargest1RTTAcked(arg0 protocol.PacketNumber) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLargest1RTTAcked", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetLargest1RTTAcked indicates an expected call of SetLargest1RTTAcked.
func (mr *MockCryptoSetupMockRecorder) SetLargest1RTTAcked(arg0 any) *MockCryptoSetupSetLargest1RTTAckedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLargest1RTTAcked", reflect.TypeOf((*MockCryptoSetup)(nil).SetLargest1RTTAcked), arg0)
	return &MockCryptoSetupSetLargest1RTTAckedCall{Call: call}
}

// MockCryptoSetupSetLargest1RTTAckedCall wrap *gomock.Call
type MockCryptoSetupSetLargest1RTTAckedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCryptoSetupSetLargest1RTTAckedCall) Return(arg0 error) *MockCryptoSetupSetLargest1RTTAckedCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCryptoSetupSetLargest1RTTAckedCall) Do(f func(protocol.PacketNumber) error) *MockCryptoSetupSetLargest1RTTAckedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCryptoSetupSetLargest1RTTAckedCall) DoAndReturn(f func(protocol.PacketNumber) error) *MockCryptoSetupSetLargest1RTTAckedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// StartHandshake mocks base method.
func (m *MockCryptoSetup) StartHandshake(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartHandshake", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartHandshake indicates an expected call of StartHandshake.
func (mr *MockCryptoSetupMockRecorder) StartHandshake(arg0 any) *MockCryptoSetupStartHandshakeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartHandshake", reflect.TypeOf((*MockCryptoSetup)(nil).StartHandshake), arg0)
	return &MockCryptoSetupStartHandshakeCall{Call: call}
}

// MockCryptoSetupStartHandshakeCall wrap *gomock.Call
type MockCryptoSetupStartHandshakeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCryptoSetupStartHandshakeCall) Return(arg0 error) *MockCryptoSetupStartHandshakeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCryptoSetupStartHandshakeCall) Do(f func(context.Context) error) *MockCryptoSetupStartHandshakeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCryptoSetupStartHandshakeCall) DoAndReturn(f func(context.Context) error) *MockCryptoSetupStartHandshakeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
