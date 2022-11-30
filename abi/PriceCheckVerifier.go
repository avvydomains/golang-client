// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PriceCheckVerifier

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

// PriceCheckVerifierMetaData contains all meta data concerning the PriceCheckVerifier contract.
var PriceCheckVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"pubSignals\",\"type\":\"uint256[]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PriceCheckVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use PriceCheckVerifierMetaData.ABI instead.
var PriceCheckVerifierABI = PriceCheckVerifierMetaData.ABI

// PriceCheckVerifier is an auto generated Go binding around an Ethereum contract.
type PriceCheckVerifier struct {
	PriceCheckVerifierCaller     // Read-only binding to the contract
	PriceCheckVerifierTransactor // Write-only binding to the contract
	PriceCheckVerifierFilterer   // Log filterer for contract events
}

// PriceCheckVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceCheckVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceCheckVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceCheckVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceCheckVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceCheckVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceCheckVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceCheckVerifierSession struct {
	Contract     *PriceCheckVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PriceCheckVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceCheckVerifierCallerSession struct {
	Contract *PriceCheckVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// PriceCheckVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceCheckVerifierTransactorSession struct {
	Contract     *PriceCheckVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// PriceCheckVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceCheckVerifierRaw struct {
	Contract *PriceCheckVerifier // Generic contract binding to access the raw methods on
}

// PriceCheckVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceCheckVerifierCallerRaw struct {
	Contract *PriceCheckVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// PriceCheckVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceCheckVerifierTransactorRaw struct {
	Contract *PriceCheckVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceCheckVerifier creates a new instance of PriceCheckVerifier, bound to a specific deployed contract.
func NewPriceCheckVerifier(address common.Address, backend bind.ContractBackend) (*PriceCheckVerifier, error) {
	contract, err := bindPriceCheckVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceCheckVerifier{PriceCheckVerifierCaller: PriceCheckVerifierCaller{contract: contract}, PriceCheckVerifierTransactor: PriceCheckVerifierTransactor{contract: contract}, PriceCheckVerifierFilterer: PriceCheckVerifierFilterer{contract: contract}}, nil
}

// NewPriceCheckVerifierCaller creates a new read-only instance of PriceCheckVerifier, bound to a specific deployed contract.
func NewPriceCheckVerifierCaller(address common.Address, caller bind.ContractCaller) (*PriceCheckVerifierCaller, error) {
	contract, err := bindPriceCheckVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceCheckVerifierCaller{contract: contract}, nil
}

// NewPriceCheckVerifierTransactor creates a new write-only instance of PriceCheckVerifier, bound to a specific deployed contract.
func NewPriceCheckVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceCheckVerifierTransactor, error) {
	contract, err := bindPriceCheckVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceCheckVerifierTransactor{contract: contract}, nil
}

// NewPriceCheckVerifierFilterer creates a new log filterer instance of PriceCheckVerifier, bound to a specific deployed contract.
func NewPriceCheckVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceCheckVerifierFilterer, error) {
	contract, err := bindPriceCheckVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceCheckVerifierFilterer{contract: contract}, nil
}

// bindPriceCheckVerifier binds a generic wrapper to an already deployed contract.
func bindPriceCheckVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PriceCheckVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceCheckVerifier *PriceCheckVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceCheckVerifier.Contract.PriceCheckVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceCheckVerifier *PriceCheckVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceCheckVerifier.Contract.PriceCheckVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceCheckVerifier *PriceCheckVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceCheckVerifier.Contract.PriceCheckVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceCheckVerifier *PriceCheckVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceCheckVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceCheckVerifier *PriceCheckVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceCheckVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceCheckVerifier *PriceCheckVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceCheckVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyProof is a free data retrieval call binding the contract method 0x1e8e1e13.
//
// Solidity: function verifyProof(bytes proof, uint256[] pubSignals) view returns(bool)
func (_PriceCheckVerifier *PriceCheckVerifierCaller) VerifyProof(opts *bind.CallOpts, proof []byte, pubSignals []*big.Int) (bool, error) {
	var out []interface{}
	err := _PriceCheckVerifier.contract.Call(opts, &out, "verifyProof", proof, pubSignals)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x1e8e1e13.
//
// Solidity: function verifyProof(bytes proof, uint256[] pubSignals) view returns(bool)
func (_PriceCheckVerifier *PriceCheckVerifierSession) VerifyProof(proof []byte, pubSignals []*big.Int) (bool, error) {
	return _PriceCheckVerifier.Contract.VerifyProof(&_PriceCheckVerifier.CallOpts, proof, pubSignals)
}

// VerifyProof is a free data retrieval call binding the contract method 0x1e8e1e13.
//
// Solidity: function verifyProof(bytes proof, uint256[] pubSignals) view returns(bool)
func (_PriceCheckVerifier *PriceCheckVerifierCallerSession) VerifyProof(proof []byte, pubSignals []*big.Int) (bool, error) {
	return _PriceCheckVerifier.Contract.VerifyProof(&_PriceCheckVerifier.CallOpts, proof, pubSignals)
}
