// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package LeasingAgentV1

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

// LeasingAgentV1MetaData contains all meta data concerning the LeasingAgentV1 contract.
var LeasingAgentV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"namespaceId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"Enabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"names\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"quantities\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"premiumStartTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"premiumEndTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"premiumPricePoints\",\"type\":\"uint256[]\"}],\"name\":\"RegistrationPremiumSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_contractRegistry\",\"outputs\":[{\"internalType\":\"contractContractRegistryInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_enabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_namespaceId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_premiumEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_premiumPricePoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_premiumStartTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"enable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"t\",\"type\":\"uint256\"}],\"name\":\"getRegistrationPremium\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"names\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"quantities\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"constraintsProofs\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"pricingProofs\",\"type\":\"bytes[]\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"names\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"quantities\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"constraintsProofs\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"pricingProofs\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"preimages\",\"type\":\"uint256[]\"}],\"name\":\"registerWithPreimage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"premiumStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premiumEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"premiumPricePoints\",\"type\":\"uint256[]\"}],\"name\":\"setRegistrationPremiumDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LeasingAgentV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use LeasingAgentV1MetaData.ABI instead.
var LeasingAgentV1ABI = LeasingAgentV1MetaData.ABI

// LeasingAgentV1 is an auto generated Go binding around an Ethereum contract.
type LeasingAgentV1 struct {
	LeasingAgentV1Caller     // Read-only binding to the contract
	LeasingAgentV1Transactor // Write-only binding to the contract
	LeasingAgentV1Filterer   // Log filterer for contract events
}

// LeasingAgentV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type LeasingAgentV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeasingAgentV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type LeasingAgentV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeasingAgentV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LeasingAgentV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeasingAgentV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LeasingAgentV1Session struct {
	Contract     *LeasingAgentV1   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LeasingAgentV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LeasingAgentV1CallerSession struct {
	Contract *LeasingAgentV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// LeasingAgentV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LeasingAgentV1TransactorSession struct {
	Contract     *LeasingAgentV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// LeasingAgentV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type LeasingAgentV1Raw struct {
	Contract *LeasingAgentV1 // Generic contract binding to access the raw methods on
}

// LeasingAgentV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LeasingAgentV1CallerRaw struct {
	Contract *LeasingAgentV1Caller // Generic read-only contract binding to access the raw methods on
}

// LeasingAgentV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LeasingAgentV1TransactorRaw struct {
	Contract *LeasingAgentV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewLeasingAgentV1 creates a new instance of LeasingAgentV1, bound to a specific deployed contract.
func NewLeasingAgentV1(address common.Address, backend bind.ContractBackend) (*LeasingAgentV1, error) {
	contract, err := bindLeasingAgentV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LeasingAgentV1{LeasingAgentV1Caller: LeasingAgentV1Caller{contract: contract}, LeasingAgentV1Transactor: LeasingAgentV1Transactor{contract: contract}, LeasingAgentV1Filterer: LeasingAgentV1Filterer{contract: contract}}, nil
}

// NewLeasingAgentV1Caller creates a new read-only instance of LeasingAgentV1, bound to a specific deployed contract.
func NewLeasingAgentV1Caller(address common.Address, caller bind.ContractCaller) (*LeasingAgentV1Caller, error) {
	contract, err := bindLeasingAgentV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LeasingAgentV1Caller{contract: contract}, nil
}

// NewLeasingAgentV1Transactor creates a new write-only instance of LeasingAgentV1, bound to a specific deployed contract.
func NewLeasingAgentV1Transactor(address common.Address, transactor bind.ContractTransactor) (*LeasingAgentV1Transactor, error) {
	contract, err := bindLeasingAgentV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LeasingAgentV1Transactor{contract: contract}, nil
}

// NewLeasingAgentV1Filterer creates a new log filterer instance of LeasingAgentV1, bound to a specific deployed contract.
func NewLeasingAgentV1Filterer(address common.Address, filterer bind.ContractFilterer) (*LeasingAgentV1Filterer, error) {
	contract, err := bindLeasingAgentV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LeasingAgentV1Filterer{contract: contract}, nil
}

