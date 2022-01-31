// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pv0_smartcontract

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Pv0SmartcontractABI is the input ABI used to generate the binding from.
const Pv0SmartcontractABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"blacklisted\",\"type\":\"bool\"}],\"name\":\"Blacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"policyID\",\"type\":\"string\"}],\"name\":\"NewPolicy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"policyID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"blacklisted\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"PolicyBlacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"policyID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalVotes\",\"type\":\"uint256\"}],\"name\":\"UnVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"policyID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalVotes\",\"type\":\"uint256\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADDRESS_REGISTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POLICY_BLACKLISTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POLICY_CREATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"policyID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bodyHash\",\"type\":\"string\"}],\"name\":\"createNewPolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"isPolicyBlacklisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastVoted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"policies\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"bodyHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"policyBlacklistedReason\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"}],\"name\":\"registerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"registered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"policyID\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"blacklisted\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"setPolicyBlacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"policyID\",\"type\":\"string\"}],\"name\":\"unVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"policyID\",\"type\":\"string\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"votedOnProposal\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"votes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Pv0Smartcontract is an auto generated Go binding around an Ethereum contract.
type Pv0Smartcontract struct {
	Pv0SmartcontractCaller     // Read-only binding to the contract
	Pv0SmartcontractTransactor // Write-only binding to the contract
	Pv0SmartcontractFilterer   // Log filterer for contract events
}

// Pv0SmartcontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type Pv0SmartcontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Pv0SmartcontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Pv0SmartcontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Pv0SmartcontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Pv0SmartcontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Pv0SmartcontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Pv0SmartcontractSession struct {
	Contract     *Pv0Smartcontract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Pv0SmartcontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Pv0SmartcontractCallerSession struct {
	Contract *Pv0SmartcontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// Pv0SmartcontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Pv0SmartcontractTransactorSession struct {
	Contract     *Pv0SmartcontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// Pv0SmartcontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type Pv0SmartcontractRaw struct {
	Contract *Pv0Smartcontract // Generic contract binding to access the raw methods on
}

// Pv0SmartcontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Pv0SmartcontractCallerRaw struct {
	Contract *Pv0SmartcontractCaller // Generic read-only contract binding to access the raw methods on
}

// Pv0SmartcontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Pv0SmartcontractTransactorRaw struct {
	Contract *Pv0SmartcontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPv0Smartcontract creates a new instance of Pv0Smartcontract, bound to a specific deployed contract.
func NewPv0Smartcontract(address common.Address, backend bind.ContractBackend) (*Pv0Smartcontract, error) {
	contract, err := bindPv0Smartcontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pv0Smartcontract{Pv0SmartcontractCaller: Pv0SmartcontractCaller{contract: contract}, Pv0SmartcontractTransactor: Pv0SmartcontractTransactor{contract: contract}, Pv0SmartcontractFilterer: Pv0SmartcontractFilterer{contract: contract}}, nil
}

// NewPv0SmartcontractCaller creates a new read-only instance of Pv0Smartcontract, bound to a specific deployed contract.
func NewPv0SmartcontractCaller(address common.Address, caller bind.ContractCaller) (*Pv0SmartcontractCaller, error) {
	contract, err := bindPv0Smartcontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Pv0SmartcontractCaller{contract: contract}, nil
}

// NewPv0SmartcontractTransactor creates a new write-only instance of Pv0Smartcontract, bound to a specific deployed contract.
func NewPv0SmartcontractTransactor(address common.Address, transactor bind.ContractTransactor) (*Pv0SmartcontractTransactor, error) {
	contract, err := bindPv0Smartcontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Pv0SmartcontractTransactor{contract: contract}, nil
}

// NewPv0SmartcontractFilterer creates a new log filterer instance of Pv0Smartcontract, bound to a specific deployed contract.
func NewPv0SmartcontractFilterer(address common.Address, filterer bind.ContractFilterer) (*Pv0SmartcontractFilterer, error) {
	contract, err := bindPv0Smartcontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Pv0SmartcontractFilterer{contract: contract}, nil
}

