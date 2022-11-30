// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package EVMReverseResolverV1

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

// EVMReverseResolverV1MetaData contains all meta data concerning the EVMReverseResolverV1 contract.
var EVMReverseResolverV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractContractRegistryInterface\",\"name\":\"_contractRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"name\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"path\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"EntrySet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"name\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"path\",\"type\":\"uint256[]\"}],\"name\":\"clear\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractRegistry\",\"outputs\":[{\"internalType\":\"contractContractRegistryInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"name\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hash\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"name\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hash\",\"type\":\"uint256\"}],\"name\":\"getEntry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"entry\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"name\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"path\",\"type\":\"uint256[]\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// EVMReverseResolverV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use EVMReverseResolverV1MetaData.ABI instead.
var EVMReverseResolverV1ABI = EVMReverseResolverV1MetaData.ABI

// EVMReverseResolverV1 is an auto generated Go binding around an Ethereum contract.
type EVMReverseResolverV1 struct {
	EVMReverseResolverV1Caller     // Read-only binding to the contract
	EVMReverseResolverV1Transactor // Write-only binding to the contract
	EVMReverseResolverV1Filterer   // Log filterer for contract events
}

// EVMReverseResolverV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type EVMReverseResolverV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EVMReverseResolverV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type EVMReverseResolverV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EVMReverseResolverV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EVMReverseResolverV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EVMReverseResolverV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EVMReverseResolverV1Session struct {
	Contract     *EVMReverseResolverV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EVMReverseResolverV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EVMReverseResolverV1CallerSession struct {
	Contract *EVMReverseResolverV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// EVMReverseResolverV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EVMReverseResolverV1TransactorSession struct {
	Contract     *EVMReverseResolverV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// EVMReverseResolverV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type EVMReverseResolverV1Raw struct {
	Contract *EVMReverseResolverV1 // Generic contract binding to access the raw methods on
}

// EVMReverseResolverV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EVMReverseResolverV1CallerRaw struct {
	Contract *EVMReverseResolverV1Caller // Generic read-only contract binding to access the raw methods on
}

// EVMReverseResolverV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EVMReverseResolverV1TransactorRaw struct {
	Contract *EVMReverseResolverV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewEVMReverseResolverV1 creates a new instance of EVMReverseResolverV1, bound to a specific deployed contract.
func NewEVMReverseResolverV1(address common.Address, backend bind.ContractBackend) (*EVMReverseResolverV1, error) {
	contract, err := bindEVMReverseResolverV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EVMReverseResolverV1{EVMReverseResolverV1Caller: EVMReverseResolverV1Caller{contract: contract}, EVMReverseResolverV1Transactor: EVMReverseResolverV1Transactor{contract: contract}, EVMReverseResolverV1Filterer: EVMReverseResolverV1Filterer{contract: contract}}, nil
}

// NewEVMReverseResolverV1Caller creates a new read-only instance of EVMReverseResolverV1, bound to a specific deployed contract.
func NewEVMReverseResolverV1Caller(address common.Address, caller bind.ContractCaller) (*EVMReverseResolverV1Caller, error) {
	contract, err := bindEVMReverseResolverV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EVMReverseResolverV1Caller{contract: contract}, nil
}

// NewEVMReverseResolverV1Transactor creates a new write-only instance of EVMReverseResolverV1, bound to a specific deployed contract.
func NewEVMReverseResolverV1Transactor(address common.Address, transactor bind.ContractTransactor) (*EVMReverseResolverV1Transactor, error) {
	contract, err := bindEVMReverseResolverV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EVMReverseResolverV1Transactor{contract: contract}, nil
}

// NewEVMReverseResolverV1Filterer creates a new log filterer instance of EVMReverseResolverV1, bound to a specific deployed contract.
func NewEVMReverseResolverV1Filterer(address common.Address, filterer bind.ContractFilterer) (*EVMReverseResolverV1Filterer, error) {
	contract, err := bindEVMReverseResolverV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EVMReverseResolverV1Filterer{contract: contract}, nil
}

// bindEVMReverseResolverV1 binds a generic wrapper to an already deployed contract.
func bindEVMReverseResolverV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EVMReverseResolverV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EVMReverseResolverV1 *EVMReverseResolverV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVMReverseResolverV1.Contract.EVMReverseResolverV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EVMReverseResolverV1 *EVMReverseResolverV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVMReverseResolverV1.Contract.EVMReverseResolverV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EVMReverseResolverV1 *EVMReverseResolverV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVMReverseResolverV1.Contract.EVMReverseResolverV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EVMReverseResolverV1 *EVMReverseResolverV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVMReverseResolverV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EVMReverseResolverV1 *EVMReverseResolverV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVMReverseResolverV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EVMReverseResolverV1 *EVMReverseResolverV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVMReverseResolverV1.Contract.contract.Transact(opts, method, params...)
}

