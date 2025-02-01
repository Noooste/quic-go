// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Noooste/quic-go/internal/mocks/logging (interfaces: Tracer)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package internal -destination internal/tracer.go github.com/Noooste/quic-go/internal/mocks/logging Tracer
//

// Package internal is a generated GoMock package.
package internal

import (
	net "net"
	reflect "reflect"

	protocol "github.com/Noooste/quic-go/internal/protocol"
	wire "github.com/Noooste/quic-go/internal/wire"
	logging "github.com/Noooste/quic-go/logging"
	gomock "go.uber.org/mock/gomock"
)

// MockTracer is a mock of Tracer interface.
type MockTracer struct {
	ctrl     *gomock.Controller
	recorder *MockTracerMockRecorder
	isgomock struct{}
}

// MockTracerMockRecorder is the mock recorder for MockTracer.
type MockTracerMockRecorder struct {
	mock *MockTracer
}

// NewMockTracer creates a new mock instance.
func NewMockTracer(ctrl *gomock.Controller) *MockTracer {
	mock := &MockTracer{ctrl: ctrl}
	mock.recorder = &MockTracerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTracer) EXPECT() *MockTracerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockTracer) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockTracerMockRecorder) Close() *MockTracerCloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockTracer)(nil).Close))
	return &MockTracerCloseCall{Call: call}
}

// MockTracerCloseCall wrap *gomock.Call
type MockTracerCloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTracerCloseCall) Return() *MockTracerCloseCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTracerCloseCall) Do(f func()) *MockTracerCloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTracerCloseCall) DoAndReturn(f func()) *MockTracerCloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Debug mocks base method.
func (m *MockTracer) Debug(name, msg string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Debug", name, msg)
}

// Debug indicates an expected call of Debug.
func (mr *MockTracerMockRecorder) Debug(name, msg any) *MockTracerDebugCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockTracer)(nil).Debug), name, msg)
	return &MockTracerDebugCall{Call: call}
}

// MockTracerDebugCall wrap *gomock.Call
type MockTracerDebugCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTracerDebugCall) Return() *MockTracerDebugCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTracerDebugCall) Do(f func(string, string)) *MockTracerDebugCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTracerDebugCall) DoAndReturn(f func(string, string)) *MockTracerDebugCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DroppedPacket mocks base method.
func (m *MockTracer) DroppedPacket(arg0 net.Addr, arg1 logging.PacketType, arg2 protocol.ByteCount, arg3 logging.PacketDropReason) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DroppedPacket", arg0, arg1, arg2, arg3)
}

// DroppedPacket indicates an expected call of DroppedPacket.
func (mr *MockTracerMockRecorder) DroppedPacket(arg0, arg1, arg2, arg3 any) *MockTracerDroppedPacketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DroppedPacket", reflect.TypeOf((*MockTracer)(nil).DroppedPacket), arg0, arg1, arg2, arg3)
	return &MockTracerDroppedPacketCall{Call: call}
}

// MockTracerDroppedPacketCall wrap *gomock.Call
type MockTracerDroppedPacketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTracerDroppedPacketCall) Return() *MockTracerDroppedPacketCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTracerDroppedPacketCall) Do(f func(net.Addr, logging.PacketType, protocol.ByteCount, logging.PacketDropReason)) *MockTracerDroppedPacketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTracerDroppedPacketCall) DoAndReturn(f func(net.Addr, logging.PacketType, protocol.ByteCount, logging.PacketDropReason)) *MockTracerDroppedPacketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SentPacket mocks base method.
func (m *MockTracer) SentPacket(arg0 net.Addr, arg1 *wire.Header, arg2 protocol.ByteCount, arg3 []logging.Frame) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SentPacket", arg0, arg1, arg2, arg3)
}

// SentPacket indicates an expected call of SentPacket.
func (mr *MockTracerMockRecorder) SentPacket(arg0, arg1, arg2, arg3 any) *MockTracerSentPacketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SentPacket", reflect.TypeOf((*MockTracer)(nil).SentPacket), arg0, arg1, arg2, arg3)
	return &MockTracerSentPacketCall{Call: call}
}

// MockTracerSentPacketCall wrap *gomock.Call
type MockTracerSentPacketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTracerSentPacketCall) Return() *MockTracerSentPacketCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTracerSentPacketCall) Do(f func(net.Addr, *wire.Header, protocol.ByteCount, []logging.Frame)) *MockTracerSentPacketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTracerSentPacketCall) DoAndReturn(f func(net.Addr, *wire.Header, protocol.ByteCount, []logging.Frame)) *MockTracerSentPacketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SentVersionNegotiationPacket mocks base method.
func (m *MockTracer) SentVersionNegotiationPacket(arg0 net.Addr, dest, src protocol.ArbitraryLenConnectionID, arg3 []protocol.Version) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SentVersionNegotiationPacket", arg0, dest, src, arg3)
}

// SentVersionNegotiationPacket indicates an expected call of SentVersionNegotiationPacket.
func (mr *MockTracerMockRecorder) SentVersionNegotiationPacket(arg0, dest, src, arg3 any) *MockTracerSentVersionNegotiationPacketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SentVersionNegotiationPacket", reflect.TypeOf((*MockTracer)(nil).SentVersionNegotiationPacket), arg0, dest, src, arg3)
	return &MockTracerSentVersionNegotiationPacketCall{Call: call}
}

// MockTracerSentVersionNegotiationPacketCall wrap *gomock.Call
type MockTracerSentVersionNegotiationPacketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTracerSentVersionNegotiationPacketCall) Return() *MockTracerSentVersionNegotiationPacketCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTracerSentVersionNegotiationPacketCall) Do(f func(net.Addr, protocol.ArbitraryLenConnectionID, protocol.ArbitraryLenConnectionID, []protocol.Version)) *MockTracerSentVersionNegotiationPacketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTracerSentVersionNegotiationPacketCall) DoAndReturn(f func(net.Addr, protocol.ArbitraryLenConnectionID, protocol.ArbitraryLenConnectionID, []protocol.Version)) *MockTracerSentVersionNegotiationPacketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
