// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PreRegistrationDistributionAgent

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

// PreRegistrationDistributionAgentMetaData contains all meta data concerning the PreRegistrationDistributionAgent contract.
var PreRegistrationDistributionAgentMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"namespaceId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"registrationLength\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"names\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"constraintsProofs\",\"type\":\"bytes[]\"}],\"name\":\"distributeNames\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PreRegistrationDistributionAgentABI is the input ABI used to generate the binding from.
// Deprecated: Use PreRegistrationDistributionAgentMetaData.ABI instead.
var PreRegistrationDistributionAgentABI = PreRegistrationDistributionAgentMetaData.ABI

// PreRegistrationDistributionAgent is an auto generated Go binding around an Ethereum contract.
type PreRegistrationDistributionAgent struct {
	PreRegistrationDistributionAgentCaller     // Read-only binding to the contract
	PreRegistrationDistributionAgentTransactor // Write-only binding to the contract
	PreRegistrationDistributionAgentFilterer   // Log filterer for contract events
}

// PreRegistrationDistributionAgentCaller is an auto generated read-only Go binding around an Ethereum contract.
type PreRegistrationDistributionAgentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreRegistrationDistributionAgentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PreRegistrationDistributionAgentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreRegistrationDistributionAgentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PreRegistrationDistributionAgentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreRegistrationDistributionAgentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PreRegistrationDistributionAgentSession struct {
	Contract     *PreRegistrationDistributionAgent // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                     // Call options to use throughout this session
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// PreRegistrationDistributionAgentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PreRegistrationDistributionAgentCallerSession struct {
	Contract *PreRegistrationDistributionAgentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                           // Call options to use throughout this session
}

// PreRegistrationDistributionAgentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PreRegistrationDistributionAgentTransactorSession struct {
	Contract     *PreRegistrationDistributionAgentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                           // Transaction auth options to use throughout this session
}

// PreRegistrationDistributionAgentRaw is an auto generated low-level Go binding around an Ethereum contract.
type PreRegistrationDistributionAgentRaw struct {
	Contract *PreRegistrationDistributionAgent // Generic contract binding to access the raw methods on
}

// PreRegistrationDistributionAgentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PreRegistrationDistributionAgentCallerRaw struct {
	Contract *PreRegistrationDistributionAgentCaller // Generic read-only contract binding to access the raw methods on
}

// PreRegistrationDistributionAgentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PreRegistrationDistributionAgentTransactorRaw struct {
	Contract *PreRegistrationDistributionAgentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPreRegistrationDistributionAgent creates a new instance of PreRegistrationDistributionAgent, bound to a specific deployed contract.
func NewPreRegistrationDistributionAgent(address common.Address, backend bind.ContractBackend) (*PreRegistrationDistributionAgent, error) {
	contract, err := bindPreRegistrationDistributionAgent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PreRegistrationDistributionAgent{PreRegistrationDistributionAgentCaller: PreRegistrationDistributionAgentCaller{contract: contract}, PreRegistrationDistributionAgentTransactor: PreRegistrationDistributionAgentTransactor{contract: contract}, PreRegistrationDistributionAgentFilterer: PreRegistrationDistributionAgentFilterer{contract: contract}}, nil
}

// NewPreRegistrationDistributionAgentCaller creates a new read-only instance of PreRegistrationDistributionAgent, bound to a specific deployed contract.
func NewPreRegistrationDistributionAgentCaller(address common.Address, caller bind.ContractCaller) (*PreRegistrationDistributionAgentCaller, error) {
	contract, err := bindPreRegistrationDistributionAgent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PreRegistrationDistributionAgentCaller{contract: contract}, nil
}

// NewPreRegistrationDistributionAgentTransactor creates a new write-only instance of PreRegistrationDistributionAgent, bound to a specific deployed contract.
func NewPreRegistrationDistributionAgentTransactor(address common.Address, transactor bind.ContractTransactor) (*PreRegistrationDistributionAgentTransactor, error) {
	contract, err := bindPreRegistrationDistributionAgent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PreRegistrationDistributionAgentTransactor{contract: contract}, nil
}

