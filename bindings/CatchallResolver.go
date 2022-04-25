// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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
)

// CatchallResolverMetaData contains all meta data concerning the CatchallResolver contract.
var CatchallResolverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractResolver\",\"name\":\"_parent\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_node\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"AddrChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"addr\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"parent\",\"outputs\":[{\"internalType\":\"contractResolver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"parentNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"3b3b57de": "addr(bytes32)",
		"60f96a8f": "parent()",
		"f3068a00": "parentNode()",
		"9061b923": "resolve(bytes,bytes)",
		"01ffc9a7": "supportsInterface(bytes4)",
	},
	Bin: "0x608060405234801561001057600080fd5b506040516105f43803806105f483398101604081905261002f91610058565b600080546001600160a01b0319166001600160a01b039390931692909217909155600155610092565b6000806040838503121561006b57600080fd5b82516001600160a01b038116811461008257600080fd5b6020939093015192949293505050565b610553806100a16000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806301ffc9a71461005c5780633b3b57de1461009557806360f96a8f146100c05780639061b923146100d3578063f3068a00146100f3575b600080fd5b61008061006a36600461029f565b6001600160e01b031916639061b92360e01b1490565b60405190151581526020015b60405180910390f35b6100a86100a33660046102d0565b61010a565b6040516001600160a01b03909116815260200161008c565b6000546100a8906001600160a01b031681565b6100e66100e13660046102ff565b61017e565b60405161008c9190610435565b6100fc60015481565b60405190815260200161008c565b60008054604051631d9dabef60e11b8152600481018490526001600160a01b0390911690633b3b57de90602401602060405180830381865afa158015610154573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101789190610468565b92915050565b606060005b60208160ff1610156101f2576001548160ff16602081106101a6576101a6610491565b1a60f81b836101b68360046104bd565b60ff16815181106101c9576101c9610491565b60200101906001600160f81b031916908160001a905350806101ea816104e2565b915050610183565b506000805460405182916001600160a01b031690610211908690610501565b600060405180830381855afa9150503d806000811461024c576040519150601f19603f3d011682016040523d82523d6000602084013e610251565b606091505b5091509150816102965760405162461bcd60e51b815260206004820152600c60248201526b1a5b9d985b1a590818d85b1b60a21b604482015260640160405180910390fd5b95945050505050565b6000602082840312156102b157600080fd5b81356001600160e01b0319811681146102c957600080fd5b9392505050565b6000602082840312156102e257600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b60008060006040848603121561031457600080fd5b833567ffffffffffffffff8082111561032c57600080fd5b818601915086601f83011261034057600080fd5b81358181111561034f57600080fd5b6020888183860101111561036257600080fd5b80840196508195508088013593508284111561037d57600080fd5b838801935088601f85011261039157600080fd5b83359150828211156103a5576103a56102e9565b604051601f8301601f19908116603f011681019084821181831017156103cd576103cd6102e9565b816040528381528a838588010111156103e557600080fd5b838387018483013760008385830101528096505050505050509250925092565b60005b83811015610420578181015183820152602001610408565b8381111561042f576000848401525b50505050565b6020815260008251806020840152610454816040850160208701610405565b601f01601f19169190910160400192915050565b60006020828403121561047a57600080fd5b81516001600160a01b03811681146102c957600080fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060ff821660ff84168060ff038211156104da576104da6104a7565b019392505050565b600060ff821660ff81036104f8576104f86104a7565b60010192915050565b60008251610513818460208701610405565b919091019291505056fea26469706673582212209f7ce091a2a2abf72ab4b8d605b705e352e512aaa0dedcfba992ba4cb7422f7d64736f6c634300080d0033",
}

// CatchallResolverABI is the input ABI used to generate the binding from.
// Deprecated: Use CatchallResolverMetaData.ABI instead.
var CatchallResolverABI = CatchallResolverMetaData.ABI

// Deprecated: Use CatchallResolverMetaData.Sigs instead.
// CatchallResolverFuncSigs maps the 4-byte function signature to its string representation.
var CatchallResolverFuncSigs = CatchallResolverMetaData.Sigs

// CatchallResolverBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CatchallResolverMetaData.Bin instead.
var CatchallResolverBin = CatchallResolverMetaData.Bin

