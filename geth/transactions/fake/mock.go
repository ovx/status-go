// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/status-im/status-go/geth/transactions/fake (interfaces: PublicTransactionPoolAPI)

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	common "github.com/ethereum/go-ethereum/common"
	hexutil "github.com/ethereum/go-ethereum/common/hexutil"
	rpc "github.com/ethereum/go-ethereum/rpc"
	gomock "github.com/golang/mock/gomock"
	big "math/big"
	reflect "reflect"
)

// MockPublicTransactionPoolAPI is a mock of PublicTransactionPoolAPI interface
type MockPublicTransactionPoolAPI struct {
	ctrl     *gomock.Controller
	recorder *MockPublicTransactionPoolAPIMockRecorder
}

// MockPublicTransactionPoolAPIMockRecorder is the mock recorder for MockPublicTransactionPoolAPI
type MockPublicTransactionPoolAPIMockRecorder struct {
	mock *MockPublicTransactionPoolAPI
}

// NewMockPublicTransactionPoolAPI creates a new mock instance
func NewMockPublicTransactionPoolAPI(ctrl *gomock.Controller) *MockPublicTransactionPoolAPI {
	mock := &MockPublicTransactionPoolAPI{ctrl: ctrl}
	mock.recorder = &MockPublicTransactionPoolAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPublicTransactionPoolAPI) EXPECT() *MockPublicTransactionPoolAPIMockRecorder {
	return m.recorder
}

// EstimateGas mocks base method
func (m *MockPublicTransactionPoolAPI) EstimateGas(arg0 context.Context, arg1 CallArgs) (hexutil.Uint64, error) {
	ret := m.ctrl.Call(m, "EstimateGas", arg0, arg1)
	ret0, _ := ret[0].(hexutil.Uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EstimateGas indicates an expected call of EstimateGas
func (mr *MockPublicTransactionPoolAPIMockRecorder) EstimateGas(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EstimateGas", reflect.TypeOf((*MockPublicTransactionPoolAPI)(nil).EstimateGas), arg0, arg1)
}

// GasPrice mocks base method
func (m *MockPublicTransactionPoolAPI) GasPrice(arg0 context.Context) (*big.Int, error) {
	ret := m.ctrl.Call(m, "GasPrice", arg0)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GasPrice indicates an expected call of GasPrice
func (mr *MockPublicTransactionPoolAPIMockRecorder) GasPrice(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GasPrice", reflect.TypeOf((*MockPublicTransactionPoolAPI)(nil).GasPrice), arg0)
}

// GetTransactionCount mocks base method
func (m *MockPublicTransactionPoolAPI) GetTransactionCount(arg0 context.Context, arg1 common.Address, arg2 rpc.BlockNumber) (*hexutil.Uint64, error) {
	ret := m.ctrl.Call(m, "GetTransactionCount", arg0, arg1, arg2)
	ret0, _ := ret[0].(*hexutil.Uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionCount indicates an expected call of GetTransactionCount
func (mr *MockPublicTransactionPoolAPIMockRecorder) GetTransactionCount(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionCount", reflect.TypeOf((*MockPublicTransactionPoolAPI)(nil).GetTransactionCount), arg0, arg1, arg2)
}

// SendRawTransaction mocks base method
func (m *MockPublicTransactionPoolAPI) SendRawTransaction(arg0 context.Context, arg1 hexutil.Bytes) (common.Hash, error) {
	ret := m.ctrl.Call(m, "SendRawTransaction", arg0, arg1)
	ret0, _ := ret[0].(common.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendRawTransaction indicates an expected call of SendRawTransaction
func (mr *MockPublicTransactionPoolAPIMockRecorder) SendRawTransaction(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendRawTransaction", reflect.TypeOf((*MockPublicTransactionPoolAPI)(nil).SendRawTransaction), arg0, arg1)
}
