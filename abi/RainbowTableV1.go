// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package RainbowTableV1

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

// RainbowTableV1MetaData contains all meta data concerning the RainbowTableV1 contract.
var RainbowTableV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractRegistryAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"hash\",\"type\":\"uint256\"}],\"name\":\"Revealed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"contractRegistry\",\"outputs\":[{\"internalType\":\"contractContractRegistryInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"hash\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"preimage\",\"type\":\"uint256[]\"}],\"name\":\"getHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"hash\",\"type\":\"uint256\"}],\"name\":\"isRevealed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"hash\",\"type\":\"uint256\"}],\"name\":\"lookup\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"preimage\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"preimage\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"hash\",\"type\":\"uint256\"}],\"name\":\"reveal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RainbowTableV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use RainbowTableV1MetaData.ABI instead.
var RainbowTableV1ABI = RainbowTableV1MetaData.ABI

// RainbowTableV1 is an auto generated Go binding around an Ethereum contract.
type RainbowTableV1 struct {
	RainbowTableV1Caller     // Read-only binding to the contract
	RainbowTableV1Transactor // Write-only binding to the contract
	RainbowTableV1Filterer   // Log filterer for contract events
}

// RainbowTableV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type RainbowTableV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RainbowTableV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type RainbowTableV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RainbowTableV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RainbowTableV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RainbowTableV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RainbowTableV1Session struct {
	Contract     *RainbowTableV1   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RainbowTableV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RainbowTableV1CallerSession struct {
	Contract *RainbowTableV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// RainbowTableV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RainbowTableV1TransactorSession struct {
	Contract     *RainbowTableV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// RainbowTableV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type RainbowTableV1Raw struct {
	Contract *RainbowTableV1 // Generic contract binding to access the raw methods on
}

// RainbowTableV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RainbowTableV1CallerRaw struct {
	Contract *RainbowTableV1Caller // Generic read-only contract binding to access the raw methods on
}

// RainbowTableV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RainbowTableV1TransactorRaw struct {
	Contract *RainbowTableV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewRainbowTableV1 creates a new instance of RainbowTableV1, bound to a specific deployed contract.
func NewRainbowTableV1(address common.Address, backend bind.ContractBackend) (*RainbowTableV1, error) {
	contract, err := bindRainbowTableV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RainbowTableV1{RainbowTableV1Caller: RainbowTableV1Caller{contract: contract}, RainbowTableV1Transactor: RainbowTableV1Transactor{contract: contract}, RainbowTableV1Filterer: RainbowTableV1Filterer{contract: contract}}, nil
}

// NewRainbowTableV1Caller creates a new read-only instance of RainbowTableV1, bound to a specific deployed contract.
func NewRainbowTableV1Caller(address common.Address, caller bind.ContractCaller) (*RainbowTableV1Caller, error) {
	contract, err := bindRainbowTableV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RainbowTableV1Caller{contract: contract}, nil
}

// NewRainbowTableV1Transactor creates a new write-only instance of RainbowTableV1, bound to a specific deployed contract.
func NewRainbowTableV1Transactor(address common.Address, transactor bind.ContractTransactor) (*RainbowTableV1Transactor, error) {
	contract, err := bindRainbowTableV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RainbowTableV1Transactor{contract: contract}, nil
}

// NewRainbowTableV1Filterer creates a new log filterer instance of RainbowTableV1, bound to a specific deployed contract.
func NewRainbowTableV1Filterer(address common.Address, filterer bind.ContractFilterer) (*RainbowTableV1Filterer, error) {
	contract, err := bindRainbowTableV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RainbowTableV1Filterer{contract: contract}, nil
}

// bindRainbowTableV1 binds a generic wrapper to an already deployed contract.
func bindRainbowTableV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RainbowTableV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RainbowTableV1 *RainbowTableV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RainbowTableV1.Contract.RainbowTableV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RainbowTableV1 *RainbowTableV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RainbowTableV1.Contract.RainbowTableV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RainbowTableV1 *RainbowTableV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RainbowTableV1.Contract.RainbowTableV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RainbowTableV1 *RainbowTableV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RainbowTableV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RainbowTableV1 *RainbowTableV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RainbowTableV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RainbowTableV1 *RainbowTableV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RainbowTableV1.Contract.contract.Transact(opts, method, params...)
}

