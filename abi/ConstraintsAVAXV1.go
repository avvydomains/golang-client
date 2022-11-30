// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ConstraintsAVAXV1

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

// ConstraintsAVAXV1MetaData contains all meta data concerning the ConstraintsAVAXV1 contract.
var ConstraintsAVAXV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractVerifierInterface\",\"name\":\"verifier\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"names\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"blocked\",\"type\":\"bool\"}],\"name\":\"NamesBlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractVerifierInterface\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"VerifierSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_verifier\",\"outputs\":[{\"internalType\":\"contractVerifierInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"names\",\"type\":\"uint256[]\"}],\"name\":\"blockNames\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"namespace\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"name\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"check\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"name\",\"type\":\"uint256\"}],\"name\":\"isNameBlocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractVerifierInterface\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"setVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"names\",\"type\":\"uint256[]\"}],\"name\":\"unblockNames\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ConstraintsAVAXV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use ConstraintsAVAXV1MetaData.ABI instead.
var ConstraintsAVAXV1ABI = ConstraintsAVAXV1MetaData.ABI

// ConstraintsAVAXV1 is an auto generated Go binding around an Ethereum contract.
type ConstraintsAVAXV1 struct {
	ConstraintsAVAXV1Caller     // Read-only binding to the contract
	ConstraintsAVAXV1Transactor // Write-only binding to the contract
	ConstraintsAVAXV1Filterer   // Log filterer for contract events
}

