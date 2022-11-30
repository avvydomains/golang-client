// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ConstraintsVerifier

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

// ConstraintsVerifierMetaData contains all meta data concerning the ConstraintsVerifier contract.
var ConstraintsVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"pubSignals\",\"type\":\"uint256[]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ConstraintsVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use ConstraintsVerifierMetaData.ABI instead.
var ConstraintsVerifierABI = ConstraintsVerifierMetaData.ABI

// ConstraintsVerifier is an auto generated Go binding around an Ethereum contract.
type ConstraintsVerifier struct {
	ConstraintsVerifierCaller     // Read-only binding to the contract
	ConstraintsVerifierTransactor // Write-only binding to the contract
	ConstraintsVerifierFilterer   // Log filterer for contract events
}

// ConstraintsVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConstraintsVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstraintsVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConstraintsVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstraintsVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConstraintsVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstraintsVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConstraintsVerifierSession struct {
	Contract     *ConstraintsVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ConstraintsVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConstraintsVerifierCallerSession struct {
	Contract *ConstraintsVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ConstraintsVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConstraintsVerifierTransactorSession struct {
	Contract     *ConstraintsVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ConstraintsVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConstraintsVerifierRaw struct {
	Contract *ConstraintsVerifier // Generic contract binding to access the raw methods on
}

// ConstraintsVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConstraintsVerifierCallerRaw struct {
	Contract *ConstraintsVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// ConstraintsVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConstraintsVerifierTransactorRaw struct {
	Contract *ConstraintsVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConstraintsVerifier creates a new instance of ConstraintsVerifier, bound to a specific deployed contract.
func NewConstraintsVerifier(address common.Address, backend bind.ContractBackend) (*ConstraintsVerifier, error) {
	contract, err := bindConstraintsVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConstraintsVerifier{ConstraintsVerifierCaller: ConstraintsVerifierCaller{contract: contract}, ConstraintsVerifierTransactor: ConstraintsVerifierTransactor{contract: contract}, ConstraintsVerifierFilterer: ConstraintsVerifierFilterer{contract: contract}}, nil
}

// NewConstraintsVerifierCaller creates a new read-only instance of ConstraintsVerifier, bound to a specific deployed contract.
func NewConstraintsVerifierCaller(address common.Address, caller bind.ContractCaller) (*ConstraintsVerifierCaller, error) {
	contract, err := bindConstraintsVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConstraintsVerifierCaller{contract: contract}, nil
}

// NewConstraintsVerifierTransactor creates a new write-only instance of ConstraintsVerifier, bound to a specific deployed contract.
func NewConstraintsVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*ConstraintsVerifierTransactor, error) {
	contract, err := bindConstraintsVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConstraintsVerifierTransactor{contract: contract}, nil
}

// NewConstraintsVerifierFilterer creates a new log filterer instance of ConstraintsVerifier, bound to a specific deployed contract.
func NewConstraintsVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*ConstraintsVerifierFilterer, error) {
	contract, err := bindConstraintsVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConstraintsVerifierFilterer{contract: contract}, nil
}

// bindConstraintsVerifier binds a generic wrapper to an already deployed contract.
func bindConstraintsVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ConstraintsVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConstraintsVerifier *ConstraintsVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConstraintsVerifier.Contract.ConstraintsVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConstraintsVerifier *ConstraintsVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConstraintsVerifier.Contract.ConstraintsVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConstraintsVerifier *ConstraintsVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConstraintsVerifier.Contract.ConstraintsVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConstraintsVerifier *ConstraintsVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConstraintsVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConstraintsVerifier *ConstraintsVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConstraintsVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConstraintsVerifier *ConstraintsVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConstraintsVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyProof is a free data retrieval call binding the contract method 0x1e8e1e13.
//
// Solidity: function verifyProof(bytes proof, uint256[] pubSignals) view returns(bool)
func (_ConstraintsVerifier *ConstraintsVerifierCaller) VerifyProof(opts *bind.CallOpts, proof []byte, pubSignals []*big.Int) (bool, error) {
	var out []interface{}
	err := _ConstraintsVerifier.contract.Call(opts, &out, "verifyProof", proof, pubSignals)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x1e8e1e13.
//
// Solidity: function verifyProof(bytes proof, uint256[] pubSignals) view returns(bool)
func (_ConstraintsVerifier *ConstraintsVerifierSession) VerifyProof(proof []byte, pubSignals []*big.Int) (bool, error) {
	return _ConstraintsVerifier.Contract.VerifyProof(&_ConstraintsVerifier.CallOpts, proof, pubSignals)
}

// VerifyProof is a free data retrieval call binding the contract method 0x1e8e1e13.
//
// Solidity: function verifyProof(bytes proof, uint256[] pubSignals) view returns(bool)
func (_ConstraintsVerifier *ConstraintsVerifierCallerSession) VerifyProof(proof []byte, pubSignals []*big.Int) (bool, error) {
	return _ConstraintsVerifier.Contract.VerifyProof(&_ConstraintsVerifier.CallOpts, proof, pubSignals)
}