// ContractRegistry is a free data retrieval call binding the contract method 0xabf410e5.
//
// Solidity: function contractRegistry() view returns(address)
func (_RainbowTableV1 *RainbowTableV1Caller) ContractRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RainbowTableV1.contract.Call(opts, &out, "contractRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractRegistry is a free data retrieval call binding the contract method 0xabf410e5.
//
// Solidity: function contractRegistry() view returns(address)
func (_RainbowTableV1 *RainbowTableV1Session) ContractRegistry() (common.Address, error) {
	return _RainbowTableV1.Contract.ContractRegistry(&_RainbowTableV1.CallOpts)
}

// ContractRegistry is a free data retrieval call binding the contract method 0xabf410e5.
//
// Solidity: function contractRegistry() view returns(address)
func (_RainbowTableV1 *RainbowTableV1CallerSession) ContractRegistry() (common.Address, error) {
	return _RainbowTableV1.Contract.ContractRegistry(&_RainbowTableV1.CallOpts)
}

// GetHash is a free data retrieval call binding the contract method 0x978d7ec2.
//
// Solidity: function getHash(uint256 hash, uint256[] preimage) view returns(uint256)
func (_RainbowTableV1 *RainbowTableV1Caller) GetHash(opts *bind.CallOpts, hash *big.Int, preimage []*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RainbowTableV1.contract.Call(opts, &out, "getHash", hash, preimage)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetHash is a free data retrieval call binding the contract method 0x978d7ec2.
//
// Solidity: function getHash(uint256 hash, uint256[] preimage) view returns(uint256)
func (_RainbowTableV1 *RainbowTableV1Session) GetHash(hash *big.Int, preimage []*big.Int) (*big.Int, error) {
	return _RainbowTableV1.Contract.GetHash(&_RainbowTableV1.CallOpts, hash, preimage)
}

// GetHash is a free data retrieval call binding the contract method 0x978d7ec2.
//
// Solidity: function getHash(uint256 hash, uint256[] preimage) view returns(uint256)
func (_RainbowTableV1 *RainbowTableV1CallerSession) GetHash(hash *big.Int, preimage []*big.Int) (*big.Int, error) {
	return _RainbowTableV1.Contract.GetHash(&_RainbowTableV1.CallOpts, hash, preimage)
}

// IsRevealed is a free data retrieval call binding the contract method 0x5055fbc3.
//
// Solidity: function isRevealed(uint256 hash) view returns(bool)
func (_RainbowTableV1 *RainbowTableV1Caller) IsRevealed(opts *bind.CallOpts, hash *big.Int) (bool, error) {
	var out []interface{}
	err := _RainbowTableV1.contract.Call(opts, &out, "isRevealed", hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRevealed is a free data retrieval call binding the contract method 0x5055fbc3.
//
// Solidity: function isRevealed(uint256 hash) view returns(bool)
func (_RainbowTableV1 *RainbowTableV1Session) IsRevealed(hash *big.Int) (bool, error) {
	return _RainbowTableV1.Contract.IsRevealed(&_RainbowTableV1.CallOpts, hash)
}

// IsRevealed is a free data retrieval call binding the contract method 0x5055fbc3.
//
// Solidity: function isRevealed(uint256 hash) view returns(bool)
func (_RainbowTableV1 *RainbowTableV1CallerSession) IsRevealed(hash *big.Int) (bool, error) {
	return _RainbowTableV1.Contract.IsRevealed(&_RainbowTableV1.CallOpts, hash)
}

// Lookup is a free data retrieval call binding the contract method 0x0a874df6.
//
// Solidity: function lookup(uint256 hash) view returns(uint256[] preimage)
func (_RainbowTableV1 *RainbowTableV1Caller) Lookup(opts *bind.CallOpts, hash *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _RainbowTableV1.contract.Call(opts, &out, "lookup", hash)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// Lookup is a free data retrieval call binding the contract method 0x0a874df6.
//
// Solidity: function lookup(uint256 hash) view returns(uint256[] preimage)
func (_RainbowTableV1 *RainbowTableV1Session) Lookup(hash *big.Int) ([]*big.Int, error) {
	return _RainbowTableV1.Contract.Lookup(&_RainbowTableV1.CallOpts, hash)
}

// Lookup is a free data retrieval call binding the contract method 0x0a874df6.
//
// Solidity: function lookup(uint256 hash) view returns(uint256[] preimage)
func (_RainbowTableV1 *RainbowTableV1CallerSession) Lookup(hash *big.Int) ([]*big.Int, error) {
	return _RainbowTableV1.Contract.Lookup(&_RainbowTableV1.CallOpts, hash)
}

// Reveal is a paid mutator transaction binding the contract method 0x07461df1.
//
// Solidity: function reveal(uint256[] preimage, uint256 hash) returns()
func (_RainbowTableV1 *RainbowTableV1Transactor) Reveal(opts *bind.TransactOpts, preimage []*big.Int, hash *big.Int) (*types.Transaction, error) {
	return _RainbowTableV1.contract.Transact(opts, "reveal", preimage, hash)
}

// Reveal is a paid mutator transaction binding the contract method 0x07461df1.
//
// Solidity: function reveal(uint256[] preimage, uint256 hash) returns()
func (_RainbowTableV1 *RainbowTableV1Session) Reveal(preimage []*big.Int, hash *big.Int) (*types.Transaction, error) {
	return _RainbowTableV1.Contract.Reveal(&_RainbowTableV1.TransactOpts, preimage, hash)
}

// Reveal is a paid mutator transaction binding the contract method 0x07461df1.
//
// Solidity: function reveal(uint256[] preimage, uint256 hash) returns()
func (_RainbowTableV1 *RainbowTableV1TransactorSession) Reveal(preimage []*big.Int, hash *big.Int) (*types.Transaction, error) {
	return _RainbowTableV1.Contract.Reveal(&_RainbowTableV1.TransactOpts, preimage, hash)
}

// RainbowTableV1RevealedIterator is returned from FilterRevealed and is used to iterate over the raw logs and unpacked data for Revealed events raised by the RainbowTableV1 contract.
type RainbowTableV1RevealedIterator struct {
	Event *RainbowTableV1Revealed // Event containing the contract specifics and raw log

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
func (it *RainbowTableV1RevealedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RainbowTableV1Revealed)
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
		it.Event = new(RainbowTableV1Revealed)
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
func (it *RainbowTableV1RevealedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RainbowTableV1RevealedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RainbowTableV1Revealed represents a Revealed event raised by the RainbowTableV1 contract.
type RainbowTableV1Revealed struct {
	Hash *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRevealed is a free log retrieval operation binding the contract event 0x15120e52505e619cbf6c2af910d5cf7f9ee1befa55801b078c33e93880b2d609.
//
// Solidity: event Revealed(uint256 indexed hash)
func (_RainbowTableV1 *RainbowTableV1Filterer) FilterRevealed(opts *bind.FilterOpts, hash []*big.Int) (*RainbowTableV1RevealedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _RainbowTableV1.contract.FilterLogs(opts, "Revealed", hashRule)
	if err != nil {
		return nil, err
	}
	return &RainbowTableV1RevealedIterator{contract: _RainbowTableV1.contract, event: "Revealed", logs: logs, sub: sub}, nil
}

// WatchRevealed is a free log subscription operation binding the contract event 0x15120e52505e619cbf6c2af910d5cf7f9ee1befa55801b078c33e93880b2d609.
//
// Solidity: event Revealed(uint256 indexed hash)
func (_RainbowTableV1 *RainbowTableV1Filterer) WatchRevealed(opts *bind.WatchOpts, sink chan<- *RainbowTableV1Revealed, hash []*big.Int) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _RainbowTableV1.contract.WatchLogs(opts, "Revealed", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RainbowTableV1Revealed)
				if err := _RainbowTableV1.contract.UnpackLog(event, "Revealed", log); err != nil {
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

// ParseRevealed is a log parse operation binding the contract event 0x15120e52505e619cbf6c2af910d5cf7f9ee1befa55801b078c33e93880b2d609.
//
// Solidity: event Revealed(uint256 indexed hash)
func (_RainbowTableV1 *RainbowTableV1Filterer) ParseRevealed(log types.Log) (*RainbowTableV1Revealed, error) {
	event := new(RainbowTableV1Revealed)
	if err := _RainbowTableV1.contract.UnpackLog(event, "Revealed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