// bindPv0Smartcontract binds a generic wrapper to an already deployed contract.
func bindPv0Smartcontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Pv0SmartcontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pv0Smartcontract *Pv0SmartcontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pv0Smartcontract.Contract.Pv0SmartcontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pv0Smartcontract *Pv0SmartcontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pv0Smartcontract.Contract.Pv0SmartcontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pv0Smartcontract *Pv0SmartcontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pv0Smartcontract.Contract.Pv0SmartcontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pv0Smartcontract *Pv0SmartcontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pv0Smartcontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pv0Smartcontract *Pv0SmartcontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pv0Smartcontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pv0Smartcontract *Pv0SmartcontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pv0Smartcontract.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSREGISTERROLE is a free data retrieval call binding the contract method 0xb464d8b7.
//
// Solidity: function ADDRESS_REGISTER_ROLE() view returns(bytes32)
func (_Pv0Smartcontract *Pv0SmartcontractCaller) ADDRESSREGISTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pv0Smartcontract.contract.Call(opts, &out, "ADDRESS_REGISTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADDRESSREGISTERROLE is a free data retrieval call binding the contract method 0xb464d8b7.
//
// Solidity: function ADDRESS_REGISTER_ROLE() view returns(bytes32)
func (_Pv0Smartcontract *Pv0SmartcontractSession) ADDRESSREGISTERROLE() ([32]byte, error) {
	return _Pv0Smartcontract.Contract.ADDRESSREGISTERROLE(&_Pv0Smartcontract.CallOpts)
}

// ADDRESSREGISTERROLE is a free data retrieval call binding the contract method 0xb464d8b7.
//
// Solidity: function ADDRESS_REGISTER_ROLE() view returns(bytes32)
func (_Pv0Smartcontract *Pv0SmartcontractCallerSession) ADDRESSREGISTERROLE() ([32]byte, error) {
	return _Pv0Smartcontract.Contract.ADDRESSREGISTERROLE(&_Pv0Smartcontract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Pv0Smartcontract *Pv0SmartcontractCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pv0Smartcontract.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Pv0Smartcontract *Pv0SmartcontractSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Pv0Smartcontract.Contract.DEFAULTADMINROLE(&_Pv0Smartcontract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Pv0Smartcontract *Pv0SmartcontractCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Pv0Smartcontract.Contract.DEFAULTADMINROLE(&_Pv0Smartcontract.CallOpts)
}

// POLICYBLACKLISTERROLE is a free data retrieval call binding the contract method 0xfc86449a.
//
// Solidity: function POLICY_BLACKLISTER_ROLE() view returns(bytes32)
func (_Pv0Smartcontract *Pv0SmartcontractCaller) POLICYBLACKLISTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pv0Smartcontract.contract.Call(opts, &out, "POLICY_BLACKLISTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// POLICYBLACKLISTERROLE is a free data retrieval call binding the contract method 0xfc86449a.
//
// Solidity: function POLICY_BLACKLISTER_ROLE() view returns(bytes32)
func (_Pv0Smartcontract *Pv0SmartcontractSession) POLICYBLACKLISTERROLE() ([32]byte, error) {
	return _Pv0Smartcontract.Contract.POLICYBLACKLISTERROLE(&_Pv0Smartcontract.CallOpts)
}

// POLICYBLACKLISTERROLE is a free data retrieval call binding the contract method 0xfc86449a.
//
// Solidity: function POLICY_BLACKLISTER_ROLE() view returns(bytes32)
func (_Pv0Smartcontract *Pv0SmartcontractCallerSession) POLICYBLACKLISTERROLE() ([32]byte, error) {
	return _Pv0Smartcontract.Contract.POLICYBLACKLISTERROLE(&_Pv0Smartcontract.CallOpts)
}

// POLICYCREATORROLE is a free data retrieval call binding the contract method 0x45f739ff.
//
// Solidity: function POLICY_CREATOR_ROLE() view returns(bytes32)
func (_Pv0Smartcontract *Pv0SmartcontractCaller) POLICYCREATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pv0Smartcontract.contract.Call(opts, &out, "POLICY_CREATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// POLICYCREATORROLE is a free data retrieval call binding the contract method 0x45f739ff.
//
// Solidity: function POLICY_CREATOR_ROLE() view returns(bytes32)
func (_Pv0Smartcontract *Pv0SmartcontractSession) POLICYCREATORROLE() ([32]byte, error) {
	return _Pv0Smartcontract.Contract.POLICYCREATORROLE(&_Pv0Smartcontract.CallOpts)
}

// POLICYCREATORROLE is a free data retrieval call binding the contract method 0x45f739ff.
//
// Solidity: function POLICY_CREATOR_ROLE() view returns(bytes32)
func (_Pv0Smartcontract *Pv0SmartcontractCallerSession) POLICYCREATORROLE() ([32]byte, error) {
	return _Pv0Smartcontract.Contract.POLICYCREATORROLE(&_Pv0Smartcontract.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Pv0Smartcontract *Pv0SmartcontractCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Pv0Smartcontract.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Pv0Smartcontract *Pv0SmartcontractSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Pv0Smartcontract.Contract.GetRoleAdmin(&_Pv0Smartcontract.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Pv0Smartcontract *Pv0SmartcontractCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Pv0Smartcontract.Contract.GetRoleAdmin(&_Pv0Smartcontract.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Pv0Smartcontract *Pv0SmartcontractCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Pv0Smartcontract.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Pv0Smartcontract *Pv0SmartcontractSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Pv0Smartcontract.Contract.HasRole(&_Pv0Smartcontract.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Pv0Smartcontract *Pv0SmartcontractCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Pv0Smartcontract.Contract.HasRole(&_Pv0Smartcontract.CallOpts, role, account)
}

// IsPolicyBlacklisted is a free data retrieval call binding the contract method 0x9131a44d.
//
// Solidity: function isPolicyBlacklisted(string ) view returns(bool)
func (_Pv0Smartcontract *Pv0SmartcontractCaller) IsPolicyBlacklisted(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _Pv0Smartcontract.contract.Call(opts, &out, "isPolicyBlacklisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPolicyBlacklisted is a free data retrieval call binding the contract method 0x9131a44d.
//
// Solidity: function isPolicyBlacklisted(string ) view returns(bool)
func (_Pv0Smartcontract *Pv0SmartcontractSession) IsPolicyBlacklisted(arg0 string) (bool, error) {
	return _Pv0Smartcontract.Contract.IsPolicyBlacklisted(&_Pv0Smartcontract.CallOpts, arg0)
}

// IsPolicyBlacklisted is a free data retrieval call binding the contract method 0x9131a44d.
//
// Solidity: function isPolicyBlacklisted(string ) view returns(bool)
func (_Pv0Smartcontract *Pv0SmartcontractCallerSession) IsPolicyBlacklisted(arg0 string) (bool, error) {
	return _Pv0Smartcontract.Contract.IsPolicyBlacklisted(&_Pv0Smartcontract.CallOpts, arg0)
}

// LastVoted is a free data retrieval call binding the contract method 0x9a61df89.
//
// Solidity: function lastVoted(address ) view returns(uint256)
func (_Pv0Smartcontract *Pv0SmartcontractCaller) LastVoted(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pv0Smartcontract.contract.Call(opts, &out, "lastVoted", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastVoted is a free data retrieval call binding the contract method 0x9a61df89.
//
// Solidity: function lastVoted(address ) view returns(uint256)
func (_Pv0Smartcontract *Pv0SmartcontractSession) LastVoted(arg0 common.Address) (*big.Int, error) {
	return _Pv0Smartcontract.Contract.LastVoted(&_Pv0Smartcontract.CallOpts, arg0)
}

// LastVoted is a free data retrieval call binding the contract method 0x9a61df89.
//
// Solidity: function lastVoted(address ) view returns(uint256)
func (_Pv0Smartcontract *Pv0SmartcontractCallerSession) LastVoted(arg0 common.Address) (*big.Int, error) {
	return _Pv0Smartcontract.Contract.LastVoted(&_Pv0Smartcontract.CallOpts, arg0)
}

// Policies is a free data retrieval call binding the contract method 0x4a7ba3e5.
//
// Solidity: function policies(string ) view returns(string bodyHash, uint256 createdAt)
func (_Pv0Smartcontract *Pv0SmartcontractCaller) Policies(opts *bind.CallOpts, arg0 string) (struct {
	BodyHash  string
	CreatedAt *big.Int
}, error) {
	var out []interface{}
	err := _Pv0Smartcontract.contract.Call(opts, &out, "policies", arg0)

	outstruct := new(struct {
		BodyHash  string
		CreatedAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BodyHash = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.CreatedAt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Policies is a free data retrieval call binding the contract method 0x4a7ba3e5.
//
// Solidity: function policies(string ) view returns(string bodyHash, uint256 createdAt)
func (_Pv0Smartcontract *Pv0SmartcontractSession) Policies(arg0 string) (struct {
	BodyHash  string
	CreatedAt *big.Int
}, error) {
	return _Pv0Smartcontract.Contract.Policies(&_Pv0Smartcontract.CallOpts, arg0)
}

// Policies is a free data retrieval call binding the contract method 0x4a7ba3e5.
//
// Solidity: function policies(string ) view returns(string bodyHash, uint256 createdAt)
func (_Pv0Smartcontract *Pv0SmartcontractCallerSession) Policies(arg0 string) (struct {
	BodyHash  string
	CreatedAt *big.Int
}, error) {
	return _Pv0Smartcontract.Contract.Policies(&_Pv0Smartcontract.CallOpts, arg0)
}

// PolicyBlacklistedReason is a free data retrieval call binding the contract method 0x707c6eaf.
//
// Solidity: function policyBlacklistedReason(string ) view returns(string)
func (_Pv0Smartcontract *Pv0SmartcontractCaller) PolicyBlacklistedReason(opts *bind.CallOpts, arg0 string) (string, error) {
	var out []interface{}
	err := _Pv0Smartcontract.contract.Call(opts, &out, "policyBlacklistedReason", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PolicyBlacklistedReason is a free data retrieval call binding the contract method 0x707c6eaf.
//
// Solidity: function policyBlacklistedReason(string ) view returns(string)
func (_Pv0Smartcontract *Pv0SmartcontractSession) PolicyBlacklistedReason(arg0 string) (string, error) {
	return _Pv0Smartcontract.Contract.PolicyBlacklistedReason(&_Pv0Smartcontract.CallOpts, arg0)
}

// PolicyBlacklistedReason is a free data retrieval call binding the contract method 0x707c6eaf.
//
// Solidity: function policyBlacklistedReason(string ) view returns(string)
func (_Pv0Smartcontract *Pv0SmartcontractCallerSession) PolicyBlacklistedReason(arg0 string) (string, error) {
	return _Pv0Smartcontract.Contract.PolicyBlacklistedReason(&_Pv0Smartcontract.CallOpts, arg0)
}

// Registered is a free data retrieval call binding the contract method 0xb2dd5c07.
//
// Solidity: function registered(address ) view returns(bool)
func (_Pv0Smartcontract *Pv0SmartcontractCaller) Registered(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Pv0Smartcontract.contract.Call(opts, &out, "registered", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Registered is a free data retrieval call binding the contract method 0xb2dd5c07.
//
// Solidity: function registered(address ) view returns(bool)
func (_Pv0Smartcontract *Pv0SmartcontractSession) Registered(arg0 common.Address) (bool, error) {
	return _Pv0Smartcontract.Contract.Registered(&_Pv0Smartcontract.CallOpts, arg0)
}

// Registered is a free data retrieval call binding the contract method 0xb2dd5c07.
//
// Solidity: function registered(address ) view returns(bool)
func (_Pv0Smartcontract *Pv0SmartcontractCallerSession) Registered(arg0 common.Address) (bool, error) {
	return _Pv0Smartcontract.Contract.Registered(&_Pv0Smartcontract.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Pv0Smartcontract *Pv0SmartcontractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Pv0Smartcontract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Pv0Smartcontract *Pv0SmartcontractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Pv0Smartcontract.Contract.SupportsInterface(&_Pv0Smartcontract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Pv0Smartcontract *Pv0SmartcontractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Pv0Smartcontract.Contract.SupportsInterface(&_Pv0Smartcontract.CallOpts, interfaceId)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Pv0Smartcontract *Pv0SmartcontractCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Pv0Smartcontract.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Pv0Smartcontract *Pv0SmartcontractSession) Version() (string, error) {
	return _Pv0Smartcontract.Contract.Version(&_Pv0Smartcontract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Pv0Smartcontract *Pv0SmartcontractCallerSession) Version() (string, error) {
	return _Pv0Smartcontract.Contract.Version(&_Pv0Smartcontract.CallOpts)
}

// VotedOnProposal is a free data retrieval call binding the contract method 0x4e8511c7.
//
// Solidity: function votedOnProposal(address ) view returns(string)
func (_Pv0Smartcontract *Pv0SmartcontractCaller) VotedOnProposal(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _Pv0Smartcontract.contract.Call(opts, &out, "votedOnProposal", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VotedOnProposal is a free data retrieval call binding the contract method 0x4e8511c7.
//
// Solidity: function votedOnProposal(address ) view returns(string)
func (_Pv0Smartcontract *Pv0SmartcontractSession) VotedOnProposal(arg0 common.Address) (string, error) {
	return _Pv0Smartcontract.Contract.VotedOnProposal(&_Pv0Smartcontract.CallOpts, arg0)
}

// VotedOnProposal is a free data retrieval call binding the contract method 0x4e8511c7.
//
// Solidity: function votedOnProposal(address ) view returns(string)
func (_Pv0Smartcontract *Pv0SmartcontractCallerSession) VotedOnProposal(arg0 common.Address) (string, error) {
	return _Pv0Smartcontract.Contract.VotedOnProposal(&_Pv0Smartcontract.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0xb99ef1fa.
//
// Solidity: function votes(string ) view returns(uint256)
func (_Pv0Smartcontract *Pv0SmartcontractCaller) Votes(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _Pv0Smartcontract.contract.Call(opts, &out, "votes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Votes is a free data retrieval call binding the contract method 0xb99ef1fa.
//
// Solidity: function votes(string ) view returns(uint256)
func (_Pv0Smartcontract *Pv0SmartcontractSession) Votes(arg0 string) (*big.Int, error) {
	return _Pv0Smartcontract.Contract.Votes(&_Pv0Smartcontract.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0xb99ef1fa.
//
// Solidity: function votes(string ) view returns(uint256)
func (_Pv0Smartcontract *Pv0SmartcontractCallerSession) Votes(arg0 string) (*big.Int, error) {
	return _Pv0Smartcontract.Contract.Votes(&_Pv0Smartcontract.CallOpts, arg0)
}

// CreateNewPolicy is a paid mutator transaction binding the contract method 0x63df9ce5.
//
// Solidity: function createNewPolicy(string policyID, string bodyHash) returns()
func (_Pv0Smartcontract *Pv0SmartcontractTransactor) CreateNewPolicy(opts *bind.TransactOpts, policyID string, bodyHash string) (*types.Transaction, error) {
	return _Pv0Smartcontract.contract.Transact(opts, "createNewPolicy", policyID, bodyHash)
}

// CreateNewPolicy is a paid mutator transaction binding the contract method 0x63df9ce5.
//
// Solidity: function createNewPolicy(string policyID, string bodyHash) returns()
func (_Pv0Smartcontract *Pv0SmartcontractSession) CreateNewPolicy(policyID string, bodyHash string) (*types.Transaction, error) {
	return _Pv0Smartcontract.Contract.CreateNewPolicy(&_Pv0Smartcontract.TransactOpts, policyID, bodyHash)
}

// CreateNewPolicy is a paid mutator transaction binding the contract method 0x63df9ce5.
//
// Solidity: function createNewPolicy(string policyID, string bodyHash) returns()
func (_Pv0Smartcontract *Pv0SmartcontractTransactorSession) CreateNewPolicy(policyID string, bodyHash string) (*types.Transaction, error) {
	return _Pv0Smartcontract.Contract.CreateNewPolicy(&_Pv0Smartcontract.TransactOpts, policyID, bodyHash)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Pv0Smartcontract *Pv0SmartcontractTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pv0Smartcontract.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Pv0Smartcontract *Pv0SmartcontractSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pv0Smartcontract.Contract.GrantRole(&_Pv0Smartcontract.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Pv0Smartcontract *Pv0SmartcontractTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pv0Smartcontract.Contract.GrantRole(&_Pv0Smartcontract.TransactOpts, role, account)
}

// RegisterAddress is a paid mutator transaction binding the contract method 0x634730da.
//
// Solidity: function registerAddress(address newAddress, bool isRegistered) returns()
func (_Pv0Smartcontract *Pv0SmartcontractTransactor) RegisterAddress(opts *bind.TransactOpts, newAddress common.Address, isRegistered bool) (*types.Transaction, error) {
	return _Pv0Smartcontract.contract.Transact(opts, "registerAddress", newAddress, isRegistered)
}

// RegisterAddress is a paid mutator transaction binding the contract method 0x634730da.
//
// Solidity: function registerAddress(address newAddress, bool isRegistered) returns()
func (_Pv0Smartcontract *Pv0SmartcontractSession) RegisterAddress(newAddress common.Address, isRegistered bool) (*types.Transaction, error) {
	return _Pv0Smartcontract.Contract.RegisterAddress(&_Pv0Smartcontract.TransactOpts, newAddress, isRegistered)
}

// RegisterAddress is a paid mutator transaction binding the contract method 0x634730da.
//
// Solidity: function registerAddress(address newAddress, bool isRegistered) returns()
func (_Pv0Smartcontract *Pv0SmartcontractTransactorSession) RegisterAddress(newAddress common.Address, isRegistered bool) (*types.Transaction, error) {
	return _Pv0Smartcontract.Contract.RegisterAddress(&_Pv0Smartcontract.TransactOpts, newAddress, isRegistered)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Pv0Smartcontract *Pv0SmartcontractTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pv0Smartcontract.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Pv0Smartcontract *Pv0SmartcontractSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pv0Smartcontract.Contract.RenounceRole(&_Pv0Smartcontract.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Pv0Smartcontract *Pv0SmartcontractTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pv0Smartcontract.Contract.RenounceRole(&_Pv0Smartcontract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Pv0Smartcontract *Pv0SmartcontractTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pv0Smartcontract.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Pv0Smartcontract *Pv0SmartcontractSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pv0Smartcontract.Contract.RevokeRole(&_Pv0Smartcontract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Pv0Smartcontract *Pv0SmartcontractTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pv0Smartcontract.Contract.RevokeRole(&_Pv0Smartcontract.TransactOpts, role, account)
}

// SetPolicyBlacklist is a paid mutator transaction binding the contract method 0xdf4c4a44.
//
// Solidity: function setPolicyBlacklist(string policyID, bool blacklisted, string reason) returns()
func (_Pv0Smartcontract *Pv0SmartcontractTransactor) SetPolicyBlacklist(opts *bind.TransactOpts, policyID string, blacklisted bool, reason string) (*types.Transaction, error) {
	return _Pv0Smartcontract.contract.Transact(opts, "setPolicyBlacklist", policyID, blacklisted, reason)
}

// SetPolicyBlacklist is a paid mutator transaction binding the contract method 0xdf4c4a44.
//
// Solidity: function setPolicyBlacklist(string policyID, bool blacklisted, string reason) returns()
func (_Pv0Smartcontract *Pv0SmartcontractSession) SetPolicyBlacklist(policyID string, blacklisted bool, reason string) (*types.Transaction, error) {
	return _Pv0Smartcontract.Contract.SetPolicyBlacklist(&_Pv0Smartcontract.TransactOpts, policyID, blacklisted, reason)
}

// SetPolicyBlacklist is a paid mutator transaction binding the contract method 0xdf4c4a44.
//
// Solidity: function setPolicyBlacklist(string policyID, bool blacklisted, string reason) returns()
func (_Pv0Smartcontract *Pv0SmartcontractTransactorSession) SetPolicyBlacklist(policyID string, blacklisted bool, reason string) (*types.Transaction, error) {
	return _Pv0Smartcontract.Contract.SetPolicyBlacklist(&_Pv0Smartcontract.TransactOpts, policyID, blacklisted, reason)
}

// UnVote is a paid mutator transaction binding the contract method 0x8cbef3f1.
//
// Solidity: function unVote(string policyID) returns()
func (_Pv0Smartcontract *Pv0SmartcontractTransactor) UnVote(opts *bind.TransactOpts, policyID string) (*types.Transaction, error) {
	return _Pv0Smartcontract.contract.Transact(opts, "unVote", policyID)
}

// UnVote is a paid mutator transaction binding the contract method 0x8cbef3f1.
//
// Solidity: function unVote(string policyID) returns()
func (_Pv0Smartcontract *Pv0SmartcontractSession) UnVote(policyID string) (*types.Transaction, error) {
	return _Pv0Smartcontract.Contract.UnVote(&_Pv0Smartcontract.TransactOpts, policyID)
}

// UnVote is a paid mutator transaction binding the contract method 0x8cbef3f1.
//
// Solidity: function unVote(string policyID) returns()
func (_Pv0Smartcontract *Pv0SmartcontractTransactorSession) UnVote(policyID string) (*types.Transaction, error) {
	return _Pv0Smartcontract.Contract.UnVote(&_Pv0Smartcontract.TransactOpts, policyID)
}

// Vote is a paid mutator transaction binding the contract method 0xfc36e15b.
//
// Solidity: function vote(string policyID) returns()
func (_Pv0Smartcontract *Pv0SmartcontractTransactor) Vote(opts *bind.TransactOpts, policyID string) (*types.Transaction, error) {
	return _Pv0Smartcontract.contract.Transact(opts, "vote", policyID)
}

// Vote is a paid mutator transaction binding the contract method 0xfc36e15b.
//
// Solidity: function vote(string policyID) returns()
func (_Pv0Smartcontract *Pv0SmartcontractSession) Vote(policyID string) (*types.Transaction, error) {
	return _Pv0Smartcontract.Contract.Vote(&_Pv0Smartcontract.TransactOpts, policyID)
}

// Vote is a paid mutator transaction binding the contract method 0xfc36e15b.
//
// Solidity: function vote(string policyID) returns()
func (_Pv0Smartcontract *Pv0SmartcontractTransactorSession) Vote(policyID string) (*types.Transaction, error) {
	return _Pv0Smartcontract.Contract.Vote(&_Pv0Smartcontract.TransactOpts, policyID)
}

// Pv0SmartcontractBlacklistedIterator is returned from FilterBlacklisted and is used to iterate over the raw logs and unpacked data for Blacklisted events raised by the Pv0Smartcontract contract.
type Pv0SmartcontractBlacklistedIterator struct {
	Event *Pv0SmartcontractBlacklisted // Event containing the contract specifics and raw log

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
func (it *Pv0SmartcontractBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pv0SmartcontractBlacklisted)
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
		it.Event = new(Pv0SmartcontractBlacklisted)
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
func (it *Pv0SmartcontractBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pv0SmartcontractBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pv0SmartcontractBlacklisted represents a Blacklisted event raised by the Pv0Smartcontract contract.
type Pv0SmartcontractBlacklisted struct {
	User        common.Address
	Blacklisted bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlacklisted is a free log retrieval operation binding the contract event 0xcf3473b85df1594d47b6958f29a32bea0abff9dd68296f7bf33443646793cfd8.
//
// Solidity: event Blacklisted(address user, bool blacklisted)
func (_Pv0Smartcontract *Pv0SmartcontractFilterer) FilterBlacklisted(opts *bind.FilterOpts) (*Pv0SmartcontractBlacklistedIterator, error) {

	logs, sub, err := _Pv0Smartcontract.contract.FilterLogs(opts, "Blacklisted")
	if err != nil {
		return nil, err
	}
	return &Pv0SmartcontractBlacklistedIterator{contract: _Pv0Smartcontract.contract, event: "Blacklisted", logs: logs, sub: sub}, nil
}

// WatchBlacklisted is a free log subscription operation binding the contract event 0xcf3473b85df1594d47b6958f29a32bea0abff9dd68296f7bf33443646793cfd8.
//
// Solidity: event Blacklisted(address user, bool blacklisted)
func (_Pv0Smartcontract *Pv0SmartcontractFilterer) WatchBlacklisted(opts *bind.WatchOpts, sink chan<- *Pv0SmartcontractBlacklisted) (event.Subscription, error) {

	logs, sub, err := _Pv0Smartcontract.contract.WatchLogs(opts, "Blacklisted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pv0SmartcontractBlacklisted)
				if err := _Pv0Smartcontract.contract.UnpackLog(event, "Blacklisted", log); err != nil {
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
func (_Pv0Smartcontract *Pv0SmartcontractFilterer) ParseBlacklisted(log types.Log) (*Pv0SmartcontractBlacklisted, error) {
	event := new(Pv0SmartcontractBlacklisted)
	if err := _Pv0Smartcontract.contract.UnpackLog(event, "Blacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Pv0SmartcontractNewPolicyIterator is returned from FilterNewPolicy and is used to iterate over the raw logs and unpacked data for NewPolicy events raised by the Pv0Smartcontract contract.
type Pv0SmartcontractNewPolicyIterator struct {
	Event *Pv0SmartcontractNewPolicy // Event containing the contract specifics and raw log

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
func (it *Pv0SmartcontractNewPolicyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pv0SmartcontractNewPolicy)
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
		it.Event = new(Pv0SmartcontractNewPolicy)
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
func (it *Pv0SmartcontractNewPolicyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pv0SmartcontractNewPolicyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pv0SmartcontractNewPolicy represents a NewPolicy event raised by the Pv0Smartcontract contract.
type Pv0SmartcontractNewPolicy struct {
	PolicyID string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewPolicy is a free log retrieval operation binding the contract event 0x95cc192aad8600ab35c3acdf1de31ac530328faf0e80b4389d3809be48ff68a1.
//
// Solidity: event NewPolicy(string policyID)
func (_Pv0Smartcontract *Pv0SmartcontractFilterer) FilterNewPolicy(opts *bind.FilterOpts) (*Pv0SmartcontractNewPolicyIterator, error) {

	logs, sub, err := _Pv0Smartcontract.contract.FilterLogs(opts, "NewPolicy")
	if err != nil {
		return nil, err
	}
	return &Pv0SmartcontractNewPolicyIterator{contract: _Pv0Smartcontract.contract, event: "NewPolicy", logs: logs, sub: sub}, nil
}

// WatchNewPolicy is a free log subscription operation binding the contract event 0x95cc192aad8600ab35c3acdf1de31ac530328faf0e80b4389d3809be48ff68a1.
//
// Solidity: event NewPolicy(string policyID)
func (_Pv0Smartcontract *Pv0SmartcontractFilterer) WatchNewPolicy(opts *bind.WatchOpts, sink chan<- *Pv0SmartcontractNewPolicy) (event.Subscription, error) {

	logs, sub, err := _Pv0Smartcontract.contract.WatchLogs(opts, "NewPolicy")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pv0SmartcontractNewPolicy)
				if err := _Pv0Smartcontract.contract.UnpackLog(event, "NewPolicy", log); err != nil {
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

// ParseNewPolicy is a log parse operation binding the contract event 0x95cc192aad8600ab35c3acdf1de31ac530328faf0e80b4389d3809be48ff68a1.
//
// Solidity: event NewPolicy(string policyID)
func (_Pv0Smartcontract *Pv0SmartcontractFilterer) ParseNewPolicy(log types.Log) (*Pv0SmartcontractNewPolicy, error) {
	event := new(Pv0SmartcontractNewPolicy)
	if err := _Pv0Smartcontract.contract.UnpackLog(event, "NewPolicy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Pv0SmartcontractPolicyBlacklistedIterator is returned from FilterPolicyBlacklisted and is used to iterate over the raw logs and unpacked data for PolicyBlacklisted events raised by the Pv0Smartcontract contract.
type Pv0SmartcontractPolicyBlacklistedIterator struct {
	Event *Pv0SmartcontractPolicyBlacklisted // Event containing the contract specifics and raw log

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
func (it *Pv0SmartcontractPolicyBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pv0SmartcontractPolicyBlacklisted)
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
		it.Event = new(Pv0SmartcontractPolicyBlacklisted)
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
func (it *Pv0SmartcontractPolicyBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pv0SmartcontractPolicyBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pv0SmartcontractPolicyBlacklisted represents a PolicyBlacklisted event raised by the Pv0Smartcontract contract.
type Pv0SmartcontractPolicyBlacklisted struct {
	PolicyID    string
	Blacklisted bool
	Reason      string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPolicyBlacklisted is a free log retrieval operation binding the contract event 0xe03285579bee722c625a02fb2b45070e578d9a320679123457906a0e90c9e6d3.
//
// Solidity: event PolicyBlacklisted(string policyID, bool blacklisted, string reason)
func (_Pv0Smartcontract *Pv0SmartcontractFilterer) FilterPolicyBlacklisted(opts *bind.FilterOpts) (*Pv0SmartcontractPolicyBlacklistedIterator, error) {

	logs, sub, err := _Pv0Smartcontract.contract.FilterLogs(opts, "PolicyBlacklisted")
	if err != nil {
		return nil, err
	}
	return &Pv0SmartcontractPolicyBlacklistedIterator{contract: _Pv0Smartcontract.contract, event: "PolicyBlacklisted", logs: logs, sub: sub}, nil
}

// WatchPolicyBlacklisted is a free log subscription operation binding the contract event 0xe03285579bee722c625a02fb2b45070e578d9a320679123457906a0e90c9e6d3.
//
// Solidity: event PolicyBlacklisted(string policyID, bool blacklisted, string reason)
func (_Pv0Smartcontract *Pv0SmartcontractFilterer) WatchPolicyBlacklisted(opts *bind.WatchOpts, sink chan<- *Pv0SmartcontractPolicyBlacklisted) (event.Subscription, error) {

	logs, sub, err := _Pv0Smartcontract.contract.WatchLogs(opts, "PolicyBlacklisted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pv0SmartcontractPolicyBlacklisted)
				if err := _Pv0Smartcontract.contract.UnpackLog(event, "PolicyBlacklisted", log); err != nil {
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
func (_Pv0Smartcontract *Pv0SmartcontractFilterer) ParsePolicyBlacklisted(log types.Log) (*Pv0SmartcontractPolicyBlacklisted, error) {
	event := new(Pv0SmartcontractPolicyBlacklisted)
	if err := _Pv0Smartcontract.contract.UnpackLog(event, "PolicyBlacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Pv0SmartcontractRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Pv0Smartcontract contract.
type Pv0SmartcontractRoleAdminChangedIterator struct {
	Event *Pv0SmartcontractRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *Pv0SmartcontractRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pv0SmartcontractRoleAdminChanged)
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
		it.Event = new(Pv0SmartcontractRoleAdminChanged)
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
func (it *Pv0SmartcontractRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pv0SmartcontractRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pv0SmartcontractRoleAdminChanged represents a RoleAdminChanged event raised by the Pv0Smartcontract contract.
type Pv0SmartcontractRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Pv0Smartcontract *Pv0SmartcontractFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*Pv0SmartcontractRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Pv0Smartcontract.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &Pv0SmartcontractRoleAdminChangedIterator{contract: _Pv0Smartcontract.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Pv0Smartcontract *Pv0SmartcontractFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *Pv0SmartcontractRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Pv0Smartcontract.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pv0SmartcontractRoleAdminChanged)
				if err := _Pv0Smartcontract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Pv0Smartcontract *Pv0SmartcontractFilterer) ParseRoleAdminChanged(log types.Log) (*Pv0SmartcontractRoleAdminChanged, error) {
	event := new(Pv0SmartcontractRoleAdminChanged)
	if err := _Pv0Smartcontract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Pv0SmartcontractRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Pv0Smartcontract contract.
type Pv0SmartcontractRoleGrantedIterator struct {
	Event *Pv0SmartcontractRoleGranted // Event containing the contract specifics and raw log

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
func (it *Pv0SmartcontractRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pv0SmartcontractRoleGranted)
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
		it.Event = new(Pv0SmartcontractRoleGranted)
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
func (it *Pv0SmartcontractRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pv0SmartcontractRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pv0SmartcontractRoleGranted represents a RoleGranted event raised by the Pv0Smartcontract contract.
type Pv0SmartcontractRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pv0Smartcontract *Pv0SmartcontractFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*Pv0SmartcontractRoleGrantedIterator, error) {

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

	logs, sub, err := _Pv0Smartcontract.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &Pv0SmartcontractRoleGrantedIterator{contract: _Pv0Smartcontract.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pv0Smartcontract *Pv0SmartcontractFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *Pv0SmartcontractRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Pv0Smartcontract.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pv0SmartcontractRoleGranted)
				if err := _Pv0Smartcontract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Pv0Smartcontract *Pv0SmartcontractFilterer) ParseRoleGranted(log types.Log) (*Pv0SmartcontractRoleGranted, error) {
	event := new(Pv0SmartcontractRoleGranted)
	if err := _Pv0Smartcontract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Pv0SmartcontractRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Pv0Smartcontract contract.
type Pv0SmartcontractRoleRevokedIterator struct {
	Event *Pv0SmartcontractRoleRevoked // Event containing the contract specifics and raw log

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
func (it *Pv0SmartcontractRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pv0SmartcontractRoleRevoked)
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
		it.Event = new(Pv0SmartcontractRoleRevoked)
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
func (it *Pv0SmartcontractRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pv0SmartcontractRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pv0SmartcontractRoleRevoked represents a RoleRevoked event raised by the Pv0Smartcontract contract.
type Pv0SmartcontractRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pv0Smartcontract *Pv0SmartcontractFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*Pv0SmartcontractRoleRevokedIterator, error) {

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

	logs, sub, err := _Pv0Smartcontract.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &Pv0SmartcontractRoleRevokedIterator{contract: _Pv0Smartcontract.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pv0Smartcontract *Pv0SmartcontractFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *Pv0SmartcontractRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Pv0Smartcontract.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pv0SmartcontractRoleRevoked)
				if err := _Pv0Smartcontract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Pv0Smartcontract *Pv0SmartcontractFilterer) ParseRoleRevoked(log types.Log) (*Pv0SmartcontractRoleRevoked, error) {
	event := new(Pv0SmartcontractRoleRevoked)
	if err := _Pv0Smartcontract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Pv0SmartcontractUnVotedIterator is returned from FilterUnVoted and is used to iterate over the raw logs and unpacked data for UnVoted events raised by the Pv0Smartcontract contract.
type Pv0SmartcontractUnVotedIterator struct {
	Event *Pv0SmartcontractUnVoted // Event containing the contract specifics and raw log

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
func (it *Pv0SmartcontractUnVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pv0SmartcontractUnVoted)
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
		it.Event = new(Pv0SmartcontractUnVoted)
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
func (it *Pv0SmartcontractUnVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pv0SmartcontractUnVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pv0SmartcontractUnVoted represents a UnVoted event raised by the Pv0Smartcontract contract.
type Pv0SmartcontractUnVoted struct {
	PolicyID   string
	TotalVotes *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUnVoted is a free log retrieval operation binding the contract event 0x9f2a13e0b25fbc63d7ba5806ace0b766e1d5d3eb8df54062d9e5171be8a37fb9.
//
// Solidity: event UnVoted(string policyID, uint256 totalVotes)
func (_Pv0Smartcontract *Pv0SmartcontractFilterer) FilterUnVoted(opts *bind.FilterOpts) (*Pv0SmartcontractUnVotedIterator, error) {

	logs, sub, err := _Pv0Smartcontract.contract.FilterLogs(opts, "UnVoted")
	if err != nil {
		return nil, err
	}
	return &Pv0SmartcontractUnVotedIterator{contract: _Pv0Smartcontract.contract, event: "UnVoted", logs: logs, sub: sub}, nil
}

// WatchUnVoted is a free log subscription operation binding the contract event 0x9f2a13e0b25fbc63d7ba5806ace0b766e1d5d3eb8df54062d9e5171be8a37fb9.
//
// Solidity: event UnVoted(string policyID, uint256 totalVotes)
func (_Pv0Smartcontract *Pv0SmartcontractFilterer) WatchUnVoted(opts *bind.WatchOpts, sink chan<- *Pv0SmartcontractUnVoted) (event.Subscription, error) {

	logs, sub, err := _Pv0Smartcontract.contract.WatchLogs(opts, "UnVoted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pv0SmartcontractUnVoted)
				if err := _Pv0Smartcontract.contract.UnpackLog(event, "UnVoted", log); err != nil {
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

// ParseUnVoted is a log parse operation binding the contract event 0x9f2a13e0b25fbc63d7ba5806ace0b766e1d5d3eb8df54062d9e5171be8a37fb9.
//
// Solidity: event UnVoted(string policyID, uint256 totalVotes)
func (_Pv0Smartcontract *Pv0SmartcontractFilterer) ParseUnVoted(log types.Log) (*Pv0SmartcontractUnVoted, error) {
	event := new(Pv0SmartcontractUnVoted)
	if err := _Pv0Smartcontract.contract.UnpackLog(event, "UnVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Pv0SmartcontractVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the Pv0Smartcontract contract.
type Pv0SmartcontractVotedIterator struct {
	Event *Pv0SmartcontractVoted // Event containing the contract specifics and raw log

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
func (it *Pv0SmartcontractVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pv0SmartcontractVoted)
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
		it.Event = new(Pv0SmartcontractVoted)
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
func (it *Pv0SmartcontractVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pv0SmartcontractVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pv0SmartcontractVoted represents a Voted event raised by the Pv0Smartcontract contract.
type Pv0SmartcontractVoted struct {
	PolicyID   string
	TotalVotes *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x82ea01a4479f27996bbe8de24f4c8d962035fe2f32a2b130c8c9ca593ab91aa8.
//
// Solidity: event Voted(string policyID, uint256 totalVotes)
func (_Pv0Smartcontract *Pv0SmartcontractFilterer) FilterVoted(opts *bind.FilterOpts) (*Pv0SmartcontractVotedIterator, error) {

	logs, sub, err := _Pv0Smartcontract.contract.FilterLogs(opts, "Voted")
	if err != nil {
		return nil, err
	}
	return &Pv0SmartcontractVotedIterator{contract: _Pv0Smartcontract.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x82ea01a4479f27996bbe8de24f4c8d962035fe2f32a2b130c8c9ca593ab91aa8.
//
// Solidity: event Voted(string policyID, uint256 totalVotes)
func (_Pv0Smartcontract *Pv0SmartcontractFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *Pv0SmartcontractVoted) (event.Subscription, error) {

	logs, sub, err := _Pv0Smartcontract.contract.WatchLogs(opts, "Voted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pv0SmartcontractVoted)
				if err := _Pv0Smartcontract.contract.UnpackLog(event, "Voted", log); err != nil {
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

// ParseVoted is a log parse operation binding the contract event 0x82ea01a4479f27996bbe8de24f4c8d962035fe2f32a2b130c8c9ca593ab91aa8.
//
// Solidity: event Voted(string policyID, uint256 totalVotes)
func (_Pv0Smartcontract *Pv0SmartcontractFilterer) ParseVoted(log types.Log) (*Pv0SmartcontractVoted, error) {
	event := new(Pv0SmartcontractVoted)
	if err := _Pv0Smartcontract.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
