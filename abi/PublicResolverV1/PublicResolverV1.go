// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PublicResolverV1

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

// PublicResolverV1MetaData contains all meta data concerning the PublicResolverV1 contract.
var PublicResolverV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractRegistryAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"name\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"hash\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"path\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"EntrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"name\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"hash\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"path\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"StandardEntrySet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"name\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hash\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"name\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"}],\"name\":\"resolveStandard\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"name\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"path\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"name\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"path\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"setStandard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PublicResolverV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use PublicResolverV1MetaData.ABI instead.
var PublicResolverV1ABI = PublicResolverV1MetaData.ABI

// PublicResolverV1 is an auto generated Go binding around an Ethereum contract.
type PublicResolverV1 struct {
	PublicResolverV1Caller     // Read-only binding to the contract
	PublicResolverV1Transactor // Write-only binding to the contract
	PublicResolverV1Filterer   // Log filterer for contract events
}

// PublicResolverV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type PublicResolverV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicResolverV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type PublicResolverV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicResolverV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PublicResolverV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicResolverV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PublicResolverV1Session struct {
	Contract     *PublicResolverV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PublicResolverV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PublicResolverV1CallerSession struct {
	Contract *PublicResolverV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// PublicResolverV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PublicResolverV1TransactorSession struct {
	Contract     *PublicResolverV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// PublicResolverV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type PublicResolverV1Raw struct {
	Contract *PublicResolverV1 // Generic contract binding to access the raw methods on
}

// PublicResolverV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PublicResolverV1CallerRaw struct {
	Contract *PublicResolverV1Caller // Generic read-only contract binding to access the raw methods on
}

// PublicResolverV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PublicResolverV1TransactorRaw struct {
	Contract *PublicResolverV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPublicResolverV1 creates a new instance of PublicResolverV1, bound to a specific deployed contract.
func NewPublicResolverV1(address common.Address, backend bind.ContractBackend) (*PublicResolverV1, error) {
	contract, err := bindPublicResolverV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PublicResolverV1{PublicResolverV1Caller: PublicResolverV1Caller{contract: contract}, PublicResolverV1Transactor: PublicResolverV1Transactor{contract: contract}, PublicResolverV1Filterer: PublicResolverV1Filterer{contract: contract}}, nil
}

// NewPublicResolverV1Caller creates a new read-only instance of PublicResolverV1, bound to a specific deployed contract.
func NewPublicResolverV1Caller(address common.Address, caller bind.ContractCaller) (*PublicResolverV1Caller, error) {
	contract, err := bindPublicResolverV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PublicResolverV1Caller{contract: contract}, nil
}

// NewPublicResolverV1Transactor creates a new write-only instance of PublicResolverV1, bound to a specific deployed contract.
func NewPublicResolverV1Transactor(address common.Address, transactor bind.ContractTransactor) (*PublicResolverV1Transactor, error) {
	contract, err := bindPublicResolverV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PublicResolverV1Transactor{contract: contract}, nil
}

// NewPublicResolverV1Filterer creates a new log filterer instance of PublicResolverV1, bound to a specific deployed contract.
func NewPublicResolverV1Filterer(address common.Address, filterer bind.ContractFilterer) (*PublicResolverV1Filterer, error) {
	contract, err := bindPublicResolverV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PublicResolverV1Filterer{contract: contract}, nil
}

// bindPublicResolverV1 binds a generic wrapper to an already deployed contract.
func bindPublicResolverV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PublicResolverV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicResolverV1 *PublicResolverV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicResolverV1.Contract.PublicResolverV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicResolverV1 *PublicResolverV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicResolverV1.Contract.PublicResolverV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicResolverV1 *PublicResolverV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicResolverV1.Contract.PublicResolverV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicResolverV1 *PublicResolverV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicResolverV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicResolverV1 *PublicResolverV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicResolverV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicResolverV1 *PublicResolverV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicResolverV1.Contract.contract.Transact(opts, method, params...)
}

