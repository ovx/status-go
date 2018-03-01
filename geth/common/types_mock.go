// Code generated by MockGen. DO NOT EDIT.
// Source: types.go

// Package common is a generated GoMock package.
package common

import (
	context "context"
	accounts "github.com/ethereum/go-ethereum/accounts"
	keystore "github.com/ethereum/go-ethereum/accounts/keystore"
	common "github.com/ethereum/go-ethereum/common"
	les "github.com/ethereum/go-ethereum/les"
	node "github.com/ethereum/go-ethereum/node"
	whisperv5 "github.com/ethereum/go-ethereum/whisper/whisperv5"
	gomock "github.com/golang/mock/gomock"
	params "github.com/status-im/status-go/geth/params"
	rpc "github.com/status-im/status-go/geth/rpc"
	reflect "reflect"
)

// MockNodeManager is a mock of NodeManager interface
type MockNodeManager struct {
	ctrl     *gomock.Controller
	recorder *MockNodeManagerMockRecorder
}

// MockNodeManagerMockRecorder is the mock recorder for MockNodeManager
type MockNodeManagerMockRecorder struct {
	mock *MockNodeManager
}

// NewMockNodeManager creates a new mock instance
func NewMockNodeManager(ctrl *gomock.Controller) *MockNodeManager {
	mock := &MockNodeManager{ctrl: ctrl}
	mock.recorder = &MockNodeManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeManager) EXPECT() *MockNodeManagerMockRecorder {
	return m.recorder
}

