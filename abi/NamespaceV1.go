// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package NamespaceV1

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

// NamespaceV1MetaData contains all meta data concerning the NamespaceV1 contract.
var NamespaceV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"namespaceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractConstraintsInterface\",\"name\":\"constraints\",\"type\":\"address\"}],\"name\":\"ConstraintsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"namespaceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gracePeriodLength\",\"type\":\"uint256\"}],\"name\":\"GracePeriodLengthSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"namespaceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"recyclePeriodLength\",\"type\":\"uint256\"}],\"name\":\"RecyclePeriodLengthSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_constraints\",\"outputs\":[{\"internalType\":\"contractConstraintsInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_initializedNamespaces\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"name\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"constraintsData\",\"type\":\"bytes\"}],\"name\":\"checkName\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getGracePeriodLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gracePeriodLength\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getRecyclePeriodLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"recyclePeriodLength\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"contractConstraintsInterface\",\"name\":\"constraints\",\"type\":\"address\"}],\"name\":\"initNamespace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"contractConstraintsInterface\",\"name\":\"constraints\",\"type\":\"address\"}],\"name\":\"setConstraints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gracePeriodLength\",\"type\":\"uint256\"}],\"name\":\"setGracePeriodLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recyclePeriodLength\",\"type\":\"uint256\"}],\"name\":\"setRecyclePeriodLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// NamespaceV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use NamespaceV1MetaData.ABI instead.
var NamespaceV1ABI = NamespaceV1MetaData.ABI

// NamespaceV1 is an auto generated Go binding around an Ethereum contract.
type NamespaceV1 struct {
	NamespaceV1Caller     // Read-only binding to the contract
	NamespaceV1Transactor // Write-only binding to the contract
	NamespaceV1Filterer   // Log filterer for contract events
}

// NamespaceV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type NamespaceV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NamespaceV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type NamespaceV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NamespaceV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NamespaceV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NamespaceV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NamespaceV1Session struct {
	Contract     *NamespaceV1      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NamespaceV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NamespaceV1CallerSession struct {
	Contract *NamespaceV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// NamespaceV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NamespaceV1TransactorSession struct {
	Contract     *NamespaceV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NamespaceV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type NamespaceV1Raw struct {
	Contract *NamespaceV1 // Generic contract binding to access the raw methods on
}

// NamespaceV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NamespaceV1CallerRaw struct {
	Contract *NamespaceV1Caller // Generic read-only contract binding to access the raw methods on
}

// NamespaceV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NamespaceV1TransactorRaw struct {
	Contract *NamespaceV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewNamespaceV1 creates a new instance of NamespaceV1, bound to a specific deployed contract.
func NewNamespaceV1(address common.Address, backend bind.ContractBackend) (*NamespaceV1, error) {
	contract, err := bindNamespaceV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NamespaceV1{NamespaceV1Caller: NamespaceV1Caller{contract: contract}, NamespaceV1Transactor: NamespaceV1Transactor{contract: contract}, NamespaceV1Filterer: NamespaceV1Filterer{contract: contract}}, nil
}

// NewNamespaceV1Caller creates a new read-only instance of NamespaceV1, bound to a specific deployed contract.
func NewNamespaceV1Caller(address common.Address, caller bind.ContractCaller) (*NamespaceV1Caller, error) {
	contract, err := bindNamespaceV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NamespaceV1Caller{contract: contract}, nil
}

// NewNamespaceV1Transactor creates a new write-only instance of NamespaceV1, bound to a specific deployed contract.
func NewNamespaceV1Transactor(address common.Address, transactor bind.ContractTransactor) (*NamespaceV1Transactor, error) {
	contract, err := bindNamespaceV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NamespaceV1Transactor{contract: contract}, nil
}

// NewNamespaceV1Filterer creates a new log filterer instance of NamespaceV1, bound to a specific deployed contract.
func NewNamespaceV1Filterer(address common.Address, filterer bind.ContractFilterer) (*NamespaceV1Filterer, error) {
	contract, err := bindNamespaceV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NamespaceV1Filterer{contract: contract}, nil
}