// Resolve is a free data retrieval call binding the contract method 0x75919f69.
//
// Solidity: function resolve(uint256 name, uint256 hash, string key) view returns(string data)
func (_PublicResolverV1 *PublicResolverV1Caller) Resolve(opts *bind.CallOpts, name *big.Int, hash *big.Int, key string) (string, error) {
	var out []interface{}
	err := _PublicResolverV1.contract.Call(opts, &out, "resolve", name, hash, key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x75919f69.
//
// Solidity: function resolve(uint256 name, uint256 hash, string key) view returns(string data)
func (_PublicResolverV1 *PublicResolverV1Session) Resolve(name *big.Int, hash *big.Int, key string) (string, error) {
	return _PublicResolverV1.Contract.Resolve(&_PublicResolverV1.CallOpts, name, hash, key)
}

// Resolve is a free data retrieval call binding the contract method 0x75919f69.
//
// Solidity: function resolve(uint256 name, uint256 hash, string key) view returns(string data)
func (_PublicResolverV1 *PublicResolverV1CallerSession) Resolve(name *big.Int, hash *big.Int, key string) (string, error) {
	return _PublicResolverV1.Contract.Resolve(&_PublicResolverV1.CallOpts, name, hash, key)
}

// ResolveStandard is a free data retrieval call binding the contract method 0xa98546f8.
//
// Solidity: function resolveStandard(uint256 name, uint256 hash, uint256 key) view returns(string data)
func (_PublicResolverV1 *PublicResolverV1Caller) ResolveStandard(opts *bind.CallOpts, name *big.Int, hash *big.Int, key *big.Int) (string, error) {
	var out []interface{}
	err := _PublicResolverV1.contract.Call(opts, &out, "resolveStandard", name, hash, key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ResolveStandard is a free data retrieval call binding the contract method 0xa98546f8.
//
// Solidity: function resolveStandard(uint256 name, uint256 hash, uint256 key) view returns(string data)
func (_PublicResolverV1 *PublicResolverV1Session) ResolveStandard(name *big.Int, hash *big.Int, key *big.Int) (string, error) {
	return _PublicResolverV1.Contract.ResolveStandard(&_PublicResolverV1.CallOpts, name, hash, key)
}

// ResolveStandard is a free data retrieval call binding the contract method 0xa98546f8.
//
// Solidity: function resolveStandard(uint256 name, uint256 hash, uint256 key) view returns(string data)
func (_PublicResolverV1 *PublicResolverV1CallerSession) ResolveStandard(name *big.Int, hash *big.Int, key *big.Int) (string, error) {
	return _PublicResolverV1.Contract.ResolveStandard(&_PublicResolverV1.CallOpts, name, hash, key)
}

// Set is a paid mutator transaction binding the contract method 0x5783e384.
//
// Solidity: function set(uint256 name, uint256[] path, string key, string data) returns()
func (_PublicResolverV1 *PublicResolverV1Transactor) Set(opts *bind.TransactOpts, name *big.Int, path []*big.Int, key string, data string) (*types.Transaction, error) {
	return _PublicResolverV1.contract.Transact(opts, "set", name, path, key, data)
}

// Set is a paid mutator transaction binding the contract method 0x5783e384.
//
// Solidity: function set(uint256 name, uint256[] path, string key, string data) returns()
func (_PublicResolverV1 *PublicResolverV1Session) Set(name *big.Int, path []*big.Int, key string, data string) (*types.Transaction, error) {
	return _PublicResolverV1.Contract.Set(&_PublicResolverV1.TransactOpts, name, path, key, data)
}

// Set is a paid mutator transaction binding the contract method 0x5783e384.
//
// Solidity: function set(uint256 name, uint256[] path, string key, string data) returns()
func (_PublicResolverV1 *PublicResolverV1TransactorSession) Set(name *big.Int, path []*big.Int, key string, data string) (*types.Transaction, error) {
	return _PublicResolverV1.Contract.Set(&_PublicResolverV1.TransactOpts, name, path, key, data)
}

// SetStandard is a paid mutator transaction binding the contract method 0xeeb6831b.
//
// Solidity: function setStandard(uint256 name, uint256[] path, uint256 key, string data) returns()
func (_PublicResolverV1 *PublicResolverV1Transactor) SetStandard(opts *bind.TransactOpts, name *big.Int, path []*big.Int, key *big.Int, data string) (*types.Transaction, error) {
	return _PublicResolverV1.contract.Transact(opts, "setStandard", name, path, key, data)
}

// SetStandard is a paid mutator transaction binding the contract method 0xeeb6831b.
//
// Solidity: function setStandard(uint256 name, uint256[] path, uint256 key, string data) returns()
func (_PublicResolverV1 *PublicResolverV1Session) SetStandard(name *big.Int, path []*big.Int, key *big.Int, data string) (*types.Transaction, error) {
	return _PublicResolverV1.Contract.SetStandard(&_PublicResolverV1.TransactOpts, name, path, key, data)
}

// SetStandard is a paid mutator transaction binding the contract method 0xeeb6831b.
//
// Solidity: function setStandard(uint256 name, uint256[] path, uint256 key, string data) returns()
func (_PublicResolverV1 *PublicResolverV1TransactorSession) SetStandard(name *big.Int, path []*big.Int, key *big.Int, data string) (*types.Transaction, error) {
	return _PublicResolverV1.Contract.SetStandard(&_PublicResolverV1.TransactOpts, name, path, key, data)
}

// PublicResolverV1EntrySetIterator is returned from FilterEntrySet and is used to iterate over the raw logs and unpacked data for EntrySet events raised by the PublicResolverV1 contract.
type PublicResolverV1EntrySetIterator struct {
	Event *PublicResolverV1EntrySet // Event containing the contract specifics and raw log

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
func (it *PublicResolverV1EntrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicResolverV1EntrySet)
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
		it.Event = new(PublicResolverV1EntrySet)
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
func (it *PublicResolverV1EntrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicResolverV1EntrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicResolverV1EntrySet represents a EntrySet event raised by the PublicResolverV1 contract.
type PublicResolverV1EntrySet struct {
	Name *big.Int
	Hash *big.Int
	Path []*big.Int
	Key  string
	Data string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterEntrySet is a free log retrieval operation binding the contract event 0x3a4da5e6dca27b6b1eae673773fc2edf95b15f6efb40c17e1e20ce534c7886d9.
//
// Solidity: event EntrySet(uint256 indexed name, uint256 indexed hash, uint256[] path, string key, string data)
func (_PublicResolverV1 *PublicResolverV1Filterer) FilterEntrySet(opts *bind.FilterOpts, name []*big.Int, hash []*big.Int) (*PublicResolverV1EntrySetIterator, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _PublicResolverV1.contract.FilterLogs(opts, "EntrySet", nameRule, hashRule)
	if err != nil {
		return nil, err
	}
	return &PublicResolverV1EntrySetIterator{contract: _PublicResolverV1.contract, event: "EntrySet", logs: logs, sub: sub}, nil
}

// WatchEntrySet is a free log subscription operation binding the contract event 0x3a4da5e6dca27b6b1eae673773fc2edf95b15f6efb40c17e1e20ce534c7886d9.
//
// Solidity: event EntrySet(uint256 indexed name, uint256 indexed hash, uint256[] path, string key, string data)
func (_PublicResolverV1 *PublicResolverV1Filterer) WatchEntrySet(opts *bind.WatchOpts, sink chan<- *PublicResolverV1EntrySet, name []*big.Int, hash []*big.Int) (event.Subscription, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _PublicResolverV1.contract.WatchLogs(opts, "EntrySet", nameRule, hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicResolverV1EntrySet)
				if err := _PublicResolverV1.contract.UnpackLog(event, "EntrySet", log); err != nil {
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

// ParseEntrySet is a log parse operation binding the contract event 0x3a4da5e6dca27b6b1eae673773fc2edf95b15f6efb40c17e1e20ce534c7886d9.
//
// Solidity: event EntrySet(uint256 indexed name, uint256 indexed hash, uint256[] path, string key, string data)
func (_PublicResolverV1 *PublicResolverV1Filterer) ParseEntrySet(log types.Log) (*PublicResolverV1EntrySet, error) {
	event := new(PublicResolverV1EntrySet)
	if err := _PublicResolverV1.contract.UnpackLog(event, "EntrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicResolverV1StandardEntrySetIterator is returned from FilterStandardEntrySet and is used to iterate over the raw logs and unpacked data for StandardEntrySet events raised by the PublicResolverV1 contract.
type PublicResolverV1StandardEntrySetIterator struct {
	Event *PublicResolverV1StandardEntrySet // Event containing the contract specifics and raw log

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
func (it *PublicResolverV1StandardEntrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicResolverV1StandardEntrySet)
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
		it.Event = new(PublicResolverV1StandardEntrySet)
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
func (it *PublicResolverV1StandardEntrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicResolverV1StandardEntrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicResolverV1StandardEntrySet represents a StandardEntrySet event raised by the PublicResolverV1 contract.
type PublicResolverV1StandardEntrySet struct {
	Name *big.Int
	Hash *big.Int
	Path []*big.Int
	Key  *big.Int
	Data string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterStandardEntrySet is a free log retrieval operation binding the contract event 0xe74cb6b565962460d347b3a97d088a45b1ab4b4666430d36e5b5868848b39bb4.
//
// Solidity: event StandardEntrySet(uint256 indexed name, uint256 indexed hash, uint256[] path, uint256 key, string data)
func (_PublicResolverV1 *PublicResolverV1Filterer) FilterStandardEntrySet(opts *bind.FilterOpts, name []*big.Int, hash []*big.Int) (*PublicResolverV1StandardEntrySetIterator, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _PublicResolverV1.contract.FilterLogs(opts, "StandardEntrySet", nameRule, hashRule)
	if err != nil {
		return nil, err
	}
	return &PublicResolverV1StandardEntrySetIterator{contract: _PublicResolverV1.contract, event: "StandardEntrySet", logs: logs, sub: sub}, nil
}

// WatchStandardEntrySet is a free log subscription operation binding the contract event 0xe74cb6b565962460d347b3a97d088a45b1ab4b4666430d36e5b5868848b39bb4.
//
// Solidity: event StandardEntrySet(uint256 indexed name, uint256 indexed hash, uint256[] path, uint256 key, string data)
func (_PublicResolverV1 *PublicResolverV1Filterer) WatchStandardEntrySet(opts *bind.WatchOpts, sink chan<- *PublicResolverV1StandardEntrySet, name []*big.Int, hash []*big.Int) (event.Subscription, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _PublicResolverV1.contract.WatchLogs(opts, "StandardEntrySet", nameRule, hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicResolverV1StandardEntrySet)
				if err := _PublicResolverV1.contract.UnpackLog(event, "StandardEntrySet", log); err != nil {
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

// ParseStandardEntrySet is a log parse operation binding the contract event 0xe74cb6b565962460d347b3a97d088a45b1ab4b4666430d36e5b5868848b39bb4.
//
// Solidity: event StandardEntrySet(uint256 indexed name, uint256 indexed hash, uint256[] path, uint256 key, string data)
func (_PublicResolverV1 *PublicResolverV1Filterer) ParseStandardEntrySet(log types.Log) (*PublicResolverV1StandardEntrySet, error) {
	event := new(PublicResolverV1StandardEntrySet)
	if err := _PublicResolverV1.contract.UnpackLog(event, "StandardEntrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