// StartNode mocks base method
func (m *MockNodeManager) StartNode(config *params.NodeConfig) error {
	ret := m.ctrl.Call(m, "StartNode", config)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartNode indicates an expected call of StartNode
func (mr *MockNodeManagerMockRecorder) StartNode(config interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartNode", reflect.TypeOf((*MockNodeManager)(nil).StartNode), config)
}

// EnsureSync mocks base method
func (m *MockNodeManager) EnsureSync(ctx context.Context) error {
	ret := m.ctrl.Call(m, "EnsureSync", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureSync indicates an expected call of EnsureSync
func (mr *MockNodeManagerMockRecorder) EnsureSync(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureSync", reflect.TypeOf((*MockNodeManager)(nil).EnsureSync), ctx)
}

// StopNode mocks base method
func (m *MockNodeManager) StopNode() error {
	ret := m.ctrl.Call(m, "StopNode")
	ret0, _ := ret[0].(error)
	return ret0
}

// StopNode indicates an expected call of StopNode
func (mr *MockNodeManagerMockRecorder) StopNode() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopNode", reflect.TypeOf((*MockNodeManager)(nil).StopNode))
}

// IsNodeRunning mocks base method
func (m *MockNodeManager) IsNodeRunning() bool {
	ret := m.ctrl.Call(m, "IsNodeRunning")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsNodeRunning indicates an expected call of IsNodeRunning
func (mr *MockNodeManagerMockRecorder) IsNodeRunning() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNodeRunning", reflect.TypeOf((*MockNodeManager)(nil).IsNodeRunning))
}

// NodeConfig mocks base method
func (m *MockNodeManager) NodeConfig() (*params.NodeConfig, error) {
	ret := m.ctrl.Call(m, "NodeConfig")
	ret0, _ := ret[0].(*params.NodeConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NodeConfig indicates an expected call of NodeConfig
func (mr *MockNodeManagerMockRecorder) NodeConfig() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeConfig", reflect.TypeOf((*MockNodeManager)(nil).NodeConfig))
}

// Node mocks base method
func (m *MockNodeManager) Node() (*node.Node, error) {
	ret := m.ctrl.Call(m, "Node")
	ret0, _ := ret[0].(*node.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Node indicates an expected call of Node
func (mr *MockNodeManagerMockRecorder) Node() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Node", reflect.TypeOf((*MockNodeManager)(nil).Node))
}

// PopulateStaticPeers mocks base method
func (m *MockNodeManager) PopulateStaticPeers() error {
	ret := m.ctrl.Call(m, "PopulateStaticPeers")
	ret0, _ := ret[0].(error)
	return ret0
}

// PopulateStaticPeers indicates an expected call of PopulateStaticPeers
func (mr *MockNodeManagerMockRecorder) PopulateStaticPeers() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PopulateStaticPeers", reflect.TypeOf((*MockNodeManager)(nil).PopulateStaticPeers))
}

// AddPeer mocks base method
func (m *MockNodeManager) AddPeer(url string) error {
	ret := m.ctrl.Call(m, "AddPeer", url)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPeer indicates an expected call of AddPeer
func (mr *MockNodeManagerMockRecorder) AddPeer(url interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPeer", reflect.TypeOf((*MockNodeManager)(nil).AddPeer), url)
}

// PeerCount mocks base method
func (m *MockNodeManager) PeerCount() int {
	ret := m.ctrl.Call(m, "PeerCount")
	ret0, _ := ret[0].(int)
	return ret0
}

// PeerCount indicates an expected call of PeerCount
func (mr *MockNodeManagerMockRecorder) PeerCount() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeerCount", reflect.TypeOf((*MockNodeManager)(nil).PeerCount))
}

// LightEthereumService mocks base method
func (m *MockNodeManager) LightEthereumService() (*les.LightEthereum, error) {
	ret := m.ctrl.Call(m, "LightEthereumService")
	ret0, _ := ret[0].(*les.LightEthereum)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LightEthereumService indicates an expected call of LightEthereumService
func (mr *MockNodeManagerMockRecorder) LightEthereumService() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LightEthereumService", reflect.TypeOf((*MockNodeManager)(nil).LightEthereumService))
}

// WhisperService mocks base method
func (m *MockNodeManager) WhisperService() (*whisperv5.Whisper, error) {
	ret := m.ctrl.Call(m, "WhisperService")
	ret0, _ := ret[0].(*whisperv5.Whisper)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WhisperService indicates an expected call of WhisperService
func (mr *MockNodeManagerMockRecorder) WhisperService() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WhisperService", reflect.TypeOf((*MockNodeManager)(nil).WhisperService))
}

// AccountManager mocks base method
func (m *MockNodeManager) AccountManager() (*accounts.Manager, error) {
	ret := m.ctrl.Call(m, "AccountManager")
	ret0, _ := ret[0].(*accounts.Manager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AccountManager indicates an expected call of AccountManager
func (mr *MockNodeManagerMockRecorder) AccountManager() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountManager", reflect.TypeOf((*MockNodeManager)(nil).AccountManager))
}

// AccountKeyStore mocks base method
func (m *MockNodeManager) AccountKeyStore() (*keystore.KeyStore, error) {
	ret := m.ctrl.Call(m, "AccountKeyStore")
	ret0, _ := ret[0].(*keystore.KeyStore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AccountKeyStore indicates an expected call of AccountKeyStore
func (mr *MockNodeManagerMockRecorder) AccountKeyStore() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountKeyStore", reflect.TypeOf((*MockNodeManager)(nil).AccountKeyStore))
}

// RPCClient mocks base method
func (m *MockNodeManager) RPCClient() *rpc.Client {
	ret := m.ctrl.Call(m, "RPCClient")
	ret0, _ := ret[0].(*rpc.Client)
	return ret0
}

// RPCClient indicates an expected call of RPCClient
func (mr *MockNodeManagerMockRecorder) RPCClient() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RPCClient", reflect.TypeOf((*MockNodeManager)(nil).RPCClient))
}

// MockAccountManager is a mock of AccountManager interface
type MockAccountManager struct {
	ctrl     *gomock.Controller
	recorder *MockAccountManagerMockRecorder
}

// MockAccountManagerMockRecorder is the mock recorder for MockAccountManager
type MockAccountManagerMockRecorder struct {
	mock *MockAccountManager
}

// NewMockAccountManager creates a new mock instance
func NewMockAccountManager(ctrl *gomock.Controller) *MockAccountManager {
	mock := &MockAccountManager{ctrl: ctrl}
	mock.recorder = &MockAccountManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccountManager) EXPECT() *MockAccountManagerMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method
func (m *MockAccountManager) CreateAccount(password string) (string, string, string, error) {
	ret := m.ctrl.Call(m, "CreateAccount", password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// CreateAccount indicates an expected call of CreateAccount
func (mr *MockAccountManagerMockRecorder) CreateAccount(password interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockAccountManager)(nil).CreateAccount), password)
}

// CreateChildAccount mocks base method
func (m *MockAccountManager) CreateChildAccount(parentAddress, password string) (string, string, error) {
	ret := m.ctrl.Call(m, "CreateChildAccount", parentAddress, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateChildAccount indicates an expected call of CreateChildAccount
func (mr *MockAccountManagerMockRecorder) CreateChildAccount(parentAddress, password interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChildAccount", reflect.TypeOf((*MockAccountManager)(nil).CreateChildAccount), parentAddress, password)
}

// RecoverAccount mocks base method
func (m *MockAccountManager) RecoverAccount(password, mnemonic string) (string, string, error) {
	ret := m.ctrl.Call(m, "RecoverAccount", password, mnemonic)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RecoverAccount indicates an expected call of RecoverAccount
func (mr *MockAccountManagerMockRecorder) RecoverAccount(password, mnemonic interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecoverAccount", reflect.TypeOf((*MockAccountManager)(nil).RecoverAccount), password, mnemonic)
}

// VerifyAccountPassword mocks base method
func (m *MockAccountManager) VerifyAccountPassword(keyStoreDir, address, password string) (*keystore.Key, error) {
	ret := m.ctrl.Call(m, "VerifyAccountPassword", keyStoreDir, address, password)
	ret0, _ := ret[0].(*keystore.Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyAccountPassword indicates an expected call of VerifyAccountPassword
func (mr *MockAccountManagerMockRecorder) VerifyAccountPassword(keyStoreDir, address, password interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyAccountPassword", reflect.TypeOf((*MockAccountManager)(nil).VerifyAccountPassword), keyStoreDir, address, password)
}

// SelectAccount mocks base method
func (m *MockAccountManager) SelectAccount(address, password string) error {
	ret := m.ctrl.Call(m, "SelectAccount", address, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// SelectAccount indicates an expected call of SelectAccount
func (mr *MockAccountManagerMockRecorder) SelectAccount(address, password interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectAccount", reflect.TypeOf((*MockAccountManager)(nil).SelectAccount), address, password)
}

// ReSelectAccount mocks base method
func (m *MockAccountManager) ReSelectAccount() error {
	ret := m.ctrl.Call(m, "ReSelectAccount")
	ret0, _ := ret[0].(error)
	return ret0
}

// ReSelectAccount indicates an expected call of ReSelectAccount
func (mr *MockAccountManagerMockRecorder) ReSelectAccount() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReSelectAccount", reflect.TypeOf((*MockAccountManager)(nil).ReSelectAccount))
}

// SelectedAccount mocks base method
func (m *MockAccountManager) SelectedAccount() (*SelectedExtKey, error) {
	ret := m.ctrl.Call(m, "SelectedAccount")
	ret0, _ := ret[0].(*SelectedExtKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectedAccount indicates an expected call of SelectedAccount
func (mr *MockAccountManagerMockRecorder) SelectedAccount() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectedAccount", reflect.TypeOf((*MockAccountManager)(nil).SelectedAccount))
}

// Logout mocks base method
func (m *MockAccountManager) Logout() error {
	ret := m.ctrl.Call(m, "Logout")
	ret0, _ := ret[0].(error)
	return ret0
}

// Logout indicates an expected call of Logout
func (mr *MockAccountManagerMockRecorder) Logout() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockAccountManager)(nil).Logout))
}

// Accounts mocks base method
func (m *MockAccountManager) Accounts() ([]common.Address, error) {
	ret := m.ctrl.Call(m, "Accounts")
	ret0, _ := ret[0].([]common.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Accounts indicates an expected call of Accounts
func (mr *MockAccountManagerMockRecorder) Accounts() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accounts", reflect.TypeOf((*MockAccountManager)(nil).Accounts))
}

// AccountsRPCHandler mocks base method
func (m *MockAccountManager) AccountsRPCHandler() rpc.Handler {
	ret := m.ctrl.Call(m, "AccountsRPCHandler")
	ret0, _ := ret[0].(rpc.Handler)
	return ret0
}

// AccountsRPCHandler indicates an expected call of AccountsRPCHandler
func (mr *MockAccountManagerMockRecorder) AccountsRPCHandler() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountsRPCHandler", reflect.TypeOf((*MockAccountManager)(nil).AccountsRPCHandler))
}

// AddressToDecryptedAccount mocks base method
func (m *MockAccountManager) AddressToDecryptedAccount(address, password string) (accounts.Account, *keystore.Key, error) {
	ret := m.ctrl.Call(m, "AddressToDecryptedAccount", address, password)
	ret0, _ := ret[0].(accounts.Account)
	ret1, _ := ret[1].(*keystore.Key)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddressToDecryptedAccount indicates an expected call of AddressToDecryptedAccount
func (mr *MockAccountManagerMockRecorder) AddressToDecryptedAccount(address, password interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddressToDecryptedAccount", reflect.TypeOf((*MockAccountManager)(nil).AddressToDecryptedAccount), address, password)
}