// DeployCatchallResolver deploys a new Ethereum contract, binding an instance of CatchallResolver to it.
func DeployCatchallResolver(auth *bind.TransactOpts, backend bind.ContractBackend, _parent common.Address, _node [32]byte) (common.Address, *types.Transaction, *CatchallResolver, error) {
	parsed, err := CatchallResolverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CatchallResolverBin), backend, _parent, _node)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CatchallResolver{CatchallResolverCaller: CatchallResolverCaller{contract: contract}, CatchallResolverTransactor: CatchallResolverTransactor{contract: contract}, CatchallResolverFilterer: CatchallResolverFilterer{contract: contract}}, nil
}

// CatchallResolver is an auto generated Go binding around an Ethereum contract.
type CatchallResolver struct {
	CatchallResolverCaller     // Read-only binding to the contract
	CatchallResolverTransactor // Write-only binding to the contract
	CatchallResolverFilterer   // Log filterer for contract events
}

// CatchallResolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type CatchallResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CatchallResolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CatchallResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CatchallResolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CatchallResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CatchallResolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CatchallResolverSession struct {
	Contract     *CatchallResolver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CatchallResolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CatchallResolverCallerSession struct {
	Contract *CatchallResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// CatchallResolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CatchallResolverTransactorSession struct {
	Contract     *CatchallResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// CatchallResolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type CatchallResolverRaw struct {
	Contract *CatchallResolver // Generic contract binding to access the raw methods on
}

// CatchallResolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CatchallResolverCallerRaw struct {
	Contract *CatchallResolverCaller // Generic read-only contract binding to access the raw methods on
}

// CatchallResolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CatchallResolverTransactorRaw struct {
	Contract *CatchallResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCatchallResolver creates a new instance of CatchallResolver, bound to a specific deployed contract.
func NewCatchallResolver(address common.Address, backend bind.ContractBackend) (*CatchallResolver, error) {
	contract, err := bindCatchallResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CatchallResolver{CatchallResolverCaller: CatchallResolverCaller{contract: contract}, CatchallResolverTransactor: CatchallResolverTransactor{contract: contract}, CatchallResolverFilterer: CatchallResolverFilterer{contract: contract}}, nil
}

// NewCatchallResolverCaller creates a new read-only instance of CatchallResolver, bound to a specific deployed contract.
func NewCatchallResolverCaller(address common.Address, caller bind.ContractCaller) (*CatchallResolverCaller, error) {
	contract, err := bindCatchallResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CatchallResolverCaller{contract: contract}, nil
}

// NewCatchallResolverTransactor creates a new write-only instance of CatchallResolver, bound to a specific deployed contract.
func NewCatchallResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*CatchallResolverTransactor, error) {
	contract, err := bindCatchallResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CatchallResolverTransactor{contract: contract}, nil
}

// NewCatchallResolverFilterer creates a new log filterer instance of CatchallResolver, bound to a specific deployed contract.
func NewCatchallResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*CatchallResolverFilterer, error) {
	contract, err := bindCatchallResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CatchallResolverFilterer{contract: contract}, nil
}