// ConstraintsAVAXV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type ConstraintsAVAXV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstraintsAVAXV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ConstraintsAVAXV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstraintsAVAXV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConstraintsAVAXV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstraintsAVAXV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConstraintsAVAXV1Session struct {
	Contract     *ConstraintsAVAXV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ConstraintsAVAXV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConstraintsAVAXV1CallerSession struct {
	Contract *ConstraintsAVAXV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ConstraintsAVAXV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConstraintsAVAXV1TransactorSession struct {
	Contract     *ConstraintsAVAXV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ConstraintsAVAXV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type ConstraintsAVAXV1Raw struct {
	Contract *ConstraintsAVAXV1 // Generic contract binding to access the raw methods on
}

// ConstraintsAVAXV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConstraintsAVAXV1CallerRaw struct {
	Contract *ConstraintsAVAXV1Caller // Generic read-only contract binding to access the raw methods on
}

// ConstraintsAVAXV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConstraintsAVAXV1TransactorRaw struct {
	Contract *ConstraintsAVAXV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewConstraintsAVAXV1 creates a new instance of ConstraintsAVAXV1, bound to a specific deployed contract.
func NewConstraintsAVAXV1(address common.Address, backend bind.ContractBackend) (*ConstraintsAVAXV1, error) {
	contract, err := bindConstraintsAVAXV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConstraintsAVAXV1{ConstraintsAVAXV1Caller: ConstraintsAVAXV1Caller{contract: contract}, ConstraintsAVAXV1Transactor: ConstraintsAVAXV1Transactor{contract: contract}, ConstraintsAVAXV1Filterer: ConstraintsAVAXV1Filterer{contract: contract}}, nil
}

// NewConstraintsAVAXV1Caller creates a new read-only instance of ConstraintsAVAXV1, bound to a specific deployed contract.
func NewConstraintsAVAXV1Caller(address common.Address, caller bind.ContractCaller) (*ConstraintsAVAXV1Caller, error) {
	contract, err := bindConstraintsAVAXV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConstraintsAVAXV1Caller{contract: contract}, nil
}

// NewConstraintsAVAXV1Transactor creates a new write-only instance of ConstraintsAVAXV1, bound to a specific deployed contract.
func NewConstraintsAVAXV1Transactor(address common.Address, transactor bind.ContractTransactor) (*ConstraintsAVAXV1Transactor, error) {
	contract, err := bindConstraintsAVAXV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConstraintsAVAXV1Transactor{contract: contract}, nil
}

// NewConstraintsAVAXV1Filterer creates a new log filterer instance of ConstraintsAVAXV1, bound to a specific deployed contract.
func NewConstraintsAVAXV1Filterer(address common.Address, filterer bind.ContractFilterer) (*ConstraintsAVAXV1Filterer, error) {
	contract, err := bindConstraintsAVAXV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConstraintsAVAXV1Filterer{contract: contract}, nil
}

// bindConstraintsAVAXV1 binds a generic wrapper to an already deployed contract.
func bindConstraintsAVAXV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ConstraintsAVAXV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConstraintsAVAXV1.Contract.ConstraintsAVAXV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConstraintsAVAXV1.Contract.ConstraintsAVAXV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConstraintsAVAXV1.Contract.ConstraintsAVAXV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConstraintsAVAXV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConstraintsAVAXV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConstraintsAVAXV1.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ConstraintsAVAXV1.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _ConstraintsAVAXV1.Contract.DEFAULTADMINROLE(&_ConstraintsAVAXV1.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ConstraintsAVAXV1.Contract.DEFAULTADMINROLE(&_ConstraintsAVAXV1.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Caller) MANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ConstraintsAVAXV1.contract.Call(opts, &out, "MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Session) MANAGERROLE() ([32]byte, error) {
	return _ConstraintsAVAXV1.Contract.MANAGERROLE(&_ConstraintsAVAXV1.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1CallerSession) MANAGERROLE() ([32]byte, error) {
	return _ConstraintsAVAXV1.Contract.MANAGERROLE(&_ConstraintsAVAXV1.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x8a172651.
//
// Solidity: function _verifier() view returns(address)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Caller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConstraintsAVAXV1.contract.Call(opts, &out, "_verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x8a172651.
//
// Solidity: function _verifier() view returns(address)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Session) Verifier() (common.Address, error) {
	return _ConstraintsAVAXV1.Contract.Verifier(&_ConstraintsAVAXV1.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x8a172651.
//
// Solidity: function _verifier() view returns(address)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1CallerSession) Verifier() (common.Address, error) {
	return _ConstraintsAVAXV1.Contract.Verifier(&_ConstraintsAVAXV1.CallOpts)
}

// Check is a free data retrieval call binding the contract method 0x7b36cdb4.
//
// Solidity: function check(uint256 namespace, uint256 name, bytes data) view returns()
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Caller) Check(opts *bind.CallOpts, namespace *big.Int, name *big.Int, data []byte) error {
	var out []interface{}
	err := _ConstraintsAVAXV1.contract.Call(opts, &out, "check", namespace, name, data)

	if err != nil {
		return err
	}

	return err

}

// Check is a free data retrieval call binding the contract method 0x7b36cdb4.
//
// Solidity: function check(uint256 namespace, uint256 name, bytes data) view returns()
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Session) Check(namespace *big.Int, name *big.Int, data []byte) error {
	return _ConstraintsAVAXV1.Contract.Check(&_ConstraintsAVAXV1.CallOpts, namespace, name, data)
}

// Check is a free data retrieval call binding the contract method 0x7b36cdb4.
//
// Solidity: function check(uint256 namespace, uint256 name, bytes data) view returns()
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1CallerSession) Check(namespace *big.Int, name *big.Int, data []byte) error {
	return _ConstraintsAVAXV1.Contract.Check(&_ConstraintsAVAXV1.CallOpts, namespace, name, data)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ConstraintsAVAXV1.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ConstraintsAVAXV1.Contract.GetRoleAdmin(&_ConstraintsAVAXV1.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ConstraintsAVAXV1.Contract.GetRoleAdmin(&_ConstraintsAVAXV1.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ConstraintsAVAXV1.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ConstraintsAVAXV1.Contract.HasRole(&_ConstraintsAVAXV1.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ConstraintsAVAXV1.Contract.HasRole(&_ConstraintsAVAXV1.CallOpts, role, account)
}

// IsNameBlocked is a free data retrieval call binding the contract method 0xb0b0d57b.
//
// Solidity: function isNameBlocked(uint256 name) view returns(bool)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Caller) IsNameBlocked(opts *bind.CallOpts, name *big.Int) (bool, error) {
	var out []interface{}
	err := _ConstraintsAVAXV1.contract.Call(opts, &out, "isNameBlocked", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNameBlocked is a free data retrieval call binding the contract method 0xb0b0d57b.
//
// Solidity: function isNameBlocked(uint256 name) view returns(bool)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Session) IsNameBlocked(name *big.Int) (bool, error) {
	return _ConstraintsAVAXV1.Contract.IsNameBlocked(&_ConstraintsAVAXV1.CallOpts, name)
}

// IsNameBlocked is a free data retrieval call binding the contract method 0xb0b0d57b.
//
// Solidity: function isNameBlocked(uint256 name) view returns(bool)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1CallerSession) IsNameBlocked(name *big.Int) (bool, error) {
	return _ConstraintsAVAXV1.Contract.IsNameBlocked(&_ConstraintsAVAXV1.CallOpts, name)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ConstraintsAVAXV1.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ConstraintsAVAXV1.Contract.SupportsInterface(&_ConstraintsAVAXV1.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ConstraintsAVAXV1.Contract.SupportsInterface(&_ConstraintsAVAXV1.CallOpts, interfaceId)
}

// BlockNames is a paid mutator transaction binding the contract method 0xf19924af.
//
// Solidity: function blockNames(uint256[] names) returns()
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Transactor) BlockNames(opts *bind.TransactOpts, names []*big.Int) (*types.Transaction, error) {
	return _ConstraintsAVAXV1.contract.Transact(opts, "blockNames", names)
}

// BlockNames is a paid mutator transaction binding the contract method 0xf19924af.
//
// Solidity: function blockNames(uint256[] names) returns()
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Session) BlockNames(names []*big.Int) (*types.Transaction, error) {
	return _ConstraintsAVAXV1.Contract.BlockNames(&_ConstraintsAVAXV1.TransactOpts, names)
}

// BlockNames is a paid mutator transaction binding the contract method 0xf19924af.
//
// Solidity: function blockNames(uint256[] names) returns()
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1TransactorSession) BlockNames(names []*big.Int) (*types.Transaction, error) {
	return _ConstraintsAVAXV1.Contract.BlockNames(&_ConstraintsAVAXV1.TransactOpts, names)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ConstraintsAVAXV1.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ConstraintsAVAXV1.Contract.GrantRole(&_ConstraintsAVAXV1.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ConstraintsAVAXV1.Contract.GrantRole(&_ConstraintsAVAXV1.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ConstraintsAVAXV1.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Session) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ConstraintsAVAXV1.Contract.RenounceRole(&_ConstraintsAVAXV1.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1TransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ConstraintsAVAXV1.Contract.RenounceRole(&_ConstraintsAVAXV1.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ConstraintsAVAXV1.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ConstraintsAVAXV1.Contract.RevokeRole(&_ConstraintsAVAXV1.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ConstraintsAVAXV1.Contract.RevokeRole(&_ConstraintsAVAXV1.TransactOpts, role, account)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address verifier) returns()
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Transactor) SetVerifier(opts *bind.TransactOpts, verifier common.Address) (*types.Transaction, error) {
	return _ConstraintsAVAXV1.contract.Transact(opts, "setVerifier", verifier)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address verifier) returns()
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Session) SetVerifier(verifier common.Address) (*types.Transaction, error) {
	return _ConstraintsAVAXV1.Contract.SetVerifier(&_ConstraintsAVAXV1.TransactOpts, verifier)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address verifier) returns()
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1TransactorSession) SetVerifier(verifier common.Address) (*types.Transaction, error) {
	return _ConstraintsAVAXV1.Contract.SetVerifier(&_ConstraintsAVAXV1.TransactOpts, verifier)
}

// UnblockNames is a paid mutator transaction binding the contract method 0xf231c653.
//
// Solidity: function unblockNames(uint256[] names) returns()
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Transactor) UnblockNames(opts *bind.TransactOpts, names []*big.Int) (*types.Transaction, error) {
	return _ConstraintsAVAXV1.contract.Transact(opts, "unblockNames", names)
}

// UnblockNames is a paid mutator transaction binding the contract method 0xf231c653.
//
// Solidity: function unblockNames(uint256[] names) returns()
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Session) UnblockNames(names []*big.Int) (*types.Transaction, error) {
	return _ConstraintsAVAXV1.Contract.UnblockNames(&_ConstraintsAVAXV1.TransactOpts, names)
}

// UnblockNames is a paid mutator transaction binding the contract method 0xf231c653.
//
// Solidity: function unblockNames(uint256[] names) returns()
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1TransactorSession) UnblockNames(names []*big.Int) (*types.Transaction, error) {
	return _ConstraintsAVAXV1.Contract.UnblockNames(&_ConstraintsAVAXV1.TransactOpts, names)
}

