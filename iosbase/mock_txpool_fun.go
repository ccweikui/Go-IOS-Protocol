package iosbase

import "reflect"

// NewMockTxPool creates a new mock instance
func NewMockTxPool(ctrl *gomock.Controller) *MockTxPool {
	mock := &MockTxPool{ctrl: ctrl}
	mock.recorder = &MockTxPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTxPool) EXPECT() *MockTxPoolMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockTxPool) Add(arg0 iosbase.Tx) error {
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockTxPoolMockRecorder) Add(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockTxPool)(nil).Add), arg0)
}

// Decode mocks base method
func (m *MockTxPool) Decode(arg0 []byte) error {
	ret := m.ctrl.Call(m, "Decode", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Decode indicates an expected call of Decode
func (mr *MockTxPoolMockRecorder) Decode(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockTxPool)(nil).Decode), arg0)
}

// Del mocks base method
func (m *MockTxPool) Del(arg0 iosbase.Tx) error {
	ret := m.ctrl.Call(m, "Del", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Del indicates an expected call of Del
func (mr *MockTxPoolMockRecorder) Del(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockTxPool)(nil).Del), arg0)
}

// Encode mocks base method
func (m *MockTxPool) Encode() []byte {
	ret := m.ctrl.Call(m, "Encode")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Encode indicates an expected call of Encode
func (mr *MockTxPoolMockRecorder) Encode() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encode", reflect.TypeOf((*MockTxPool)(nil).Encode))
}

// Find mocks base method
func (m *MockTxPool) Find(arg0 []byte) (iosbase.Tx, error) {
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].(iosbase.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockTxPoolMockRecorder) Find(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockTxPool)(nil).Find), arg0)
}

// GetSlice mocks base method
func (m *MockTxPool) GetSlice() ([]iosbase.Tx, error) {
	ret := m.ctrl.Call(m, "GetSlice")
	ret0, _ := ret[0].([]iosbase.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSlice indicates an expected call of GetSlice
func (mr *MockTxPoolMockRecorder) GetSlice() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSlice", reflect.TypeOf((*MockTxPool)(nil).GetSlice))
}

// Has mocks base method
func (m *MockTxPool) Has(arg0 []byte) (bool, error) {
	ret := m.ctrl.Call(m, "Has", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Has indicates an expected call of Has
func (mr *MockTxPoolMockRecorder) Has(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockTxPool)(nil).Has), arg0)
}

// Hash mocks base method
func (m *MockTxPool) Hash() []byte {
	ret := m.ctrl.Call(m, "Hash")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Hash indicates an expected call of Hash
func (mr *MockTxPoolMockRecorder) Hash() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hash", reflect.TypeOf((*MockTxPool)(nil).Hash))
}

// Size mocks base method
func (m *MockTxPool) Size() int {
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size
func (mr *MockTxPoolMockRecorder) Size() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockTxPool)(nil).Size))
}