// bindLeasingAgentV1 binds a generic wrapper to an already deployed contract.
func bindLeasingAgentV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LeasingAgentV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LeasingAgentV1 *LeasingAgentV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LeasingAgentV1.Contract.LeasingAgentV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LeasingAgentV1 *LeasingAgentV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LeasingAgentV1.Contract.LeasingAgentV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LeasingAgentV1 *LeasingAgentV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LeasingAgentV1.Contract.LeasingAgentV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LeasingAgentV1 *LeasingAgentV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LeasingAgentV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LeasingAgentV1 *LeasingAgentV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LeasingAgentV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LeasingAgentV1 *LeasingAgentV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LeasingAgentV1.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LeasingAgentV1 *LeasingAgentV1Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LeasingAgentV1.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LeasingAgentV1 *LeasingAgentV1Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _LeasingAgentV1.Contract.DEFAULTADMINROLE(&_LeasingAgentV1.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LeasingAgentV1 *LeasingAgentV1CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LeasingAgentV1.Contract.DEFAULTADMINROLE(&_LeasingAgentV1.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_LeasingAgentV1 *LeasingAgentV1Caller) MANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LeasingAgentV1.contract.Call(opts, &out, "MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_LeasingAgentV1 *LeasingAgentV1Session) MANAGERROLE() ([32]byte, error) {
	return _LeasingAgentV1.Contract.MANAGERROLE(&_LeasingAgentV1.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_LeasingAgentV1 *LeasingAgentV1CallerSession) MANAGERROLE() ([32]byte, error) {
	return _LeasingAgentV1.Contract.MANAGERROLE(&_LeasingAgentV1.CallOpts)
}

// ContractRegistry is a free data retrieval call binding the contract method 0xc1058b8f.
//
// Solidity: function _contractRegistry() view returns(address)
func (_LeasingAgentV1 *LeasingAgentV1Caller) ContractRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LeasingAgentV1.contract.Call(opts, &out, "_contractRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractRegistry is a free data retrieval call binding the contract method 0xc1058b8f.
//
// Solidity: function _contractRegistry() view returns(address)
func (_LeasingAgentV1 *LeasingAgentV1Session) ContractRegistry() (common.Address, error) {
	return _LeasingAgentV1.Contract.ContractRegistry(&_LeasingAgentV1.CallOpts)
}

// ContractRegistry is a free data retrieval call binding the contract method 0xc1058b8f.
//
// Solidity: function _contractRegistry() view returns(address)
func (_LeasingAgentV1 *LeasingAgentV1CallerSession) ContractRegistry() (common.Address, error) {
	return _LeasingAgentV1.Contract.ContractRegistry(&_LeasingAgentV1.CallOpts)
}

// Enabled is a free data retrieval call binding the contract method 0xc36be496.
//
// Solidity: function _enabled() view returns(bool)
func (_LeasingAgentV1 *LeasingAgentV1Caller) Enabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LeasingAgentV1.contract.Call(opts, &out, "_enabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Enabled is a free data retrieval call binding the contract method 0xc36be496.
//
// Solidity: function _enabled() view returns(bool)
func (_LeasingAgentV1 *LeasingAgentV1Session) Enabled() (bool, error) {
	return _LeasingAgentV1.Contract.Enabled(&_LeasingAgentV1.CallOpts)
}

// Enabled is a free data retrieval call binding the contract method 0xc36be496.
//
// Solidity: function _enabled() view returns(bool)
func (_LeasingAgentV1 *LeasingAgentV1CallerSession) Enabled() (bool, error) {
	return _LeasingAgentV1.Contract.Enabled(&_LeasingAgentV1.CallOpts)
}

// NamespaceId is a free data retrieval call binding the contract method 0xf1be1461.
//
// Solidity: function _namespaceId() view returns(uint256)
func (_LeasingAgentV1 *LeasingAgentV1Caller) NamespaceId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LeasingAgentV1.contract.Call(opts, &out, "_namespaceId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NamespaceId is a free data retrieval call binding the contract method 0xf1be1461.
//
// Solidity: function _namespaceId() view returns(uint256)
func (_LeasingAgentV1 *LeasingAgentV1Session) NamespaceId() (*big.Int, error) {
	return _LeasingAgentV1.Contract.NamespaceId(&_LeasingAgentV1.CallOpts)
}

// NamespaceId is a free data retrieval call binding the contract method 0xf1be1461.
//
// Solidity: function _namespaceId() view returns(uint256)
func (_LeasingAgentV1 *LeasingAgentV1CallerSession) NamespaceId() (*big.Int, error) {
	return _LeasingAgentV1.Contract.NamespaceId(&_LeasingAgentV1.CallOpts)
}

// PremiumEndTime is a free data retrieval call binding the contract method 0x4630d113.
//
// Solidity: function _premiumEndTime() view returns(uint256)
func (_LeasingAgentV1 *LeasingAgentV1Caller) PremiumEndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LeasingAgentV1.contract.Call(opts, &out, "_premiumEndTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PremiumEndTime is a free data retrieval call binding the contract method 0x4630d113.
//
// Solidity: function _premiumEndTime() view returns(uint256)
func (_LeasingAgentV1 *LeasingAgentV1Session) PremiumEndTime() (*big.Int, error) {
	return _LeasingAgentV1.Contract.PremiumEndTime(&_LeasingAgentV1.CallOpts)
}

// PremiumEndTime is a free data retrieval call binding the contract method 0x4630d113.
//
// Solidity: function _premiumEndTime() view returns(uint256)
func (_LeasingAgentV1 *LeasingAgentV1CallerSession) PremiumEndTime() (*big.Int, error) {
	return _LeasingAgentV1.Contract.PremiumEndTime(&_LeasingAgentV1.CallOpts)
}

// PremiumPricePoints is a free data retrieval call binding the contract method 0x26634d20.
//
// Solidity: function _premiumPricePoints(uint256 ) view returns(uint256)
func (_LeasingAgentV1 *LeasingAgentV1Caller) PremiumPricePoints(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LeasingAgentV1.contract.Call(opts, &out, "_premiumPricePoints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PremiumPricePoints is a free data retrieval call binding the contract method 0x26634d20.
//
// Solidity: function _premiumPricePoints(uint256 ) view returns(uint256)
func (_LeasingAgentV1 *LeasingAgentV1Session) PremiumPricePoints(arg0 *big.Int) (*big.Int, error) {
	return _LeasingAgentV1.Contract.PremiumPricePoints(&_LeasingAgentV1.CallOpts, arg0)
}

// PremiumPricePoints is a free data retrieval call binding the contract method 0x26634d20.
//
// Solidity: function _premiumPricePoints(uint256 ) view returns(uint256)
func (_LeasingAgentV1 *LeasingAgentV1CallerSession) PremiumPricePoints(arg0 *big.Int) (*big.Int, error) {
	return _LeasingAgentV1.Contract.PremiumPricePoints(&_LeasingAgentV1.CallOpts, arg0)
}

// PremiumStartTime is a free data retrieval call binding the contract method 0x037a561d.
//
// Solidity: function _premiumStartTime() view returns(uint256)
func (_LeasingAgentV1 *LeasingAgentV1Caller) PremiumStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LeasingAgentV1.contract.Call(opts, &out, "_premiumStartTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PremiumStartTime is a free data retrieval call binding the contract method 0x037a561d.
//
// Solidity: function _premiumStartTime() view returns(uint256)
func (_LeasingAgentV1 *LeasingAgentV1Session) PremiumStartTime() (*big.Int, error) {
	return _LeasingAgentV1.Contract.PremiumStartTime(&_LeasingAgentV1.CallOpts)
}

// PremiumStartTime is a free data retrieval call binding the contract method 0x037a561d.
//
// Solidity: function _premiumStartTime() view returns(uint256)
func (_LeasingAgentV1 *LeasingAgentV1CallerSession) PremiumStartTime() (*big.Int, error) {
	return _LeasingAgentV1.Contract.PremiumStartTime(&_LeasingAgentV1.CallOpts)
}

// GetRegistrationPremium is a free data retrieval call binding the contract method 0x67b76924.
//
// Solidity: function getRegistrationPremium(uint256 t) view returns(uint256)
func (_LeasingAgentV1 *LeasingAgentV1Caller) GetRegistrationPremium(opts *bind.CallOpts, t *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LeasingAgentV1.contract.Call(opts, &out, "getRegistrationPremium", t)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRegistrationPremium is a free data retrieval call binding the contract method 0x67b76924.
//
// Solidity: function getRegistrationPremium(uint256 t) view returns(uint256)
func (_LeasingAgentV1 *LeasingAgentV1Session) GetRegistrationPremium(t *big.Int) (*big.Int, error) {
	return _LeasingAgentV1.Contract.GetRegistrationPremium(&_LeasingAgentV1.CallOpts, t)
}

// GetRegistrationPremium is a free data retrieval call binding the contract method 0x67b76924.
//
// Solidity: function getRegistrationPremium(uint256 t) view returns(uint256)
func (_LeasingAgentV1 *LeasingAgentV1CallerSession) GetRegistrationPremium(t *big.Int) (*big.Int, error) {
	return _LeasingAgentV1.Contract.GetRegistrationPremium(&_LeasingAgentV1.CallOpts, t)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LeasingAgentV1 *LeasingAgentV1Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _LeasingAgentV1.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LeasingAgentV1 *LeasingAgentV1Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LeasingAgentV1.Contract.GetRoleAdmin(&_LeasingAgentV1.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LeasingAgentV1 *LeasingAgentV1CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LeasingAgentV1.Contract.GetRoleAdmin(&_LeasingAgentV1.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LeasingAgentV1 *LeasingAgentV1Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _LeasingAgentV1.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LeasingAgentV1 *LeasingAgentV1Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LeasingAgentV1.Contract.HasRole(&_LeasingAgentV1.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LeasingAgentV1 *LeasingAgentV1CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LeasingAgentV1.Contract.HasRole(&_LeasingAgentV1.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LeasingAgentV1 *LeasingAgentV1Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LeasingAgentV1.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LeasingAgentV1 *LeasingAgentV1Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LeasingAgentV1.Contract.SupportsInterface(&_LeasingAgentV1.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LeasingAgentV1 *LeasingAgentV1CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LeasingAgentV1.Contract.SupportsInterface(&_LeasingAgentV1.CallOpts, interfaceId)
}

// Enable is a paid mutator transaction binding the contract method 0x4df68ada.
//
// Solidity: function enable(bool enabled) returns()
func (_LeasingAgentV1 *LeasingAgentV1Transactor) Enable(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _LeasingAgentV1.contract.Transact(opts, "enable", enabled)
}

// Enable is a paid mutator transaction binding the contract method 0x4df68ada.
//
// Solidity: function enable(bool enabled) returns()
func (_LeasingAgentV1 *LeasingAgentV1Session) Enable(enabled bool) (*types.Transaction, error) {
	return _LeasingAgentV1.Contract.Enable(&_LeasingAgentV1.TransactOpts, enabled)
}

// Enable is a paid mutator transaction binding the contract method 0x4df68ada.
//
// Solidity: function enable(bool enabled) returns()
func (_LeasingAgentV1 *LeasingAgentV1TransactorSession) Enable(enabled bool) (*types.Transaction, error) {
	return _LeasingAgentV1.Contract.Enable(&_LeasingAgentV1.TransactOpts, enabled)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LeasingAgentV1 *LeasingAgentV1Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LeasingAgentV1.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LeasingAgentV1 *LeasingAgentV1Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LeasingAgentV1.Contract.GrantRole(&_LeasingAgentV1.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LeasingAgentV1 *LeasingAgentV1TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LeasingAgentV1.Contract.GrantRole(&_LeasingAgentV1.TransactOpts, role, account)
}

// Register is a paid mutator transaction binding the contract method 0x74c9f18f.
//
// Solidity: function register(uint256[] names, uint256[] quantities, bytes[] constraintsProofs, bytes[] pricingProofs) payable returns()
func (_LeasingAgentV1 *LeasingAgentV1Transactor) Register(opts *bind.TransactOpts, names []*big.Int, quantities []*big.Int, constraintsProofs [][]byte, pricingProofs [][]byte) (*types.Transaction, error) {
	return _LeasingAgentV1.contract.Transact(opts, "register", names, quantities, constraintsProofs, pricingProofs)
}

// Register is a paid mutator transaction binding the contract method 0x74c9f18f.
//
// Solidity: function register(uint256[] names, uint256[] quantities, bytes[] constraintsProofs, bytes[] pricingProofs) payable returns()
func (_LeasingAgentV1 *LeasingAgentV1Session) Register(names []*big.Int, quantities []*big.Int, constraintsProofs [][]byte, pricingProofs [][]byte) (*types.Transaction, error) {
	return _LeasingAgentV1.Contract.Register(&_LeasingAgentV1.TransactOpts, names, quantities, constraintsProofs, pricingProofs)
}

// Register is a paid mutator transaction binding the contract method 0x74c9f18f.
//
// Solidity: function register(uint256[] names, uint256[] quantities, bytes[] constraintsProofs, bytes[] pricingProofs) payable returns()
func (_LeasingAgentV1 *LeasingAgentV1TransactorSession) Register(names []*big.Int, quantities []*big.Int, constraintsProofs [][]byte, pricingProofs [][]byte) (*types.Transaction, error) {
	return _LeasingAgentV1.Contract.Register(&_LeasingAgentV1.TransactOpts, names, quantities, constraintsProofs, pricingProofs)
}

// RegisterWithPreimage is a paid mutator transaction binding the contract method 0xa1da5024.
//
// Solidity: function registerWithPreimage(uint256[] names, uint256[] quantities, bytes[] constraintsProofs, bytes[] pricingProofs, uint256[] preimages) payable returns()
func (_LeasingAgentV1 *LeasingAgentV1Transactor) RegisterWithPreimage(opts *bind.TransactOpts, names []*big.Int, quantities []*big.Int, constraintsProofs [][]byte, pricingProofs [][]byte, preimages []*big.Int) (*types.Transaction, error) {
	return _LeasingAgentV1.contract.Transact(opts, "registerWithPreimage", names, quantities, constraintsProofs, pricingProofs, preimages)
}

// RegisterWithPreimage is a paid mutator transaction binding the contract method 0xa1da5024.
//
// Solidity: function registerWithPreimage(uint256[] names, uint256[] quantities, bytes[] constraintsProofs, bytes[] pricingProofs, uint256[] preimages) payable returns()
func (_LeasingAgentV1 *LeasingAgentV1Session) RegisterWithPreimage(names []*big.Int, quantities []*big.Int, constraintsProofs [][]byte, pricingProofs [][]byte, preimages []*big.Int) (*types.Transaction, error) {
	return _LeasingAgentV1.Contract.RegisterWithPreimage(&_LeasingAgentV1.TransactOpts, names, quantities, constraintsProofs, pricingProofs, preimages)
}

// RegisterWithPreimage is a paid mutator transaction binding the contract method 0xa1da5024.
//
// Solidity: function registerWithPreimage(uint256[] names, uint256[] quantities, bytes[] constraintsProofs, bytes[] pricingProofs, uint256[] preimages) payable returns()
func (_LeasingAgentV1 *LeasingAgentV1TransactorSession) RegisterWithPreimage(names []*big.Int, quantities []*big.Int, constraintsProofs [][]byte, pricingProofs [][]byte, preimages []*big.Int) (*types.Transaction, error) {
	return _LeasingAgentV1.Contract.RegisterWithPreimage(&_LeasingAgentV1.TransactOpts, names, quantities, constraintsProofs, pricingProofs, preimages)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_LeasingAgentV1 *LeasingAgentV1Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LeasingAgentV1.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_LeasingAgentV1 *LeasingAgentV1Session) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LeasingAgentV1.Contract.RenounceRole(&_LeasingAgentV1.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_LeasingAgentV1 *LeasingAgentV1TransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LeasingAgentV1.Contract.RenounceRole(&_LeasingAgentV1.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LeasingAgentV1 *LeasingAgentV1Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LeasingAgentV1.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LeasingAgentV1 *LeasingAgentV1Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LeasingAgentV1.Contract.RevokeRole(&_LeasingAgentV1.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LeasingAgentV1 *LeasingAgentV1TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LeasingAgentV1.Contract.RevokeRole(&_LeasingAgentV1.TransactOpts, role, account)
}

// SetRegistrationPremiumDetails is a paid mutator transaction binding the contract method 0x17ae3213.
//
// Solidity: function setRegistrationPremiumDetails(uint256 premiumStartTime, uint256 premiumEndTime, uint256[] premiumPricePoints) returns()
func (_LeasingAgentV1 *LeasingAgentV1Transactor) SetRegistrationPremiumDetails(opts *bind.TransactOpts, premiumStartTime *big.Int, premiumEndTime *big.Int, premiumPricePoints []*big.Int) (*types.Transaction, error) {
	return _LeasingAgentV1.contract.Transact(opts, "setRegistrationPremiumDetails", premiumStartTime, premiumEndTime, premiumPricePoints)
}

// SetRegistrationPremiumDetails is a paid mutator transaction binding the contract method 0x17ae3213.
//
// Solidity: function setRegistrationPremiumDetails(uint256 premiumStartTime, uint256 premiumEndTime, uint256[] premiumPricePoints) returns()
func (_LeasingAgentV1 *LeasingAgentV1Session) SetRegistrationPremiumDetails(premiumStartTime *big.Int, premiumEndTime *big.Int, premiumPricePoints []*big.Int) (*types.Transaction, error) {
	return _LeasingAgentV1.Contract.SetRegistrationPremiumDetails(&_LeasingAgentV1.TransactOpts, premiumStartTime, premiumEndTime, premiumPricePoints)
}

// SetRegistrationPremiumDetails is a paid mutator transaction binding the contract method 0x17ae3213.
//
// Solidity: function setRegistrationPremiumDetails(uint256 premiumStartTime, uint256 premiumEndTime, uint256[] premiumPricePoints) returns()
func (_LeasingAgentV1 *LeasingAgentV1TransactorSession) SetRegistrationPremiumDetails(premiumStartTime *big.Int, premiumEndTime *big.Int, premiumPricePoints []*big.Int) (*types.Transaction, error) {
	return _LeasingAgentV1.Contract.SetRegistrationPremiumDetails(&_LeasingAgentV1.TransactOpts, premiumStartTime, premiumEndTime, premiumPricePoints)
}

// LeasingAgentV1EnabledIterator is returned from FilterEnabled and is used to iterate over the raw logs and unpacked data for Enabled events raised by the LeasingAgentV1 contract.
type LeasingAgentV1EnabledIterator struct {
	Event *LeasingAgentV1Enabled // Event containing the contract specifics and raw log

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
func (it *LeasingAgentV1EnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeasingAgentV1Enabled)
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
		it.Event = new(LeasingAgentV1Enabled)
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
func (it *LeasingAgentV1EnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeasingAgentV1EnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeasingAgentV1Enabled represents a Enabled event raised by the LeasingAgentV1 contract.
type LeasingAgentV1Enabled struct {
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEnabled is a free log retrieval operation binding the contract event 0x54ca1f89aee2b8d11c89b7813c6aa99caa0f8c55c8eccf8b70c3bb42029fa134.
//
// Solidity: event Enabled(bool enabled)
func (_LeasingAgentV1 *LeasingAgentV1Filterer) FilterEnabled(opts *bind.FilterOpts) (*LeasingAgentV1EnabledIterator, error) {

	logs, sub, err := _LeasingAgentV1.contract.FilterLogs(opts, "Enabled")
	if err != nil {
		return nil, err
	}
	return &LeasingAgentV1EnabledIterator{contract: _LeasingAgentV1.contract, event: "Enabled", logs: logs, sub: sub}, nil
}

// WatchEnabled is a free log subscription operation binding the contract event 0x54ca1f89aee2b8d11c89b7813c6aa99caa0f8c55c8eccf8b70c3bb42029fa134.
//
// Solidity: event Enabled(bool enabled)
func (_LeasingAgentV1 *LeasingAgentV1Filterer) WatchEnabled(opts *bind.WatchOpts, sink chan<- *LeasingAgentV1Enabled) (event.Subscription, error) {

	logs, sub, err := _LeasingAgentV1.contract.WatchLogs(opts, "Enabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeasingAgentV1Enabled)
				if err := _LeasingAgentV1.contract.UnpackLog(event, "Enabled", log); err != nil {
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

// ParseEnabled is a log parse operation binding the contract event 0x54ca1f89aee2b8d11c89b7813c6aa99caa0f8c55c8eccf8b70c3bb42029fa134.
//
// Solidity: event Enabled(bool enabled)
func (_LeasingAgentV1 *LeasingAgentV1Filterer) ParseEnabled(log types.Log) (*LeasingAgentV1Enabled, error) {
	event := new(LeasingAgentV1Enabled)
	if err := _LeasingAgentV1.contract.UnpackLog(event, "Enabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeasingAgentV1RegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the LeasingAgentV1 contract.
type LeasingAgentV1RegisteredIterator struct {
	Event *LeasingAgentV1Registered // Event containing the contract specifics and raw log

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
func (it *LeasingAgentV1RegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeasingAgentV1Registered)
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
		it.Event = new(LeasingAgentV1Registered)
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
func (it *LeasingAgentV1RegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeasingAgentV1RegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeasingAgentV1Registered represents a Registered event raised by the LeasingAgentV1 contract.
type LeasingAgentV1Registered struct {
	Names      []*big.Int
	Quantities []*big.Int
	Payment    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0x56e633e09b0a624ca11db9196c2f48c5a7b4421dc876324d3774c1df2efed650.
//
// Solidity: event Registered(uint256[] names, uint256[] quantities, uint256 payment)
func (_LeasingAgentV1 *LeasingAgentV1Filterer) FilterRegistered(opts *bind.FilterOpts) (*LeasingAgentV1RegisteredIterator, error) {

	logs, sub, err := _LeasingAgentV1.contract.FilterLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return &LeasingAgentV1RegisteredIterator{contract: _LeasingAgentV1.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0x56e633e09b0a624ca11db9196c2f48c5a7b4421dc876324d3774c1df2efed650.
//
// Solidity: event Registered(uint256[] names, uint256[] quantities, uint256 payment)
func (_LeasingAgentV1 *LeasingAgentV1Filterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *LeasingAgentV1Registered) (event.Subscription, error) {

	logs, sub, err := _LeasingAgentV1.contract.WatchLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeasingAgentV1Registered)
				if err := _LeasingAgentV1.contract.UnpackLog(event, "Registered", log); err != nil {
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

// ParseRegistered is a log parse operation binding the contract event 0x56e633e09b0a624ca11db9196c2f48c5a7b4421dc876324d3774c1df2efed650.
//
// Solidity: event Registered(uint256[] names, uint256[] quantities, uint256 payment)
func (_LeasingAgentV1 *LeasingAgentV1Filterer) ParseRegistered(log types.Log) (*LeasingAgentV1Registered, error) {
	event := new(LeasingAgentV1Registered)
	if err := _LeasingAgentV1.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeasingAgentV1RegistrationPremiumSetIterator is returned from FilterRegistrationPremiumSet and is used to iterate over the raw logs and unpacked data for RegistrationPremiumSet events raised by the LeasingAgentV1 contract.
type LeasingAgentV1RegistrationPremiumSetIterator struct {
	Event *LeasingAgentV1RegistrationPremiumSet // Event containing the contract specifics and raw log

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
func (it *LeasingAgentV1RegistrationPremiumSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeasingAgentV1RegistrationPremiumSet)
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
		it.Event = new(LeasingAgentV1RegistrationPremiumSet)
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
func (it *LeasingAgentV1RegistrationPremiumSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeasingAgentV1RegistrationPremiumSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeasingAgentV1RegistrationPremiumSet represents a RegistrationPremiumSet event raised by the LeasingAgentV1 contract.
type LeasingAgentV1RegistrationPremiumSet struct {
	PremiumStartTime   *big.Int
	PremiumEndTime     *big.Int
	PremiumPricePoints []*big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRegistrationPremiumSet is a free log retrieval operation binding the contract event 0x3bed3244e6743d5945578f85ef626f0df265038562d9376fe3182a00bcd7adc0.
//
// Solidity: event RegistrationPremiumSet(uint256 premiumStartTime, uint256 premiumEndTime, uint256[] premiumPricePoints)
func (_LeasingAgentV1 *LeasingAgentV1Filterer) FilterRegistrationPremiumSet(opts *bind.FilterOpts) (*LeasingAgentV1RegistrationPremiumSetIterator, error) {

	logs, sub, err := _LeasingAgentV1.contract.FilterLogs(opts, "RegistrationPremiumSet")
	if err != nil {
		return nil, err
	}
	return &LeasingAgentV1RegistrationPremiumSetIterator{contract: _LeasingAgentV1.contract, event: "RegistrationPremiumSet", logs: logs, sub: sub}, nil
}

// WatchRegistrationPremiumSet is a free log subscription operation binding the contract event 0x3bed3244e6743d5945578f85ef626f0df265038562d9376fe3182a00bcd7adc0.
//
// Solidity: event RegistrationPremiumSet(uint256 premiumStartTime, uint256 premiumEndTime, uint256[] premiumPricePoints)
func (_LeasingAgentV1 *LeasingAgentV1Filterer) WatchRegistrationPremiumSet(opts *bind.WatchOpts, sink chan<- *LeasingAgentV1RegistrationPremiumSet) (event.Subscription, error) {

	logs, sub, err := _LeasingAgentV1.contract.WatchLogs(opts, "RegistrationPremiumSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeasingAgentV1RegistrationPremiumSet)
				if err := _LeasingAgentV1.contract.UnpackLog(event, "RegistrationPremiumSet", log); err != nil {
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

// ParseRegistrationPremiumSet is a log parse operation binding the contract event 0x3bed3244e6743d5945578f85ef626f0df265038562d9376fe3182a00bcd7adc0.
//
// Solidity: event RegistrationPremiumSet(uint256 premiumStartTime, uint256 premiumEndTime, uint256[] premiumPricePoints)
func (_LeasingAgentV1 *LeasingAgentV1Filterer) ParseRegistrationPremiumSet(log types.Log) (*LeasingAgentV1RegistrationPremiumSet, error) {
	event := new(LeasingAgentV1RegistrationPremiumSet)
	if err := _LeasingAgentV1.contract.UnpackLog(event, "RegistrationPremiumSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeasingAgentV1RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the LeasingAgentV1 contract.
type LeasingAgentV1RoleAdminChangedIterator struct {
	Event *LeasingAgentV1RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *LeasingAgentV1RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeasingAgentV1RoleAdminChanged)
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
		it.Event = new(LeasingAgentV1RoleAdminChanged)
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
func (it *LeasingAgentV1RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeasingAgentV1RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeasingAgentV1RoleAdminChanged represents a RoleAdminChanged event raised by the LeasingAgentV1 contract.
type LeasingAgentV1RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LeasingAgentV1 *LeasingAgentV1Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*LeasingAgentV1RoleAdminChangedIterator, error) {

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

	logs, sub, err := _LeasingAgentV1.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &LeasingAgentV1RoleAdminChangedIterator{contract: _LeasingAgentV1.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LeasingAgentV1 *LeasingAgentV1Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *LeasingAgentV1RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _LeasingAgentV1.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeasingAgentV1RoleAdminChanged)
				if err := _LeasingAgentV1.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_LeasingAgentV1 *LeasingAgentV1Filterer) ParseRoleAdminChanged(log types.Log) (*LeasingAgentV1RoleAdminChanged, error) {
	event := new(LeasingAgentV1RoleAdminChanged)
	if err := _LeasingAgentV1.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeasingAgentV1RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the LeasingAgentV1 contract.
type LeasingAgentV1RoleGrantedIterator struct {
	Event *LeasingAgentV1RoleGranted // Event containing the contract specifics and raw log

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
func (it *LeasingAgentV1RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeasingAgentV1RoleGranted)
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
		it.Event = new(LeasingAgentV1RoleGranted)
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
func (it *LeasingAgentV1RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeasingAgentV1RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeasingAgentV1RoleGranted represents a RoleGranted event raised by the LeasingAgentV1 contract.
type LeasingAgentV1RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LeasingAgentV1 *LeasingAgentV1Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LeasingAgentV1RoleGrantedIterator, error) {

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

	logs, sub, err := _LeasingAgentV1.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LeasingAgentV1RoleGrantedIterator{contract: _LeasingAgentV1.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LeasingAgentV1 *LeasingAgentV1Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *LeasingAgentV1RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _LeasingAgentV1.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeasingAgentV1RoleGranted)
				if err := _LeasingAgentV1.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_LeasingAgentV1 *LeasingAgentV1Filterer) ParseRoleGranted(log types.Log) (*LeasingAgentV1RoleGranted, error) {
	event := new(LeasingAgentV1RoleGranted)
	if err := _LeasingAgentV1.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeasingAgentV1RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the LeasingAgentV1 contract.
type LeasingAgentV1RoleRevokedIterator struct {
	Event *LeasingAgentV1RoleRevoked // Event containing the contract specifics and raw log

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
func (it *LeasingAgentV1RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeasingAgentV1RoleRevoked)
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
		it.Event = new(LeasingAgentV1RoleRevoked)
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
func (it *LeasingAgentV1RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeasingAgentV1RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeasingAgentV1RoleRevoked represents a RoleRevoked event raised by the LeasingAgentV1 contract.
type LeasingAgentV1RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LeasingAgentV1 *LeasingAgentV1Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LeasingAgentV1RoleRevokedIterator, error) {

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

	logs, sub, err := _LeasingAgentV1.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LeasingAgentV1RoleRevokedIterator{contract: _LeasingAgentV1.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LeasingAgentV1 *LeasingAgentV1Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *LeasingAgentV1RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _LeasingAgentV1.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeasingAgentV1RoleRevoked)
				if err := _LeasingAgentV1.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_LeasingAgentV1 *LeasingAgentV1Filterer) ParseRoleRevoked(log types.Log) (*LeasingAgentV1RoleRevoked, error) {
	event := new(LeasingAgentV1RoleRevoked)
	if err := _LeasingAgentV1.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
