// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Domain

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

// DomainMetaData contains all meta data concerning the Domain contract.
var DomainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"contractContractRegistryInterface\",\"name\":\"contractRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"ContractURISet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"registrant\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"name\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"leaseLength\",\"type\":\"uint256\"}],\"name\":\"Recycle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"registrant\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"name\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"leaseLength\",\"type\":\"uint256\"}],\"name\":\"Register\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"name\",\"type\":\"uint256\"}],\"name\":\"Revoke\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"name\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"suspended\",\"type\":\"bool\"}],\"name\":\"Suspend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"TokenBaseURISet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_AGENT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LEASING_AGENT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RECYCLING_AGENT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REVOCATION_AGENT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUSPENSION_AGENT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_contractRegistry\",\"outputs\":[{\"internalType\":\"contractContractRegistryInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"name\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"name\",\"type\":\"uint256\"}],\"name\":\"getDomainExpiry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"domainId\",\"type\":\"uint256\"}],\"name\":\"getNamespaceId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"namespaceId\",\"type\":\"uint256\"}],\"name\":\"getRoleForNamespace\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"name\",\"type\":\"uint256\"}],\"name\":\"isSuspended\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"suspended\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"namespaceId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"name\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leaseLength\",\"type\":\"uint256\"}],\"name\":\"recycle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registrant\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"namespaceId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"name\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leaseLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"constraintsData\",\"type\":\"bytes\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"namespaceId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"name\",\"type\":\"uint256\"}],\"name\":\"revoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"setBaseTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"setContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"namespaceId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"name\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"suspended\",\"type\":\"bool\"}],\"name\":\"suspend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DomainABI is the input ABI used to generate the binding from.
// Deprecated: Use DomainMetaData.ABI instead.
var DomainABI = DomainMetaData.ABI

// Domain is an auto generated Go binding around an Ethereum contract.
type Domain struct {
	DomainCaller     // Read-only binding to the contract
	DomainTransactor // Write-only binding to the contract
	DomainFilterer   // Log filterer for contract events
}