// bindNamespaceV1 binds a generic wrapper to an already deployed contract.
func bindNamespaceV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NamespaceV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NamespaceV1 *NamespaceV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NamespaceV1.Contract.NamespaceV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NamespaceV1 *NamespaceV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NamespaceV1.Contract.NamespaceV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NamespaceV1 *NamespaceV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NamespaceV1.Contract.NamespaceV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NamespaceV1 *NamespaceV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NamespaceV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NamespaceV1 *NamespaceV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NamespaceV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NamespaceV1 *NamespaceV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NamespaceV1.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NamespaceV1 *NamespaceV1Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NamespaceV1.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NamespaceV1 *NamespaceV1Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _NamespaceV1.Contract.DEFAULTADMINROLE(&_NamespaceV1.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NamespaceV1 *NamespaceV1CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _NamespaceV1.Contract.DEFAULTADMINROLE(&_NamespaceV1.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_NamespaceV1 *NamespaceV1Caller) MANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NamespaceV1.contract.Call(opts, &out, "MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_NamespaceV1 *NamespaceV1Session) MANAGERROLE() ([32]byte, error) {
	return _NamespaceV1.Contract.MANAGERROLE(&_NamespaceV1.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_NamespaceV1 *NamespaceV1CallerSession) MANAGERROLE() ([32]byte, error) {
	return _NamespaceV1.Contract.MANAGERROLE(&_NamespaceV1.CallOpts)
}

// Constraints is a free data retrieval call binding the contract method 0xe81d839d.
//
// Solidity: function _constraints(uint256 ) view returns(address)
func (_NamespaceV1 *NamespaceV1Caller) Constraints(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NamespaceV1.contract.Call(opts, &out, "_constraints", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Constraints is a free data retrieval call binding the contract method 0xe81d839d.
//
// Solidity: function _constraints(uint256 ) view returns(address)
func (_NamespaceV1 *NamespaceV1Session) Constraints(arg0 *big.Int) (common.Address, error) {
	return _NamespaceV1.Contract.Constraints(&_NamespaceV1.CallOpts, arg0)
}

// Constraints is a free data retrieval call binding the contract method 0xe81d839d.
//
// Solidity: function _constraints(uint256 ) view returns(address)
func (_NamespaceV1 *NamespaceV1CallerSession) Constraints(arg0 *big.Int) (common.Address, error) {
	return _NamespaceV1.Contract.Constraints(&_NamespaceV1.CallOpts, arg0)
}

// InitializedNamespaces is a free data retrieval call binding the contract method 0xecb901b0.
//
// Solidity: function _initializedNamespaces(uint256 ) view returns(bool)
func (_NamespaceV1 *NamespaceV1Caller) InitializedNamespaces(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _NamespaceV1.contract.Call(opts, &out, "_initializedNamespaces", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InitializedNamespaces is a free data retrieval call binding the contract method 0xecb901b0.
//
// Solidity: function _initializedNamespaces(uint256 ) view returns(bool)
func (_NamespaceV1 *NamespaceV1Session) InitializedNamespaces(arg0 *big.Int) (bool, error) {
	return _NamespaceV1.Contract.InitializedNamespaces(&_NamespaceV1.CallOpts, arg0)
}

// InitializedNamespaces is a free data retrieval call binding the contract method 0xecb901b0.
//
// Solidity: function _initializedNamespaces(uint256 ) view returns(bool)
func (_NamespaceV1 *NamespaceV1CallerSession) InitializedNamespaces(arg0 *big.Int) (bool, error) {
	return _NamespaceV1.Contract.InitializedNamespaces(&_NamespaceV1.CallOpts, arg0)
}

// CheckName is a free data retrieval call binding the contract method 0x3b566707.
//
// Solidity: function checkName(uint256 id, uint256 name, bytes constraintsData) view returns()
func (_NamespaceV1 *NamespaceV1Caller) CheckName(opts *bind.CallOpts, id *big.Int, name *big.Int, constraintsData []byte) error {
	var out []interface{}
	err := _NamespaceV1.contract.Call(opts, &out, "checkName", id, name, constraintsData)

	if err != nil {
		return err
	}

	return err

}

// CheckName is a free data retrieval call binding the contract method 0x3b566707.
//
// Solidity: function checkName(uint256 id, uint256 name, bytes constraintsData) view returns()
func (_NamespaceV1 *NamespaceV1Session) CheckName(id *big.Int, name *big.Int, constraintsData []byte) error {
	return _NamespaceV1.Contract.CheckName(&_NamespaceV1.CallOpts, id, name, constraintsData)
}

// CheckName is a free data retrieval call binding the contract method 0x3b566707.
//
// Solidity: function checkName(uint256 id, uint256 name, bytes constraintsData) view returns()
func (_NamespaceV1 *NamespaceV1CallerSession) CheckName(id *big.Int, name *big.Int, constraintsData []byte) error {
	return _NamespaceV1.Contract.CheckName(&_NamespaceV1.CallOpts, id, name, constraintsData)
}

// GetGracePeriodLength is a free data retrieval call binding the contract method 0x5575a455.
//
// Solidity: function getGracePeriodLength(uint256 id) view returns(uint256 gracePeriodLength)
func (_NamespaceV1 *NamespaceV1Caller) GetGracePeriodLength(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NamespaceV1.contract.Call(opts, &out, "getGracePeriodLength", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGracePeriodLength is a free data retrieval call binding the contract method 0x5575a455.
//
// Solidity: function getGracePeriodLength(uint256 id) view returns(uint256 gracePeriodLength)
func (_NamespaceV1 *NamespaceV1Session) GetGracePeriodLength(id *big.Int) (*big.Int, error) {
	return _NamespaceV1.Contract.GetGracePeriodLength(&_NamespaceV1.CallOpts, id)
}

// GetGracePeriodLength is a free data retrieval call binding the contract method 0x5575a455.
//
// Solidity: function getGracePeriodLength(uint256 id) view returns(uint256 gracePeriodLength)
func (_NamespaceV1 *NamespaceV1CallerSession) GetGracePeriodLength(id *big.Int) (*big.Int, error) {
	return _NamespaceV1.Contract.GetGracePeriodLength(&_NamespaceV1.CallOpts, id)
}

// GetRecyclePeriodLength is a free data retrieval call binding the contract method 0x93059e2e.
//
// Solidity: function getRecyclePeriodLength(uint256 id) view returns(uint256 recyclePeriodLength)
func (_NamespaceV1 *NamespaceV1Caller) GetRecyclePeriodLength(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NamespaceV1.contract.Call(opts, &out, "getRecyclePeriodLength", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRecyclePeriodLength is a free data retrieval call binding the contract method 0x93059e2e.
//
// Solidity: function getRecyclePeriodLength(uint256 id) view returns(uint256 recyclePeriodLength)
func (_NamespaceV1 *NamespaceV1Session) GetRecyclePeriodLength(id *big.Int) (*big.Int, error) {
	return _NamespaceV1.Contract.GetRecyclePeriodLength(&_NamespaceV1.CallOpts, id)
}

// GetRecyclePeriodLength is a free data retrieval call binding the contract method 0x93059e2e.
//
// Solidity: function getRecyclePeriodLength(uint256 id) view returns(uint256 recyclePeriodLength)
func (_NamespaceV1 *NamespaceV1CallerSession) GetRecyclePeriodLength(id *big.Int) (*big.Int, error) {
	return _NamespaceV1.Contract.GetRecyclePeriodLength(&_NamespaceV1.CallOpts, id)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NamespaceV1 *NamespaceV1Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _NamespaceV1.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NamespaceV1 *NamespaceV1Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _NamespaceV1.Contract.GetRoleAdmin(&_NamespaceV1.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NamespaceV1 *NamespaceV1CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _NamespaceV1.Contract.GetRoleAdmin(&_NamespaceV1.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NamespaceV1 *NamespaceV1Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _NamespaceV1.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NamespaceV1 *NamespaceV1Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _NamespaceV1.Contract.HasRole(&_NamespaceV1.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NamespaceV1 *NamespaceV1CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _NamespaceV1.Contract.HasRole(&_NamespaceV1.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NamespaceV1 *NamespaceV1Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _NamespaceV1.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NamespaceV1 *NamespaceV1Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NamespaceV1.Contract.SupportsInterface(&_NamespaceV1.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NamespaceV1 *NamespaceV1CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NamespaceV1.Contract.SupportsInterface(&_NamespaceV1.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NamespaceV1 *NamespaceV1Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NamespaceV1.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NamespaceV1 *NamespaceV1Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NamespaceV1.Contract.GrantRole(&_NamespaceV1.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NamespaceV1 *NamespaceV1TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NamespaceV1.Contract.GrantRole(&_NamespaceV1.TransactOpts, role, account)
}

// InitNamespace is a paid mutator transaction binding the contract method 0xa52b4b7c.
//
// Solidity: function initNamespace(uint256 id, address constraints) returns()
func (_NamespaceV1 *NamespaceV1Transactor) InitNamespace(opts *bind.TransactOpts, id *big.Int, constraints common.Address) (*types.Transaction, error) {
	return _NamespaceV1.contract.Transact(opts, "initNamespace", id, constraints)
}

// InitNamespace is a paid mutator transaction binding the contract method 0xa52b4b7c.
//
// Solidity: function initNamespace(uint256 id, address constraints) returns()
func (_NamespaceV1 *NamespaceV1Session) InitNamespace(id *big.Int, constraints common.Address) (*types.Transaction, error) {
	return _NamespaceV1.Contract.InitNamespace(&_NamespaceV1.TransactOpts, id, constraints)
}

// InitNamespace is a paid mutator transaction binding the contract method 0xa52b4b7c.
//
// Solidity: function initNamespace(uint256 id, address constraints) returns()
func (_NamespaceV1 *NamespaceV1TransactorSession) InitNamespace(id *big.Int, constraints common.Address) (*types.Transaction, error) {
	return _NamespaceV1.Contract.InitNamespace(&_NamespaceV1.TransactOpts, id, constraints)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_NamespaceV1 *NamespaceV1Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NamespaceV1.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_NamespaceV1 *NamespaceV1Session) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NamespaceV1.Contract.RenounceRole(&_NamespaceV1.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_NamespaceV1 *NamespaceV1TransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NamespaceV1.Contract.RenounceRole(&_NamespaceV1.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NamespaceV1 *NamespaceV1Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NamespaceV1.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NamespaceV1 *NamespaceV1Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NamespaceV1.Contract.RevokeRole(&_NamespaceV1.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NamespaceV1 *NamespaceV1TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NamespaceV1.Contract.RevokeRole(&_NamespaceV1.TransactOpts, role, account)
}

// SetConstraints is a paid mutator transaction binding the contract method 0x087c927e.
//
// Solidity: function setConstraints(uint256 id, address constraints) returns()
func (_NamespaceV1 *NamespaceV1Transactor) SetConstraints(opts *bind.TransactOpts, id *big.Int, constraints common.Address) (*types.Transaction, error) {
	return _NamespaceV1.contract.Transact(opts, "setConstraints", id, constraints)
}

// SetConstraints is a paid mutator transaction binding the contract method 0x087c927e.
//
// Solidity: function setConstraints(uint256 id, address constraints) returns()
func (_NamespaceV1 *NamespaceV1Session) SetConstraints(id *big.Int, constraints common.Address) (*types.Transaction, error) {
	return _NamespaceV1.Contract.SetConstraints(&_NamespaceV1.TransactOpts, id, constraints)
}

// SetConstraints is a paid mutator transaction binding the contract method 0x087c927e.
//
// Solidity: function setConstraints(uint256 id, address constraints) returns()
func (_NamespaceV1 *NamespaceV1TransactorSession) SetConstraints(id *big.Int, constraints common.Address) (*types.Transaction, error) {
	return _NamespaceV1.Contract.SetConstraints(&_NamespaceV1.TransactOpts, id, constraints)
}

// SetGracePeriodLength is a paid mutator transaction binding the contract method 0x1dddb957.
//
// Solidity: function setGracePeriodLength(uint256 id, uint256 gracePeriodLength) returns()
func (_NamespaceV1 *NamespaceV1Transactor) SetGracePeriodLength(opts *bind.TransactOpts, id *big.Int, gracePeriodLength *big.Int) (*types.Transaction, error) {
	return _NamespaceV1.contract.Transact(opts, "setGracePeriodLength", id, gracePeriodLength)
}

// SetGracePeriodLength is a paid mutator transaction binding the contract method 0x1dddb957.
//
// Solidity: function setGracePeriodLength(uint256 id, uint256 gracePeriodLength) returns()
func (_NamespaceV1 *NamespaceV1Session) SetGracePeriodLength(id *big.Int, gracePeriodLength *big.Int) (*types.Transaction, error) {
	return _NamespaceV1.Contract.SetGracePeriodLength(&_NamespaceV1.TransactOpts, id, gracePeriodLength)
}

// SetGracePeriodLength is a paid mutator transaction binding the contract method 0x1dddb957.
//
// Solidity: function setGracePeriodLength(uint256 id, uint256 gracePeriodLength) returns()
func (_NamespaceV1 *NamespaceV1TransactorSession) SetGracePeriodLength(id *big.Int, gracePeriodLength *big.Int) (*types.Transaction, error) {
	return _NamespaceV1.Contract.SetGracePeriodLength(&_NamespaceV1.TransactOpts, id, gracePeriodLength)
}

// SetRecyclePeriodLength is a paid mutator transaction binding the contract method 0x13e60025.
//
// Solidity: function setRecyclePeriodLength(uint256 id, uint256 recyclePeriodLength) returns()
func (_NamespaceV1 *NamespaceV1Transactor) SetRecyclePeriodLength(opts *bind.TransactOpts, id *big.Int, recyclePeriodLength *big.Int) (*types.Transaction, error) {
	return _NamespaceV1.contract.Transact(opts, "setRecyclePeriodLength", id, recyclePeriodLength)
}

// SetRecyclePeriodLength is a paid mutator transaction binding the contract method 0x13e60025.
//
// Solidity: function setRecyclePeriodLength(uint256 id, uint256 recyclePeriodLength) returns()
func (_NamespaceV1 *NamespaceV1Session) SetRecyclePeriodLength(id *big.Int, recyclePeriodLength *big.Int) (*types.Transaction, error) {
	return _NamespaceV1.Contract.SetRecyclePeriodLength(&_NamespaceV1.TransactOpts, id, recyclePeriodLength)
}

// SetRecyclePeriodLength is a paid mutator transaction binding the contract method 0x13e60025.
//
// Solidity: function setRecyclePeriodLength(uint256 id, uint256 recyclePeriodLength) returns()
func (_NamespaceV1 *NamespaceV1TransactorSession) SetRecyclePeriodLength(id *big.Int, recyclePeriodLength *big.Int) (*types.Transaction, error) {
	return _NamespaceV1.Contract.SetRecyclePeriodLength(&_NamespaceV1.TransactOpts, id, recyclePeriodLength)
}

// NamespaceV1ConstraintsSetIterator is returned from FilterConstraintsSet and is used to iterate over the raw logs and unpacked data for ConstraintsSet events raised by the NamespaceV1 contract.
type NamespaceV1ConstraintsSetIterator struct {
	Event *NamespaceV1ConstraintsSet // Event containing the contract specifics and raw log

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
func (it *NamespaceV1ConstraintsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NamespaceV1ConstraintsSet)
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
		it.Event = new(NamespaceV1ConstraintsSet)
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
func (it *NamespaceV1ConstraintsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NamespaceV1ConstraintsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NamespaceV1ConstraintsSet represents a ConstraintsSet event raised by the NamespaceV1 contract.
type NamespaceV1ConstraintsSet struct {
	NamespaceId *big.Int
	Constraints common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConstraintsSet is a free log retrieval operation binding the contract event 0x125621b1bbb30858c4155ac8f386022f8e4c51fc38a48b90162f8be0f96b0e64.
//
// Solidity: event ConstraintsSet(uint256 indexed namespaceId, address constraints)
func (_NamespaceV1 *NamespaceV1Filterer) FilterConstraintsSet(opts *bind.FilterOpts, namespaceId []*big.Int) (*NamespaceV1ConstraintsSetIterator, error) {

	var namespaceIdRule []interface{}
	for _, namespaceIdItem := range namespaceId {
		namespaceIdRule = append(namespaceIdRule, namespaceIdItem)
	}

	logs, sub, err := _NamespaceV1.contract.FilterLogs(opts, "ConstraintsSet", namespaceIdRule)
	if err != nil {
		return nil, err
	}
	return &NamespaceV1ConstraintsSetIterator{contract: _NamespaceV1.contract, event: "ConstraintsSet", logs: logs, sub: sub}, nil
}

// WatchConstraintsSet is a free log subscription operation binding the contract event 0x125621b1bbb30858c4155ac8f386022f8e4c51fc38a48b90162f8be0f96b0e64.
//
// Solidity: event ConstraintsSet(uint256 indexed namespaceId, address constraints)
func (_NamespaceV1 *NamespaceV1Filterer) WatchConstraintsSet(opts *bind.WatchOpts, sink chan<- *NamespaceV1ConstraintsSet, namespaceId []*big.Int) (event.Subscription, error) {

	var namespaceIdRule []interface{}
	for _, namespaceIdItem := range namespaceId {
		namespaceIdRule = append(namespaceIdRule, namespaceIdItem)
	}

	logs, sub, err := _NamespaceV1.contract.WatchLogs(opts, "ConstraintsSet", namespaceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NamespaceV1ConstraintsSet)
				if err := _NamespaceV1.contract.UnpackLog(event, "ConstraintsSet", log); err != nil {
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

// ParseConstraintsSet is a log parse operation binding the contract event 0x125621b1bbb30858c4155ac8f386022f8e4c51fc38a48b90162f8be0f96b0e64.
//
// Solidity: event ConstraintsSet(uint256 indexed namespaceId, address constraints)
func (_NamespaceV1 *NamespaceV1Filterer) ParseConstraintsSet(log types.Log) (*NamespaceV1ConstraintsSet, error) {
	event := new(NamespaceV1ConstraintsSet)
	if err := _NamespaceV1.contract.UnpackLog(event, "ConstraintsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NamespaceV1GracePeriodLengthSetIterator is returned from FilterGracePeriodLengthSet and is used to iterate over the raw logs and unpacked data for GracePeriodLengthSet events raised by the NamespaceV1 contract.
type NamespaceV1GracePeriodLengthSetIterator struct {
	Event *NamespaceV1GracePeriodLengthSet // Event containing the contract specifics and raw log

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
func (it *NamespaceV1GracePeriodLengthSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NamespaceV1GracePeriodLengthSet)
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
		it.Event = new(NamespaceV1GracePeriodLengthSet)
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
func (it *NamespaceV1GracePeriodLengthSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NamespaceV1GracePeriodLengthSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NamespaceV1GracePeriodLengthSet represents a GracePeriodLengthSet event raised by the NamespaceV1 contract.
type NamespaceV1GracePeriodLengthSet struct {
	NamespaceId       *big.Int
	GracePeriodLength *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGracePeriodLengthSet is a free log retrieval operation binding the contract event 0x34494e3e48a45557695ad94e28284a6bc0bf4e2047f35068996c1d3de6b7048f.
//
// Solidity: event GracePeriodLengthSet(uint256 indexed namespaceId, uint256 gracePeriodLength)
func (_NamespaceV1 *NamespaceV1Filterer) FilterGracePeriodLengthSet(opts *bind.FilterOpts, namespaceId []*big.Int) (*NamespaceV1GracePeriodLengthSetIterator, error) {

	var namespaceIdRule []interface{}
	for _, namespaceIdItem := range namespaceId {
		namespaceIdRule = append(namespaceIdRule, namespaceIdItem)
	}

	logs, sub, err := _NamespaceV1.contract.FilterLogs(opts, "GracePeriodLengthSet", namespaceIdRule)
	if err != nil {
		return nil, err
	}
	return &NamespaceV1GracePeriodLengthSetIterator{contract: _NamespaceV1.contract, event: "GracePeriodLengthSet", logs: logs, sub: sub}, nil
}

// WatchGracePeriodLengthSet is a free log subscription operation binding the contract event 0x34494e3e48a45557695ad94e28284a6bc0bf4e2047f35068996c1d3de6b7048f.
//
// Solidity: event GracePeriodLengthSet(uint256 indexed namespaceId, uint256 gracePeriodLength)
func (_NamespaceV1 *NamespaceV1Filterer) WatchGracePeriodLengthSet(opts *bind.WatchOpts, sink chan<- *NamespaceV1GracePeriodLengthSet, namespaceId []*big.Int) (event.Subscription, error) {

	var namespaceIdRule []interface{}
	for _, namespaceIdItem := range namespaceId {
		namespaceIdRule = append(namespaceIdRule, namespaceIdItem)
	}

	logs, sub, err := _NamespaceV1.contract.WatchLogs(opts, "GracePeriodLengthSet", namespaceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NamespaceV1GracePeriodLengthSet)
				if err := _NamespaceV1.contract.UnpackLog(event, "GracePeriodLengthSet", log); err != nil {
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

// ParseGracePeriodLengthSet is a log parse operation binding the contract event 0x34494e3e48a45557695ad94e28284a6bc0bf4e2047f35068996c1d3de6b7048f.
//
// Solidity: event GracePeriodLengthSet(uint256 indexed namespaceId, uint256 gracePeriodLength)
func (_NamespaceV1 *NamespaceV1Filterer) ParseGracePeriodLengthSet(log types.Log) (*NamespaceV1GracePeriodLengthSet, error) {
	event := new(NamespaceV1GracePeriodLengthSet)
	if err := _NamespaceV1.contract.UnpackLog(event, "GracePeriodLengthSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NamespaceV1RecyclePeriodLengthSetIterator is returned from FilterRecyclePeriodLengthSet and is used to iterate over the raw logs and unpacked data for RecyclePeriodLengthSet events raised by the NamespaceV1 contract.
type NamespaceV1RecyclePeriodLengthSetIterator struct {
	Event *NamespaceV1RecyclePeriodLengthSet // Event containing the contract specifics and raw log

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
func (it *NamespaceV1RecyclePeriodLengthSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NamespaceV1RecyclePeriodLengthSet)
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
		it.Event = new(NamespaceV1RecyclePeriodLengthSet)
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
func (it *NamespaceV1RecyclePeriodLengthSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NamespaceV1RecyclePeriodLengthSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NamespaceV1RecyclePeriodLengthSet represents a RecyclePeriodLengthSet event raised by the NamespaceV1 contract.
type NamespaceV1RecyclePeriodLengthSet struct {
	NamespaceId         *big.Int
	RecyclePeriodLength *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRecyclePeriodLengthSet is a free log retrieval operation binding the contract event 0x175d807df355224a08d172bf2ae0a940570a3c3b8b2ca12266d48888fcc137c9.
//
// Solidity: event RecyclePeriodLengthSet(uint256 indexed namespaceId, uint256 recyclePeriodLength)
func (_NamespaceV1 *NamespaceV1Filterer) FilterRecyclePeriodLengthSet(opts *bind.FilterOpts, namespaceId []*big.Int) (*NamespaceV1RecyclePeriodLengthSetIterator, error) {

	var namespaceIdRule []interface{}
	for _, namespaceIdItem := range namespaceId {
		namespaceIdRule = append(namespaceIdRule, namespaceIdItem)
	}

	logs, sub, err := _NamespaceV1.contract.FilterLogs(opts, "RecyclePeriodLengthSet", namespaceIdRule)
	if err != nil {
		return nil, err
	}
	return &NamespaceV1RecyclePeriodLengthSetIterator{contract: _NamespaceV1.contract, event: "RecyclePeriodLengthSet", logs: logs, sub: sub}, nil
}

// WatchRecyclePeriodLengthSet is a free log subscription operation binding the contract event 0x175d807df355224a08d172bf2ae0a940570a3c3b8b2ca12266d48888fcc137c9.
//
// Solidity: event RecyclePeriodLengthSet(uint256 indexed namespaceId, uint256 recyclePeriodLength)
func (_NamespaceV1 *NamespaceV1Filterer) WatchRecyclePeriodLengthSet(opts *bind.WatchOpts, sink chan<- *NamespaceV1RecyclePeriodLengthSet, namespaceId []*big.Int) (event.Subscription, error) {

	var namespaceIdRule []interface{}
	for _, namespaceIdItem := range namespaceId {
		namespaceIdRule = append(namespaceIdRule, namespaceIdItem)
	}

	logs, sub, err := _NamespaceV1.contract.WatchLogs(opts, "RecyclePeriodLengthSet", namespaceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NamespaceV1RecyclePeriodLengthSet)
				if err := _NamespaceV1.contract.UnpackLog(event, "RecyclePeriodLengthSet", log); err != nil {
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

// ParseRecyclePeriodLengthSet is a log parse operation binding the contract event 0x175d807df355224a08d172bf2ae0a940570a3c3b8b2ca12266d48888fcc137c9.
//
// Solidity: event RecyclePeriodLengthSet(uint256 indexed namespaceId, uint256 recyclePeriodLength)
func (_NamespaceV1 *NamespaceV1Filterer) ParseRecyclePeriodLengthSet(log types.Log) (*NamespaceV1RecyclePeriodLengthSet, error) {
	event := new(NamespaceV1RecyclePeriodLengthSet)
	if err := _NamespaceV1.contract.UnpackLog(event, "RecyclePeriodLengthSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NamespaceV1RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the NamespaceV1 contract.
type NamespaceV1RoleAdminChangedIterator struct {
	Event *NamespaceV1RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *NamespaceV1RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NamespaceV1RoleAdminChanged)
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
		it.Event = new(NamespaceV1RoleAdminChanged)
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
func (it *NamespaceV1RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NamespaceV1RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NamespaceV1RoleAdminChanged represents a RoleAdminChanged event raised by the NamespaceV1 contract.
type NamespaceV1RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_NamespaceV1 *NamespaceV1Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*NamespaceV1RoleAdminChangedIterator, error) {

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

	logs, sub, err := _NamespaceV1.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &NamespaceV1RoleAdminChangedIterator{contract: _NamespaceV1.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_NamespaceV1 *NamespaceV1Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *NamespaceV1RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _NamespaceV1.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NamespaceV1RoleAdminChanged)
				if err := _NamespaceV1.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_NamespaceV1 *NamespaceV1Filterer) ParseRoleAdminChanged(log types.Log) (*NamespaceV1RoleAdminChanged, error) {
	event := new(NamespaceV1RoleAdminChanged)
	if err := _NamespaceV1.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NamespaceV1RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the NamespaceV1 contract.
type NamespaceV1RoleGrantedIterator struct {
	Event *NamespaceV1RoleGranted // Event containing the contract specifics and raw log

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
func (it *NamespaceV1RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NamespaceV1RoleGranted)
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
		it.Event = new(NamespaceV1RoleGranted)
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
func (it *NamespaceV1RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NamespaceV1RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NamespaceV1RoleGranted represents a RoleGranted event raised by the NamespaceV1 contract.
type NamespaceV1RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_NamespaceV1 *NamespaceV1Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NamespaceV1RoleGrantedIterator, error) {

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

	logs, sub, err := _NamespaceV1.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NamespaceV1RoleGrantedIterator{contract: _NamespaceV1.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_NamespaceV1 *NamespaceV1Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *NamespaceV1RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _NamespaceV1.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NamespaceV1RoleGranted)
				if err := _NamespaceV1.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_NamespaceV1 *NamespaceV1Filterer) ParseRoleGranted(log types.Log) (*NamespaceV1RoleGranted, error) {
	event := new(NamespaceV1RoleGranted)
	if err := _NamespaceV1.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NamespaceV1RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the NamespaceV1 contract.
type NamespaceV1RoleRevokedIterator struct {
	Event *NamespaceV1RoleRevoked // Event containing the contract specifics and raw log

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
func (it *NamespaceV1RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NamespaceV1RoleRevoked)
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
		it.Event = new(NamespaceV1RoleRevoked)
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
func (it *NamespaceV1RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NamespaceV1RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NamespaceV1RoleRevoked represents a RoleRevoked event raised by the NamespaceV1 contract.
type NamespaceV1RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_NamespaceV1 *NamespaceV1Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NamespaceV1RoleRevokedIterator, error) {

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

	logs, sub, err := _NamespaceV1.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NamespaceV1RoleRevokedIterator{contract: _NamespaceV1.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_NamespaceV1 *NamespaceV1Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *NamespaceV1RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _NamespaceV1.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NamespaceV1RoleRevoked)
				if err := _NamespaceV1.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_NamespaceV1 *NamespaceV1Filterer) ParseRoleRevoked(log types.Log) (*NamespaceV1RoleRevoked, error) {
	event := new(NamespaceV1RoleRevoked)
	if err := _NamespaceV1.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
