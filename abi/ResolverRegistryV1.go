// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ResolverRegistryV1

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

// ResolverRegistryV1MetaData contains all meta data concerning the ResolverRegistryV1 contract.
var ResolverRegistryV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractRegistryAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"name\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"hash\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"path\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"datasetId\",\"type\":\"uint256\"}],\"name\":\"ResolverSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"name\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hash\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"datasetId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"name\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"path\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"datasetId\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ResolverRegistryV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use ResolverRegistryV1MetaData.ABI instead.
var ResolverRegistryV1ABI = ResolverRegistryV1MetaData.ABI

// ResolverRegistryV1 is an auto generated Go binding around an Ethereum contract.
type ResolverRegistryV1 struct {
	ResolverRegistryV1Caller     // Read-only binding to the contract
	ResolverRegistryV1Transactor // Write-only binding to the contract
	ResolverRegistryV1Filterer   // Log filterer for contract events
}

// ResolverRegistryV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type ResolverRegistryV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverRegistryV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ResolverRegistryV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverRegistryV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ResolverRegistryV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverRegistryV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ResolverRegistryV1Session struct {
	Contract     *ResolverRegistryV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ResolverRegistryV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ResolverRegistryV1CallerSession struct {
	Contract *ResolverRegistryV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ResolverRegistryV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ResolverRegistryV1TransactorSession struct {
	Contract     *ResolverRegistryV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ResolverRegistryV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type ResolverRegistryV1Raw struct {
	Contract *ResolverRegistryV1 // Generic contract binding to access the raw methods on
}

// ResolverRegistryV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ResolverRegistryV1CallerRaw struct {
	Contract *ResolverRegistryV1Caller // Generic read-only contract binding to access the raw methods on
}

// ResolverRegistryV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ResolverRegistryV1TransactorRaw struct {
	Contract *ResolverRegistryV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewResolverRegistryV1 creates a new instance of ResolverRegistryV1, bound to a specific deployed contract.
func NewResolverRegistryV1(address common.Address, backend bind.ContractBackend) (*ResolverRegistryV1, error) {
	contract, err := bindResolverRegistryV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ResolverRegistryV1{ResolverRegistryV1Caller: ResolverRegistryV1Caller{contract: contract}, ResolverRegistryV1Transactor: ResolverRegistryV1Transactor{contract: contract}, ResolverRegistryV1Filterer: ResolverRegistryV1Filterer{contract: contract}}, nil
}

// NewResolverRegistryV1Caller creates a new read-only instance of ResolverRegistryV1, bound to a specific deployed contract.
func NewResolverRegistryV1Caller(address common.Address, caller bind.ContractCaller) (*ResolverRegistryV1Caller, error) {
	contract, err := bindResolverRegistryV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ResolverRegistryV1Caller{contract: contract}, nil
}

// NewResolverRegistryV1Transactor creates a new write-only instance of ResolverRegistryV1, bound to a specific deployed contract.
func NewResolverRegistryV1Transactor(address common.Address, transactor bind.ContractTransactor) (*ResolverRegistryV1Transactor, error) {
	contract, err := bindResolverRegistryV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ResolverRegistryV1Transactor{contract: contract}, nil
}

// NewResolverRegistryV1Filterer creates a new log filterer instance of ResolverRegistryV1, bound to a specific deployed contract.
func NewResolverRegistryV1Filterer(address common.Address, filterer bind.ContractFilterer) (*ResolverRegistryV1Filterer, error) {
	contract, err := bindResolverRegistryV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ResolverRegistryV1Filterer{contract: contract}, nil
}

// bindResolverRegistryV1 binds a generic wrapper to an already deployed contract.
func bindResolverRegistryV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ResolverRegistryV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ResolverRegistryV1 *ResolverRegistryV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ResolverRegistryV1.Contract.ResolverRegistryV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ResolverRegistryV1 *ResolverRegistryV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ResolverRegistryV1.Contract.ResolverRegistryV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ResolverRegistryV1 *ResolverRegistryV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ResolverRegistryV1.Contract.ResolverRegistryV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ResolverRegistryV1 *ResolverRegistryV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ResolverRegistryV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ResolverRegistryV1 *ResolverRegistryV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ResolverRegistryV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ResolverRegistryV1 *ResolverRegistryV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ResolverRegistryV1.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x669e48aa.
//
// Solidity: function get(uint256 name, uint256 hash) view returns(address resolver, uint256 datasetId)
func (_ResolverRegistryV1 *ResolverRegistryV1Caller) Get(opts *bind.CallOpts, name *big.Int, hash *big.Int) (struct {
	Resolver  common.Address
	DatasetId *big.Int
}, error) {
	var out []interface{}
	err := _ResolverRegistryV1.contract.Call(opts, &out, "get", name, hash)

	outstruct := new(struct {
		Resolver  common.Address
		DatasetId *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Resolver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.DatasetId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Get is a free data retrieval call binding the contract method 0x669e48aa.
//
// Solidity: function get(uint256 name, uint256 hash) view returns(address resolver, uint256 datasetId)
func (_ResolverRegistryV1 *ResolverRegistryV1Session) Get(name *big.Int, hash *big.Int) (struct {
	Resolver  common.Address
	DatasetId *big.Int
}, error) {
	return _ResolverRegistryV1.Contract.Get(&_ResolverRegistryV1.CallOpts, name, hash)
}

// Get is a free data retrieval call binding the contract method 0x669e48aa.
//
// Solidity: function get(uint256 name, uint256 hash) view returns(address resolver, uint256 datasetId)
func (_ResolverRegistryV1 *ResolverRegistryV1CallerSession) Get(name *big.Int, hash *big.Int) (struct {
	Resolver  common.Address
	DatasetId *big.Int
}, error) {
	return _ResolverRegistryV1.Contract.Get(&_ResolverRegistryV1.CallOpts, name, hash)
}

// Set is a paid mutator transaction binding the contract method 0x6db48875.
//
// Solidity: function set(uint256 name, uint256[] path, address resolver, uint256 datasetId) returns()
func (_ResolverRegistryV1 *ResolverRegistryV1Transactor) Set(opts *bind.TransactOpts, name *big.Int, path []*big.Int, resolver common.Address, datasetId *big.Int) (*types.Transaction, error) {
	return _ResolverRegistryV1.contract.Transact(opts, "set", name, path, resolver, datasetId)
}

// Set is a paid mutator transaction binding the contract method 0x6db48875.
//
// Solidity: function set(uint256 name, uint256[] path, address resolver, uint256 datasetId) returns()
func (_ResolverRegistryV1 *ResolverRegistryV1Session) Set(name *big.Int, path []*big.Int, resolver common.Address, datasetId *big.Int) (*types.Transaction, error) {
	return _ResolverRegistryV1.Contract.Set(&_ResolverRegistryV1.TransactOpts, name, path, resolver, datasetId)
}

// Set is a paid mutator transaction binding the contract method 0x6db48875.
//
// Solidity: function set(uint256 name, uint256[] path, address resolver, uint256 datasetId) returns()
func (_ResolverRegistryV1 *ResolverRegistryV1TransactorSession) Set(name *big.Int, path []*big.Int, resolver common.Address, datasetId *big.Int) (*types.Transaction, error) {
	return _ResolverRegistryV1.Contract.Set(&_ResolverRegistryV1.TransactOpts, name, path, resolver, datasetId)
}

// ResolverRegistryV1ResolverSetIterator is returned from FilterResolverSet and is used to iterate over the raw logs and unpacked data for ResolverSet events raised by the ResolverRegistryV1 contract.
type ResolverRegistryV1ResolverSetIterator struct {
	Event *ResolverRegistryV1ResolverSet // Event containing the contract specifics and raw log

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
func (it *ResolverRegistryV1ResolverSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverRegistryV1ResolverSet)
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
		it.Event = new(ResolverRegistryV1ResolverSet)
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
func (it *ResolverRegistryV1ResolverSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverRegistryV1ResolverSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverRegistryV1ResolverSet represents a ResolverSet event raised by the ResolverRegistryV1 contract.
type ResolverRegistryV1ResolverSet struct {
	Name      *big.Int
	Hash      *big.Int
	Path      []*big.Int
	Resolver  common.Address
	DatasetId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterResolverSet is a free log retrieval operation binding the contract event 0x8cd646ba74863cac4df68574be48d8eb1c1bec797eb22c3e27535504e06d7d74.
//
// Solidity: event ResolverSet(uint256 indexed name, uint256 indexed hash, uint256[] path, address resolver, uint256 datasetId)
func (_ResolverRegistryV1 *ResolverRegistryV1Filterer) FilterResolverSet(opts *bind.FilterOpts, name []*big.Int, hash []*big.Int) (*ResolverRegistryV1ResolverSetIterator, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _ResolverRegistryV1.contract.FilterLogs(opts, "ResolverSet", nameRule, hashRule)
	if err != nil {
		return nil, err
	}
	return &ResolverRegistryV1ResolverSetIterator{contract: _ResolverRegistryV1.contract, event: "ResolverSet", logs: logs, sub: sub}, nil
}

// WatchResolverSet is a free log subscription operation binding the contract event 0x8cd646ba74863cac4df68574be48d8eb1c1bec797eb22c3e27535504e06d7d74.
//
// Solidity: event ResolverSet(uint256 indexed name, uint256 indexed hash, uint256[] path, address resolver, uint256 datasetId)
func (_ResolverRegistryV1 *ResolverRegistryV1Filterer) WatchResolverSet(opts *bind.WatchOpts, sink chan<- *ResolverRegistryV1ResolverSet, name []*big.Int, hash []*big.Int) (event.Subscription, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _ResolverRegistryV1.contract.WatchLogs(opts, "ResolverSet", nameRule, hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverRegistryV1ResolverSet)
				if err := _ResolverRegistryV1.contract.UnpackLog(event, "ResolverSet", log); err != nil {
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

// ParseResolverSet is a log parse operation binding the contract event 0x8cd646ba74863cac4df68574be48d8eb1c1bec797eb22c3e27535504e06d7d74.
//
// Solidity: event ResolverSet(uint256 indexed name, uint256 indexed hash, uint256[] path, address resolver, uint256 datasetId)
func (_ResolverRegistryV1 *ResolverRegistryV1Filterer) ParseResolverSet(log types.Log) (*ResolverRegistryV1ResolverSet, error) {
	event := new(ResolverRegistryV1ResolverSet)
	if err := _ResolverRegistryV1.contract.UnpackLog(event, "ResolverSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