// bindCatchallResolver binds a generic wrapper to an already deployed contract.
func bindCatchallResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CatchallResolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CatchallResolver *CatchallResolverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CatchallResolver.Contract.CatchallResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CatchallResolver *CatchallResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CatchallResolver.Contract.CatchallResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CatchallResolver *CatchallResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CatchallResolver.Contract.CatchallResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CatchallResolver *CatchallResolverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CatchallResolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CatchallResolver *CatchallResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CatchallResolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CatchallResolver *CatchallResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CatchallResolver.Contract.contract.Transact(opts, method, params...)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_CatchallResolver *CatchallResolverCaller) Addr(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var out []interface{}
	err := _CatchallResolver.contract.Call(opts, &out, "addr", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_CatchallResolver *CatchallResolverSession) Addr(node [32]byte) (common.Address, error) {
	return _CatchallResolver.Contract.Addr(&_CatchallResolver.CallOpts, node)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_CatchallResolver *CatchallResolverCallerSession) Addr(node [32]byte) (common.Address, error) {
	return _CatchallResolver.Contract.Addr(&_CatchallResolver.CallOpts, node)
}

// Parent is a free data retrieval call binding the contract method 0x60f96a8f.
//
// Solidity: function parent() view returns(address)
func (_CatchallResolver *CatchallResolverCaller) Parent(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CatchallResolver.contract.Call(opts, &out, "parent")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Parent is a free data retrieval call binding the contract method 0x60f96a8f.
//
// Solidity: function parent() view returns(address)
func (_CatchallResolver *CatchallResolverSession) Parent() (common.Address, error) {
	return _CatchallResolver.Contract.Parent(&_CatchallResolver.CallOpts)
}

// Parent is a free data retrieval call binding the contract method 0x60f96a8f.
//
// Solidity: function parent() view returns(address)
func (_CatchallResolver *CatchallResolverCallerSession) Parent() (common.Address, error) {
	return _CatchallResolver.Contract.Parent(&_CatchallResolver.CallOpts)
}

// ParentNode is a free data retrieval call binding the contract method 0xf3068a00.
//
// Solidity: function parentNode() view returns(bytes32)
func (_CatchallResolver *CatchallResolverCaller) ParentNode(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CatchallResolver.contract.Call(opts, &out, "parentNode")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ParentNode is a free data retrieval call binding the contract method 0xf3068a00.
//
// Solidity: function parentNode() view returns(bytes32)
func (_CatchallResolver *CatchallResolverSession) ParentNode() ([32]byte, error) {
	return _CatchallResolver.Contract.ParentNode(&_CatchallResolver.CallOpts)
}

// ParentNode is a free data retrieval call binding the contract method 0xf3068a00.
//
// Solidity: function parentNode() view returns(bytes32)
func (_CatchallResolver *CatchallResolverCallerSession) ParentNode() ([32]byte, error) {
	return _CatchallResolver.Contract.ParentNode(&_CatchallResolver.CallOpts)
}

// Resolve is a free data retrieval call binding the contract method 0x9061b923.
//
// Solidity: function resolve(bytes , bytes data) view returns(bytes)
func (_CatchallResolver *CatchallResolverCaller) Resolve(opts *bind.CallOpts, arg0 []byte, data []byte) ([]byte, error) {
	var out []interface{}
	err := _CatchallResolver.contract.Call(opts, &out, "resolve", arg0, data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x9061b923.
//
// Solidity: function resolve(bytes , bytes data) view returns(bytes)
func (_CatchallResolver *CatchallResolverSession) Resolve(arg0 []byte, data []byte) ([]byte, error) {
	return _CatchallResolver.Contract.Resolve(&_CatchallResolver.CallOpts, arg0, data)
}

// Resolve is a free data retrieval call binding the contract method 0x9061b923.
//
// Solidity: function resolve(bytes , bytes data) view returns(bytes)
func (_CatchallResolver *CatchallResolverCallerSession) Resolve(arg0 []byte, data []byte) ([]byte, error) {
	return _CatchallResolver.Contract.Resolve(&_CatchallResolver.CallOpts, arg0, data)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_CatchallResolver *CatchallResolverCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _CatchallResolver.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_CatchallResolver *CatchallResolverSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _CatchallResolver.Contract.SupportsInterface(&_CatchallResolver.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_CatchallResolver *CatchallResolverCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _CatchallResolver.Contract.SupportsInterface(&_CatchallResolver.CallOpts, interfaceID)
}

// CatchallResolverAddrChangedIterator is returned from FilterAddrChanged and is used to iterate over the raw logs and unpacked data for AddrChanged events raised by the CatchallResolver contract.
type CatchallResolverAddrChangedIterator struct {
	Event *CatchallResolverAddrChanged // Event containing the contract specifics and raw log

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
func (it *CatchallResolverAddrChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CatchallResolverAddrChanged)
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
		it.Event = new(CatchallResolverAddrChanged)
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
func (it *CatchallResolverAddrChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CatchallResolverAddrChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CatchallResolverAddrChanged represents a AddrChanged event raised by the CatchallResolver contract.
type CatchallResolverAddrChanged struct {
	Node [32]byte
	A    common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddrChanged is a free log retrieval operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(bytes32 indexed node, address a)
func (_CatchallResolver *CatchallResolverFilterer) FilterAddrChanged(opts *bind.FilterOpts, node [][32]byte) (*CatchallResolverAddrChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _CatchallResolver.contract.FilterLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &CatchallResolverAddrChangedIterator{contract: _CatchallResolver.contract, event: "AddrChanged", logs: logs, sub: sub}, nil
}

// WatchAddrChanged is a free log subscription operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(bytes32 indexed node, address a)
func (_CatchallResolver *CatchallResolverFilterer) WatchAddrChanged(opts *bind.WatchOpts, sink chan<- *CatchallResolverAddrChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _CatchallResolver.contract.WatchLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CatchallResolverAddrChanged)
				if err := _CatchallResolver.contract.UnpackLog(event, "AddrChanged", log); err != nil {
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

// ParseAddrChanged is a log parse operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(bytes32 indexed node, address a)
func (_CatchallResolver *CatchallResolverFilterer) ParseAddrChanged(log types.Log) (*CatchallResolverAddrChanged, error) {
	event := new(CatchallResolverAddrChanged)
	if err := _CatchallResolver.contract.UnpackLog(event, "AddrChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAddrResolverMetaData contains all meta data concerning the IAddrResolver contract.
var IAddrResolverMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"AddrChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"addr\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"3b3b57de": "addr(bytes32)",
	},
}

// IAddrResolverABI is the input ABI used to generate the binding from.
// Deprecated: Use IAddrResolverMetaData.ABI instead.
var IAddrResolverABI = IAddrResolverMetaData.ABI

// Deprecated: Use IAddrResolverMetaData.Sigs instead.
// IAddrResolverFuncSigs maps the 4-byte function signature to its string representation.
var IAddrResolverFuncSigs = IAddrResolverMetaData.Sigs

// IAddrResolver is an auto generated Go binding around an Ethereum contract.
type IAddrResolver struct {
	IAddrResolverCaller     // Read-only binding to the contract
	IAddrResolverTransactor // Write-only binding to the contract
	IAddrResolverFilterer   // Log filterer for contract events
}

// IAddrResolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAddrResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAddrResolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAddrResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAddrResolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAddrResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAddrResolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAddrResolverSession struct {
	Contract     *IAddrResolver    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAddrResolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAddrResolverCallerSession struct {
	Contract *IAddrResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IAddrResolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAddrResolverTransactorSession struct {
	Contract     *IAddrResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IAddrResolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAddrResolverRaw struct {
	Contract *IAddrResolver // Generic contract binding to access the raw methods on
}

// IAddrResolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAddrResolverCallerRaw struct {
	Contract *IAddrResolverCaller // Generic read-only contract binding to access the raw methods on
}

// IAddrResolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAddrResolverTransactorRaw struct {
	Contract *IAddrResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAddrResolver creates a new instance of IAddrResolver, bound to a specific deployed contract.
func NewIAddrResolver(address common.Address, backend bind.ContractBackend) (*IAddrResolver, error) {
	contract, err := bindIAddrResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAddrResolver{IAddrResolverCaller: IAddrResolverCaller{contract: contract}, IAddrResolverTransactor: IAddrResolverTransactor{contract: contract}, IAddrResolverFilterer: IAddrResolverFilterer{contract: contract}}, nil
}

// NewIAddrResolverCaller creates a new read-only instance of IAddrResolver, bound to a specific deployed contract.
func NewIAddrResolverCaller(address common.Address, caller bind.ContractCaller) (*IAddrResolverCaller, error) {
	contract, err := bindIAddrResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAddrResolverCaller{contract: contract}, nil
}

// NewIAddrResolverTransactor creates a new write-only instance of IAddrResolver, bound to a specific deployed contract.
func NewIAddrResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*IAddrResolverTransactor, error) {
	contract, err := bindIAddrResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAddrResolverTransactor{contract: contract}, nil
}

// NewIAddrResolverFilterer creates a new log filterer instance of IAddrResolver, bound to a specific deployed contract.
func NewIAddrResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*IAddrResolverFilterer, error) {
	contract, err := bindIAddrResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAddrResolverFilterer{contract: contract}, nil
}

// bindIAddrResolver binds a generic wrapper to an already deployed contract.
func bindIAddrResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAddrResolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAddrResolver *IAddrResolverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAddrResolver.Contract.IAddrResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAddrResolver *IAddrResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAddrResolver.Contract.IAddrResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAddrResolver *IAddrResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAddrResolver.Contract.IAddrResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAddrResolver *IAddrResolverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAddrResolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAddrResolver *IAddrResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAddrResolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAddrResolver *IAddrResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAddrResolver.Contract.contract.Transact(opts, method, params...)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_IAddrResolver *IAddrResolverCaller) Addr(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IAddrResolver.contract.Call(opts, &out, "addr", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_IAddrResolver *IAddrResolverSession) Addr(node [32]byte) (common.Address, error) {
	return _IAddrResolver.Contract.Addr(&_IAddrResolver.CallOpts, node)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_IAddrResolver *IAddrResolverCallerSession) Addr(node [32]byte) (common.Address, error) {
	return _IAddrResolver.Contract.Addr(&_IAddrResolver.CallOpts, node)
}

// IAddrResolverAddrChangedIterator is returned from FilterAddrChanged and is used to iterate over the raw logs and unpacked data for AddrChanged events raised by the IAddrResolver contract.
type IAddrResolverAddrChangedIterator struct {
	Event *IAddrResolverAddrChanged // Event containing the contract specifics and raw log

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
func (it *IAddrResolverAddrChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAddrResolverAddrChanged)
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
		it.Event = new(IAddrResolverAddrChanged)
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
func (it *IAddrResolverAddrChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAddrResolverAddrChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAddrResolverAddrChanged represents a AddrChanged event raised by the IAddrResolver contract.
type IAddrResolverAddrChanged struct {
	Node [32]byte
	A    common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddrChanged is a free log retrieval operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(bytes32 indexed node, address a)
func (_IAddrResolver *IAddrResolverFilterer) FilterAddrChanged(opts *bind.FilterOpts, node [][32]byte) (*IAddrResolverAddrChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _IAddrResolver.contract.FilterLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &IAddrResolverAddrChangedIterator{contract: _IAddrResolver.contract, event: "AddrChanged", logs: logs, sub: sub}, nil
}

// WatchAddrChanged is a free log subscription operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(bytes32 indexed node, address a)
func (_IAddrResolver *IAddrResolverFilterer) WatchAddrChanged(opts *bind.WatchOpts, sink chan<- *IAddrResolverAddrChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _IAddrResolver.contract.WatchLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAddrResolverAddrChanged)
				if err := _IAddrResolver.contract.UnpackLog(event, "AddrChanged", log); err != nil {
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

// ParseAddrChanged is a log parse operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(bytes32 indexed node, address a)
func (_IAddrResolver *IAddrResolverFilterer) ParseAddrChanged(log types.Log) (*IAddrResolverAddrChanged, error) {
	event := new(IAddrResolverAddrChanged)
	if err := _IAddrResolver.contract.UnpackLog(event, "AddrChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IExtendedResolverMetaData contains all meta data concerning the IExtendedResolver contract.
var IExtendedResolverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"9061b923": "resolve(bytes,bytes)",
	},
}

// IExtendedResolverABI is the input ABI used to generate the binding from.
// Deprecated: Use IExtendedResolverMetaData.ABI instead.
var IExtendedResolverABI = IExtendedResolverMetaData.ABI

// Deprecated: Use IExtendedResolverMetaData.Sigs instead.
// IExtendedResolverFuncSigs maps the 4-byte function signature to its string representation.
var IExtendedResolverFuncSigs = IExtendedResolverMetaData.Sigs

// IExtendedResolver is an auto generated Go binding around an Ethereum contract.
type IExtendedResolver struct {
	IExtendedResolverCaller     // Read-only binding to the contract
	IExtendedResolverTransactor // Write-only binding to the contract
	IExtendedResolverFilterer   // Log filterer for contract events
}

// IExtendedResolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IExtendedResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExtendedResolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IExtendedResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExtendedResolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IExtendedResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExtendedResolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IExtendedResolverSession struct {
	Contract     *IExtendedResolver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IExtendedResolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IExtendedResolverCallerSession struct {
	Contract *IExtendedResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IExtendedResolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IExtendedResolverTransactorSession struct {
	Contract     *IExtendedResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IExtendedResolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IExtendedResolverRaw struct {
	Contract *IExtendedResolver // Generic contract binding to access the raw methods on
}

// IExtendedResolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IExtendedResolverCallerRaw struct {
	Contract *IExtendedResolverCaller // Generic read-only contract binding to access the raw methods on
}

// IExtendedResolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IExtendedResolverTransactorRaw struct {
	Contract *IExtendedResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIExtendedResolver creates a new instance of IExtendedResolver, bound to a specific deployed contract.
func NewIExtendedResolver(address common.Address, backend bind.ContractBackend) (*IExtendedResolver, error) {
	contract, err := bindIExtendedResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IExtendedResolver{IExtendedResolverCaller: IExtendedResolverCaller{contract: contract}, IExtendedResolverTransactor: IExtendedResolverTransactor{contract: contract}, IExtendedResolverFilterer: IExtendedResolverFilterer{contract: contract}}, nil
}

// NewIExtendedResolverCaller creates a new read-only instance of IExtendedResolver, bound to a specific deployed contract.
func NewIExtendedResolverCaller(address common.Address, caller bind.ContractCaller) (*IExtendedResolverCaller, error) {
	contract, err := bindIExtendedResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IExtendedResolverCaller{contract: contract}, nil
}

// NewIExtendedResolverTransactor creates a new write-only instance of IExtendedResolver, bound to a specific deployed contract.
func NewIExtendedResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*IExtendedResolverTransactor, error) {
	contract, err := bindIExtendedResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IExtendedResolverTransactor{contract: contract}, nil
}

// NewIExtendedResolverFilterer creates a new log filterer instance of IExtendedResolver, bound to a specific deployed contract.
func NewIExtendedResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*IExtendedResolverFilterer, error) {
	contract, err := bindIExtendedResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IExtendedResolverFilterer{contract: contract}, nil
}

// bindIExtendedResolver binds a generic wrapper to an already deployed contract.
func bindIExtendedResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IExtendedResolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IExtendedResolver *IExtendedResolverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IExtendedResolver.Contract.IExtendedResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IExtendedResolver *IExtendedResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExtendedResolver.Contract.IExtendedResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IExtendedResolver *IExtendedResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IExtendedResolver.Contract.IExtendedResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IExtendedResolver *IExtendedResolverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IExtendedResolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IExtendedResolver *IExtendedResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExtendedResolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IExtendedResolver *IExtendedResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IExtendedResolver.Contract.contract.Transact(opts, method, params...)
}

// Resolve is a free data retrieval call binding the contract method 0x9061b923.
//
// Solidity: function resolve(bytes name, bytes data) view returns(bytes)
func (_IExtendedResolver *IExtendedResolverCaller) Resolve(opts *bind.CallOpts, name []byte, data []byte) ([]byte, error) {
	var out []interface{}
	err := _IExtendedResolver.contract.Call(opts, &out, "resolve", name, data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x9061b923.
//
// Solidity: function resolve(bytes name, bytes data) view returns(bytes)
func (_IExtendedResolver *IExtendedResolverSession) Resolve(name []byte, data []byte) ([]byte, error) {
	return _IExtendedResolver.Contract.Resolve(&_IExtendedResolver.CallOpts, name, data)
}

// Resolve is a free data retrieval call binding the contract method 0x9061b923.
//
// Solidity: function resolve(bytes name, bytes data) view returns(bytes)
func (_IExtendedResolver *IExtendedResolverCallerSession) Resolve(name []byte, data []byte) ([]byte, error) {
	return _IExtendedResolver.Contract.Resolve(&_IExtendedResolver.CallOpts, name, data)
}

// ISupportsInterfaceMetaData contains all meta data concerning the ISupportsInterface contract.
var ISupportsInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"01ffc9a7": "supportsInterface(bytes4)",
	},
}

// ISupportsInterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use ISupportsInterfaceMetaData.ABI instead.
var ISupportsInterfaceABI = ISupportsInterfaceMetaData.ABI

// Deprecated: Use ISupportsInterfaceMetaData.Sigs instead.
// ISupportsInterfaceFuncSigs maps the 4-byte function signature to its string representation.
var ISupportsInterfaceFuncSigs = ISupportsInterfaceMetaData.Sigs

// ISupportsInterface is an auto generated Go binding around an Ethereum contract.
type ISupportsInterface struct {
	ISupportsInterfaceCaller     // Read-only binding to the contract
	ISupportsInterfaceTransactor // Write-only binding to the contract
	ISupportsInterfaceFilterer   // Log filterer for contract events
}

// ISupportsInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISupportsInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISupportsInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISupportsInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISupportsInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISupportsInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISupportsInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISupportsInterfaceSession struct {
	Contract     *ISupportsInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ISupportsInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISupportsInterfaceCallerSession struct {
	Contract *ISupportsInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ISupportsInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISupportsInterfaceTransactorSession struct {
	Contract     *ISupportsInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ISupportsInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISupportsInterfaceRaw struct {
	Contract *ISupportsInterface // Generic contract binding to access the raw methods on
}

// ISupportsInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISupportsInterfaceCallerRaw struct {
	Contract *ISupportsInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// ISupportsInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISupportsInterfaceTransactorRaw struct {
	Contract *ISupportsInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISupportsInterface creates a new instance of ISupportsInterface, bound to a specific deployed contract.
func NewISupportsInterface(address common.Address, backend bind.ContractBackend) (*ISupportsInterface, error) {
	contract, err := bindISupportsInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISupportsInterface{ISupportsInterfaceCaller: ISupportsInterfaceCaller{contract: contract}, ISupportsInterfaceTransactor: ISupportsInterfaceTransactor{contract: contract}, ISupportsInterfaceFilterer: ISupportsInterfaceFilterer{contract: contract}}, nil
}

// NewISupportsInterfaceCaller creates a new read-only instance of ISupportsInterface, bound to a specific deployed contract.
func NewISupportsInterfaceCaller(address common.Address, caller bind.ContractCaller) (*ISupportsInterfaceCaller, error) {
	contract, err := bindISupportsInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISupportsInterfaceCaller{contract: contract}, nil
}

// NewISupportsInterfaceTransactor creates a new write-only instance of ISupportsInterface, bound to a specific deployed contract.
func NewISupportsInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*ISupportsInterfaceTransactor, error) {
	contract, err := bindISupportsInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISupportsInterfaceTransactor{contract: contract}, nil
}

// NewISupportsInterfaceFilterer creates a new log filterer instance of ISupportsInterface, bound to a specific deployed contract.
func NewISupportsInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*ISupportsInterfaceFilterer, error) {
	contract, err := bindISupportsInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISupportsInterfaceFilterer{contract: contract}, nil
}

// bindISupportsInterface binds a generic wrapper to an already deployed contract.
func bindISupportsInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISupportsInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISupportsInterface *ISupportsInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISupportsInterface.Contract.ISupportsInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISupportsInterface *ISupportsInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISupportsInterface.Contract.ISupportsInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISupportsInterface *ISupportsInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISupportsInterface.Contract.ISupportsInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISupportsInterface *ISupportsInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISupportsInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISupportsInterface *ISupportsInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISupportsInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISupportsInterface *ISupportsInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISupportsInterface.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_ISupportsInterface *ISupportsInterfaceCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _ISupportsInterface.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_ISupportsInterface *ISupportsInterfaceSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _ISupportsInterface.Contract.SupportsInterface(&_ISupportsInterface.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_ISupportsInterface *ISupportsInterfaceCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _ISupportsInterface.Contract.SupportsInterface(&_ISupportsInterface.CallOpts, interfaceID)
}

// ResolverMetaData contains all meta data concerning the Resolver contract.
var ResolverMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"AddrChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"addr\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"3b3b57de": "addr(bytes32)",
		"01ffc9a7": "supportsInterface(bytes4)",
	},
}

// ResolverABI is the input ABI used to generate the binding from.
// Deprecated: Use ResolverMetaData.ABI instead.
var ResolverABI = ResolverMetaData.ABI

// Deprecated: Use ResolverMetaData.Sigs instead.
// ResolverFuncSigs maps the 4-byte function signature to its string representation.
var ResolverFuncSigs = ResolverMetaData.Sigs

// Resolver is an auto generated Go binding around an Ethereum contract.
type Resolver struct {
	ResolverCaller     // Read-only binding to the contract
	ResolverTransactor // Write-only binding to the contract
	ResolverFilterer   // Log filterer for contract events
}

// ResolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ResolverSession struct {
	Contract     *Resolver         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ResolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ResolverCallerSession struct {
	Contract *ResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ResolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ResolverTransactorSession struct {
	Contract     *ResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ResolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ResolverRaw struct {
	Contract *Resolver // Generic contract binding to access the raw methods on
}

// ResolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ResolverCallerRaw struct {
	Contract *ResolverCaller // Generic read-only contract binding to access the raw methods on
}

// ResolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ResolverTransactorRaw struct {
	Contract *ResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewResolver creates a new instance of Resolver, bound to a specific deployed contract.
func NewResolver(address common.Address, backend bind.ContractBackend) (*Resolver, error) {
	contract, err := bindResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Resolver{ResolverCaller: ResolverCaller{contract: contract}, ResolverTransactor: ResolverTransactor{contract: contract}, ResolverFilterer: ResolverFilterer{contract: contract}}, nil
}

// NewResolverCaller creates a new read-only instance of Resolver, bound to a specific deployed contract.
func NewResolverCaller(address common.Address, caller bind.ContractCaller) (*ResolverCaller, error) {
	contract, err := bindResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ResolverCaller{contract: contract}, nil
}

// NewResolverTransactor creates a new write-only instance of Resolver, bound to a specific deployed contract.
func NewResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*ResolverTransactor, error) {
	contract, err := bindResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ResolverTransactor{contract: contract}, nil
}

// NewResolverFilterer creates a new log filterer instance of Resolver, bound to a specific deployed contract.
func NewResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*ResolverFilterer, error) {
	contract, err := bindResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ResolverFilterer{contract: contract}, nil
}

// bindResolver binds a generic wrapper to an already deployed contract.
func bindResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ResolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Resolver *ResolverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Resolver.Contract.ResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Resolver *ResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Resolver.Contract.ResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Resolver *ResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Resolver.Contract.ResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Resolver *ResolverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Resolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Resolver *ResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Resolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Resolver *ResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Resolver.Contract.contract.Transact(opts, method, params...)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_Resolver *ResolverCaller) Addr(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Resolver.contract.Call(opts, &out, "addr", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_Resolver *ResolverSession) Addr(node [32]byte) (common.Address, error) {
	return _Resolver.Contract.Addr(&_Resolver.CallOpts, node)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_Resolver *ResolverCallerSession) Addr(node [32]byte) (common.Address, error) {
	return _Resolver.Contract.Addr(&_Resolver.CallOpts, node)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_Resolver *ResolverCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _Resolver.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_Resolver *ResolverSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _Resolver.Contract.SupportsInterface(&_Resolver.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_Resolver *ResolverCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _Resolver.Contract.SupportsInterface(&_Resolver.CallOpts, interfaceID)
}

// ResolverAddrChangedIterator is returned from FilterAddrChanged and is used to iterate over the raw logs and unpacked data for AddrChanged events raised by the Resolver contract.
type ResolverAddrChangedIterator struct {
	Event *ResolverAddrChanged // Event containing the contract specifics and raw log

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
func (it *ResolverAddrChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverAddrChanged)
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
		it.Event = new(ResolverAddrChanged)
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
func (it *ResolverAddrChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverAddrChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverAddrChanged represents a AddrChanged event raised by the Resolver contract.
type ResolverAddrChanged struct {
	Node [32]byte
	A    common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddrChanged is a free log retrieval operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(bytes32 indexed node, address a)
func (_Resolver *ResolverFilterer) FilterAddrChanged(opts *bind.FilterOpts, node [][32]byte) (*ResolverAddrChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Resolver.contract.FilterLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ResolverAddrChangedIterator{contract: _Resolver.contract, event: "AddrChanged", logs: logs, sub: sub}, nil
}

// WatchAddrChanged is a free log subscription operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(bytes32 indexed node, address a)
func (_Resolver *ResolverFilterer) WatchAddrChanged(opts *bind.WatchOpts, sink chan<- *ResolverAddrChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Resolver.contract.WatchLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverAddrChanged)
				if err := _Resolver.contract.UnpackLog(event, "AddrChanged", log); err != nil {
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

// ParseAddrChanged is a log parse operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(bytes32 indexed node, address a)
func (_Resolver *ResolverFilterer) ParseAddrChanged(log types.Log) (*ResolverAddrChanged, error) {
	event := new(ResolverAddrChanged)
	if err := _Resolver.contract.UnpackLog(event, "AddrChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
