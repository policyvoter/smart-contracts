// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pv1_smartcontract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// Pv1SmartcontractMetaData contains all meta data concerning the Pv1Smartcontract contract.
var Pv1SmartcontractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"blacklisted\",\"type\":\"bool\"}],\"name\":\"Blacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"policyID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"groupID\",\"type\":\"uint256\"}],\"name\":\"NewPolicy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"policyID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"blacklisted\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"PolicyBlacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"policyID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"groupID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalVotes\",\"type\":\"uint256\"}],\"name\":\"UnVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"policyID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"groupID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalVotes\",\"type\":\"uint256\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADDRESS_REGISTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POLICY_BLACKLISTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POLICY_CREATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"policyID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"gID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"bodyHash\",\"type\":\"string\"}],\"name\":\"createNewPolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"groupForAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"isPolicyBlacklisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastVoted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minIntervalBetweenVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"policies\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"bodyHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"groupID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"policyBlacklistedReason\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"groupID\",\"type\":\"uint256\"}],\"name\":\"registerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"registered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMinInterval\",\"type\":\"uint256\"}],\"name\":\"setMinIntervalBetweenVotes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"policyID\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"blacklisted\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"setPolicyBlacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"policyID\",\"type\":\"string\"}],\"name\":\"unVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"policyID\",\"type\":\"string\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"votedOnProposal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"votes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Pv1SmartcontractABI is the input ABI used to generate the binding from.
// Deprecated: Use Pv1SmartcontractMetaData.ABI instead.
var Pv1SmartcontractABI = Pv1SmartcontractMetaData.ABI

// Pv1Smartcontract is an auto generated Go binding around an Ethereum contract.
type Pv1Smartcontract struct {
	Pv1SmartcontractCaller     // Read-only binding to the contract
	Pv1SmartcontractTransactor // Write-only binding to the contract
	Pv1SmartcontractFilterer   // Log filterer for contract events
}

// Pv1SmartcontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type Pv1SmartcontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Pv1SmartcontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Pv1SmartcontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Pv1SmartcontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Pv1SmartcontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Pv1SmartcontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Pv1SmartcontractSession struct {
	Contract     *Pv1Smartcontract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Pv1SmartcontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Pv1SmartcontractCallerSession struct {
	Contract *Pv1SmartcontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// Pv1SmartcontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Pv1SmartcontractTransactorSession struct {
	Contract     *Pv1SmartcontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// Pv1SmartcontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type Pv1SmartcontractRaw struct {
	Contract *Pv1Smartcontract // Generic contract binding to access the raw methods on
}

// Pv1SmartcontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Pv1SmartcontractCallerRaw struct {
	Contract *Pv1SmartcontractCaller // Generic read-only contract binding to access the raw methods on
}

// Pv1SmartcontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Pv1SmartcontractTransactorRaw struct {
	Contract *Pv1SmartcontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPv1Smartcontract creates a new instance of Pv1Smartcontract, bound to a specific deployed contract.
func NewPv1Smartcontract(address common.Address, backend bind.ContractBackend) (*Pv1Smartcontract, error) {
	contract, err := bindPv1Smartcontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pv1Smartcontract{Pv1SmartcontractCaller: Pv1SmartcontractCaller{contract: contract}, Pv1SmartcontractTransactor: Pv1SmartcontractTransactor{contract: contract}, Pv1SmartcontractFilterer: Pv1SmartcontractFilterer{contract: contract}}, nil
}

// NewPv1SmartcontractCaller creates a new read-only instance of Pv1Smartcontract, bound to a specific deployed contract.
func NewPv1SmartcontractCaller(address common.Address, caller bind.ContractCaller) (*Pv1SmartcontractCaller, error) {
	contract, err := bindPv1Smartcontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Pv1SmartcontractCaller{contract: contract}, nil
}

// NewPv1SmartcontractTransactor creates a new write-only instance of Pv1Smartcontract, bound to a specific deployed contract.
func NewPv1SmartcontractTransactor(address common.Address, transactor bind.ContractTransactor) (*Pv1SmartcontractTransactor, error) {
	contract, err := bindPv1Smartcontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Pv1SmartcontractTransactor{contract: contract}, nil
}

// NewPv1SmartcontractFilterer creates a new log filterer instance of Pv1Smartcontract, bound to a specific deployed contract.
func NewPv1SmartcontractFilterer(address common.Address, filterer bind.ContractFilterer) (*Pv1SmartcontractFilterer, error) {
	contract, err := bindPv1Smartcontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Pv1SmartcontractFilterer{contract: contract}, nil
}

