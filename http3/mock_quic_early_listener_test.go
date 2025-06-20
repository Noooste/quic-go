// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Noooste/uquic-go/http3 (interfaces: QUICEarlyListener)
//
// Generated by this command:
//
//	mockgen -typed -package http3 -destination mock_quic_early_listener_test.go github.com/Noooste/uquic-go/http3 QUICEarlyListener
//

// Package http3 is a generated GoMock package.
package http3

import (
	context "context"
	net "net"
	reflect "reflect"

	quic "github.com/Noooste/uquic-go"
	gomock "go.uber.org/mock/gomock"
)

// MockQUICEarlyListener is a mock of QUICEarlyListener interface.
type MockQUICEarlyListener struct {
	ctrl     *gomock.Controller
	recorder *MockQUICEarlyListenerMockRecorder
	isgomock struct{}
}

// MockQUICEarlyListenerMockRecorder is the mock recorder for MockQUICEarlyListener.
type MockQUICEarlyListenerMockRecorder struct {
	mock *MockQUICEarlyListener
}

// NewMockQUICEarlyListener creates a new mock instance.
func NewMockQUICEarlyListener(ctrl *gomock.Controller) *MockQUICEarlyListener {
	mock := &MockQUICEarlyListener{ctrl: ctrl}
	mock.recorder = &MockQUICEarlyListenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQUICEarlyListener) EXPECT() *MockQUICEarlyListenerMockRecorder {
	return m.recorder
}

// Accept mocks base method.
func (m *MockQUICEarlyListener) Accept(arg0 context.Context) (quic.EarlyConnection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept", arg0)
	ret0, _ := ret[0].(quic.EarlyConnection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Accept indicates an expected call of Accept.
func (mr *MockQUICEarlyListenerMockRecorder) Accept(arg0 any) *MockQUICEarlyListenerAcceptCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockQUICEarlyListener)(nil).Accept), arg0)
	return &MockQUICEarlyListenerAcceptCall{Call: call}
}

// MockQUICEarlyListenerAcceptCall wrap *gomock.Call
type MockQUICEarlyListenerAcceptCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockQUICEarlyListenerAcceptCall) Return(arg0 quic.EarlyConnection, arg1 error) *MockQUICEarlyListenerAcceptCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockQUICEarlyListenerAcceptCall) Do(f func(context.Context) (quic.EarlyConnection, error)) *MockQUICEarlyListenerAcceptCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockQUICEarlyListenerAcceptCall) DoAndReturn(f func(context.Context) (quic.EarlyConnection, error)) *MockQUICEarlyListenerAcceptCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Addr mocks base method.
func (m *MockQUICEarlyListener) Addr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Addr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// Addr indicates an expected call of Addr.
func (mr *MockQUICEarlyListenerMockRecorder) Addr() *MockQUICEarlyListenerAddrCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Addr", reflect.TypeOf((*MockQUICEarlyListener)(nil).Addr))
	return &MockQUICEarlyListenerAddrCall{Call: call}
}

// MockQUICEarlyListenerAddrCall wrap *gomock.Call
type MockQUICEarlyListenerAddrCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockQUICEarlyListenerAddrCall) Return(arg0 net.Addr) *MockQUICEarlyListenerAddrCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockQUICEarlyListenerAddrCall) Do(f func() net.Addr) *MockQUICEarlyListenerAddrCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockQUICEarlyListenerAddrCall) DoAndReturn(f func() net.Addr) *MockQUICEarlyListenerAddrCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Close mocks base method.
func (m *MockQUICEarlyListener) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockQUICEarlyListenerMockRecorder) Close() *MockQUICEarlyListenerCloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockQUICEarlyListener)(nil).Close))
	return &MockQUICEarlyListenerCloseCall{Call: call}
}

// MockQUICEarlyListenerCloseCall wrap *gomock.Call
type MockQUICEarlyListenerCloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockQUICEarlyListenerCloseCall) Return(arg0 error) *MockQUICEarlyListenerCloseCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockQUICEarlyListenerCloseCall) Do(f func() error) *MockQUICEarlyListenerCloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockQUICEarlyListenerCloseCall) DoAndReturn(f func() error) *MockQUICEarlyListenerCloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
