// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PricingOracleV1

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

// PricingOracleV1MetaData contains all meta data concerning the PricingOracleV1 contract.
var PricingOracleV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractVerifierInterface\",\"name\":\"verifier\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_verifier\",\"outputs\":[{\"internalType\":\"contractVerifierInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"convertWeiToUsdCents\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"usdCents\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"name\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"getPriceForName\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceCentsUsd\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceFeed\",\"outputs\":[{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PricingOracleV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use PricingOracleV1MetaData.ABI instead.
var PricingOracleV1ABI = PricingOracleV1MetaData.ABI

// PricingOracleV1 is an auto generated Go binding around an Ethereum contract.
type PricingOracleV1 struct {
	PricingOracleV1Caller     // Read-only binding to the contract
	PricingOracleV1Transactor // Write-only binding to the contract
	PricingOracleV1Filterer   // Log filterer for contract events
}

// PricingOracleV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type PricingOracleV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PricingOracleV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type PricingOracleV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PricingOracleV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PricingOracleV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PricingOracleV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PricingOracleV1Session struct {
	Contract     *PricingOracleV1  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PricingOracleV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PricingOracleV1CallerSession struct {
	Contract *PricingOracleV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PricingOracleV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PricingOracleV1TransactorSession struct {
	Contract     *PricingOracleV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PricingOracleV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type PricingOracleV1Raw struct {
	Contract *PricingOracleV1 // Generic contract binding to access the raw methods on
}

// PricingOracleV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PricingOracleV1CallerRaw struct {
	Contract *PricingOracleV1Caller // Generic read-only contract binding to access the raw methods on
}

// PricingOracleV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PricingOracleV1TransactorRaw struct {
	Contract *PricingOracleV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPricingOracleV1 creates a new instance of PricingOracleV1, bound to a specific deployed contract.
func NewPricingOracleV1(address common.Address, backend bind.ContractBackend) (*PricingOracleV1, error) {
	contract, err := bindPricingOracleV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PricingOracleV1{PricingOracleV1Caller: PricingOracleV1Caller{contract: contract}, PricingOracleV1Transactor: PricingOracleV1Transactor{contract: contract}, PricingOracleV1Filterer: PricingOracleV1Filterer{contract: contract}}, nil
}

// NewPricingOracleV1Caller creates a new read-only instance of PricingOracleV1, bound to a specific deployed contract.
func NewPricingOracleV1Caller(address common.Address, caller bind.ContractCaller) (*PricingOracleV1Caller, error) {
	contract, err := bindPricingOracleV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PricingOracleV1Caller{contract: contract}, nil
}

// NewPricingOracleV1Transactor creates a new write-only instance of PricingOracleV1, bound to a specific deployed contract.
func NewPricingOracleV1Transactor(address common.Address, transactor bind.ContractTransactor) (*PricingOracleV1Transactor, error) {
	contract, err := bindPricingOracleV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PricingOracleV1Transactor{contract: contract}, nil
}

// NewPricingOracleV1Filterer creates a new log filterer instance of PricingOracleV1, bound to a specific deployed contract.
func NewPricingOracleV1Filterer(address common.Address, filterer bind.ContractFilterer) (*PricingOracleV1Filterer, error) {
	contract, err := bindPricingOracleV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PricingOracleV1Filterer{contract: contract}, nil
}

// bindPricingOracleV1 binds a generic wrapper to an already deployed contract.
func bindPricingOracleV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PricingOracleV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PricingOracleV1 *PricingOracleV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PricingOracleV1.Contract.PricingOracleV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PricingOracleV1 *PricingOracleV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PricingOracleV1.Contract.PricingOracleV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PricingOracleV1 *PricingOracleV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PricingOracleV1.Contract.PricingOracleV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PricingOracleV1 *PricingOracleV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PricingOracleV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PricingOracleV1 *PricingOracleV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PricingOracleV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PricingOracleV1 *PricingOracleV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PricingOracleV1.Contract.contract.Transact(opts, method, params...)
}