// bindPv1Smartcontract binds a generic wrapper to an already deployed contract.
func bindPv1Smartcontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Pv1SmartcontractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pv1Smartcontract *Pv1SmartcontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pv1Smartcontract.Contract.Pv1SmartcontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pv1Smartcontract *Pv1SmartcontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pv1Smartcontract.Contract.Pv1SmartcontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pv1Smartcontract *Pv1SmartcontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pv1Smartcontract.Contract.Pv1SmartcontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pv1Smartcontract *Pv1SmartcontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pv1Smartcontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pv1Smartcontract *Pv1SmartcontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pv1Smartcontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pv1Smartcontract *Pv1SmartcontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pv1Smartcontract.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSREGISTERROLE is a free data retrieval call binding the contract method 0xb464d8b7.
//
// Solidity: function ADDRESS_REGISTER_ROLE() view returns(bytes32)
func (_Pv1Smartcontract *Pv1SmartcontractCaller) ADDRESSREGISTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pv1Smartcontract.contract.Call(opts, &out, "ADDRESS_REGISTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADDRESSREGISTERROLE is a free data retrieval call binding the contract method 0xb464d8b7.
//
// Solidity: function ADDRESS_REGISTER_ROLE() view returns(bytes32)
func (_Pv1Smartcontract *Pv1SmartcontractSession) ADDRESSREGISTERROLE() ([32]byte, error) {
	return _Pv1Smartcontract.Contract.ADDRESSREGISTERROLE(&_Pv1Smartcontract.CallOpts)
}

// ADDRESSREGISTERROLE is a free data retrieval call binding the contract method 0xb464d8b7.
//
// Solidity: function ADDRESS_REGISTER_ROLE() view returns(bytes32)
func (_Pv1Smartcontract *Pv1SmartcontractCallerSession) ADDRESSREGISTERROLE() ([32]byte, error) {
	return _Pv1Smartcontract.Contract.ADDRESSREGISTERROLE(&_Pv1Smartcontract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Pv1Smartcontract *Pv1SmartcontractCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pv1Smartcontract.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Pv1Smartcontract *Pv1SmartcontractSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Pv1Smartcontract.Contract.DEFAULTADMINROLE(&_Pv1Smartcontract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Pv1Smartcontract *Pv1SmartcontractCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Pv1Smartcontract.Contract.DEFAULTADMINROLE(&_Pv1Smartcontract.CallOpts)
}

// POLICYBLACKLISTERROLE is a free data retrieval call binding the contract method 0xfc86449a.
//
// Solidity: function POLICY_BLACKLISTER_ROLE() view returns(bytes32)
func (_Pv1Smartcontract *Pv1SmartcontractCaller) POLICYBLACKLISTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pv1Smartcontract.contract.Call(opts, &out, "POLICY_BLACKLISTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// POLICYBLACKLISTERROLE is a free data retrieval call binding the contract method 0xfc86449a.
//
// Solidity: function POLICY_BLACKLISTER_ROLE() view returns(bytes32)
func (_Pv1Smartcontract *Pv1SmartcontractSession) POLICYBLACKLISTERROLE() ([32]byte, error) {
	return _Pv1Smartcontract.Contract.POLICYBLACKLISTERROLE(&_Pv1Smartcontract.CallOpts)
}

// POLICYBLACKLISTERROLE is a free data retrieval call binding the contract method 0xfc86449a.
//
// Solidity: function POLICY_BLACKLISTER_ROLE() view returns(bytes32)
func (_Pv1Smartcontract *Pv1SmartcontractCallerSession) POLICYBLACKLISTERROLE() ([32]byte, error) {
	return _Pv1Smartcontract.Contract.POLICYBLACKLISTERROLE(&_Pv1Smartcontract.CallOpts)
}

// POLICYCREATORROLE is a free data retrieval call binding the contract method 0x45f739ff.
//
// Solidity: function POLICY_CREATOR_ROLE() view returns(bytes32)
func (_Pv1Smartcontract *Pv1SmartcontractCaller) POLICYCREATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pv1Smartcontract.contract.Call(opts, &out, "POLICY_CREATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// POLICYCREATORROLE is a free data retrieval call binding the contract method 0x45f739ff.
//
// Solidity: function POLICY_CREATOR_ROLE() view returns(bytes32)
func (_Pv1Smartcontract *Pv1SmartcontractSession) POLICYCREATORROLE() ([32]byte, error) {
	return _Pv1Smartcontract.Contract.POLICYCREATORROLE(&_Pv1Smartcontract.CallOpts)
}

// POLICYCREATORROLE is a free data retrieval call binding the contract method 0x45f739ff.
//
// Solidity: function POLICY_CREATOR_ROLE() view returns(bytes32)
func (_Pv1Smartcontract *Pv1SmartcontractCallerSession) POLICYCREATORROLE() ([32]byte, error) {
	return _Pv1Smartcontract.Contract.POLICYCREATORROLE(&_Pv1Smartcontract.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Pv1Smartcontract *Pv1SmartcontractCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Pv1Smartcontract.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Pv1Smartcontract *Pv1SmartcontractSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Pv1Smartcontract.Contract.GetRoleAdmin(&_Pv1Smartcontract.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Pv1Smartcontract *Pv1SmartcontractCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Pv1Smartcontract.Contract.GetRoleAdmin(&_Pv1Smartcontract.CallOpts, role)
}

// GroupForAddress is a free data retrieval call binding the contract method 0x4beafca5.
//
// Solidity: function groupForAddress(address ) view returns(uint256)
func (_Pv1Smartcontract *Pv1SmartcontractCaller) GroupForAddress(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pv1Smartcontract.contract.Call(opts, &out, "groupForAddress", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GroupForAddress is a free data retrieval call binding the contract method 0x4beafca5.
//
// Solidity: function groupForAddress(address ) view returns(uint256)
func (_Pv1Smartcontract *Pv1SmartcontractSession) GroupForAddress(arg0 common.Address) (*big.Int, error) {
	return _Pv1Smartcontract.Contract.GroupForAddress(&_Pv1Smartcontract.CallOpts, arg0)
}

// GroupForAddress is a free data retrieval call binding the contract method 0x4beafca5.
//
// Solidity: function groupForAddress(address ) view returns(uint256)
func (_Pv1Smartcontract *Pv1SmartcontractCallerSession) GroupForAddress(arg0 common.Address) (*big.Int, error) {
	return _Pv1Smartcontract.Contract.GroupForAddress(&_Pv1Smartcontract.CallOpts, arg0)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Pv1Smartcontract *Pv1SmartcontractCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Pv1Smartcontract.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Pv1Smartcontract *Pv1SmartcontractSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Pv1Smartcontract.Contract.HasRole(&_Pv1Smartcontract.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Pv1Smartcontract *Pv1SmartcontractCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Pv1Smartcontract.Contract.HasRole(&_Pv1Smartcontract.CallOpts, role, account)
}

// IsPolicyBlacklisted is a free data retrieval call binding the contract method 0x9131a44d.
//
// Solidity: function isPolicyBlacklisted(string ) view returns(bool)
func (_Pv1Smartcontract *Pv1SmartcontractCaller) IsPolicyBlacklisted(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _Pv1Smartcontract.contract.Call(opts, &out, "isPolicyBlacklisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPolicyBlacklisted is a free data retrieval call binding the contract method 0x9131a44d.
//
// Solidity: function isPolicyBlacklisted(string ) view returns(bool)
func (_Pv1Smartcontract *Pv1SmartcontractSession) IsPolicyBlacklisted(arg0 string) (bool, error) {
	return _Pv1Smartcontract.Contract.IsPolicyBlacklisted(&_Pv1Smartcontract.CallOpts, arg0)
}

// IsPolicyBlacklisted is a free data retrieval call binding the contract method 0x9131a44d.
//
// Solidity: function isPolicyBlacklisted(string ) view returns(bool)
func (_Pv1Smartcontract *Pv1SmartcontractCallerSession) IsPolicyBlacklisted(arg0 string) (bool, error) {
	return _Pv1Smartcontract.Contract.IsPolicyBlacklisted(&_Pv1Smartcontract.CallOpts, arg0)
}

// LastVoted is a free data retrieval call binding the contract method 0x9a61df89.
//
// Solidity: function lastVoted(address ) view returns(uint256)
func (_Pv1Smartcontract *Pv1SmartcontractCaller) LastVoted(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pv1Smartcontract.contract.Call(opts, &out, "lastVoted", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastVoted is a free data retrieval call binding the contract method 0x9a61df89.
//
// Solidity: function lastVoted(address ) view returns(uint256)
func (_Pv1Smartcontract *Pv1SmartcontractSession) LastVoted(arg0 common.Address) (*big.Int, error) {
	return _Pv1Smartcontract.Contract.LastVoted(&_Pv1Smartcontract.CallOpts, arg0)
}

// LastVoted is a free data retrieval call binding the contract method 0x9a61df89.
//
// Solidity: function lastVoted(address ) view returns(uint256)
func (_Pv1Smartcontract *Pv1SmartcontractCallerSession) LastVoted(arg0 common.Address) (*big.Int, error) {
	return _Pv1Smartcontract.Contract.LastVoted(&_Pv1Smartcontract.CallOpts, arg0)
}

// MinIntervalBetweenVotes is a free data retrieval call binding the contract method 0xf99ac2c2.
//
// Solidity: function minIntervalBetweenVotes() view returns(uint256)
func (_Pv1Smartcontract *Pv1SmartcontractCaller) MinIntervalBetweenVotes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pv1Smartcontract.contract.Call(opts, &out, "minIntervalBetweenVotes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinIntervalBetweenVotes is a free data retrieval call binding the contract method 0xf99ac2c2.
//
// Solidity: function minIntervalBetweenVotes() view returns(uint256)
func (_Pv1Smartcontract *Pv1SmartcontractSession) MinIntervalBetweenVotes() (*big.Int, error) {
	return _Pv1Smartcontract.Contract.MinIntervalBetweenVotes(&_Pv1Smartcontract.CallOpts)
}

// MinIntervalBetweenVotes is a free data retrieval call binding the contract method 0xf99ac2c2.
//
// Solidity: function minIntervalBetweenVotes() view returns(uint256)
func (_Pv1Smartcontract *Pv1SmartcontractCallerSession) MinIntervalBetweenVotes() (*big.Int, error) {
	return _Pv1Smartcontract.Contract.MinIntervalBetweenVotes(&_Pv1Smartcontract.CallOpts)
}

// Policies is a free data retrieval call binding the contract method 0x4a7ba3e5.
//
// Solidity: function policies(string ) view returns(string bodyHash, uint256 groupID, uint256 createdAt)
func (_Pv1Smartcontract *Pv1SmartcontractCaller) Policies(opts *bind.CallOpts, arg0 string) (struct {
	BodyHash  string
	GroupID   *big.Int
	CreatedAt *big.Int
}, error) {
	var out []interface{}
	err := _Pv1Smartcontract.contract.Call(opts, &out, "policies", arg0)

	outstruct := new(struct {
		BodyHash  string
		GroupID   *big.Int
		CreatedAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BodyHash = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.GroupID = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CreatedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Policies is a free data retrieval call binding the contract method 0x4a7ba3e5.
//
// Solidity: function policies(string ) view returns(string bodyHash, uint256 groupID, uint256 createdAt)
func (_Pv1Smartcontract *Pv1SmartcontractSession) Policies(arg0 string) (struct {
	BodyHash  string
	GroupID   *big.Int
	CreatedAt *big.Int
}, error) {
	return _Pv1Smartcontract.Contract.Policies(&_Pv1Smartcontract.CallOpts, arg0)
}

// Policies is a free data retrieval call binding the contract method 0x4a7ba3e5.
//
// Solidity: function policies(string ) view returns(string bodyHash, uint256 groupID, uint256 createdAt)
func (_Pv1Smartcontract *Pv1SmartcontractCallerSession) Policies(arg0 string) (struct {
	BodyHash  string
	GroupID   *big.Int
	CreatedAt *big.Int
}, error) {
	return _Pv1Smartcontract.Contract.Policies(&_Pv1Smartcontract.CallOpts, arg0)
}

// PolicyBlacklistedReason is a free data retrieval call binding the contract method 0x707c6eaf.
//
// Solidity: function policyBlacklistedReason(string ) view returns(string)
func (_Pv1Smartcontract *Pv1SmartcontractCaller) PolicyBlacklistedReason(opts *bind.CallOpts, arg0 string) (string, error) {
	var out []interface{}
	err := _Pv1Smartcontract.contract.Call(opts, &out, "policyBlacklistedReason", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PolicyBlacklistedReason is a free data retrieval call binding the contract method 0x707c6eaf.
//
// Solidity: function policyBlacklistedReason(string ) view returns(string)
func (_Pv1Smartcontract *Pv1SmartcontractSession) PolicyBlacklistedReason(arg0 string) (string, error) {
	return _Pv1Smartcontract.Contract.PolicyBlacklistedReason(&_Pv1Smartcontract.CallOpts, arg0)
}

// PolicyBlacklistedReason is a free data retrieval call binding the contract method 0x707c6eaf.
//
// Solidity: function policyBlacklistedReason(string ) view returns(string)
func (_Pv1Smartcontract *Pv1SmartcontractCallerSession) PolicyBlacklistedReason(arg0 string) (string, error) {
	return _Pv1Smartcontract.Contract.PolicyBlacklistedReason(&_Pv1Smartcontract.CallOpts, arg0)
}

// Registered is a free data retrieval call binding the contract method 0xb2dd5c07.
//
// Solidity: function registered(address ) view returns(bool)
func (_Pv1Smartcontract *Pv1SmartcontractCaller) Registered(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Pv1Smartcontract.contract.Call(opts, &out, "registered", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Registered is a free data retrieval call binding the contract method 0xb2dd5c07.
//
// Solidity: function registered(address ) view returns(bool)
func (_Pv1Smartcontract *Pv1SmartcontractSession) Registered(arg0 common.Address) (bool, error) {
	return _Pv1Smartcontract.Contract.Registered(&_Pv1Smartcontract.CallOpts, arg0)
}

// Registered is a free data retrieval call binding the contract method 0xb2dd5c07.
//
// Solidity: function registered(address ) view returns(bool)
func (_Pv1Smartcontract *Pv1SmartcontractCallerSession) Registered(arg0 common.Address) (bool, error) {
	return _Pv1Smartcontract.Contract.Registered(&_Pv1Smartcontract.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Pv1Smartcontract *Pv1SmartcontractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Pv1Smartcontract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Pv1Smartcontract *Pv1SmartcontractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Pv1Smartcontract.Contract.SupportsInterface(&_Pv1Smartcontract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Pv1Smartcontract *Pv1SmartcontractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Pv1Smartcontract.Contract.SupportsInterface(&_Pv1Smartcontract.CallOpts, interfaceId)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Pv1Smartcontract *Pv1SmartcontractCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Pv1Smartcontract.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Pv1Smartcontract *Pv1SmartcontractSession) Version() (string, error) {
	return _Pv1Smartcontract.Contract.Version(&_Pv1Smartcontract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Pv1Smartcontract *Pv1SmartcontractCallerSession) Version() (string, error) {
	return _Pv1Smartcontract.Contract.Version(&_Pv1Smartcontract.CallOpts)
}

// VotedOnProposal is a free data retrieval call binding the contract method 0xfa606467.
//
// Solidity: function votedOnProposal(bytes32 ) view returns(bool)
func (_Pv1Smartcontract *Pv1SmartcontractCaller) VotedOnProposal(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Pv1Smartcontract.contract.Call(opts, &out, "votedOnProposal", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VotedOnProposal is a free data retrieval call binding the contract method 0xfa606467.
//
// Solidity: function votedOnProposal(bytes32 ) view returns(bool)
func (_Pv1Smartcontract *Pv1SmartcontractSession) VotedOnProposal(arg0 [32]byte) (bool, error) {
	return _Pv1Smartcontract.Contract.VotedOnProposal(&_Pv1Smartcontract.CallOpts, arg0)
}

// VotedOnProposal is a free data retrieval call binding the contract method 0xfa606467.
//
// Solidity: function votedOnProposal(bytes32 ) view returns(bool)
func (_Pv1Smartcontract *Pv1SmartcontractCallerSession) VotedOnProposal(arg0 [32]byte) (bool, error) {
	return _Pv1Smartcontract.Contract.VotedOnProposal(&_Pv1Smartcontract.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0xb99ef1fa.
//
// Solidity: function votes(string ) view returns(uint256)
func (_Pv1Smartcontract *Pv1SmartcontractCaller) Votes(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _Pv1Smartcontract.contract.Call(opts, &out, "votes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Votes is a free data retrieval call binding the contract method 0xb99ef1fa.
//
// Solidity: function votes(string ) view returns(uint256)
func (_Pv1Smartcontract *Pv1SmartcontractSession) Votes(arg0 string) (*big.Int, error) {
	return _Pv1Smartcontract.Contract.Votes(&_Pv1Smartcontract.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0xb99ef1fa.
//
// Solidity: function votes(string ) view returns(uint256)
func (_Pv1Smartcontract *Pv1SmartcontractCallerSession) Votes(arg0 string) (*big.Int, error) {
	return _Pv1Smartcontract.Contract.Votes(&_Pv1Smartcontract.CallOpts, arg0)
}

// CreateNewPolicy is a paid mutator transaction binding the contract method 0xd188ccc6.
//
// Solidity: function createNewPolicy(string policyID, uint256 gID, string bodyHash) returns()
func (_Pv1Smartcontract *Pv1SmartcontractTransactor) CreateNewPolicy(opts *bind.TransactOpts, policyID string, gID *big.Int, bodyHash string) (*types.Transaction, error) {
	return _Pv1Smartcontract.contract.Transact(opts, "createNewPolicy", policyID, gID, bodyHash)
}

// CreateNewPolicy is a paid mutator transaction binding the contract method 0xd188ccc6.
//
// Solidity: function createNewPolicy(string policyID, uint256 gID, string bodyHash) returns()
func (_Pv1Smartcontract *Pv1SmartcontractSession) CreateNewPolicy(policyID string, gID *big.Int, bodyHash string) (*types.Transaction, error) {
	return _Pv1Smartcontract.Contract.CreateNewPolicy(&_Pv1Smartcontract.TransactOpts, policyID, gID, bodyHash)
}

// CreateNewPolicy is a paid mutator transaction binding the contract method 0xd188ccc6.
//
// Solidity: function createNewPolicy(string policyID, uint256 gID, string bodyHash) returns()
func (_Pv1Smartcontract *Pv1SmartcontractTransactorSession) CreateNewPolicy(policyID string, gID *big.Int, bodyHash string) (*types.Transaction, error) {
	return _Pv1Smartcontract.Contract.CreateNewPolicy(&_Pv1Smartcontract.TransactOpts, policyID, gID, bodyHash)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Pv1Smartcontract *Pv1SmartcontractTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pv1Smartcontract.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Pv1Smartcontract *Pv1SmartcontractSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pv1Smartcontract.Contract.GrantRole(&_Pv1Smartcontract.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Pv1Smartcontract *Pv1SmartcontractTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pv1Smartcontract.Contract.GrantRole(&_Pv1Smartcontract.TransactOpts, role, account)
}

// RegisterAddress is a paid mutator transaction binding the contract method 0xded79106.
//
// Solidity: function registerAddress(address newAddress, bool isRegistered, uint256 groupID) returns()
func (_Pv1Smartcontract *Pv1SmartcontractTransactor) RegisterAddress(opts *bind.TransactOpts, newAddress common.Address, isRegistered bool, groupID *big.Int) (*types.Transaction, error) {
	return _Pv1Smartcontract.contract.Transact(opts, "registerAddress", newAddress, isRegistered, groupID)
}

// RegisterAddress is a paid mutator transaction binding the contract method 0xded79106.
//
// Solidity: function registerAddress(address newAddress, bool isRegistered, uint256 groupID) returns()
func (_Pv1Smartcontract *Pv1SmartcontractSession) RegisterAddress(newAddress common.Address, isRegistered bool, groupID *big.Int) (*types.Transaction, error) {
	return _Pv1Smartcontract.Contract.RegisterAddress(&_Pv1Smartcontract.TransactOpts, newAddress, isRegistered, groupID)
}

// RegisterAddress is a paid mutator transaction binding the contract method 0xded79106.
//
// Solidity: function registerAddress(address newAddress, bool isRegistered, uint256 groupID) returns()
func (_Pv1Smartcontract *Pv1SmartcontractTransactorSession) RegisterAddress(newAddress common.Address, isRegistered bool, groupID *big.Int) (*types.Transaction, error) {
	return _Pv1Smartcontract.Contract.RegisterAddress(&_Pv1Smartcontract.TransactOpts, newAddress, isRegistered, groupID)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Pv1Smartcontract *Pv1SmartcontractTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pv1Smartcontract.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Pv1Smartcontract *Pv1SmartcontractSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pv1Smartcontract.Contract.RenounceRole(&_Pv1Smartcontract.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Pv1Smartcontract *Pv1SmartcontractTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pv1Smartcontract.Contract.RenounceRole(&_Pv1Smartcontract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Pv1Smartcontract *Pv1SmartcontractTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pv1Smartcontract.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Pv1Smartcontract *Pv1SmartcontractSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pv1Smartcontract.Contract.RevokeRole(&_Pv1Smartcontract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Pv1Smartcontract *Pv1SmartcontractTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pv1Smartcontract.Contract.RevokeRole(&_Pv1Smartcontract.TransactOpts, role, account)
}

// SetMinIntervalBetweenVotes is a paid mutator transaction binding the contract method 0x2714e22f.
//
// Solidity: function setMinIntervalBetweenVotes(uint256 newMinInterval) returns()
func (_Pv1Smartcontract *Pv1SmartcontractTransactor) SetMinIntervalBetweenVotes(opts *bind.TransactOpts, newMinInterval *big.Int) (*types.Transaction, error) {
	return _Pv1Smartcontract.contract.Transact(opts, "setMinIntervalBetweenVotes", newMinInterval)
}

// SetMinIntervalBetweenVotes is a paid mutator transaction binding the contract method 0x2714e22f.
//
// Solidity: function setMinIntervalBetweenVotes(uint256 newMinInterval) returns()
func (_Pv1Smartcontract *Pv1SmartcontractSession) SetMinIntervalBetweenVotes(newMinInterval *big.Int) (*types.Transaction, error) {
	return _Pv1Smartcontract.Contract.SetMinIntervalBetweenVotes(&_Pv1Smartcontract.TransactOpts, newMinInterval)
}

// SetMinIntervalBetweenVotes is a paid mutator transaction binding the contract method 0x2714e22f.
//
// Solidity: function setMinIntervalBetweenVotes(uint256 newMinInterval) returns()
func (_Pv1Smartcontract *Pv1SmartcontractTransactorSession) SetMinIntervalBetweenVotes(newMinInterval *big.Int) (*types.Transaction, error) {
	return _Pv1Smartcontract.Contract.SetMinIntervalBetweenVotes(&_Pv1Smartcontract.TransactOpts, newMinInterval)
}

// SetPolicyBlacklist is a paid mutator transaction binding the contract method 0xdf4c4a44.
//
// Solidity: function setPolicyBlacklist(string policyID, bool blacklisted, string reason) returns()
func (_Pv1Smartcontract *Pv1SmartcontractTransactor) SetPolicyBlacklist(opts *bind.TransactOpts, policyID string, blacklisted bool, reason string) (*types.Transaction, error) {
	return _Pv1Smartcontract.contract.Transact(opts, "setPolicyBlacklist", policyID, blacklisted, reason)
}

// SetPolicyBlacklist is a paid mutator transaction binding the contract method 0xdf4c4a44.
//
// Solidity: function setPolicyBlacklist(string policyID, bool blacklisted, string reason) returns()
func (_Pv1Smartcontract *Pv1SmartcontractSession) SetPolicyBlacklist(policyID string, blacklisted bool, reason string) (*types.Transaction, error) {
	return _Pv1Smartcontract.Contract.SetPolicyBlacklist(&_Pv1Smartcontract.TransactOpts, policyID, blacklisted, reason)
}

// SetPolicyBlacklist is a paid mutator transaction binding the contract method 0xdf4c4a44.
//
// Solidity: function setPolicyBlacklist(string policyID, bool blacklisted, string reason) returns()
func (_Pv1Smartcontract *Pv1SmartcontractTransactorSession) SetPolicyBlacklist(policyID string, blacklisted bool, reason string) (*types.Transaction, error) {
	return _Pv1Smartcontract.Contract.SetPolicyBlacklist(&_Pv1Smartcontract.TransactOpts, policyID, blacklisted, reason)
}

// UnVote is a paid mutator transaction binding the contract method 0x8cbef3f1.
//
// Solidity: function unVote(string policyID) returns()
func (_Pv1Smartcontract *Pv1SmartcontractTransactor) UnVote(opts *bind.TransactOpts, policyID string) (*types.Transaction, error) {
	return _Pv1Smartcontract.contract.Transact(opts, "unVote", policyID)
}

// UnVote is a paid mutator transaction binding the contract method 0x8cbef3f1.
//
// Solidity: function unVote(string policyID) returns()
func (_Pv1Smartcontract *Pv1SmartcontractSession) UnVote(policyID string) (*types.Transaction, error) {
	return _Pv1Smartcontract.Contract.UnVote(&_Pv1Smartcontract.TransactOpts, policyID)
}

// UnVote is a paid mutator transaction binding the contract method 0x8cbef3f1.
//
// Solidity: function unVote(string policyID) returns()
func (_Pv1Smartcontract *Pv1SmartcontractTransactorSession) UnVote(policyID string) (*types.Transaction, error) {
	return _Pv1Smartcontract.Contract.UnVote(&_Pv1Smartcontract.TransactOpts, policyID)
}

// Vote is a paid mutator transaction binding the contract method 0xfc36e15b.
//
// Solidity: function vote(string policyID) returns()
func (_Pv1Smartcontract *Pv1SmartcontractTransactor) Vote(opts *bind.TransactOpts, policyID string) (*types.Transaction, error) {
	return _Pv1Smartcontract.contract.Transact(opts, "vote", policyID)
}

// Vote is a paid mutator transaction binding the contract method 0xfc36e15b.
//
// Solidity: function vote(string policyID) returns()
func (_Pv1Smartcontract *Pv1SmartcontractSession) Vote(policyID string) (*types.Transaction, error) {
	return _Pv1Smartcontract.Contract.Vote(&_Pv1Smartcontract.TransactOpts, policyID)
}

// Vote is a paid mutator transaction binding the contract method 0xfc36e15b.
//
// Solidity: function vote(string policyID) returns()
func (_Pv1Smartcontract *Pv1SmartcontractTransactorSession) Vote(policyID string) (*types.Transaction, error) {
	return _Pv1Smartcontract.Contract.Vote(&_Pv1Smartcontract.TransactOpts, policyID)
}

// Pv1SmartcontractBlacklistedIterator is returned from FilterBlacklisted and is used to iterate over the raw logs and unpacked data for Blacklisted events raised by the Pv1Smartcontract contract.
type Pv1SmartcontractBlacklistedIterator struct {
	Event *Pv1SmartcontractBlacklisted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Pv1SmartcontractBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pv1SmartcontractBlacklisted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Pv1SmartcontractBlacklisted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Pv1SmartcontractBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pv1SmartcontractBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pv1SmartcontractBlacklisted represents a Blacklisted event raised by the Pv1Smartcontract contract.
type Pv1SmartcontractBlacklisted struct {
	User        common.Address
	Blacklisted bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlacklisted is a free log retrieval operation binding the contract event 0xcf3473b85df1594d47b6958f29a32bea0abff9dd68296f7bf33443646793cfd8.
//
// Solidity: event Blacklisted(address user, bool blacklisted)
func (_Pv1Smartcontract *Pv1SmartcontractFilterer) FilterBlacklisted(opts *bind.FilterOpts) (*Pv1SmartcontractBlacklistedIterator, error) {

	logs, sub, err := _Pv1Smartcontract.contract.FilterLogs(opts, "Blacklisted")
	if err != nil {
		return nil, err
	}
	return &Pv1SmartcontractBlacklistedIterator{contract: _Pv1Smartcontract.contract, event: "Blacklisted", logs: logs, sub: sub}, nil
}

// WatchBlacklisted is a free log subscription operation binding the contract event 0xcf3473b85df1594d47b6958f29a32bea0abff9dd68296f7bf33443646793cfd8.
//
// Solidity: event Blacklisted(address user, bool blacklisted)
func (_Pv1Smartcontract *Pv1SmartcontractFilterer) WatchBlacklisted(opts *bind.WatchOpts, sink chan<- *Pv1SmartcontractBlacklisted) (event.Subscription, error) {

	logs, sub, err := _Pv1Smartcontract.contract.WatchLogs(opts, "Blacklisted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pv1SmartcontractBlacklisted)
				if err := _Pv1Smartcontract.contract.UnpackLog(event, "Blacklisted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBlacklisted is a log parse operation binding the contract event 0xcf3473b85df1594d47b6958f29a32bea0abff9dd68296f7bf33443646793cfd8.
//
// Solidity: event Blacklisted(address user, bool blacklisted)
func (_Pv1Smartcontract *Pv1SmartcontractFilterer) ParseBlacklisted(log types.Log) (*Pv1SmartcontractBlacklisted, error) {
	event := new(Pv1SmartcontractBlacklisted)
	if err := _Pv1Smartcontract.contract.UnpackLog(event, "Blacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Pv1SmartcontractNewPolicyIterator is returned from FilterNewPolicy and is used to iterate over the raw logs and unpacked data for NewPolicy events raised by the Pv1Smartcontract contract.
type Pv1SmartcontractNewPolicyIterator struct {
	Event *Pv1SmartcontractNewPolicy // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Pv1SmartcontractNewPolicyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pv1SmartcontractNewPolicy)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Pv1SmartcontractNewPolicy)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Pv1SmartcontractNewPolicyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pv1SmartcontractNewPolicyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pv1SmartcontractNewPolicy represents a NewPolicy event raised by the Pv1Smartcontract contract.
type Pv1SmartcontractNewPolicy struct {
	PolicyID string
	GroupID  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewPolicy is a free log retrieval operation binding the contract event 0x20e4a76a7d395731d66918bcdf3594e099b03d3077ba0343867cfdc7fce70abf.
//
// Solidity: event NewPolicy(string policyID, uint256 groupID)
func (_Pv1Smartcontract *Pv1SmartcontractFilterer) FilterNewPolicy(opts *bind.FilterOpts) (*Pv1SmartcontractNewPolicyIterator, error) {

	logs, sub, err := _Pv1Smartcontract.contract.FilterLogs(opts, "NewPolicy")
	if err != nil {
		return nil, err
	}
	return &Pv1SmartcontractNewPolicyIterator{contract: _Pv1Smartcontract.contract, event: "NewPolicy", logs: logs, sub: sub}, nil
}

// WatchNewPolicy is a free log subscription operation binding the contract event 0x20e4a76a7d395731d66918bcdf3594e099b03d3077ba0343867cfdc7fce70abf.
//
// Solidity: event NewPolicy(string policyID, uint256 groupID)
func (_Pv1Smartcontract *Pv1SmartcontractFilterer) WatchNewPolicy(opts *bind.WatchOpts, sink chan<- *Pv1SmartcontractNewPolicy) (event.Subscription, error) {

	logs, sub, err := _Pv1Smartcontract.contract.WatchLogs(opts, "NewPolicy")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pv1SmartcontractNewPolicy)
				if err := _Pv1Smartcontract.contract.UnpackLog(event, "NewPolicy", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewPolicy is a log parse operation binding the contract event 0x20e4a76a7d395731d66918bcdf3594e099b03d3077ba0343867cfdc7fce70abf.
//
// Solidity: event NewPolicy(string policyID, uint256 groupID)
func (_Pv1Smartcontract *Pv1SmartcontractFilterer) ParseNewPolicy(log types.Log) (*Pv1SmartcontractNewPolicy, error) {
	event := new(Pv1SmartcontractNewPolicy)
	if err := _Pv1Smartcontract.contract.UnpackLog(event, "NewPolicy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Pv1SmartcontractPolicyBlacklistedIterator is returned from FilterPolicyBlacklisted and is used to iterate over the raw logs and unpacked data for PolicyBlacklisted events raised by the Pv1Smartcontract contract.
type Pv1SmartcontractPolicyBlacklistedIterator struct {
	Event *Pv1SmartcontractPolicyBlacklisted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Pv1SmartcontractPolicyBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pv1SmartcontractPolicyBlacklisted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Pv1SmartcontractPolicyBlacklisted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Pv1SmartcontractPolicyBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pv1SmartcontractPolicyBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pv1SmartcontractPolicyBlacklisted represents a PolicyBlacklisted event raised by the Pv1Smartcontract contract.
type Pv1SmartcontractPolicyBlacklisted struct {
	PolicyID    string
	Blacklisted bool
	Reason      string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPolicyBlacklisted is a free log retrieval operation binding the contract event 0xe03285579bee722c625a02fb2b45070e578d9a320679123457906a0e90c9e6d3.
//
// Solidity: event PolicyBlacklisted(string policyID, bool blacklisted, string reason)
func (_Pv1Smartcontract *Pv1SmartcontractFilterer) FilterPolicyBlacklisted(opts *bind.FilterOpts) (*Pv1SmartcontractPolicyBlacklistedIterator, error) {

	logs, sub, err := _Pv1Smartcontract.contract.FilterLogs(opts, "PolicyBlacklisted")
	if err != nil {
		return nil, err
	}
	return &Pv1SmartcontractPolicyBlacklistedIterator{contract: _Pv1Smartcontract.contract, event: "PolicyBlacklisted", logs: logs, sub: sub}, nil
}

// WatchPolicyBlacklisted is a free log subscription operation binding the contract event 0xe03285579bee722c625a02fb2b45070e578d9a320679123457906a0e90c9e6d3.
//
// Solidity: event PolicyBlacklisted(string policyID, bool blacklisted, string reason)
func (_Pv1Smartcontract *Pv1SmartcontractFilterer) WatchPolicyBlacklisted(opts *bind.WatchOpts, sink chan<- *Pv1SmartcontractPolicyBlacklisted) (event.Subscription, error) {

	logs, sub, err := _Pv1Smartcontract.contract.WatchLogs(opts, "PolicyBlacklisted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pv1SmartcontractPolicyBlacklisted)
				if err := _Pv1Smartcontract.contract.UnpackLog(event, "PolicyBlacklisted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePolicyBlacklisted is a log parse operation binding the contract event 0xe03285579bee722c625a02fb2b45070e578d9a320679123457906a0e90c9e6d3.
//
// Solidity: event PolicyBlacklisted(string policyID, bool blacklisted, string reason)
func (_Pv1Smartcontract *Pv1SmartcontractFilterer) ParsePolicyBlacklisted(log types.Log) (*Pv1SmartcontractPolicyBlacklisted, error) {
	event := new(Pv1SmartcontractPolicyBlacklisted)
	if err := _Pv1Smartcontract.contract.UnpackLog(event, "PolicyBlacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Pv1SmartcontractRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Pv1Smartcontract contract.
type Pv1SmartcontractRoleAdminChangedIterator struct {
	Event *Pv1SmartcontractRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Pv1SmartcontractRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pv1SmartcontractRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Pv1SmartcontractRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Pv1SmartcontractRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pv1SmartcontractRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pv1SmartcontractRoleAdminChanged represents a RoleAdminChanged event raised by the Pv1Smartcontract contract.
type Pv1SmartcontractRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Pv1Smartcontract *Pv1SmartcontractFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*Pv1SmartcontractRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Pv1Smartcontract.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &Pv1SmartcontractRoleAdminChangedIterator{contract: _Pv1Smartcontract.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Pv1Smartcontract *Pv1SmartcontractFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *Pv1SmartcontractRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Pv1Smartcontract.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pv1SmartcontractRoleAdminChanged)
				if err := _Pv1Smartcontract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Pv1Smartcontract *Pv1SmartcontractFilterer) ParseRoleAdminChanged(log types.Log) (*Pv1SmartcontractRoleAdminChanged, error) {
	event := new(Pv1SmartcontractRoleAdminChanged)
	if err := _Pv1Smartcontract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Pv1SmartcontractRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Pv1Smartcontract contract.
type Pv1SmartcontractRoleGrantedIterator struct {
	Event *Pv1SmartcontractRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Pv1SmartcontractRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pv1SmartcontractRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Pv1SmartcontractRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Pv1SmartcontractRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pv1SmartcontractRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pv1SmartcontractRoleGranted represents a RoleGranted event raised by the Pv1Smartcontract contract.
type Pv1SmartcontractRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pv1Smartcontract *Pv1SmartcontractFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*Pv1SmartcontractRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Pv1Smartcontract.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &Pv1SmartcontractRoleGrantedIterator{contract: _Pv1Smartcontract.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pv1Smartcontract *Pv1SmartcontractFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *Pv1SmartcontractRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Pv1Smartcontract.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pv1SmartcontractRoleGranted)
				if err := _Pv1Smartcontract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pv1Smartcontract *Pv1SmartcontractFilterer) ParseRoleGranted(log types.Log) (*Pv1SmartcontractRoleGranted, error) {
	event := new(Pv1SmartcontractRoleGranted)
	if err := _Pv1Smartcontract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Pv1SmartcontractRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Pv1Smartcontract contract.
type Pv1SmartcontractRoleRevokedIterator struct {
	Event *Pv1SmartcontractRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Pv1SmartcontractRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pv1SmartcontractRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Pv1SmartcontractRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Pv1SmartcontractRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pv1SmartcontractRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pv1SmartcontractRoleRevoked represents a RoleRevoked event raised by the Pv1Smartcontract contract.
type Pv1SmartcontractRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pv1Smartcontract *Pv1SmartcontractFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*Pv1SmartcontractRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Pv1Smartcontract.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &Pv1SmartcontractRoleRevokedIterator{contract: _Pv1Smartcontract.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pv1Smartcontract *Pv1SmartcontractFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *Pv1SmartcontractRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Pv1Smartcontract.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pv1SmartcontractRoleRevoked)
				if err := _Pv1Smartcontract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pv1Smartcontract *Pv1SmartcontractFilterer) ParseRoleRevoked(log types.Log) (*Pv1SmartcontractRoleRevoked, error) {
	event := new(Pv1SmartcontractRoleRevoked)
	if err := _Pv1Smartcontract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Pv1SmartcontractUnVotedIterator is returned from FilterUnVoted and is used to iterate over the raw logs and unpacked data for UnVoted events raised by the Pv1Smartcontract contract.
type Pv1SmartcontractUnVotedIterator struct {
	Event *Pv1SmartcontractUnVoted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Pv1SmartcontractUnVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pv1SmartcontractUnVoted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Pv1SmartcontractUnVoted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Pv1SmartcontractUnVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pv1SmartcontractUnVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pv1SmartcontractUnVoted represents a UnVoted event raised by the Pv1Smartcontract contract.
type Pv1SmartcontractUnVoted struct {
	PolicyID   string
	GroupID    *big.Int
	TotalVotes *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUnVoted is a free log retrieval operation binding the contract event 0xa5ab64c941a274cfae04b116dfc0777bab498844134daccc294eeb81578e8b4e.
//
// Solidity: event UnVoted(string policyID, uint256 groupID, uint256 totalVotes)
func (_Pv1Smartcontract *Pv1SmartcontractFilterer) FilterUnVoted(opts *bind.FilterOpts) (*Pv1SmartcontractUnVotedIterator, error) {

	logs, sub, err := _Pv1Smartcontract.contract.FilterLogs(opts, "UnVoted")
	if err != nil {
		return nil, err
	}
	return &Pv1SmartcontractUnVotedIterator{contract: _Pv1Smartcontract.contract, event: "UnVoted", logs: logs, sub: sub}, nil
}

// WatchUnVoted is a free log subscription operation binding the contract event 0xa5ab64c941a274cfae04b116dfc0777bab498844134daccc294eeb81578e8b4e.
//
// Solidity: event UnVoted(string policyID, uint256 groupID, uint256 totalVotes)
func (_Pv1Smartcontract *Pv1SmartcontractFilterer) WatchUnVoted(opts *bind.WatchOpts, sink chan<- *Pv1SmartcontractUnVoted) (event.Subscription, error) {

	logs, sub, err := _Pv1Smartcontract.contract.WatchLogs(opts, "UnVoted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pv1SmartcontractUnVoted)
				if err := _Pv1Smartcontract.contract.UnpackLog(event, "UnVoted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnVoted is a log parse operation binding the contract event 0xa5ab64c941a274cfae04b116dfc0777bab498844134daccc294eeb81578e8b4e.
//
// Solidity: event UnVoted(string policyID, uint256 groupID, uint256 totalVotes)
func (_Pv1Smartcontract *Pv1SmartcontractFilterer) ParseUnVoted(log types.Log) (*Pv1SmartcontractUnVoted, error) {
	event := new(Pv1SmartcontractUnVoted)
	if err := _Pv1Smartcontract.contract.UnpackLog(event, "UnVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Pv1SmartcontractVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the Pv1Smartcontract contract.
type Pv1SmartcontractVotedIterator struct {
	Event *Pv1SmartcontractVoted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Pv1SmartcontractVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pv1SmartcontractVoted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Pv1SmartcontractVoted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Pv1SmartcontractVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pv1SmartcontractVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pv1SmartcontractVoted represents a Voted event raised by the Pv1Smartcontract contract.
type Pv1SmartcontractVoted struct {
	PolicyID   string
	GroupID    *big.Int
	TotalVotes *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x757fc7d2692c5e5698afb5c7da2ebdc75e0f7418b820f9114894ac7f462be686.
//
// Solidity: event Voted(string policyID, uint256 groupID, uint256 totalVotes)
func (_Pv1Smartcontract *Pv1SmartcontractFilterer) FilterVoted(opts *bind.FilterOpts) (*Pv1SmartcontractVotedIterator, error) {

	logs, sub, err := _Pv1Smartcontract.contract.FilterLogs(opts, "Voted")
	if err != nil {
		return nil, err
	}
	return &Pv1SmartcontractVotedIterator{contract: _Pv1Smartcontract.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x757fc7d2692c5e5698afb5c7da2ebdc75e0f7418b820f9114894ac7f462be686.
//
// Solidity: event Voted(string policyID, uint256 groupID, uint256 totalVotes)
func (_Pv1Smartcontract *Pv1SmartcontractFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *Pv1SmartcontractVoted) (event.Subscription, error) {

	logs, sub, err := _Pv1Smartcontract.contract.WatchLogs(opts, "Voted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pv1SmartcontractVoted)
				if err := _Pv1Smartcontract.contract.UnpackLog(event, "Voted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVoted is a log parse operation binding the contract event 0x757fc7d2692c5e5698afb5c7da2ebdc75e0f7418b820f9114894ac7f462be686.
//
// Solidity: event Voted(string policyID, uint256 groupID, uint256 totalVotes)
func (_Pv1Smartcontract *Pv1SmartcontractFilterer) ParseVoted(log types.Log) (*Pv1SmartcontractVoted, error) {
	event := new(Pv1SmartcontractVoted)
	if err := _Pv1Smartcontract.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