// ConstraintsAVAXV1NamesBlockedIterator is returned from FilterNamesBlocked and is used to iterate over the raw logs and unpacked data for NamesBlocked events raised by the ConstraintsAVAXV1 contract.
type ConstraintsAVAXV1NamesBlockedIterator struct {
	Event *ConstraintsAVAXV1NamesBlocked // Event containing the contract specifics and raw log

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
func (it *ConstraintsAVAXV1NamesBlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConstraintsAVAXV1NamesBlocked)
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
		it.Event = new(ConstraintsAVAXV1NamesBlocked)
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
func (it *ConstraintsAVAXV1NamesBlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConstraintsAVAXV1NamesBlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConstraintsAVAXV1NamesBlocked represents a NamesBlocked event raised by the ConstraintsAVAXV1 contract.
type ConstraintsAVAXV1NamesBlocked struct {
	Names   []*big.Int
	Blocked bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNamesBlocked is a free log retrieval operation binding the contract event 0x51ae321260e9ad4c2fe9c1efdb6dedf346f800bb828b5ebba5ed77b7dafda2f3.
//
// Solidity: event NamesBlocked(uint256[] names, bool blocked)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Filterer) FilterNamesBlocked(opts *bind.FilterOpts) (*ConstraintsAVAXV1NamesBlockedIterator, error) {

	logs, sub, err := _ConstraintsAVAXV1.contract.FilterLogs(opts, "NamesBlocked")
	if err != nil {
		return nil, err
	}
	return &ConstraintsAVAXV1NamesBlockedIterator{contract: _ConstraintsAVAXV1.contract, event: "NamesBlocked", logs: logs, sub: sub}, nil
}

// WatchNamesBlocked is a free log subscription operation binding the contract event 0x51ae321260e9ad4c2fe9c1efdb6dedf346f800bb828b5ebba5ed77b7dafda2f3.
//
// Solidity: event NamesBlocked(uint256[] names, bool blocked)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Filterer) WatchNamesBlocked(opts *bind.WatchOpts, sink chan<- *ConstraintsAVAXV1NamesBlocked) (event.Subscription, error) {

	logs, sub, err := _ConstraintsAVAXV1.contract.WatchLogs(opts, "NamesBlocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConstraintsAVAXV1NamesBlocked)
				if err := _ConstraintsAVAXV1.contract.UnpackLog(event, "NamesBlocked", log); err != nil {
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

// ParseNamesBlocked is a log parse operation binding the contract event 0x51ae321260e9ad4c2fe9c1efdb6dedf346f800bb828b5ebba5ed77b7dafda2f3.
//
// Solidity: event NamesBlocked(uint256[] names, bool blocked)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Filterer) ParseNamesBlocked(log types.Log) (*ConstraintsAVAXV1NamesBlocked, error) {
	event := new(ConstraintsAVAXV1NamesBlocked)
	if err := _ConstraintsAVAXV1.contract.UnpackLog(event, "NamesBlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConstraintsAVAXV1RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ConstraintsAVAXV1 contract.
type ConstraintsAVAXV1RoleAdminChangedIterator struct {
	Event *ConstraintsAVAXV1RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ConstraintsAVAXV1RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConstraintsAVAXV1RoleAdminChanged)
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
		it.Event = new(ConstraintsAVAXV1RoleAdminChanged)
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
func (it *ConstraintsAVAXV1RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConstraintsAVAXV1RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConstraintsAVAXV1RoleAdminChanged represents a RoleAdminChanged event raised by the ConstraintsAVAXV1 contract.
type ConstraintsAVAXV1RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ConstraintsAVAXV1RoleAdminChangedIterator, error) {

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

	logs, sub, err := _ConstraintsAVAXV1.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ConstraintsAVAXV1RoleAdminChangedIterator{contract: _ConstraintsAVAXV1.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ConstraintsAVAXV1RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ConstraintsAVAXV1.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConstraintsAVAXV1RoleAdminChanged)
				if err := _ConstraintsAVAXV1.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Filterer) ParseRoleAdminChanged(log types.Log) (*ConstraintsAVAXV1RoleAdminChanged, error) {
	event := new(ConstraintsAVAXV1RoleAdminChanged)
	if err := _ConstraintsAVAXV1.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConstraintsAVAXV1RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ConstraintsAVAXV1 contract.
type ConstraintsAVAXV1RoleGrantedIterator struct {
	Event *ConstraintsAVAXV1RoleGranted // Event containing the contract specifics and raw log

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
func (it *ConstraintsAVAXV1RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConstraintsAVAXV1RoleGranted)
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
		it.Event = new(ConstraintsAVAXV1RoleGranted)
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
func (it *ConstraintsAVAXV1RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConstraintsAVAXV1RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConstraintsAVAXV1RoleGranted represents a RoleGranted event raised by the ConstraintsAVAXV1 contract.
type ConstraintsAVAXV1RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ConstraintsAVAXV1RoleGrantedIterator, error) {

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

	logs, sub, err := _ConstraintsAVAXV1.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ConstraintsAVAXV1RoleGrantedIterator{contract: _ConstraintsAVAXV1.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ConstraintsAVAXV1RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ConstraintsAVAXV1.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConstraintsAVAXV1RoleGranted)
				if err := _ConstraintsAVAXV1.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Filterer) ParseRoleGranted(log types.Log) (*ConstraintsAVAXV1RoleGranted, error) {
	event := new(ConstraintsAVAXV1RoleGranted)
	if err := _ConstraintsAVAXV1.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConstraintsAVAXV1RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ConstraintsAVAXV1 contract.
type ConstraintsAVAXV1RoleRevokedIterator struct {
	Event *ConstraintsAVAXV1RoleRevoked // Event containing the contract specifics and raw log

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
func (it *ConstraintsAVAXV1RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConstraintsAVAXV1RoleRevoked)
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
		it.Event = new(ConstraintsAVAXV1RoleRevoked)
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
func (it *ConstraintsAVAXV1RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConstraintsAVAXV1RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConstraintsAVAXV1RoleRevoked represents a RoleRevoked event raised by the ConstraintsAVAXV1 contract.
type ConstraintsAVAXV1RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ConstraintsAVAXV1RoleRevokedIterator, error) {

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

	logs, sub, err := _ConstraintsAVAXV1.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ConstraintsAVAXV1RoleRevokedIterator{contract: _ConstraintsAVAXV1.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ConstraintsAVAXV1RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ConstraintsAVAXV1.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConstraintsAVAXV1RoleRevoked)
				if err := _ConstraintsAVAXV1.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Filterer) ParseRoleRevoked(log types.Log) (*ConstraintsAVAXV1RoleRevoked, error) {
	event := new(ConstraintsAVAXV1RoleRevoked)
	if err := _ConstraintsAVAXV1.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConstraintsAVAXV1VerifierSetIterator is returned from FilterVerifierSet and is used to iterate over the raw logs and unpacked data for VerifierSet events raised by the ConstraintsAVAXV1 contract.
type ConstraintsAVAXV1VerifierSetIterator struct {
	Event *ConstraintsAVAXV1VerifierSet // Event containing the contract specifics and raw log

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
func (it *ConstraintsAVAXV1VerifierSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConstraintsAVAXV1VerifierSet)
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
		it.Event = new(ConstraintsAVAXV1VerifierSet)
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
func (it *ConstraintsAVAXV1VerifierSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConstraintsAVAXV1VerifierSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConstraintsAVAXV1VerifierSet represents a VerifierSet event raised by the ConstraintsAVAXV1 contract.
type ConstraintsAVAXV1VerifierSet struct {
	Verifier common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVerifierSet is a free log retrieval operation binding the contract event 0x480b37e3d134e44cb444c9703493c7db564c707cb8a18cecea165f340431da1f.
//
// Solidity: event VerifierSet(address verifier)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Filterer) FilterVerifierSet(opts *bind.FilterOpts) (*ConstraintsAVAXV1VerifierSetIterator, error) {

	logs, sub, err := _ConstraintsAVAXV1.contract.FilterLogs(opts, "VerifierSet")
	if err != nil {
		return nil, err
	}
	return &ConstraintsAVAXV1VerifierSetIterator{contract: _ConstraintsAVAXV1.contract, event: "VerifierSet", logs: logs, sub: sub}, nil
}

// WatchVerifierSet is a free log subscription operation binding the contract event 0x480b37e3d134e44cb444c9703493c7db564c707cb8a18cecea165f340431da1f.
//
// Solidity: event VerifierSet(address verifier)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Filterer) WatchVerifierSet(opts *bind.WatchOpts, sink chan<- *ConstraintsAVAXV1VerifierSet) (event.Subscription, error) {

	logs, sub, err := _ConstraintsAVAXV1.contract.WatchLogs(opts, "VerifierSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConstraintsAVAXV1VerifierSet)
				if err := _ConstraintsAVAXV1.contract.UnpackLog(event, "VerifierSet", log); err != nil {
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

// ParseVerifierSet is a log parse operation binding the contract event 0x480b37e3d134e44cb444c9703493c7db564c707cb8a18cecea165f340431da1f.
//
// Solidity: event VerifierSet(address verifier)
func (_ConstraintsAVAXV1 *ConstraintsAVAXV1Filterer) ParseVerifierSet(log types.Log) (*ConstraintsAVAXV1VerifierSet, error) {
	event := new(ConstraintsAVAXV1VerifierSet)
	if err := _ConstraintsAVAXV1.contract.UnpackLog(event, "VerifierSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