// Verifier is a free data retrieval call binding the contract method 0x8a172651.
//
// Solidity: function _verifier() view returns(address)
func (_PricingOracleV1 *PricingOracleV1Caller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PricingOracleV1.contract.Call(opts, &out, "_verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x8a172651.
//
// Solidity: function _verifier() view returns(address)
func (_PricingOracleV1 *PricingOracleV1Session) Verifier() (common.Address, error) {
	return _PricingOracleV1.Contract.Verifier(&_PricingOracleV1.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x8a172651.
//
// Solidity: function _verifier() view returns(address)
func (_PricingOracleV1 *PricingOracleV1CallerSession) Verifier() (common.Address, error) {
	return _PricingOracleV1.Contract.Verifier(&_PricingOracleV1.CallOpts)
}

// ConvertWeiToUsdCents is a free data retrieval call binding the contract method 0xfe024b4a.
//
// Solidity: function convertWeiToUsdCents(uint256 amount) view returns(uint256 usdCents)
func (_PricingOracleV1 *PricingOracleV1Caller) ConvertWeiToUsdCents(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PricingOracleV1.contract.Call(opts, &out, "convertWeiToUsdCents", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertWeiToUsdCents is a free data retrieval call binding the contract method 0xfe024b4a.
//
// Solidity: function convertWeiToUsdCents(uint256 amount) view returns(uint256 usdCents)
func (_PricingOracleV1 *PricingOracleV1Session) ConvertWeiToUsdCents(amount *big.Int) (*big.Int, error) {
	return _PricingOracleV1.Contract.ConvertWeiToUsdCents(&_PricingOracleV1.CallOpts, amount)
}

// ConvertWeiToUsdCents is a free data retrieval call binding the contract method 0xfe024b4a.
//
// Solidity: function convertWeiToUsdCents(uint256 amount) view returns(uint256 usdCents)
func (_PricingOracleV1 *PricingOracleV1CallerSession) ConvertWeiToUsdCents(amount *big.Int) (*big.Int, error) {
	return _PricingOracleV1.Contract.ConvertWeiToUsdCents(&_PricingOracleV1.CallOpts, amount)
}

// GetPriceForName is a free data retrieval call binding the contract method 0x8fb25887.
//
// Solidity: function getPriceForName(uint256 name, bytes data) view returns(uint256 price, uint256 priceCentsUsd)
func (_PricingOracleV1 *PricingOracleV1Caller) GetPriceForName(opts *bind.CallOpts, name *big.Int, data []byte) (struct {
	Price         *big.Int
	PriceCentsUsd *big.Int
}, error) {
	var out []interface{}
	err := _PricingOracleV1.contract.Call(opts, &out, "getPriceForName", name, data)

	outstruct := new(struct {
		Price         *big.Int
		PriceCentsUsd *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PriceCentsUsd = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPriceForName is a free data retrieval call binding the contract method 0x8fb25887.
//
// Solidity: function getPriceForName(uint256 name, bytes data) view returns(uint256 price, uint256 priceCentsUsd)
func (_PricingOracleV1 *PricingOracleV1Session) GetPriceForName(name *big.Int, data []byte) (struct {
	Price         *big.Int
	PriceCentsUsd *big.Int
}, error) {
	return _PricingOracleV1.Contract.GetPriceForName(&_PricingOracleV1.CallOpts, name, data)
}

// GetPriceForName is a free data retrieval call binding the contract method 0x8fb25887.
//
// Solidity: function getPriceForName(uint256 name, bytes data) view returns(uint256 price, uint256 priceCentsUsd)
func (_PricingOracleV1 *PricingOracleV1CallerSession) GetPriceForName(name *big.Int, data []byte) (struct {
	Price         *big.Int
	PriceCentsUsd *big.Int
}, error) {
	return _PricingOracleV1.Contract.GetPriceForName(&_PricingOracleV1.CallOpts, name, data)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_PricingOracleV1 *PricingOracleV1Caller) PriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PricingOracleV1.contract.Call(opts, &out, "priceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_PricingOracleV1 *PricingOracleV1Session) PriceFeed() (common.Address, error) {
	return _PricingOracleV1.Contract.PriceFeed(&_PricingOracleV1.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_PricingOracleV1 *PricingOracleV1CallerSession) PriceFeed() (common.Address, error) {
	return _PricingOracleV1.Contract.PriceFeed(&_PricingOracleV1.CallOpts)
}
