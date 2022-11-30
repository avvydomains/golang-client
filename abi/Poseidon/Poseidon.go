// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Poseidon

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

// PoseidonMetaData contains all meta data concerning the Poseidon contract.
var PoseidonMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"input\",\"type\":\"bytes32[3]\"}],\"name\":\"poseidon\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"input\",\"type\":\"uint256[3]\"}],\"name\":\"poseidon\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// PoseidonABI is the input ABI used to generate the binding from.
// Deprecated: Use PoseidonMetaData.ABI instead.
var PoseidonABI = PoseidonMetaData.ABI

// Poseidon is an auto generated Go binding around an Ethereum contract.
type Poseidon struct {
	PoseidonCaller     // Read-only binding to the contract
	PoseidonTransactor // Write-only binding to the contract
	PoseidonFilterer   // Log filterer for contract events
}

// PoseidonCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoseidonCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoseidonTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoseidonTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoseidonFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoseidonFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoseidonSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoseidonSession struct {
	Contract     *Poseidon         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoseidonCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoseidonCallerSession struct {
	Contract *PoseidonCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PoseidonTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoseidonTransactorSession struct {
	Contract     *PoseidonTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PoseidonRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoseidonRaw struct {
	Contract *Poseidon // Generic contract binding to access the raw methods on
}

// PoseidonCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoseidonCallerRaw struct {
	Contract *PoseidonCaller // Generic read-only contract binding to access the raw methods on
}

// PoseidonTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoseidonTransactorRaw struct {
	Contract *PoseidonTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoseidon creates a new instance of Poseidon, bound to a specific deployed contract.
func NewPoseidon(address common.Address, backend bind.ContractBackend) (*Poseidon, error) {
	contract, err := bindPoseidon(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Poseidon{PoseidonCaller: PoseidonCaller{contract: contract}, PoseidonTransactor: PoseidonTransactor{contract: contract}, PoseidonFilterer: PoseidonFilterer{contract: contract}}, nil
}

// NewPoseidonCaller creates a new read-only instance of Poseidon, bound to a specific deployed contract.
func NewPoseidonCaller(address common.Address, caller bind.ContractCaller) (*PoseidonCaller, error) {
	contract, err := bindPoseidon(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoseidonCaller{contract: contract}, nil
}

// NewPoseidonTransactor creates a new write-only instance of Poseidon, bound to a specific deployed contract.
func NewPoseidonTransactor(address common.Address, transactor bind.ContractTransactor) (*PoseidonTransactor, error) {
	contract, err := bindPoseidon(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoseidonTransactor{contract: contract}, nil
}

// NewPoseidonFilterer creates a new log filterer instance of Poseidon, bound to a specific deployed contract.
func NewPoseidonFilterer(address common.Address, filterer bind.ContractFilterer) (*PoseidonFilterer, error) {
	contract, err := bindPoseidon(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoseidonFilterer{contract: contract}, nil
}

// bindPoseidon binds a generic wrapper to an already deployed contract.
func bindPoseidon(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoseidonMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Poseidon *PoseidonRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Poseidon.Contract.PoseidonCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Poseidon *PoseidonRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poseidon.Contract.PoseidonTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Poseidon *PoseidonRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Poseidon.Contract.PoseidonTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Poseidon *PoseidonCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Poseidon.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Poseidon *PoseidonTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poseidon.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Poseidon *PoseidonTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Poseidon.Contract.contract.Transact(opts, method, params...)
}

// Poseidon is a free data retrieval call binding the contract method 0x5a53025d.
//
// Solidity: function poseidon(bytes32[3] input) pure returns(bytes32)
func (_Poseidon *PoseidonCaller) Poseidon(opts *bind.CallOpts, input [3][32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Poseidon.contract.Call(opts, &out, "poseidon", input)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Poseidon is a free data retrieval call binding the contract method 0x5a53025d.
//
// Solidity: function poseidon(bytes32[3] input) pure returns(bytes32)
func (_Poseidon *PoseidonSession) Poseidon(input [3][32]byte) ([32]byte, error) {
	return _Poseidon.Contract.Poseidon(&_Poseidon.CallOpts, input)
}

// Poseidon is a free data retrieval call binding the contract method 0x5a53025d.
//
// Solidity: function poseidon(bytes32[3] input) pure returns(bytes32)
func (_Poseidon *PoseidonCallerSession) Poseidon(input [3][32]byte) ([32]byte, error) {
	return _Poseidon.Contract.Poseidon(&_Poseidon.CallOpts, input)
}

// Poseidon0 is a free data retrieval call binding the contract method 0x25cc70e8.
//
// Solidity: function poseidon(uint256[3] input) pure returns(uint256)
func (_Poseidon *PoseidonCaller) Poseidon0(opts *bind.CallOpts, input [3]*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Poseidon.contract.Call(opts, &out, "poseidon0", input)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Poseidon0 is a free data retrieval call binding the contract method 0x25cc70e8.
//
// Solidity: function poseidon(uint256[3] input) pure returns(uint256)
func (_Poseidon *PoseidonSession) Poseidon0(input [3]*big.Int) (*big.Int, error) {
	return _Poseidon.Contract.Poseidon0(&_Poseidon.CallOpts, input)
}

// Poseidon0 is a free data retrieval call binding the contract method 0x25cc70e8.
//
// Solidity: function poseidon(uint256[3] input) pure returns(uint256)
func (_Poseidon *PoseidonCallerSession) Poseidon0(input [3]*big.Int) (*big.Int, error) {
	return _Poseidon.Contract.Poseidon0(&_Poseidon.CallOpts, input)
}