// NewPreRegistrationDistributionAgentFilterer creates a new log filterer instance of PreRegistrationDistributionAgent, bound to a specific deployed contract.
func NewPreRegistrationDistributionAgentFilterer(address common.Address, filterer bind.ContractFilterer) (*PreRegistrationDistributionAgentFilterer, error) {
	contract, err := bindPreRegistrationDistributionAgent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PreRegistrationDistributionAgentFilterer{contract: contract}, nil
}

// bindPreRegistrationDistributionAgent binds a generic wrapper to an already deployed contract.
func bindPreRegistrationDistributionAgent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PreRegistrationDistributionAgentMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PreRegistrationDistributionAgent.Contract.PreRegistrationDistributionAgentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PreRegistrationDistributionAgent.Contract.PreRegistrationDistributionAgentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PreRegistrationDistributionAgent.Contract.PreRegistrationDistributionAgentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PreRegistrationDistributionAgent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PreRegistrationDistributionAgent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PreRegistrationDistributionAgent.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PreRegistrationDistributionAgent.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PreRegistrationDistributionAgent.Contract.DEFAULTADMINROLE(&_PreRegistrationDistributionAgent.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PreRegistrationDistributionAgent.Contract.DEFAULTADMINROLE(&_PreRegistrationDistributionAgent.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PreRegistrationDistributionAgent.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PreRegistrationDistributionAgent.Contract.GetRoleAdmin(&_PreRegistrationDistributionAgent.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PreRegistrationDistributionAgent.Contract.GetRoleAdmin(&_PreRegistrationDistributionAgent.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _PreRegistrationDistributionAgent.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PreRegistrationDistributionAgent.Contract.HasRole(&_PreRegistrationDistributionAgent.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PreRegistrationDistributionAgent.Contract.HasRole(&_PreRegistrationDistributionAgent.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PreRegistrationDistributionAgent.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PreRegistrationDistributionAgent.Contract.SupportsInterface(&_PreRegistrationDistributionAgent.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PreRegistrationDistributionAgent.Contract.SupportsInterface(&_PreRegistrationDistributionAgent.CallOpts, interfaceId)
}

// DistributeNames is a paid mutator transaction binding the contract method 0x6757fb69.
//
// Solidity: function distributeNames(address[] recipients, uint256[] names, bytes[] constraintsProofs) returns()
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentTransactor) DistributeNames(opts *bind.TransactOpts, recipients []common.Address, names []*big.Int, constraintsProofs [][]byte) (*types.Transaction, error) {
	return _PreRegistrationDistributionAgent.contract.Transact(opts, "distributeNames", recipients, names, constraintsProofs)
}

// DistributeNames is a paid mutator transaction binding the contract method 0x6757fb69.
//
// Solidity: function distributeNames(address[] recipients, uint256[] names, bytes[] constraintsProofs) returns()
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentSession) DistributeNames(recipients []common.Address, names []*big.Int, constraintsProofs [][]byte) (*types.Transaction, error) {
	return _PreRegistrationDistributionAgent.Contract.DistributeNames(&_PreRegistrationDistributionAgent.TransactOpts, recipients, names, constraintsProofs)
}

// DistributeNames is a paid mutator transaction binding the contract method 0x6757fb69.
//
// Solidity: function distributeNames(address[] recipients, uint256[] names, bytes[] constraintsProofs) returns()
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentTransactorSession) DistributeNames(recipients []common.Address, names []*big.Int, constraintsProofs [][]byte) (*types.Transaction, error) {
	return _PreRegistrationDistributionAgent.Contract.DistributeNames(&_PreRegistrationDistributionAgent.TransactOpts, recipients, names, constraintsProofs)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PreRegistrationDistributionAgent.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PreRegistrationDistributionAgent.Contract.GrantRole(&_PreRegistrationDistributionAgent.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PreRegistrationDistributionAgent.Contract.GrantRole(&_PreRegistrationDistributionAgent.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PreRegistrationDistributionAgent.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PreRegistrationDistributionAgent.Contract.RenounceRole(&_PreRegistrationDistributionAgent.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PreRegistrationDistributionAgent.Contract.RenounceRole(&_PreRegistrationDistributionAgent.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PreRegistrationDistributionAgent.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PreRegistrationDistributionAgent.Contract.RevokeRole(&_PreRegistrationDistributionAgent.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PreRegistrationDistributionAgent.Contract.RevokeRole(&_PreRegistrationDistributionAgent.TransactOpts, role, account)
}

// PreRegistrationDistributionAgentRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the PreRegistrationDistributionAgent contract.
type PreRegistrationDistributionAgentRoleAdminChangedIterator struct {
	Event *PreRegistrationDistributionAgentRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PreRegistrationDistributionAgentRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreRegistrationDistributionAgentRoleAdminChanged)
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
		it.Event = new(PreRegistrationDistributionAgentRoleAdminChanged)
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
func (it *PreRegistrationDistributionAgentRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreRegistrationDistributionAgentRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreRegistrationDistributionAgentRoleAdminChanged represents a RoleAdminChanged event raised by the PreRegistrationDistributionAgent contract.
type PreRegistrationDistributionAgentRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PreRegistrationDistributionAgentRoleAdminChangedIterator, error) {

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

	logs, sub, err := _PreRegistrationDistributionAgent.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PreRegistrationDistributionAgentRoleAdminChangedIterator{contract: _PreRegistrationDistributionAgent.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PreRegistrationDistributionAgentRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _PreRegistrationDistributionAgent.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreRegistrationDistributionAgentRoleAdminChanged)
				if err := _PreRegistrationDistributionAgent.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentFilterer) ParseRoleAdminChanged(log types.Log) (*PreRegistrationDistributionAgentRoleAdminChanged, error) {
	event := new(PreRegistrationDistributionAgentRoleAdminChanged)
	if err := _PreRegistrationDistributionAgent.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PreRegistrationDistributionAgentRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the PreRegistrationDistributionAgent contract.
type PreRegistrationDistributionAgentRoleGrantedIterator struct {
	Event *PreRegistrationDistributionAgentRoleGranted // Event containing the contract specifics and raw log

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
func (it *PreRegistrationDistributionAgentRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreRegistrationDistributionAgentRoleGranted)
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
		it.Event = new(PreRegistrationDistributionAgentRoleGranted)
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
func (it *PreRegistrationDistributionAgentRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreRegistrationDistributionAgentRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreRegistrationDistributionAgentRoleGranted represents a RoleGranted event raised by the PreRegistrationDistributionAgent contract.
type PreRegistrationDistributionAgentRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PreRegistrationDistributionAgentRoleGrantedIterator, error) {

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

	logs, sub, err := _PreRegistrationDistributionAgent.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PreRegistrationDistributionAgentRoleGrantedIterator{contract: _PreRegistrationDistributionAgent.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PreRegistrationDistributionAgentRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PreRegistrationDistributionAgent.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreRegistrationDistributionAgentRoleGranted)
				if err := _PreRegistrationDistributionAgent.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentFilterer) ParseRoleGranted(log types.Log) (*PreRegistrationDistributionAgentRoleGranted, error) {
	event := new(PreRegistrationDistributionAgentRoleGranted)
	if err := _PreRegistrationDistributionAgent.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PreRegistrationDistributionAgentRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the PreRegistrationDistributionAgent contract.
type PreRegistrationDistributionAgentRoleRevokedIterator struct {
	Event *PreRegistrationDistributionAgentRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PreRegistrationDistributionAgentRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreRegistrationDistributionAgentRoleRevoked)
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
		it.Event = new(PreRegistrationDistributionAgentRoleRevoked)
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
func (it *PreRegistrationDistributionAgentRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreRegistrationDistributionAgentRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreRegistrationDistributionAgentRoleRevoked represents a RoleRevoked event raised by the PreRegistrationDistributionAgent contract.
type PreRegistrationDistributionAgentRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PreRegistrationDistributionAgentRoleRevokedIterator, error) {

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

	logs, sub, err := _PreRegistrationDistributionAgent.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PreRegistrationDistributionAgentRoleRevokedIterator{contract: _PreRegistrationDistributionAgent.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PreRegistrationDistributionAgentRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PreRegistrationDistributionAgent.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreRegistrationDistributionAgentRoleRevoked)
				if err := _PreRegistrationDistributionAgent.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_PreRegistrationDistributionAgent *PreRegistrationDistributionAgentFilterer) ParseRoleRevoked(log types.Log) (*PreRegistrationDistributionAgentRoleRevoked, error) {
	event := new(PreRegistrationDistributionAgentRoleRevoked)
	if err := _PreRegistrationDistributionAgent.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