// ContractRegistry is a free data retrieval call binding the contract method 0xabf410e5.
//
// Solidity: function contractRegistry() view returns(address)
func (_EVMReverseResolverV1 *EVMReverseResolverV1Caller) ContractRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVMReverseResolverV1.contract.Call(opts, &out, "contractRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractRegistry is a free data retrieval call binding the contract method 0xabf410e5.
//
// Solidity: function contractRegistry() view returns(address)
func (_EVMReverseResolverV1 *EVMReverseResolverV1Session) ContractRegistry() (common.Address, error) {
	return _EVMReverseResolverV1.Contract.ContractRegistry(&_EVMReverseResolverV1.CallOpts)
}

// ContractRegistry is a free data retrieval call binding the contract method 0xabf410e5.
//
// Solidity: function contractRegistry() view returns(address)
func (_EVMReverseResolverV1 *EVMReverseResolverV1CallerSession) ContractRegistry() (common.Address, error) {
	return _EVMReverseResolverV1.Contract.ContractRegistry(&_EVMReverseResolverV1.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address target) view returns(uint256 name, uint256 hash)
func (_EVMReverseResolverV1 *EVMReverseResolverV1Caller) Get(opts *bind.CallOpts, target common.Address) (struct {
	Name *big.Int
	Hash *big.Int
}, error) {
	var out []interface{}
	err := _EVMReverseResolverV1.contract.Call(opts, &out, "get", target)

	outstruct := new(struct {
		Name *big.Int
		Hash *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Hash = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address target) view returns(uint256 name, uint256 hash)
func (_EVMReverseResolverV1 *EVMReverseResolverV1Session) Get(target common.Address) (struct {
	Name *big.Int
	Hash *big.Int
}, error) {
	return _EVMReverseResolverV1.Contract.Get(&_EVMReverseResolverV1.CallOpts, target)
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address target) view returns(uint256 name, uint256 hash)
func (_EVMReverseResolverV1 *EVMReverseResolverV1CallerSession) Get(target common.Address) (struct {
	Name *big.Int
	Hash *big.Int
}, error) {
	return _EVMReverseResolverV1.Contract.Get(&_EVMReverseResolverV1.CallOpts, target)
}

// GetEntry is a free data retrieval call binding the contract method 0x98ba676d.
//
// Solidity: function getEntry(uint256 name, uint256 hash) view returns(address entry)
func (_EVMReverseResolverV1 *EVMReverseResolverV1Caller) GetEntry(opts *bind.CallOpts, name *big.Int, hash *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EVMReverseResolverV1.contract.Call(opts, &out, "getEntry", name, hash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEntry is a free data retrieval call binding the contract method 0x98ba676d.
//
// Solidity: function getEntry(uint256 name, uint256 hash) view returns(address entry)
func (_EVMReverseResolverV1 *EVMReverseResolverV1Session) GetEntry(name *big.Int, hash *big.Int) (common.Address, error) {
	return _EVMReverseResolverV1.Contract.GetEntry(&_EVMReverseResolverV1.CallOpts, name, hash)
}

// GetEntry is a free data retrieval call binding the contract method 0x98ba676d.
//
// Solidity: function getEntry(uint256 name, uint256 hash) view returns(address entry)
func (_EVMReverseResolverV1 *EVMReverseResolverV1CallerSession) GetEntry(name *big.Int, hash *big.Int) (common.Address, error) {
	return _EVMReverseResolverV1.Contract.GetEntry(&_EVMReverseResolverV1.CallOpts, name, hash)
}

// Clear is a paid mutator transaction binding the contract method 0x2236d8c4.
//
// Solidity: function clear(uint256 name, uint256[] path) returns()
func (_EVMReverseResolverV1 *EVMReverseResolverV1Transactor) Clear(opts *bind.TransactOpts, name *big.Int, path []*big.Int) (*types.Transaction, error) {
	return _EVMReverseResolverV1.contract.Transact(opts, "clear", name, path)
}

// Clear is a paid mutator transaction binding the contract method 0x2236d8c4.
//
// Solidity: function clear(uint256 name, uint256[] path) returns()
func (_EVMReverseResolverV1 *EVMReverseResolverV1Session) Clear(name *big.Int, path []*big.Int) (*types.Transaction, error) {
	return _EVMReverseResolverV1.Contract.Clear(&_EVMReverseResolverV1.TransactOpts, name, path)
}

// Clear is a paid mutator transaction binding the contract method 0x2236d8c4.
//
// Solidity: function clear(uint256 name, uint256[] path) returns()
func (_EVMReverseResolverV1 *EVMReverseResolverV1TransactorSession) Clear(name *big.Int, path []*big.Int) (*types.Transaction, error) {
	return _EVMReverseResolverV1.Contract.Clear(&_EVMReverseResolverV1.TransactOpts, name, path)
}

// Set is a paid mutator transaction binding the contract method 0x946aadc6.
//
// Solidity: function set(uint256 name, uint256[] path) returns()
func (_EVMReverseResolverV1 *EVMReverseResolverV1Transactor) Set(opts *bind.TransactOpts, name *big.Int, path []*big.Int) (*types.Transaction, error) {
	return _EVMReverseResolverV1.contract.Transact(opts, "set", name, path)
}

// Set is a paid mutator transaction binding the contract method 0x946aadc6.
//
// Solidity: function set(uint256 name, uint256[] path) returns()
func (_EVMReverseResolverV1 *EVMReverseResolverV1Session) Set(name *big.Int, path []*big.Int) (*types.Transaction, error) {
	return _EVMReverseResolverV1.Contract.Set(&_EVMReverseResolverV1.TransactOpts, name, path)
}

// Set is a paid mutator transaction binding the contract method 0x946aadc6.
//
// Solidity: function set(uint256 name, uint256[] path) returns()
func (_EVMReverseResolverV1 *EVMReverseResolverV1TransactorSession) Set(name *big.Int, path []*big.Int) (*types.Transaction, error) {
	return _EVMReverseResolverV1.Contract.Set(&_EVMReverseResolverV1.TransactOpts, name, path)
}

// EVMReverseResolverV1EntrySetIterator is returned from FilterEntrySet and is used to iterate over the raw logs and unpacked data for EntrySet events raised by the EVMReverseResolverV1 contract.
type EVMReverseResolverV1EntrySetIterator struct {
	Event *EVMReverseResolverV1EntrySet // Event containing the contract specifics and raw log

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
func (it *EVMReverseResolverV1EntrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVMReverseResolverV1EntrySet)
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
		it.Event = new(EVMReverseResolverV1EntrySet)
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
func (it *EVMReverseResolverV1EntrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EVMReverseResolverV1EntrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EVMReverseResolverV1EntrySet represents a EntrySet event raised by the EVMReverseResolverV1 contract.
type EVMReverseResolverV1EntrySet struct {
	Name   *big.Int
	Path   []*big.Int
	Target common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEntrySet is a free log retrieval operation binding the contract event 0x6fb25384a611c5efa0955dd23a92c63694b342f22364c35d966e1b05ed158925.
//
// Solidity: event EntrySet(uint256 indexed name, uint256[] path, address target)
func (_EVMReverseResolverV1 *EVMReverseResolverV1Filterer) FilterEntrySet(opts *bind.FilterOpts, name []*big.Int) (*EVMReverseResolverV1EntrySetIterator, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _EVMReverseResolverV1.contract.FilterLogs(opts, "EntrySet", nameRule)
	if err != nil {
		return nil, err
	}
	return &EVMReverseResolverV1EntrySetIterator{contract: _EVMReverseResolverV1.contract, event: "EntrySet", logs: logs, sub: sub}, nil
}

// WatchEntrySet is a free log subscription operation binding the contract event 0x6fb25384a611c5efa0955dd23a92c63694b342f22364c35d966e1b05ed158925.
//
// Solidity: event EntrySet(uint256 indexed name, uint256[] path, address target)
func (_EVMReverseResolverV1 *EVMReverseResolverV1Filterer) WatchEntrySet(opts *bind.WatchOpts, sink chan<- *EVMReverseResolverV1EntrySet, name []*big.Int) (event.Subscription, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _EVMReverseResolverV1.contract.WatchLogs(opts, "EntrySet", nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EVMReverseResolverV1EntrySet)
				if err := _EVMReverseResolverV1.contract.UnpackLog(event, "EntrySet", log); err != nil {
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

// ParseEntrySet is a log parse operation binding the contract event 0x6fb25384a611c5efa0955dd23a92c63694b342f22364c35d966e1b05ed158925.
//
// Solidity: event EntrySet(uint256 indexed name, uint256[] path, address target)
func (_EVMReverseResolverV1 *EVMReverseResolverV1Filterer) ParseEntrySet(log types.Log) (*EVMReverseResolverV1EntrySet, error) {
	event := new(EVMReverseResolverV1EntrySet)
	if err := _EVMReverseResolverV1.contract.UnpackLog(event, "EntrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