// DomainCaller is an auto generated read-only Go binding around an Ethereum contract.
type DomainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DomainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DomainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DomainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DomainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DomainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DomainSession struct {
	Contract     *Domain           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DomainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DomainCallerSession struct {
	Contract *DomainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DomainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DomainTransactorSession struct {
	Contract     *DomainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DomainRaw is an auto generated low-level Go binding around an Ethereum contract.
type DomainRaw struct {
	Contract *Domain // Generic contract binding to access the raw methods on
}

// DomainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DomainCallerRaw struct {
	Contract *DomainCaller // Generic read-only contract binding to access the raw methods on
}

// DomainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DomainTransactorRaw struct {
	Contract *DomainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDomain creates a new instance of Domain, bound to a specific deployed contract.
func NewDomain(address common.Address, backend bind.ContractBackend) (*Domain, error) {
	contract, err := bindDomain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Domain{DomainCaller: DomainCaller{contract: contract}, DomainTransactor: DomainTransactor{contract: contract}, DomainFilterer: DomainFilterer{contract: contract}}, nil
}

// NewDomainCaller creates a new read-only instance of Domain, bound to a specific deployed contract.
func NewDomainCaller(address common.Address, caller bind.ContractCaller) (*DomainCaller, error) {
	contract, err := bindDomain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DomainCaller{contract: contract}, nil
}

// NewDomainTransactor creates a new write-only instance of Domain, bound to a specific deployed contract.
func NewDomainTransactor(address common.Address, transactor bind.ContractTransactor) (*DomainTransactor, error) {
	contract, err := bindDomain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DomainTransactor{contract: contract}, nil
}

// NewDomainFilterer creates a new log filterer instance of Domain, bound to a specific deployed contract.
func NewDomainFilterer(address common.Address, filterer bind.ContractFilterer) (*DomainFilterer, error) {
	contract, err := bindDomain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DomainFilterer{contract: contract}, nil
}

// bindDomain binds a generic wrapper to an already deployed contract.
func bindDomain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DomainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Domain *DomainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Domain.Contract.DomainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Domain *DomainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Domain.Contract.DomainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Domain *DomainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Domain.Contract.DomainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Domain *DomainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Domain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Domain *DomainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Domain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Domain *DomainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Domain.Contract.contract.Transact(opts, method, params...)
}

// ADMINAGENT is a free data retrieval call binding the contract method 0xe5bdd1a8.
//
// Solidity: function ADMIN_AGENT() view returns(bytes32)
func (_Domain *DomainCaller) ADMINAGENT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Domain.contract.Call(opts, &out, "ADMIN_AGENT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINAGENT is a free data retrieval call binding the contract method 0xe5bdd1a8.
//
// Solidity: function ADMIN_AGENT() view returns(bytes32)
func (_Domain *DomainSession) ADMINAGENT() ([32]byte, error) {
	return _Domain.Contract.ADMINAGENT(&_Domain.CallOpts)
}

// ADMINAGENT is a free data retrieval call binding the contract method 0xe5bdd1a8.
//
// Solidity: function ADMIN_AGENT() view returns(bytes32)
func (_Domain *DomainCallerSession) ADMINAGENT() ([32]byte, error) {
	return _Domain.Contract.ADMINAGENT(&_Domain.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Domain *DomainCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Domain.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Domain *DomainSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Domain.Contract.DEFAULTADMINROLE(&_Domain.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Domain *DomainCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Domain.Contract.DEFAULTADMINROLE(&_Domain.CallOpts)
}

// LEASINGAGENT is a free data retrieval call binding the contract method 0x94b629fb.
//
// Solidity: function LEASING_AGENT() view returns(bytes32)
func (_Domain *DomainCaller) LEASINGAGENT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Domain.contract.Call(opts, &out, "LEASING_AGENT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LEASINGAGENT is a free data retrieval call binding the contract method 0x94b629fb.
//
// Solidity: function LEASING_AGENT() view returns(bytes32)
func (_Domain *DomainSession) LEASINGAGENT() ([32]byte, error) {
	return _Domain.Contract.LEASINGAGENT(&_Domain.CallOpts)
}

// LEASINGAGENT is a free data retrieval call binding the contract method 0x94b629fb.
//
// Solidity: function LEASING_AGENT() view returns(bytes32)
func (_Domain *DomainCallerSession) LEASINGAGENT() ([32]byte, error) {
	return _Domain.Contract.LEASINGAGENT(&_Domain.CallOpts)
}

// RECYCLINGAGENT is a free data retrieval call binding the contract method 0x441c0ca3.
//
// Solidity: function RECYCLING_AGENT() view returns(bytes32)
func (_Domain *DomainCaller) RECYCLINGAGENT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Domain.contract.Call(opts, &out, "RECYCLING_AGENT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RECYCLINGAGENT is a free data retrieval call binding the contract method 0x441c0ca3.
//
// Solidity: function RECYCLING_AGENT() view returns(bytes32)
func (_Domain *DomainSession) RECYCLINGAGENT() ([32]byte, error) {
	return _Domain.Contract.RECYCLINGAGENT(&_Domain.CallOpts)
}

// RECYCLINGAGENT is a free data retrieval call binding the contract method 0x441c0ca3.
//
// Solidity: function RECYCLING_AGENT() view returns(bytes32)
func (_Domain *DomainCallerSession) RECYCLINGAGENT() ([32]byte, error) {
	return _Domain.Contract.RECYCLINGAGENT(&_Domain.CallOpts)
}

// REVOCATIONAGENT is a free data retrieval call binding the contract method 0x0138b3bf.
//
// Solidity: function REVOCATION_AGENT() view returns(bytes32)
func (_Domain *DomainCaller) REVOCATIONAGENT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Domain.contract.Call(opts, &out, "REVOCATION_AGENT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REVOCATIONAGENT is a free data retrieval call binding the contract method 0x0138b3bf.
//
// Solidity: function REVOCATION_AGENT() view returns(bytes32)
func (_Domain *DomainSession) REVOCATIONAGENT() ([32]byte, error) {
	return _Domain.Contract.REVOCATIONAGENT(&_Domain.CallOpts)
}

// REVOCATIONAGENT is a free data retrieval call binding the contract method 0x0138b3bf.
//
// Solidity: function REVOCATION_AGENT() view returns(bytes32)
func (_Domain *DomainCallerSession) REVOCATIONAGENT() ([32]byte, error) {
	return _Domain.Contract.REVOCATIONAGENT(&_Domain.CallOpts)
}

// SUSPENSIONAGENT is a free data retrieval call binding the contract method 0x99d786ec.
//
// Solidity: function SUSPENSION_AGENT() view returns(bytes32)
func (_Domain *DomainCaller) SUSPENSIONAGENT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Domain.contract.Call(opts, &out, "SUSPENSION_AGENT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SUSPENSIONAGENT is a free data retrieval call binding the contract method 0x99d786ec.
//
// Solidity: function SUSPENSION_AGENT() view returns(bytes32)
func (_Domain *DomainSession) SUSPENSIONAGENT() ([32]byte, error) {
	return _Domain.Contract.SUSPENSIONAGENT(&_Domain.CallOpts)
}

// SUSPENSIONAGENT is a free data retrieval call binding the contract method 0x99d786ec.
//
// Solidity: function SUSPENSION_AGENT() view returns(bytes32)
func (_Domain *DomainCallerSession) SUSPENSIONAGENT() ([32]byte, error) {
	return _Domain.Contract.SUSPENSIONAGENT(&_Domain.CallOpts)
}

// ContractRegistry is a free data retrieval call binding the contract method 0xc1058b8f.
//
// Solidity: function _contractRegistry() view returns(address)
func (_Domain *DomainCaller) ContractRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Domain.contract.Call(opts, &out, "_contractRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractRegistry is a free data retrieval call binding the contract method 0xc1058b8f.
//
// Solidity: function _contractRegistry() view returns(address)
func (_Domain *DomainSession) ContractRegistry() (common.Address, error) {
	return _Domain.Contract.ContractRegistry(&_Domain.CallOpts)
}

// ContractRegistry is a free data retrieval call binding the contract method 0xc1058b8f.
//
// Solidity: function _contractRegistry() view returns(address)
func (_Domain *DomainCallerSession) ContractRegistry() (common.Address, error) {
	return _Domain.Contract.ContractRegistry(&_Domain.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Domain *DomainCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Domain.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Domain *DomainSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Domain.Contract.BalanceOf(&_Domain.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Domain *DomainCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Domain.Contract.BalanceOf(&_Domain.CallOpts, owner)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Domain *DomainCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Domain.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Domain *DomainSession) ContractURI() (string, error) {
	return _Domain.Contract.ContractURI(&_Domain.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Domain *DomainCallerSession) ContractURI() (string, error) {
	return _Domain.Contract.ContractURI(&_Domain.CallOpts)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 name) view returns(bool)
func (_Domain *DomainCaller) Exists(opts *bind.CallOpts, name *big.Int) (bool, error) {
	var out []interface{}
	err := _Domain.contract.Call(opts, &out, "exists", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 name) view returns(bool)
func (_Domain *DomainSession) Exists(name *big.Int) (bool, error) {
	return _Domain.Contract.Exists(&_Domain.CallOpts, name)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 name) view returns(bool)
func (_Domain *DomainCallerSession) Exists(name *big.Int) (bool, error) {
	return _Domain.Contract.Exists(&_Domain.CallOpts, name)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Domain *DomainCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Domain.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Domain *DomainSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Domain.Contract.GetApproved(&_Domain.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Domain *DomainCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Domain.Contract.GetApproved(&_Domain.CallOpts, tokenId)
}

// GetDomainExpiry is a free data retrieval call binding the contract method 0x9c9e3698.
//
// Solidity: function getDomainExpiry(uint256 name) view returns(uint256)
func (_Domain *DomainCaller) GetDomainExpiry(opts *bind.CallOpts, name *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Domain.contract.Call(opts, &out, "getDomainExpiry", name)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDomainExpiry is a free data retrieval call binding the contract method 0x9c9e3698.
//
// Solidity: function getDomainExpiry(uint256 name) view returns(uint256)
func (_Domain *DomainSession) GetDomainExpiry(name *big.Int) (*big.Int, error) {
	return _Domain.Contract.GetDomainExpiry(&_Domain.CallOpts, name)
}

// GetDomainExpiry is a free data retrieval call binding the contract method 0x9c9e3698.
//
// Solidity: function getDomainExpiry(uint256 name) view returns(uint256)
func (_Domain *DomainCallerSession) GetDomainExpiry(name *big.Int) (*big.Int, error) {
	return _Domain.Contract.GetDomainExpiry(&_Domain.CallOpts, name)
}

// GetNamespaceId is a free data retrieval call binding the contract method 0x881ea9d1.
//
// Solidity: function getNamespaceId(uint256 domainId) view returns(uint256)
func (_Domain *DomainCaller) GetNamespaceId(opts *bind.CallOpts, domainId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Domain.contract.Call(opts, &out, "getNamespaceId", domainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNamespaceId is a free data retrieval call binding the contract method 0x881ea9d1.
//
// Solidity: function getNamespaceId(uint256 domainId) view returns(uint256)
func (_Domain *DomainSession) GetNamespaceId(domainId *big.Int) (*big.Int, error) {
	return _Domain.Contract.GetNamespaceId(&_Domain.CallOpts, domainId)
}

// GetNamespaceId is a free data retrieval call binding the contract method 0x881ea9d1.
//
// Solidity: function getNamespaceId(uint256 domainId) view returns(uint256)
func (_Domain *DomainCallerSession) GetNamespaceId(domainId *big.Int) (*big.Int, error) {
	return _Domain.Contract.GetNamespaceId(&_Domain.CallOpts, domainId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Domain *DomainCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Domain.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Domain *DomainSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Domain.Contract.GetRoleAdmin(&_Domain.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Domain *DomainCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Domain.Contract.GetRoleAdmin(&_Domain.CallOpts, role)
}

// GetRoleForNamespace is a free data retrieval call binding the contract method 0x831ff462.
//
// Solidity: function getRoleForNamespace(bytes32 role, uint256 namespaceId) pure returns(bytes32)
func (_Domain *DomainCaller) GetRoleForNamespace(opts *bind.CallOpts, role [32]byte, namespaceId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Domain.contract.Call(opts, &out, "getRoleForNamespace", role, namespaceId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleForNamespace is a free data retrieval call binding the contract method 0x831ff462.
//
// Solidity: function getRoleForNamespace(bytes32 role, uint256 namespaceId) pure returns(bytes32)
func (_Domain *DomainSession) GetRoleForNamespace(role [32]byte, namespaceId *big.Int) ([32]byte, error) {
	return _Domain.Contract.GetRoleForNamespace(&_Domain.CallOpts, role, namespaceId)
}

// GetRoleForNamespace is a free data retrieval call binding the contract method 0x831ff462.
//
// Solidity: function getRoleForNamespace(bytes32 role, uint256 namespaceId) pure returns(bytes32)
func (_Domain *DomainCallerSession) GetRoleForNamespace(role [32]byte, namespaceId *big.Int) ([32]byte, error) {
	return _Domain.Contract.GetRoleForNamespace(&_Domain.CallOpts, role, namespaceId)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Domain *DomainCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Domain.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Domain *DomainSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Domain.Contract.HasRole(&_Domain.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Domain *DomainCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Domain.Contract.HasRole(&_Domain.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Domain *DomainCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Domain.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Domain *DomainSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Domain.Contract.IsApprovedForAll(&_Domain.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Domain *DomainCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Domain.Contract.IsApprovedForAll(&_Domain.CallOpts, owner, operator)
}

// IsSuspended is a free data retrieval call binding the contract method 0x0bae1bd0.
//
// Solidity: function isSuspended(uint256 name) view returns(bool suspended)
func (_Domain *DomainCaller) IsSuspended(opts *bind.CallOpts, name *big.Int) (bool, error) {
	var out []interface{}
	err := _Domain.contract.Call(opts, &out, "isSuspended", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSuspended is a free data retrieval call binding the contract method 0x0bae1bd0.
//
// Solidity: function isSuspended(uint256 name) view returns(bool suspended)
func (_Domain *DomainSession) IsSuspended(name *big.Int) (bool, error) {
	return _Domain.Contract.IsSuspended(&_Domain.CallOpts, name)
}

// IsSuspended is a free data retrieval call binding the contract method 0x0bae1bd0.
//
// Solidity: function isSuspended(uint256 name) view returns(bool suspended)
func (_Domain *DomainCallerSession) IsSuspended(name *big.Int) (bool, error) {
	return _Domain.Contract.IsSuspended(&_Domain.CallOpts, name)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Domain *DomainCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Domain.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Domain *DomainSession) Name() (string, error) {
	return _Domain.Contract.Name(&_Domain.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Domain *DomainCallerSession) Name() (string, error) {
	return _Domain.Contract.Name(&_Domain.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Domain *DomainCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Domain.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Domain *DomainSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Domain.Contract.OwnerOf(&_Domain.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Domain *DomainCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Domain.Contract.OwnerOf(&_Domain.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Domain *DomainCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Domain.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Domain *DomainSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Domain.Contract.SupportsInterface(&_Domain.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Domain *DomainCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Domain.Contract.SupportsInterface(&_Domain.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Domain *DomainCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Domain.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Domain *DomainSession) Symbol() (string, error) {
	return _Domain.Contract.Symbol(&_Domain.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Domain *DomainCallerSession) Symbol() (string, error) {
	return _Domain.Contract.Symbol(&_Domain.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Domain *DomainCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Domain.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Domain *DomainSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Domain.Contract.TokenByIndex(&_Domain.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Domain *DomainCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Domain.Contract.TokenByIndex(&_Domain.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Domain *DomainCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Domain.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Domain *DomainSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Domain.Contract.TokenOfOwnerByIndex(&_Domain.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Domain *DomainCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Domain.Contract.TokenOfOwnerByIndex(&_Domain.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Domain *DomainCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Domain.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Domain *DomainSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Domain.Contract.TokenURI(&_Domain.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Domain *DomainCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Domain.Contract.TokenURI(&_Domain.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Domain *DomainCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Domain.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Domain *DomainSession) TotalSupply() (*big.Int, error) {
	return _Domain.Contract.TotalSupply(&_Domain.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Domain *DomainCallerSession) TotalSupply() (*big.Int, error) {
	return _Domain.Contract.TotalSupply(&_Domain.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Domain *DomainTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Domain.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Domain *DomainSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Domain.Contract.Approve(&_Domain.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Domain *DomainTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Domain.Contract.Approve(&_Domain.TransactOpts, to, tokenId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Domain *DomainTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Domain.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Domain *DomainSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Domain.Contract.GrantRole(&_Domain.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Domain *DomainTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Domain.Contract.GrantRole(&_Domain.TransactOpts, role, account)
}

// Recycle is a paid mutator transaction binding the contract method 0xb012290b.
//
// Solidity: function recycle(address to, uint256 namespaceId, uint256 name, uint256 leaseLength) returns()
func (_Domain *DomainTransactor) Recycle(opts *bind.TransactOpts, to common.Address, namespaceId *big.Int, name *big.Int, leaseLength *big.Int) (*types.Transaction, error) {
	return _Domain.contract.Transact(opts, "recycle", to, namespaceId, name, leaseLength)
}

// Recycle is a paid mutator transaction binding the contract method 0xb012290b.
//
// Solidity: function recycle(address to, uint256 namespaceId, uint256 name, uint256 leaseLength) returns()
func (_Domain *DomainSession) Recycle(to common.Address, namespaceId *big.Int, name *big.Int, leaseLength *big.Int) (*types.Transaction, error) {
	return _Domain.Contract.Recycle(&_Domain.TransactOpts, to, namespaceId, name, leaseLength)
}

// Recycle is a paid mutator transaction binding the contract method 0xb012290b.
//
// Solidity: function recycle(address to, uint256 namespaceId, uint256 name, uint256 leaseLength) returns()
func (_Domain *DomainTransactorSession) Recycle(to common.Address, namespaceId *big.Int, name *big.Int, leaseLength *big.Int) (*types.Transaction, error) {
	return _Domain.Contract.Recycle(&_Domain.TransactOpts, to, namespaceId, name, leaseLength)
}

// Register is a paid mutator transaction binding the contract method 0xb3e482fa.
//
// Solidity: function register(address registrant, uint256 namespaceId, uint256 name, uint256 leaseLength, bytes constraintsData) returns()
func (_Domain *DomainTransactor) Register(opts *bind.TransactOpts, registrant common.Address, namespaceId *big.Int, name *big.Int, leaseLength *big.Int, constraintsData []byte) (*types.Transaction, error) {
	return _Domain.contract.Transact(opts, "register", registrant, namespaceId, name, leaseLength, constraintsData)
}

// Register is a paid mutator transaction binding the contract method 0xb3e482fa.
//
// Solidity: function register(address registrant, uint256 namespaceId, uint256 name, uint256 leaseLength, bytes constraintsData) returns()
func (_Domain *DomainSession) Register(registrant common.Address, namespaceId *big.Int, name *big.Int, leaseLength *big.Int, constraintsData []byte) (*types.Transaction, error) {
	return _Domain.Contract.Register(&_Domain.TransactOpts, registrant, namespaceId, name, leaseLength, constraintsData)
}

// Register is a paid mutator transaction binding the contract method 0xb3e482fa.
//
// Solidity: function register(address registrant, uint256 namespaceId, uint256 name, uint256 leaseLength, bytes constraintsData) returns()
func (_Domain *DomainTransactorSession) Register(registrant common.Address, namespaceId *big.Int, name *big.Int, leaseLength *big.Int, constraintsData []byte) (*types.Transaction, error) {
	return _Domain.Contract.Register(&_Domain.TransactOpts, registrant, namespaceId, name, leaseLength, constraintsData)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Domain *DomainTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Domain.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Domain *DomainSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Domain.Contract.RenounceRole(&_Domain.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Domain *DomainTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Domain.Contract.RenounceRole(&_Domain.TransactOpts, role, account)
}

// Revoke is a paid mutator transaction binding the contract method 0x990e4c51.
//
// Solidity: function revoke(address to, uint256 namespaceId, uint256 name) returns()
func (_Domain *DomainTransactor) Revoke(opts *bind.TransactOpts, to common.Address, namespaceId *big.Int, name *big.Int) (*types.Transaction, error) {
	return _Domain.contract.Transact(opts, "revoke", to, namespaceId, name)
}

// Revoke is a paid mutator transaction binding the contract method 0x990e4c51.
//
// Solidity: function revoke(address to, uint256 namespaceId, uint256 name) returns()
func (_Domain *DomainSession) Revoke(to common.Address, namespaceId *big.Int, name *big.Int) (*types.Transaction, error) {
	return _Domain.Contract.Revoke(&_Domain.TransactOpts, to, namespaceId, name)
}

// Revoke is a paid mutator transaction binding the contract method 0x990e4c51.
//
// Solidity: function revoke(address to, uint256 namespaceId, uint256 name) returns()
func (_Domain *DomainTransactorSession) Revoke(to common.Address, namespaceId *big.Int, name *big.Int) (*types.Transaction, error) {
	return _Domain.Contract.Revoke(&_Domain.TransactOpts, to, namespaceId, name)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Domain *DomainTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Domain.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Domain *DomainSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Domain.Contract.RevokeRole(&_Domain.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Domain *DomainTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Domain.Contract.RevokeRole(&_Domain.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Domain *DomainTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Domain.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Domain *DomainSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Domain.Contract.SafeTransferFrom(&_Domain.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Domain *DomainTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Domain.Contract.SafeTransferFrom(&_Domain.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Domain *DomainTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Domain.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Domain *DomainSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Domain.Contract.SafeTransferFrom0(&_Domain.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Domain *DomainTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Domain.Contract.SafeTransferFrom0(&_Domain.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Domain *DomainTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Domain.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Domain *DomainSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Domain.Contract.SetApprovalForAll(&_Domain.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Domain *DomainTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Domain.Contract.SetApprovalForAll(&_Domain.TransactOpts, operator, approved)
}

// SetBaseTokenURI is a paid mutator transaction binding the contract method 0x30176e13.
//
// Solidity: function setBaseTokenURI(string uri) returns()
func (_Domain *DomainTransactor) SetBaseTokenURI(opts *bind.TransactOpts, uri string) (*types.Transaction, error) {
	return _Domain.contract.Transact(opts, "setBaseTokenURI", uri)
}

// SetBaseTokenURI is a paid mutator transaction binding the contract method 0x30176e13.
//
// Solidity: function setBaseTokenURI(string uri) returns()
func (_Domain *DomainSession) SetBaseTokenURI(uri string) (*types.Transaction, error) {
	return _Domain.Contract.SetBaseTokenURI(&_Domain.TransactOpts, uri)
}

// SetBaseTokenURI is a paid mutator transaction binding the contract method 0x30176e13.
//
// Solidity: function setBaseTokenURI(string uri) returns()
func (_Domain *DomainTransactorSession) SetBaseTokenURI(uri string) (*types.Transaction, error) {
	return _Domain.Contract.SetBaseTokenURI(&_Domain.TransactOpts, uri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string uri) returns()
func (_Domain *DomainTransactor) SetContractURI(opts *bind.TransactOpts, uri string) (*types.Transaction, error) {
	return _Domain.contract.Transact(opts, "setContractURI", uri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string uri) returns()
func (_Domain *DomainSession) SetContractURI(uri string) (*types.Transaction, error) {
	return _Domain.Contract.SetContractURI(&_Domain.TransactOpts, uri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string uri) returns()
func (_Domain *DomainTransactorSession) SetContractURI(uri string) (*types.Transaction, error) {
	return _Domain.Contract.SetContractURI(&_Domain.TransactOpts, uri)
}

// Suspend is a paid mutator transaction binding the contract method 0x2d8fbcad.
//
// Solidity: function suspend(uint256 namespaceId, uint256 name, bool suspended) returns()
func (_Domain *DomainTransactor) Suspend(opts *bind.TransactOpts, namespaceId *big.Int, name *big.Int, suspended bool) (*types.Transaction, error) {
	return _Domain.contract.Transact(opts, "suspend", namespaceId, name, suspended)
}

// Suspend is a paid mutator transaction binding the contract method 0x2d8fbcad.
//
// Solidity: function suspend(uint256 namespaceId, uint256 name, bool suspended) returns()
func (_Domain *DomainSession) Suspend(namespaceId *big.Int, name *big.Int, suspended bool) (*types.Transaction, error) {
	return _Domain.Contract.Suspend(&_Domain.TransactOpts, namespaceId, name, suspended)
}

// Suspend is a paid mutator transaction binding the contract method 0x2d8fbcad.
//
// Solidity: function suspend(uint256 namespaceId, uint256 name, bool suspended) returns()
func (_Domain *DomainTransactorSession) Suspend(namespaceId *big.Int, name *big.Int, suspended bool) (*types.Transaction, error) {
	return _Domain.Contract.Suspend(&_Domain.TransactOpts, namespaceId, name, suspended)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Domain *DomainTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Domain.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Domain *DomainSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Domain.Contract.TransferFrom(&_Domain.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Domain *DomainTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Domain.Contract.TransferFrom(&_Domain.TransactOpts, from, to, tokenId)
}

// DomainApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Domain contract.
type DomainApprovalIterator struct {
	Event *DomainApproval // Event containing the contract specifics and raw log

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
func (it *DomainApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DomainApproval)
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
		it.Event = new(DomainApproval)
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
func (it *DomainApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DomainApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DomainApproval represents a Approval event raised by the Domain contract.
type DomainApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Domain *DomainFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*DomainApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Domain.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DomainApprovalIterator{contract: _Domain.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Domain *DomainFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DomainApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Domain.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DomainApproval)
				if err := _Domain.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Domain *DomainFilterer) ParseApproval(log types.Log) (*DomainApproval, error) {
	event := new(DomainApproval)
	if err := _Domain.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DomainApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Domain contract.
type DomainApprovalForAllIterator struct {
	Event *DomainApprovalForAll // Event containing the contract specifics and raw log

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
func (it *DomainApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DomainApprovalForAll)
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
		it.Event = new(DomainApprovalForAll)
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
func (it *DomainApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DomainApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DomainApprovalForAll represents a ApprovalForAll event raised by the Domain contract.
type DomainApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Domain *DomainFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*DomainApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Domain.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &DomainApprovalForAllIterator{contract: _Domain.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Domain *DomainFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *DomainApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Domain.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DomainApprovalForAll)
				if err := _Domain.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Domain *DomainFilterer) ParseApprovalForAll(log types.Log) (*DomainApprovalForAll, error) {
	event := new(DomainApprovalForAll)
	if err := _Domain.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DomainContractURISetIterator is returned from FilterContractURISet and is used to iterate over the raw logs and unpacked data for ContractURISet events raised by the Domain contract.
type DomainContractURISetIterator struct {
	Event *DomainContractURISet // Event containing the contract specifics and raw log

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
func (it *DomainContractURISetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DomainContractURISet)
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
		it.Event = new(DomainContractURISet)
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
func (it *DomainContractURISetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DomainContractURISetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DomainContractURISet represents a ContractURISet event raised by the Domain contract.
type DomainContractURISet struct {
	Uri string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterContractURISet is a free log retrieval operation binding the contract event 0xaf497693a87db12ca89131a31edbb3db4bb5702dfb284e8ae7427d185f09112d.
//
// Solidity: event ContractURISet(string uri)
func (_Domain *DomainFilterer) FilterContractURISet(opts *bind.FilterOpts) (*DomainContractURISetIterator, error) {

	logs, sub, err := _Domain.contract.FilterLogs(opts, "ContractURISet")
	if err != nil {
		return nil, err
	}
	return &DomainContractURISetIterator{contract: _Domain.contract, event: "ContractURISet", logs: logs, sub: sub}, nil
}

// WatchContractURISet is a free log subscription operation binding the contract event 0xaf497693a87db12ca89131a31edbb3db4bb5702dfb284e8ae7427d185f09112d.
//
// Solidity: event ContractURISet(string uri)
func (_Domain *DomainFilterer) WatchContractURISet(opts *bind.WatchOpts, sink chan<- *DomainContractURISet) (event.Subscription, error) {

	logs, sub, err := _Domain.contract.WatchLogs(opts, "ContractURISet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DomainContractURISet)
				if err := _Domain.contract.UnpackLog(event, "ContractURISet", log); err != nil {
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

// ParseContractURISet is a log parse operation binding the contract event 0xaf497693a87db12ca89131a31edbb3db4bb5702dfb284e8ae7427d185f09112d.
//
// Solidity: event ContractURISet(string uri)
func (_Domain *DomainFilterer) ParseContractURISet(log types.Log) (*DomainContractURISet, error) {
	event := new(DomainContractURISet)
	if err := _Domain.contract.UnpackLog(event, "ContractURISet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DomainRecycleIterator is returned from FilterRecycle and is used to iterate over the raw logs and unpacked data for Recycle events raised by the Domain contract.
type DomainRecycleIterator struct {
	Event *DomainRecycle // Event containing the contract specifics and raw log

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
func (it *DomainRecycleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DomainRecycle)
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
		it.Event = new(DomainRecycle)
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
func (it *DomainRecycleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DomainRecycleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DomainRecycle represents a Recycle event raised by the Domain contract.
type DomainRecycle struct {
	Agent       common.Address
	Registrant  common.Address
	Name        *big.Int
	LeaseLength *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRecycle is a free log retrieval operation binding the contract event 0x14295458e7a580d6a197fb2ba945cb8e733fa6a9642a29909cdef9301bc3fa3e.
//
// Solidity: event Recycle(address agent, address indexed registrant, uint256 indexed name, uint256 leaseLength)
func (_Domain *DomainFilterer) FilterRecycle(opts *bind.FilterOpts, registrant []common.Address, name []*big.Int) (*DomainRecycleIterator, error) {

	var registrantRule []interface{}
	for _, registrantItem := range registrant {
		registrantRule = append(registrantRule, registrantItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _Domain.contract.FilterLogs(opts, "Recycle", registrantRule, nameRule)
	if err != nil {
		return nil, err
	}
	return &DomainRecycleIterator{contract: _Domain.contract, event: "Recycle", logs: logs, sub: sub}, nil
}

// WatchRecycle is a free log subscription operation binding the contract event 0x14295458e7a580d6a197fb2ba945cb8e733fa6a9642a29909cdef9301bc3fa3e.
//
// Solidity: event Recycle(address agent, address indexed registrant, uint256 indexed name, uint256 leaseLength)
func (_Domain *DomainFilterer) WatchRecycle(opts *bind.WatchOpts, sink chan<- *DomainRecycle, registrant []common.Address, name []*big.Int) (event.Subscription, error) {

	var registrantRule []interface{}
	for _, registrantItem := range registrant {
		registrantRule = append(registrantRule, registrantItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _Domain.contract.WatchLogs(opts, "Recycle", registrantRule, nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DomainRecycle)
				if err := _Domain.contract.UnpackLog(event, "Recycle", log); err != nil {
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

// ParseRecycle is a log parse operation binding the contract event 0x14295458e7a580d6a197fb2ba945cb8e733fa6a9642a29909cdef9301bc3fa3e.
//
// Solidity: event Recycle(address agent, address indexed registrant, uint256 indexed name, uint256 leaseLength)
func (_Domain *DomainFilterer) ParseRecycle(log types.Log) (*DomainRecycle, error) {
	event := new(DomainRecycle)
	if err := _Domain.contract.UnpackLog(event, "Recycle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DomainRegisterIterator is returned from FilterRegister and is used to iterate over the raw logs and unpacked data for Register events raised by the Domain contract.
type DomainRegisterIterator struct {
	Event *DomainRegister // Event containing the contract specifics and raw log

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
func (it *DomainRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DomainRegister)
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
		it.Event = new(DomainRegister)
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
func (it *DomainRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DomainRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DomainRegister represents a Register event raised by the Domain contract.
type DomainRegister struct {
	Agent       common.Address
	Registrant  common.Address
	Name        *big.Int
	LeaseLength *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRegister is a free log retrieval operation binding the contract event 0x2e99fd66dad9611ac11d08dcbfdb417ac840074a8659831e928c70a92a6215e7.
//
// Solidity: event Register(address agent, address indexed registrant, uint256 indexed name, uint256 leaseLength)
func (_Domain *DomainFilterer) FilterRegister(opts *bind.FilterOpts, registrant []common.Address, name []*big.Int) (*DomainRegisterIterator, error) {

	var registrantRule []interface{}
	for _, registrantItem := range registrant {
		registrantRule = append(registrantRule, registrantItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _Domain.contract.FilterLogs(opts, "Register", registrantRule, nameRule)
	if err != nil {
		return nil, err
	}
	return &DomainRegisterIterator{contract: _Domain.contract, event: "Register", logs: logs, sub: sub}, nil
}

// WatchRegister is a free log subscription operation binding the contract event 0x2e99fd66dad9611ac11d08dcbfdb417ac840074a8659831e928c70a92a6215e7.
//
// Solidity: event Register(address agent, address indexed registrant, uint256 indexed name, uint256 leaseLength)
func (_Domain *DomainFilterer) WatchRegister(opts *bind.WatchOpts, sink chan<- *DomainRegister, registrant []common.Address, name []*big.Int) (event.Subscription, error) {

	var registrantRule []interface{}
	for _, registrantItem := range registrant {
		registrantRule = append(registrantRule, registrantItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _Domain.contract.WatchLogs(opts, "Register", registrantRule, nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DomainRegister)
				if err := _Domain.contract.UnpackLog(event, "Register", log); err != nil {
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

// ParseRegister is a log parse operation binding the contract event 0x2e99fd66dad9611ac11d08dcbfdb417ac840074a8659831e928c70a92a6215e7.
//
// Solidity: event Register(address agent, address indexed registrant, uint256 indexed name, uint256 leaseLength)
func (_Domain *DomainFilterer) ParseRegister(log types.Log) (*DomainRegister, error) {
	event := new(DomainRegister)
	if err := _Domain.contract.UnpackLog(event, "Register", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DomainRevokeIterator is returned from FilterRevoke and is used to iterate over the raw logs and unpacked data for Revoke events raised by the Domain contract.
type DomainRevokeIterator struct {
	Event *DomainRevoke // Event containing the contract specifics and raw log

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
func (it *DomainRevokeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DomainRevoke)
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
		it.Event = new(DomainRevoke)
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
func (it *DomainRevokeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DomainRevokeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DomainRevoke represents a Revoke event raised by the Domain contract.
type DomainRevoke struct {
	Agent  common.Address
	Holder common.Address
	Name   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRevoke is a free log retrieval operation binding the contract event 0xb698e31a2abee5824d0d7bcfd2339aead7f9e9ae413fba50bf554ff3fa470b7b.
//
// Solidity: event Revoke(address agent, address indexed holder, uint256 indexed name)
func (_Domain *DomainFilterer) FilterRevoke(opts *bind.FilterOpts, holder []common.Address, name []*big.Int) (*DomainRevokeIterator, error) {

	var holderRule []interface{}
	for _, holderItem := range holder {
		holderRule = append(holderRule, holderItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _Domain.contract.FilterLogs(opts, "Revoke", holderRule, nameRule)
	if err != nil {
		return nil, err
	}
	return &DomainRevokeIterator{contract: _Domain.contract, event: "Revoke", logs: logs, sub: sub}, nil
}

// WatchRevoke is a free log subscription operation binding the contract event 0xb698e31a2abee5824d0d7bcfd2339aead7f9e9ae413fba50bf554ff3fa470b7b.
//
// Solidity: event Revoke(address agent, address indexed holder, uint256 indexed name)
func (_Domain *DomainFilterer) WatchRevoke(opts *bind.WatchOpts, sink chan<- *DomainRevoke, holder []common.Address, name []*big.Int) (event.Subscription, error) {

	var holderRule []interface{}
	for _, holderItem := range holder {
		holderRule = append(holderRule, holderItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _Domain.contract.WatchLogs(opts, "Revoke", holderRule, nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DomainRevoke)
				if err := _Domain.contract.UnpackLog(event, "Revoke", log); err != nil {
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

// ParseRevoke is a log parse operation binding the contract event 0xb698e31a2abee5824d0d7bcfd2339aead7f9e9ae413fba50bf554ff3fa470b7b.
//
// Solidity: event Revoke(address agent, address indexed holder, uint256 indexed name)
func (_Domain *DomainFilterer) ParseRevoke(log types.Log) (*DomainRevoke, error) {
	event := new(DomainRevoke)
	if err := _Domain.contract.UnpackLog(event, "Revoke", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DomainRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Domain contract.
type DomainRoleAdminChangedIterator struct {
	Event *DomainRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *DomainRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DomainRoleAdminChanged)
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
		it.Event = new(DomainRoleAdminChanged)
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
func (it *DomainRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DomainRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DomainRoleAdminChanged represents a RoleAdminChanged event raised by the Domain contract.
type DomainRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Domain *DomainFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*DomainRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Domain.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &DomainRoleAdminChangedIterator{contract: _Domain.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Domain *DomainFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *DomainRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Domain.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DomainRoleAdminChanged)
				if err := _Domain.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Domain *DomainFilterer) ParseRoleAdminChanged(log types.Log) (*DomainRoleAdminChanged, error) {
	event := new(DomainRoleAdminChanged)
	if err := _Domain.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DomainRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Domain contract.
type DomainRoleGrantedIterator struct {
	Event *DomainRoleGranted // Event containing the contract specifics and raw log

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
func (it *DomainRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DomainRoleGranted)
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
		it.Event = new(DomainRoleGranted)
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
func (it *DomainRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DomainRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DomainRoleGranted represents a RoleGranted event raised by the Domain contract.
type DomainRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Domain *DomainFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DomainRoleGrantedIterator, error) {

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

	logs, sub, err := _Domain.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DomainRoleGrantedIterator{contract: _Domain.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Domain *DomainFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *DomainRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Domain.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DomainRoleGranted)
				if err := _Domain.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Domain *DomainFilterer) ParseRoleGranted(log types.Log) (*DomainRoleGranted, error) {
	event := new(DomainRoleGranted)
	if err := _Domain.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DomainRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Domain contract.
type DomainRoleRevokedIterator struct {
	Event *DomainRoleRevoked // Event containing the contract specifics and raw log

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
func (it *DomainRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DomainRoleRevoked)
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
		it.Event = new(DomainRoleRevoked)
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
func (it *DomainRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DomainRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DomainRoleRevoked represents a RoleRevoked event raised by the Domain contract.
type DomainRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Domain *DomainFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DomainRoleRevokedIterator, error) {

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

	logs, sub, err := _Domain.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DomainRoleRevokedIterator{contract: _Domain.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Domain *DomainFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *DomainRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Domain.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DomainRoleRevoked)
				if err := _Domain.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Domain *DomainFilterer) ParseRoleRevoked(log types.Log) (*DomainRoleRevoked, error) {
	event := new(DomainRoleRevoked)
	if err := _Domain.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DomainSuspendIterator is returned from FilterSuspend and is used to iterate over the raw logs and unpacked data for Suspend events raised by the Domain contract.
type DomainSuspendIterator struct {
	Event *DomainSuspend // Event containing the contract specifics and raw log

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
func (it *DomainSuspendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DomainSuspend)
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
		it.Event = new(DomainSuspend)
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
func (it *DomainSuspendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DomainSuspendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DomainSuspend represents a Suspend event raised by the Domain contract.
type DomainSuspend struct {
	Agent     common.Address
	Name      *big.Int
	Suspended bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSuspend is a free log retrieval operation binding the contract event 0xcd26ab8516a4725be7b2ea7ae931dfb377ac5cf653a5cecafe7b030e4007dbae.
//
// Solidity: event Suspend(address agent, uint256 indexed name, bool suspended)
func (_Domain *DomainFilterer) FilterSuspend(opts *bind.FilterOpts, name []*big.Int) (*DomainSuspendIterator, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _Domain.contract.FilterLogs(opts, "Suspend", nameRule)
	if err != nil {
		return nil, err
	}
	return &DomainSuspendIterator{contract: _Domain.contract, event: "Suspend", logs: logs, sub: sub}, nil
}

// WatchSuspend is a free log subscription operation binding the contract event 0xcd26ab8516a4725be7b2ea7ae931dfb377ac5cf653a5cecafe7b030e4007dbae.
//
// Solidity: event Suspend(address agent, uint256 indexed name, bool suspended)
func (_Domain *DomainFilterer) WatchSuspend(opts *bind.WatchOpts, sink chan<- *DomainSuspend, name []*big.Int) (event.Subscription, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _Domain.contract.WatchLogs(opts, "Suspend", nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DomainSuspend)
				if err := _Domain.contract.UnpackLog(event, "Suspend", log); err != nil {
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

// ParseSuspend is a log parse operation binding the contract event 0xcd26ab8516a4725be7b2ea7ae931dfb377ac5cf653a5cecafe7b030e4007dbae.
//
// Solidity: event Suspend(address agent, uint256 indexed name, bool suspended)
func (_Domain *DomainFilterer) ParseSuspend(log types.Log) (*DomainSuspend, error) {
	event := new(DomainSuspend)
	if err := _Domain.contract.UnpackLog(event, "Suspend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DomainTokenBaseURISetIterator is returned from FilterTokenBaseURISet and is used to iterate over the raw logs and unpacked data for TokenBaseURISet events raised by the Domain contract.
type DomainTokenBaseURISetIterator struct {
	Event *DomainTokenBaseURISet // Event containing the contract specifics and raw log

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
func (it *DomainTokenBaseURISetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DomainTokenBaseURISet)
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
		it.Event = new(DomainTokenBaseURISet)
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
func (it *DomainTokenBaseURISetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DomainTokenBaseURISetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DomainTokenBaseURISet represents a TokenBaseURISet event raised by the Domain contract.
type DomainTokenBaseURISet struct {
	Uri string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTokenBaseURISet is a free log retrieval operation binding the contract event 0x6ae07e2a08f067a8d5ece04d32173e8ecaac9a2273ca0e3db850c5460a5d6160.
//
// Solidity: event TokenBaseURISet(string uri)
func (_Domain *DomainFilterer) FilterTokenBaseURISet(opts *bind.FilterOpts) (*DomainTokenBaseURISetIterator, error) {

	logs, sub, err := _Domain.contract.FilterLogs(opts, "TokenBaseURISet")
	if err != nil {
		return nil, err
	}
	return &DomainTokenBaseURISetIterator{contract: _Domain.contract, event: "TokenBaseURISet", logs: logs, sub: sub}, nil
}

// WatchTokenBaseURISet is a free log subscription operation binding the contract event 0x6ae07e2a08f067a8d5ece04d32173e8ecaac9a2273ca0e3db850c5460a5d6160.
//
// Solidity: event TokenBaseURISet(string uri)
func (_Domain *DomainFilterer) WatchTokenBaseURISet(opts *bind.WatchOpts, sink chan<- *DomainTokenBaseURISet) (event.Subscription, error) {

	logs, sub, err := _Domain.contract.WatchLogs(opts, "TokenBaseURISet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DomainTokenBaseURISet)
				if err := _Domain.contract.UnpackLog(event, "TokenBaseURISet", log); err != nil {
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

// ParseTokenBaseURISet is a log parse operation binding the contract event 0x6ae07e2a08f067a8d5ece04d32173e8ecaac9a2273ca0e3db850c5460a5d6160.
//
// Solidity: event TokenBaseURISet(string uri)
func (_Domain *DomainFilterer) ParseTokenBaseURISet(log types.Log) (*DomainTokenBaseURISet, error) {
	event := new(DomainTokenBaseURISet)
	if err := _Domain.contract.UnpackLog(event, "TokenBaseURISet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DomainTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Domain contract.
type DomainTransferIterator struct {
	Event *DomainTransfer // Event containing the contract specifics and raw log

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
func (it *DomainTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DomainTransfer)
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
		it.Event = new(DomainTransfer)
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
func (it *DomainTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DomainTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DomainTransfer represents a Transfer event raised by the Domain contract.
type DomainTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Domain *DomainFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*DomainTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Domain.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DomainTransferIterator{contract: _Domain.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Domain *DomainFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DomainTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Domain.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DomainTransfer)
				if err := _Domain.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Domain *DomainFilterer) ParseTransfer(log types.Log) (*DomainTransfer, error) {
	event := new(DomainTransfer)
	if err := _Domain.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
