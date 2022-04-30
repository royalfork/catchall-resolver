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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"indexedKey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"TextChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractENS\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"resolver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_node\",\"type\":\"bytes32\"},{\"internalType\":\"contractResolver\",\"name\":\"_resolver\",\"type\":\"address\"}],\"name\":\"setResolver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"text\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"7b103999": "registry()",
		"9061b923": "resolve(bytes,bytes)",
		"79e062f6": "resolver(bytes)",
		"1896f70a": "setResolver(bytes32,address)",
		"01ffc9a7": "supportsInterface(bytes4)",
		"59d1d43c": "text(bytes32,string)",
	},
	Bin: "0x60a060405234801561001057600080fd5b50604051610b39380380610b3983398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b608051610aa16100986000396000818160ea0152818161029901526103450152610aa16000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806301ffc9a7146100675780631896f70a1461008f57806359d1d43c146100a457806379e062f6146100c45780637b103999146100e55780639061b92314610124575b600080fd5b61007a610075366004610612565b61013b565b60405190151581526020015b60405180910390f35b6100a261009d36600461065b565b610280565b005b6100b76100b23660046106cd565b6103e8565b6040516100869190610775565b6100d76100d2366004610788565b610477565b6040516100869291906107ca565b61010c7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610086565b6100b761013236600461085d565b60609392505050565b60006001600160e01b03198216639061b92360e01b148061016c57506001600160e01b031982166301ffc9a760e01b145b8061018757506001600160e01b03198216631101d5ab60e11b145b806101a257506001600160e01b031982166378e5bf0360e11b145b806101bd57506001600160e01b03198216631d9dabef60e11b145b806101d857506001600160e01b0319821663bc1c58d160e01b145b806101f357506001600160e01b0319821663547d2b4160e11b145b8061020e57506001600160e01b03198216635c98042b60e01b145b8061022957506001600160e01b031982166304928c6760e21b145b8061024457506001600160e01b0319821663691f343160e01b145b8061025f57506001600160e01b0319821663c869023360e01b145b8061027a57506001600160e01b03198216631674750f60e21b145b92915050565b6040516302571be360e01b8152600481018390526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906302571be390602401602060405180830381865afa1580156102e8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061030c9190610909565b90506001600160a01b0381163314806103b0575060405163e985e9c560e01b81526001600160a01b0382811660048301523360248301527f0000000000000000000000000000000000000000000000000000000000000000169063e985e9c590604401602060405180830381865afa15801561038c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103b09190610926565b6103b957600080fd5b5060009182526020829052604090912080546001600160a01b0319166001600160a01b03909216919091179055565b60008381526020819052604090819020549051631674750f60e21b81526060916001600160a01b0316906359d1d43c9061042a90879087908790600401610948565b600060405180830381865afa158015610447573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261046f919081019061097e565b949350505050565b6000606060008061048a868660006104e7565b50919350915082905061049f8683818a6109f5565b81818080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250959950919750505050505050505b9250929050565b600080600080600087878781811061050157610501610a1f565b919091013560f81c9150506000819003610528575060009350859250839150819050610609565b60006105348288610a35565b61053f906001610a35565b9050600089896105508a6001610a35565b61055c928592906109f5565b60405161056a929190610a5b565b604051809103902090506000806000806105858e8e886104e7565b9350935093509350600081866040516020016105ab929190918252602082015260400190565b60408051601f1981840301815291815281516020928301206000818152928390529120549091506001600160a01b031680156105f8579b508c9a5098508897506106099650505050505050565b50939a509198509650909450505050505b93509350935093565b60006020828403121561062457600080fd5b81356001600160e01b03198116811461063c57600080fd5b9392505050565b6001600160a01b038116811461065857600080fd5b50565b6000806040838503121561066e57600080fd5b82359150602083013561068081610643565b809150509250929050565b60008083601f84011261069d57600080fd5b50813567ffffffffffffffff8111156106b557600080fd5b6020830191508360208285010111156104e057600080fd5b6000806000604084860312156106e257600080fd5b83359250602084013567ffffffffffffffff81111561070057600080fd5b61070c8682870161068b565b9497909650939450505050565b60005b8381101561073457818101518382015260200161071c565b83811115610743576000848401525b50505050565b60008151808452610761816020860160208601610719565b601f01601f19169290920160200192915050565b60208152600061063c6020830184610749565b6000806020838503121561079b57600080fd5b823567ffffffffffffffff8111156107b257600080fd5b6107be8582860161068b565b90969095509350505050565b6001600160a01b038316815260406020820181905260009061046f90830184610749565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561082d5761082d6107ee565b604052919050565b600067ffffffffffffffff82111561084f5761084f6107ee565b50601f01601f191660200190565b60008060006040848603121561087257600080fd5b833567ffffffffffffffff8082111561088a57600080fd5b6108968783880161068b565b909550935060208601359150808211156108af57600080fd5b508401601f810186136108c157600080fd5b80356108d46108cf82610835565b610804565b8181528760208385010111156108e957600080fd5b816020840160208301376000602083830101528093505050509250925092565b60006020828403121561091b57600080fd5b815161063c81610643565b60006020828403121561093857600080fd5b8151801515811461063c57600080fd5b83815260406020820152816040820152818360608301376000818301606090810191909152601f909201601f1916010192915050565b60006020828403121561099057600080fd5b815167ffffffffffffffff8111156109a757600080fd5b8201601f810184136109b857600080fd5b80516109c66108cf82610835565b8181528560208385010111156109db57600080fd5b6109ec826020830160208601610719565b95945050505050565b60008085851115610a0557600080fd5b83861115610a1257600080fd5b5050820193919092039150565b634e487b7160e01b600052603260045260246000fd5b60008219821115610a5657634e487b7160e01b600052601160045260246000fd5b500190565b818382376000910190815291905056fea2646970667358221220f13b479dfb141c8b8be473a452664f372695081d51e12982bb8d01407d1edafc64736f6c634300080d0033",
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
func DeployCatchallResolver(auth *bind.TransactOpts, backend bind.ContractBackend, _registry common.Address) (common.Address, *types.Transaction, *CatchallResolver, error) {
	parsed, err := CatchallResolverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CatchallResolverBin), backend, _registry)
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

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_CatchallResolver *CatchallResolverCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CatchallResolver.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_CatchallResolver *CatchallResolverSession) Registry() (common.Address, error) {
	return _CatchallResolver.Contract.Registry(&_CatchallResolver.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_CatchallResolver *CatchallResolverCallerSession) Registry() (common.Address, error) {
	return _CatchallResolver.Contract.Registry(&_CatchallResolver.CallOpts)
}

// Resolve is a free data retrieval call binding the contract method 0x9061b923.
//
// Solidity: function resolve(bytes name, bytes data) view returns(bytes)
func (_CatchallResolver *CatchallResolverCaller) Resolve(opts *bind.CallOpts, name []byte, data []byte) ([]byte, error) {
	var out []interface{}
	err := _CatchallResolver.contract.Call(opts, &out, "resolve", name, data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x9061b923.
//
// Solidity: function resolve(bytes name, bytes data) view returns(bytes)
func (_CatchallResolver *CatchallResolverSession) Resolve(name []byte, data []byte) ([]byte, error) {
	return _CatchallResolver.Contract.Resolve(&_CatchallResolver.CallOpts, name, data)
}

// Resolve is a free data retrieval call binding the contract method 0x9061b923.
//
// Solidity: function resolve(bytes name, bytes data) view returns(bytes)
func (_CatchallResolver *CatchallResolverCallerSession) Resolve(name []byte, data []byte) ([]byte, error) {
	return _CatchallResolver.Contract.Resolve(&_CatchallResolver.CallOpts, name, data)
}

// Resolver is a free data retrieval call binding the contract method 0x79e062f6.
//
// Solidity: function resolver(bytes name) view returns(address, bytes)
func (_CatchallResolver *CatchallResolverCaller) Resolver(opts *bind.CallOpts, name []byte) (common.Address, []byte, error) {
	var out []interface{}
	err := _CatchallResolver.contract.Call(opts, &out, "resolver", name)

	if err != nil {
		return *new(common.Address), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// Resolver is a free data retrieval call binding the contract method 0x79e062f6.
//
// Solidity: function resolver(bytes name) view returns(address, bytes)
func (_CatchallResolver *CatchallResolverSession) Resolver(name []byte) (common.Address, []byte, error) {
	return _CatchallResolver.Contract.Resolver(&_CatchallResolver.CallOpts, name)
}

// Resolver is a free data retrieval call binding the contract method 0x79e062f6.
//
// Solidity: function resolver(bytes name) view returns(address, bytes)
func (_CatchallResolver *CatchallResolverCallerSession) Resolver(name []byte) (common.Address, []byte, error) {
	return _CatchallResolver.Contract.Resolver(&_CatchallResolver.CallOpts, name)
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

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(bytes32 node, string key) view returns(string)
func (_CatchallResolver *CatchallResolverCaller) Text(opts *bind.CallOpts, node [32]byte, key string) (string, error) {
	var out []interface{}
	err := _CatchallResolver.contract.Call(opts, &out, "text", node, key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(bytes32 node, string key) view returns(string)
func (_CatchallResolver *CatchallResolverSession) Text(node [32]byte, key string) (string, error) {
	return _CatchallResolver.Contract.Text(&_CatchallResolver.CallOpts, node, key)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(bytes32 node, string key) view returns(string)
func (_CatchallResolver *CatchallResolverCallerSession) Text(node [32]byte, key string) (string, error) {
	return _CatchallResolver.Contract.Text(&_CatchallResolver.CallOpts, node, key)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 _node, address _resolver) returns()
func (_CatchallResolver *CatchallResolverTransactor) SetResolver(opts *bind.TransactOpts, _node [32]byte, _resolver common.Address) (*types.Transaction, error) {
	return _CatchallResolver.contract.Transact(opts, "setResolver", _node, _resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 _node, address _resolver) returns()
func (_CatchallResolver *CatchallResolverSession) SetResolver(_node [32]byte, _resolver common.Address) (*types.Transaction, error) {
	return _CatchallResolver.Contract.SetResolver(&_CatchallResolver.TransactOpts, _node, _resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 _node, address _resolver) returns()
func (_CatchallResolver *CatchallResolverTransactorSession) SetResolver(_node [32]byte, _resolver common.Address) (*types.Transaction, error) {
	return _CatchallResolver.Contract.SetResolver(&_CatchallResolver.TransactOpts, _node, _resolver)
}

// CatchallResolverTextChangedIterator is returned from FilterTextChanged and is used to iterate over the raw logs and unpacked data for TextChanged events raised by the CatchallResolver contract.
type CatchallResolverTextChangedIterator struct {
	Event *CatchallResolverTextChanged // Event containing the contract specifics and raw log

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
func (it *CatchallResolverTextChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CatchallResolverTextChanged)
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
		it.Event = new(CatchallResolverTextChanged)
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
func (it *CatchallResolverTextChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CatchallResolverTextChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CatchallResolverTextChanged represents a TextChanged event raised by the CatchallResolver contract.
type CatchallResolverTextChanged struct {
	Node       [32]byte
	IndexedKey common.Hash
	Key        string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTextChanged is a free log retrieval operation binding the contract event 0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550.
//
// Solidity: event TextChanged(bytes32 indexed node, string indexed indexedKey, string key)
func (_CatchallResolver *CatchallResolverFilterer) FilterTextChanged(opts *bind.FilterOpts, node [][32]byte, indexedKey []string) (*CatchallResolverTextChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _CatchallResolver.contract.FilterLogs(opts, "TextChanged", nodeRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return &CatchallResolverTextChangedIterator{contract: _CatchallResolver.contract, event: "TextChanged", logs: logs, sub: sub}, nil
}

// WatchTextChanged is a free log subscription operation binding the contract event 0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550.
//
// Solidity: event TextChanged(bytes32 indexed node, string indexed indexedKey, string key)
func (_CatchallResolver *CatchallResolverFilterer) WatchTextChanged(opts *bind.WatchOpts, sink chan<- *CatchallResolverTextChanged, node [][32]byte, indexedKey []string) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _CatchallResolver.contract.WatchLogs(opts, "TextChanged", nodeRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CatchallResolverTextChanged)
				if err := _CatchallResolver.contract.UnpackLog(event, "TextChanged", log); err != nil {
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

// ParseTextChanged is a log parse operation binding the contract event 0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550.
//
// Solidity: event TextChanged(bytes32 indexed node, string indexed indexedKey, string key)
func (_CatchallResolver *CatchallResolverFilterer) ParseTextChanged(log types.Log) (*CatchallResolverTextChanged, error) {
	event := new(CatchallResolverTextChanged)
	if err := _CatchallResolver.contract.UnpackLog(event, "TextChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ENSMetaData contains all meta data concerning the ENS contract.
var ENSMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"NewResolver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"NewTTL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"recordExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"resolver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"setRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"setResolver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setSubnodeOwner\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"setSubnodeRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"setTTL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"ttl\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e985e9c5": "isApprovedForAll(address,address)",
		"02571be3": "owner(bytes32)",
		"f79fe538": "recordExists(bytes32)",
		"0178b8bf": "resolver(bytes32)",
		"a22cb465": "setApprovalForAll(address,bool)",
		"5b0fc9c3": "setOwner(bytes32,address)",
		"cf408823": "setRecord(bytes32,address,address,uint64)",
		"1896f70a": "setResolver(bytes32,address)",
		"06ab5923": "setSubnodeOwner(bytes32,bytes32,address)",
		"5ef2c7f0": "setSubnodeRecord(bytes32,bytes32,address,address,uint64)",
		"14ab9038": "setTTL(bytes32,uint64)",
		"16a25cbd": "ttl(bytes32)",
	},
}

// ENSABI is the input ABI used to generate the binding from.
// Deprecated: Use ENSMetaData.ABI instead.
var ENSABI = ENSMetaData.ABI

// Deprecated: Use ENSMetaData.Sigs instead.
// ENSFuncSigs maps the 4-byte function signature to its string representation.
var ENSFuncSigs = ENSMetaData.Sigs

// ENS is an auto generated Go binding around an Ethereum contract.
type ENS struct {
	ENSCaller     // Read-only binding to the contract
	ENSTransactor // Write-only binding to the contract
	ENSFilterer   // Log filterer for contract events
}

// ENSCaller is an auto generated read-only Go binding around an Ethereum contract.
type ENSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ENSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ENSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ENSSession struct {
	Contract     *ENS              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ENSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ENSCallerSession struct {
	Contract *ENSCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ENSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ENSTransactorSession struct {
	Contract     *ENSTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ENSRaw is an auto generated low-level Go binding around an Ethereum contract.
type ENSRaw struct {
	Contract *ENS // Generic contract binding to access the raw methods on
}

// ENSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ENSCallerRaw struct {
	Contract *ENSCaller // Generic read-only contract binding to access the raw methods on
}

// ENSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ENSTransactorRaw struct {
	Contract *ENSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewENS creates a new instance of ENS, bound to a specific deployed contract.
func NewENS(address common.Address, backend bind.ContractBackend) (*ENS, error) {
	contract, err := bindENS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ENS{ENSCaller: ENSCaller{contract: contract}, ENSTransactor: ENSTransactor{contract: contract}, ENSFilterer: ENSFilterer{contract: contract}}, nil
}

// NewENSCaller creates a new read-only instance of ENS, bound to a specific deployed contract.
func NewENSCaller(address common.Address, caller bind.ContractCaller) (*ENSCaller, error) {
	contract, err := bindENS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ENSCaller{contract: contract}, nil
}

// NewENSTransactor creates a new write-only instance of ENS, bound to a specific deployed contract.
func NewENSTransactor(address common.Address, transactor bind.ContractTransactor) (*ENSTransactor, error) {
	contract, err := bindENS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ENSTransactor{contract: contract}, nil
}

// NewENSFilterer creates a new log filterer instance of ENS, bound to a specific deployed contract.
func NewENSFilterer(address common.Address, filterer bind.ContractFilterer) (*ENSFilterer, error) {
	contract, err := bindENS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ENSFilterer{contract: contract}, nil
}

// bindENS binds a generic wrapper to an already deployed contract.
func bindENS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ENSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ENS *ENSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ENS.Contract.ENSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ENS *ENSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ENS.Contract.ENSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ENS *ENSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ENS.Contract.ENSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ENS *ENSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ENS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ENS *ENSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ENS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ENS *ENSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ENS.Contract.contract.Transact(opts, method, params...)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ENS *ENSCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ENS.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ENS *ENSSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ENS.Contract.IsApprovedForAll(&_ENS.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ENS *ENSCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ENS.Contract.IsApprovedForAll(&_ENS.CallOpts, owner, operator)
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(bytes32 node) view returns(address)
func (_ENS *ENSCaller) Owner(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ENS.contract.Call(opts, &out, "owner", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(bytes32 node) view returns(address)
func (_ENS *ENSSession) Owner(node [32]byte) (common.Address, error) {
	return _ENS.Contract.Owner(&_ENS.CallOpts, node)
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(bytes32 node) view returns(address)
func (_ENS *ENSCallerSession) Owner(node [32]byte) (common.Address, error) {
	return _ENS.Contract.Owner(&_ENS.CallOpts, node)
}

// RecordExists is a free data retrieval call binding the contract method 0xf79fe538.
//
// Solidity: function recordExists(bytes32 node) view returns(bool)
func (_ENS *ENSCaller) RecordExists(opts *bind.CallOpts, node [32]byte) (bool, error) {
	var out []interface{}
	err := _ENS.contract.Call(opts, &out, "recordExists", node)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RecordExists is a free data retrieval call binding the contract method 0xf79fe538.
//
// Solidity: function recordExists(bytes32 node) view returns(bool)
func (_ENS *ENSSession) RecordExists(node [32]byte) (bool, error) {
	return _ENS.Contract.RecordExists(&_ENS.CallOpts, node)
}

// RecordExists is a free data retrieval call binding the contract method 0xf79fe538.
//
// Solidity: function recordExists(bytes32 node) view returns(bool)
func (_ENS *ENSCallerSession) RecordExists(node [32]byte) (bool, error) {
	return _ENS.Contract.RecordExists(&_ENS.CallOpts, node)
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(bytes32 node) view returns(address)
func (_ENS *ENSCaller) Resolver(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ENS.contract.Call(opts, &out, "resolver", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(bytes32 node) view returns(address)
func (_ENS *ENSSession) Resolver(node [32]byte) (common.Address, error) {
	return _ENS.Contract.Resolver(&_ENS.CallOpts, node)
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(bytes32 node) view returns(address)
func (_ENS *ENSCallerSession) Resolver(node [32]byte) (common.Address, error) {
	return _ENS.Contract.Resolver(&_ENS.CallOpts, node)
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(bytes32 node) view returns(uint64)
func (_ENS *ENSCaller) Ttl(opts *bind.CallOpts, node [32]byte) (uint64, error) {
	var out []interface{}
	err := _ENS.contract.Call(opts, &out, "ttl", node)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(bytes32 node) view returns(uint64)
func (_ENS *ENSSession) Ttl(node [32]byte) (uint64, error) {
	return _ENS.Contract.Ttl(&_ENS.CallOpts, node)
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(bytes32 node) view returns(uint64)
func (_ENS *ENSCallerSession) Ttl(node [32]byte) (uint64, error) {
	return _ENS.Contract.Ttl(&_ENS.CallOpts, node)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ENS *ENSTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ENS.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ENS *ENSSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ENS.Contract.SetApprovalForAll(&_ENS.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ENS *ENSTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ENS.Contract.SetApprovalForAll(&_ENS.TransactOpts, operator, approved)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(bytes32 node, address owner) returns()
func (_ENS *ENSTransactor) SetOwner(opts *bind.TransactOpts, node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENS.contract.Transact(opts, "setOwner", node, owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(bytes32 node, address owner) returns()
func (_ENS *ENSSession) SetOwner(node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENS.Contract.SetOwner(&_ENS.TransactOpts, node, owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(bytes32 node, address owner) returns()
func (_ENS *ENSTransactorSession) SetOwner(node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENS.Contract.SetOwner(&_ENS.TransactOpts, node, owner)
}

// SetRecord is a paid mutator transaction binding the contract method 0xcf408823.
//
// Solidity: function setRecord(bytes32 node, address owner, address resolver, uint64 ttl) returns()
func (_ENS *ENSTransactor) SetRecord(opts *bind.TransactOpts, node [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _ENS.contract.Transact(opts, "setRecord", node, owner, resolver, ttl)
}

// SetRecord is a paid mutator transaction binding the contract method 0xcf408823.
//
// Solidity: function setRecord(bytes32 node, address owner, address resolver, uint64 ttl) returns()
func (_ENS *ENSSession) SetRecord(node [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _ENS.Contract.SetRecord(&_ENS.TransactOpts, node, owner, resolver, ttl)
}

// SetRecord is a paid mutator transaction binding the contract method 0xcf408823.
//
// Solidity: function setRecord(bytes32 node, address owner, address resolver, uint64 ttl) returns()
func (_ENS *ENSTransactorSession) SetRecord(node [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _ENS.Contract.SetRecord(&_ENS.TransactOpts, node, owner, resolver, ttl)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 node, address resolver) returns()
func (_ENS *ENSTransactor) SetResolver(opts *bind.TransactOpts, node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _ENS.contract.Transact(opts, "setResolver", node, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 node, address resolver) returns()
func (_ENS *ENSSession) SetResolver(node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _ENS.Contract.SetResolver(&_ENS.TransactOpts, node, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 node, address resolver) returns()
func (_ENS *ENSTransactorSession) SetResolver(node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _ENS.Contract.SetResolver(&_ENS.TransactOpts, node, resolver)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(bytes32 node, bytes32 label, address owner) returns(bytes32)
func (_ENS *ENSTransactor) SetSubnodeOwner(opts *bind.TransactOpts, node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENS.contract.Transact(opts, "setSubnodeOwner", node, label, owner)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(bytes32 node, bytes32 label, address owner) returns(bytes32)
func (_ENS *ENSSession) SetSubnodeOwner(node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENS.Contract.SetSubnodeOwner(&_ENS.TransactOpts, node, label, owner)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(bytes32 node, bytes32 label, address owner) returns(bytes32)
func (_ENS *ENSTransactorSession) SetSubnodeOwner(node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENS.Contract.SetSubnodeOwner(&_ENS.TransactOpts, node, label, owner)
}

// SetSubnodeRecord is a paid mutator transaction binding the contract method 0x5ef2c7f0.
//
// Solidity: function setSubnodeRecord(bytes32 node, bytes32 label, address owner, address resolver, uint64 ttl) returns()
func (_ENS *ENSTransactor) SetSubnodeRecord(opts *bind.TransactOpts, node [32]byte, label [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _ENS.contract.Transact(opts, "setSubnodeRecord", node, label, owner, resolver, ttl)
}

// SetSubnodeRecord is a paid mutator transaction binding the contract method 0x5ef2c7f0.
//
// Solidity: function setSubnodeRecord(bytes32 node, bytes32 label, address owner, address resolver, uint64 ttl) returns()
func (_ENS *ENSSession) SetSubnodeRecord(node [32]byte, label [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _ENS.Contract.SetSubnodeRecord(&_ENS.TransactOpts, node, label, owner, resolver, ttl)
}

// SetSubnodeRecord is a paid mutator transaction binding the contract method 0x5ef2c7f0.
//
// Solidity: function setSubnodeRecord(bytes32 node, bytes32 label, address owner, address resolver, uint64 ttl) returns()
func (_ENS *ENSTransactorSession) SetSubnodeRecord(node [32]byte, label [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _ENS.Contract.SetSubnodeRecord(&_ENS.TransactOpts, node, label, owner, resolver, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 node, uint64 ttl) returns()
func (_ENS *ENSTransactor) SetTTL(opts *bind.TransactOpts, node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _ENS.contract.Transact(opts, "setTTL", node, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 node, uint64 ttl) returns()
func (_ENS *ENSSession) SetTTL(node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _ENS.Contract.SetTTL(&_ENS.TransactOpts, node, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 node, uint64 ttl) returns()
func (_ENS *ENSTransactorSession) SetTTL(node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _ENS.Contract.SetTTL(&_ENS.TransactOpts, node, ttl)
}

// ENSApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ENS contract.
type ENSApprovalForAllIterator struct {
	Event *ENSApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ENSApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ENSApprovalForAll)
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
		it.Event = new(ENSApprovalForAll)
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
func (it *ENSApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ENSApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ENSApprovalForAll represents a ApprovalForAll event raised by the ENS contract.
type ENSApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ENS *ENSFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ENSApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ENS.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ENSApprovalForAllIterator{contract: _ENS.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ENS *ENSFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ENSApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ENS.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ENSApprovalForAll)
				if err := _ENS.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ENS *ENSFilterer) ParseApprovalForAll(log types.Log) (*ENSApprovalForAll, error) {
	event := new(ENSApprovalForAll)
	if err := _ENS.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ENSNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the ENS contract.
type ENSNewOwnerIterator struct {
	Event *ENSNewOwner // Event containing the contract specifics and raw log

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
func (it *ENSNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ENSNewOwner)
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
		it.Event = new(ENSNewOwner)
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
func (it *ENSNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ENSNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ENSNewOwner represents a NewOwner event raised by the ENS contract.
type ENSNewOwner struct {
	Node  [32]byte
	Label [32]byte
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0xce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82.
//
// Solidity: event NewOwner(bytes32 indexed node, bytes32 indexed label, address owner)
func (_ENS *ENSFilterer) FilterNewOwner(opts *bind.FilterOpts, node [][32]byte, label [][32]byte) (*ENSNewOwnerIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _ENS.contract.FilterLogs(opts, "NewOwner", nodeRule, labelRule)
	if err != nil {
		return nil, err
	}
	return &ENSNewOwnerIterator{contract: _ENS.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0xce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82.
//
// Solidity: event NewOwner(bytes32 indexed node, bytes32 indexed label, address owner)
func (_ENS *ENSFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *ENSNewOwner, node [][32]byte, label [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _ENS.contract.WatchLogs(opts, "NewOwner", nodeRule, labelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ENSNewOwner)
				if err := _ENS.contract.UnpackLog(event, "NewOwner", log); err != nil {
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

// ParseNewOwner is a log parse operation binding the contract event 0xce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82.
//
// Solidity: event NewOwner(bytes32 indexed node, bytes32 indexed label, address owner)
func (_ENS *ENSFilterer) ParseNewOwner(log types.Log) (*ENSNewOwner, error) {
	event := new(ENSNewOwner)
	if err := _ENS.contract.UnpackLog(event, "NewOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ENSNewResolverIterator is returned from FilterNewResolver and is used to iterate over the raw logs and unpacked data for NewResolver events raised by the ENS contract.
type ENSNewResolverIterator struct {
	Event *ENSNewResolver // Event containing the contract specifics and raw log

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
func (it *ENSNewResolverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ENSNewResolver)
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
		it.Event = new(ENSNewResolver)
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
func (it *ENSNewResolverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ENSNewResolverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ENSNewResolver represents a NewResolver event raised by the ENS contract.
type ENSNewResolver struct {
	Node     [32]byte
	Resolver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewResolver is a free log retrieval operation binding the contract event 0x335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0.
//
// Solidity: event NewResolver(bytes32 indexed node, address resolver)
func (_ENS *ENSFilterer) FilterNewResolver(opts *bind.FilterOpts, node [][32]byte) (*ENSNewResolverIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENS.contract.FilterLogs(opts, "NewResolver", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ENSNewResolverIterator{contract: _ENS.contract, event: "NewResolver", logs: logs, sub: sub}, nil
}

// WatchNewResolver is a free log subscription operation binding the contract event 0x335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0.
//
// Solidity: event NewResolver(bytes32 indexed node, address resolver)
func (_ENS *ENSFilterer) WatchNewResolver(opts *bind.WatchOpts, sink chan<- *ENSNewResolver, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENS.contract.WatchLogs(opts, "NewResolver", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ENSNewResolver)
				if err := _ENS.contract.UnpackLog(event, "NewResolver", log); err != nil {
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

// ParseNewResolver is a log parse operation binding the contract event 0x335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0.
//
// Solidity: event NewResolver(bytes32 indexed node, address resolver)
func (_ENS *ENSFilterer) ParseNewResolver(log types.Log) (*ENSNewResolver, error) {
	event := new(ENSNewResolver)
	if err := _ENS.contract.UnpackLog(event, "NewResolver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ENSNewTTLIterator is returned from FilterNewTTL and is used to iterate over the raw logs and unpacked data for NewTTL events raised by the ENS contract.
type ENSNewTTLIterator struct {
	Event *ENSNewTTL // Event containing the contract specifics and raw log

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
func (it *ENSNewTTLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ENSNewTTL)
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
		it.Event = new(ENSNewTTL)
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
func (it *ENSNewTTLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ENSNewTTLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ENSNewTTL represents a NewTTL event raised by the ENS contract.
type ENSNewTTL struct {
	Node [32]byte
	Ttl  uint64
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNewTTL is a free log retrieval operation binding the contract event 0x1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68.
//
// Solidity: event NewTTL(bytes32 indexed node, uint64 ttl)
func (_ENS *ENSFilterer) FilterNewTTL(opts *bind.FilterOpts, node [][32]byte) (*ENSNewTTLIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENS.contract.FilterLogs(opts, "NewTTL", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ENSNewTTLIterator{contract: _ENS.contract, event: "NewTTL", logs: logs, sub: sub}, nil
}

// WatchNewTTL is a free log subscription operation binding the contract event 0x1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68.
//
// Solidity: event NewTTL(bytes32 indexed node, uint64 ttl)
func (_ENS *ENSFilterer) WatchNewTTL(opts *bind.WatchOpts, sink chan<- *ENSNewTTL, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENS.contract.WatchLogs(opts, "NewTTL", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ENSNewTTL)
				if err := _ENS.contract.UnpackLog(event, "NewTTL", log); err != nil {
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

// ParseNewTTL is a log parse operation binding the contract event 0x1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68.
//
// Solidity: event NewTTL(bytes32 indexed node, uint64 ttl)
func (_ENS *ENSFilterer) ParseNewTTL(log types.Log) (*ENSNewTTL, error) {
	event := new(ENSNewTTL)
	if err := _ENS.contract.UnpackLog(event, "NewTTL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ENSTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ENS contract.
type ENSTransferIterator struct {
	Event *ENSTransfer // Event containing the contract specifics and raw log

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
func (it *ENSTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ENSTransfer)
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
		it.Event = new(ENSTransfer)
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
func (it *ENSTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ENSTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ENSTransfer represents a Transfer event raised by the ENS contract.
type ENSTransfer struct {
	Node  [32]byte
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266.
//
// Solidity: event Transfer(bytes32 indexed node, address owner)
func (_ENS *ENSFilterer) FilterTransfer(opts *bind.FilterOpts, node [][32]byte) (*ENSTransferIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENS.contract.FilterLogs(opts, "Transfer", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ENSTransferIterator{contract: _ENS.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266.
//
// Solidity: event Transfer(bytes32 indexed node, address owner)
func (_ENS *ENSFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ENSTransfer, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENS.contract.WatchLogs(opts, "Transfer", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ENSTransfer)
				if err := _ENS.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266.
//
// Solidity: event Transfer(bytes32 indexed node, address owner)
func (_ENS *ENSFilterer) ParseTransfer(log types.Log) (*ENSTransfer, error) {
	event := new(ENSTransfer)
	if err := _ENS.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IABIResolverMetaData contains all meta data concerning the IABIResolver contract.
var IABIResolverMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"contentType\",\"type\":\"uint256\"}],\"name\":\"ABIChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"contentTypes\",\"type\":\"uint256\"}],\"name\":\"ABI\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"2203ab56": "ABI(bytes32,uint256)",
	},
}

// IABIResolverABI is the input ABI used to generate the binding from.
// Deprecated: Use IABIResolverMetaData.ABI instead.
var IABIResolverABI = IABIResolverMetaData.ABI

// Deprecated: Use IABIResolverMetaData.Sigs instead.
// IABIResolverFuncSigs maps the 4-byte function signature to its string representation.
var IABIResolverFuncSigs = IABIResolverMetaData.Sigs

// IABIResolver is an auto generated Go binding around an Ethereum contract.
type IABIResolver struct {
	IABIResolverCaller     // Read-only binding to the contract
	IABIResolverTransactor // Write-only binding to the contract
	IABIResolverFilterer   // Log filterer for contract events
}

// IABIResolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IABIResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IABIResolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IABIResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IABIResolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IABIResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IABIResolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IABIResolverSession struct {
	Contract     *IABIResolver     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IABIResolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IABIResolverCallerSession struct {
	Contract *IABIResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IABIResolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IABIResolverTransactorSession struct {
	Contract     *IABIResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IABIResolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IABIResolverRaw struct {
	Contract *IABIResolver // Generic contract binding to access the raw methods on
}

// IABIResolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IABIResolverCallerRaw struct {
	Contract *IABIResolverCaller // Generic read-only contract binding to access the raw methods on
}

// IABIResolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IABIResolverTransactorRaw struct {
	Contract *IABIResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIABIResolver creates a new instance of IABIResolver, bound to a specific deployed contract.
func NewIABIResolver(address common.Address, backend bind.ContractBackend) (*IABIResolver, error) {
	contract, err := bindIABIResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IABIResolver{IABIResolverCaller: IABIResolverCaller{contract: contract}, IABIResolverTransactor: IABIResolverTransactor{contract: contract}, IABIResolverFilterer: IABIResolverFilterer{contract: contract}}, nil
}

// NewIABIResolverCaller creates a new read-only instance of IABIResolver, bound to a specific deployed contract.
func NewIABIResolverCaller(address common.Address, caller bind.ContractCaller) (*IABIResolverCaller, error) {
	contract, err := bindIABIResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IABIResolverCaller{contract: contract}, nil
}

// NewIABIResolverTransactor creates a new write-only instance of IABIResolver, bound to a specific deployed contract.
func NewIABIResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*IABIResolverTransactor, error) {
	contract, err := bindIABIResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IABIResolverTransactor{contract: contract}, nil
}

// NewIABIResolverFilterer creates a new log filterer instance of IABIResolver, bound to a specific deployed contract.
func NewIABIResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*IABIResolverFilterer, error) {
	contract, err := bindIABIResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IABIResolverFilterer{contract: contract}, nil
}

// bindIABIResolver binds a generic wrapper to an already deployed contract.
func bindIABIResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IABIResolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IABIResolver *IABIResolverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IABIResolver.Contract.IABIResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IABIResolver *IABIResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IABIResolver.Contract.IABIResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IABIResolver *IABIResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IABIResolver.Contract.IABIResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IABIResolver *IABIResolverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IABIResolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IABIResolver *IABIResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IABIResolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IABIResolver *IABIResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IABIResolver.Contract.contract.Transact(opts, method, params...)
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(bytes32 node, uint256 contentTypes) view returns(uint256, bytes)
func (_IABIResolver *IABIResolverCaller) ABI(opts *bind.CallOpts, node [32]byte, contentTypes *big.Int) (*big.Int, []byte, error) {
	var out []interface{}
	err := _IABIResolver.contract.Call(opts, &out, "ABI", node, contentTypes)

	if err != nil {
		return *new(*big.Int), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(bytes32 node, uint256 contentTypes) view returns(uint256, bytes)
func (_IABIResolver *IABIResolverSession) ABI(node [32]byte, contentTypes *big.Int) (*big.Int, []byte, error) {
	return _IABIResolver.Contract.ABI(&_IABIResolver.CallOpts, node, contentTypes)
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(bytes32 node, uint256 contentTypes) view returns(uint256, bytes)
func (_IABIResolver *IABIResolverCallerSession) ABI(node [32]byte, contentTypes *big.Int) (*big.Int, []byte, error) {
	return _IABIResolver.Contract.ABI(&_IABIResolver.CallOpts, node, contentTypes)
}

// IABIResolverABIChangedIterator is returned from FilterABIChanged and is used to iterate over the raw logs and unpacked data for ABIChanged events raised by the IABIResolver contract.
type IABIResolverABIChangedIterator struct {
	Event *IABIResolverABIChanged // Event containing the contract specifics and raw log

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
func (it *IABIResolverABIChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IABIResolverABIChanged)
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
		it.Event = new(IABIResolverABIChanged)
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
func (it *IABIResolverABIChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IABIResolverABIChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IABIResolverABIChanged represents a ABIChanged event raised by the IABIResolver contract.
type IABIResolverABIChanged struct {
	Node        [32]byte
	ContentType *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterABIChanged is a free log retrieval operation binding the contract event 0xaa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe3.
//
// Solidity: event ABIChanged(bytes32 indexed node, uint256 indexed contentType)
func (_IABIResolver *IABIResolverFilterer) FilterABIChanged(opts *bind.FilterOpts, node [][32]byte, contentType []*big.Int) (*IABIResolverABIChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var contentTypeRule []interface{}
	for _, contentTypeItem := range contentType {
		contentTypeRule = append(contentTypeRule, contentTypeItem)
	}

	logs, sub, err := _IABIResolver.contract.FilterLogs(opts, "ABIChanged", nodeRule, contentTypeRule)
	if err != nil {
		return nil, err
	}
	return &IABIResolverABIChangedIterator{contract: _IABIResolver.contract, event: "ABIChanged", logs: logs, sub: sub}, nil
}

// WatchABIChanged is a free log subscription operation binding the contract event 0xaa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe3.
//
// Solidity: event ABIChanged(bytes32 indexed node, uint256 indexed contentType)
func (_IABIResolver *IABIResolverFilterer) WatchABIChanged(opts *bind.WatchOpts, sink chan<- *IABIResolverABIChanged, node [][32]byte, contentType []*big.Int) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var contentTypeRule []interface{}
	for _, contentTypeItem := range contentType {
		contentTypeRule = append(contentTypeRule, contentTypeItem)
	}

	logs, sub, err := _IABIResolver.contract.WatchLogs(opts, "ABIChanged", nodeRule, contentTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IABIResolverABIChanged)
				if err := _IABIResolver.contract.UnpackLog(event, "ABIChanged", log); err != nil {
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

// ParseABIChanged is a log parse operation binding the contract event 0xaa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe3.
//
// Solidity: event ABIChanged(bytes32 indexed node, uint256 indexed contentType)
func (_IABIResolver *IABIResolverFilterer) ParseABIChanged(log types.Log) (*IABIResolverABIChanged, error) {
	event := new(IABIResolverABIChanged)
	if err := _IABIResolver.contract.UnpackLog(event, "ABIChanged", log); err != nil {
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

// IAddressResolverMetaData contains all meta data concerning the IAddressResolver contract.
var IAddressResolverMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newAddress\",\"type\":\"bytes\"}],\"name\":\"AddressChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"}],\"name\":\"addr\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f1cb7e06": "addr(bytes32,uint256)",
	},
}

// IAddressResolverABI is the input ABI used to generate the binding from.
// Deprecated: Use IAddressResolverMetaData.ABI instead.
var IAddressResolverABI = IAddressResolverMetaData.ABI

// Deprecated: Use IAddressResolverMetaData.Sigs instead.
// IAddressResolverFuncSigs maps the 4-byte function signature to its string representation.
var IAddressResolverFuncSigs = IAddressResolverMetaData.Sigs

// IAddressResolver is an auto generated Go binding around an Ethereum contract.
type IAddressResolver struct {
	IAddressResolverCaller     // Read-only binding to the contract
	IAddressResolverTransactor // Write-only binding to the contract
	IAddressResolverFilterer   // Log filterer for contract events
}

// IAddressResolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAddressResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAddressResolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAddressResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAddressResolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAddressResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAddressResolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAddressResolverSession struct {
	Contract     *IAddressResolver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAddressResolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAddressResolverCallerSession struct {
	Contract *IAddressResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IAddressResolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAddressResolverTransactorSession struct {
	Contract     *IAddressResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IAddressResolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAddressResolverRaw struct {
	Contract *IAddressResolver // Generic contract binding to access the raw methods on
}

// IAddressResolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAddressResolverCallerRaw struct {
	Contract *IAddressResolverCaller // Generic read-only contract binding to access the raw methods on
}

// IAddressResolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAddressResolverTransactorRaw struct {
	Contract *IAddressResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAddressResolver creates a new instance of IAddressResolver, bound to a specific deployed contract.
func NewIAddressResolver(address common.Address, backend bind.ContractBackend) (*IAddressResolver, error) {
	contract, err := bindIAddressResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAddressResolver{IAddressResolverCaller: IAddressResolverCaller{contract: contract}, IAddressResolverTransactor: IAddressResolverTransactor{contract: contract}, IAddressResolverFilterer: IAddressResolverFilterer{contract: contract}}, nil
}

// NewIAddressResolverCaller creates a new read-only instance of IAddressResolver, bound to a specific deployed contract.
func NewIAddressResolverCaller(address common.Address, caller bind.ContractCaller) (*IAddressResolverCaller, error) {
	contract, err := bindIAddressResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAddressResolverCaller{contract: contract}, nil
}

// NewIAddressResolverTransactor creates a new write-only instance of IAddressResolver, bound to a specific deployed contract.
func NewIAddressResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*IAddressResolverTransactor, error) {
	contract, err := bindIAddressResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAddressResolverTransactor{contract: contract}, nil
}

// NewIAddressResolverFilterer creates a new log filterer instance of IAddressResolver, bound to a specific deployed contract.
func NewIAddressResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*IAddressResolverFilterer, error) {
	contract, err := bindIAddressResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAddressResolverFilterer{contract: contract}, nil
}

// bindIAddressResolver binds a generic wrapper to an already deployed contract.
func bindIAddressResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAddressResolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAddressResolver *IAddressResolverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAddressResolver.Contract.IAddressResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAddressResolver *IAddressResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAddressResolver.Contract.IAddressResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAddressResolver *IAddressResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAddressResolver.Contract.IAddressResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAddressResolver *IAddressResolverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAddressResolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAddressResolver *IAddressResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAddressResolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAddressResolver *IAddressResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAddressResolver.Contract.contract.Transact(opts, method, params...)
}

// Addr is a free data retrieval call binding the contract method 0xf1cb7e06.
//
// Solidity: function addr(bytes32 node, uint256 coinType) view returns(bytes)
func (_IAddressResolver *IAddressResolverCaller) Addr(opts *bind.CallOpts, node [32]byte, coinType *big.Int) ([]byte, error) {
	var out []interface{}
	err := _IAddressResolver.contract.Call(opts, &out, "addr", node, coinType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Addr is a free data retrieval call binding the contract method 0xf1cb7e06.
//
// Solidity: function addr(bytes32 node, uint256 coinType) view returns(bytes)
func (_IAddressResolver *IAddressResolverSession) Addr(node [32]byte, coinType *big.Int) ([]byte, error) {
	return _IAddressResolver.Contract.Addr(&_IAddressResolver.CallOpts, node, coinType)
}

// Addr is a free data retrieval call binding the contract method 0xf1cb7e06.
//
// Solidity: function addr(bytes32 node, uint256 coinType) view returns(bytes)
func (_IAddressResolver *IAddressResolverCallerSession) Addr(node [32]byte, coinType *big.Int) ([]byte, error) {
	return _IAddressResolver.Contract.Addr(&_IAddressResolver.CallOpts, node, coinType)
}

// IAddressResolverAddressChangedIterator is returned from FilterAddressChanged and is used to iterate over the raw logs and unpacked data for AddressChanged events raised by the IAddressResolver contract.
type IAddressResolverAddressChangedIterator struct {
	Event *IAddressResolverAddressChanged // Event containing the contract specifics and raw log

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
func (it *IAddressResolverAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAddressResolverAddressChanged)
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
		it.Event = new(IAddressResolverAddressChanged)
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
func (it *IAddressResolverAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAddressResolverAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAddressResolverAddressChanged represents a AddressChanged event raised by the IAddressResolver contract.
type IAddressResolverAddressChanged struct {
	Node       [32]byte
	CoinType   *big.Int
	NewAddress []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddressChanged is a free log retrieval operation binding the contract event 0x65412581168e88a1e60c6459d7f44ae83ad0832e670826c05a4e2476b57af752.
//
// Solidity: event AddressChanged(bytes32 indexed node, uint256 coinType, bytes newAddress)
func (_IAddressResolver *IAddressResolverFilterer) FilterAddressChanged(opts *bind.FilterOpts, node [][32]byte) (*IAddressResolverAddressChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _IAddressResolver.contract.FilterLogs(opts, "AddressChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &IAddressResolverAddressChangedIterator{contract: _IAddressResolver.contract, event: "AddressChanged", logs: logs, sub: sub}, nil
}

// WatchAddressChanged is a free log subscription operation binding the contract event 0x65412581168e88a1e60c6459d7f44ae83ad0832e670826c05a4e2476b57af752.
//
// Solidity: event AddressChanged(bytes32 indexed node, uint256 coinType, bytes newAddress)
func (_IAddressResolver *IAddressResolverFilterer) WatchAddressChanged(opts *bind.WatchOpts, sink chan<- *IAddressResolverAddressChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _IAddressResolver.contract.WatchLogs(opts, "AddressChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAddressResolverAddressChanged)
				if err := _IAddressResolver.contract.UnpackLog(event, "AddressChanged", log); err != nil {
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

// ParseAddressChanged is a log parse operation binding the contract event 0x65412581168e88a1e60c6459d7f44ae83ad0832e670826c05a4e2476b57af752.
//
// Solidity: event AddressChanged(bytes32 indexed node, uint256 coinType, bytes newAddress)
func (_IAddressResolver *IAddressResolverFilterer) ParseAddressChanged(log types.Log) (*IAddressResolverAddressChanged, error) {
	event := new(IAddressResolverAddressChanged)
	if err := _IAddressResolver.contract.UnpackLog(event, "AddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IContentHashResolverMetaData contains all meta data concerning the IContentHashResolver contract.
var IContentHashResolverMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"hash\",\"type\":\"bytes\"}],\"name\":\"ContenthashChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"contenthash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"bc1c58d1": "contenthash(bytes32)",
	},
}

// IContentHashResolverABI is the input ABI used to generate the binding from.
// Deprecated: Use IContentHashResolverMetaData.ABI instead.
var IContentHashResolverABI = IContentHashResolverMetaData.ABI

// Deprecated: Use IContentHashResolverMetaData.Sigs instead.
// IContentHashResolverFuncSigs maps the 4-byte function signature to its string representation.
var IContentHashResolverFuncSigs = IContentHashResolverMetaData.Sigs

// IContentHashResolver is an auto generated Go binding around an Ethereum contract.
type IContentHashResolver struct {
	IContentHashResolverCaller     // Read-only binding to the contract
	IContentHashResolverTransactor // Write-only binding to the contract
	IContentHashResolverFilterer   // Log filterer for contract events
}

// IContentHashResolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IContentHashResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IContentHashResolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IContentHashResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IContentHashResolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IContentHashResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IContentHashResolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IContentHashResolverSession struct {
	Contract     *IContentHashResolver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IContentHashResolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IContentHashResolverCallerSession struct {
	Contract *IContentHashResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// IContentHashResolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IContentHashResolverTransactorSession struct {
	Contract     *IContentHashResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IContentHashResolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IContentHashResolverRaw struct {
	Contract *IContentHashResolver // Generic contract binding to access the raw methods on
}

// IContentHashResolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IContentHashResolverCallerRaw struct {
	Contract *IContentHashResolverCaller // Generic read-only contract binding to access the raw methods on
}

// IContentHashResolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IContentHashResolverTransactorRaw struct {
	Contract *IContentHashResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIContentHashResolver creates a new instance of IContentHashResolver, bound to a specific deployed contract.
func NewIContentHashResolver(address common.Address, backend bind.ContractBackend) (*IContentHashResolver, error) {
	contract, err := bindIContentHashResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IContentHashResolver{IContentHashResolverCaller: IContentHashResolverCaller{contract: contract}, IContentHashResolverTransactor: IContentHashResolverTransactor{contract: contract}, IContentHashResolverFilterer: IContentHashResolverFilterer{contract: contract}}, nil
}

// NewIContentHashResolverCaller creates a new read-only instance of IContentHashResolver, bound to a specific deployed contract.
func NewIContentHashResolverCaller(address common.Address, caller bind.ContractCaller) (*IContentHashResolverCaller, error) {
	contract, err := bindIContentHashResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IContentHashResolverCaller{contract: contract}, nil
}

// NewIContentHashResolverTransactor creates a new write-only instance of IContentHashResolver, bound to a specific deployed contract.
func NewIContentHashResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*IContentHashResolverTransactor, error) {
	contract, err := bindIContentHashResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IContentHashResolverTransactor{contract: contract}, nil
}

// NewIContentHashResolverFilterer creates a new log filterer instance of IContentHashResolver, bound to a specific deployed contract.
func NewIContentHashResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*IContentHashResolverFilterer, error) {
	contract, err := bindIContentHashResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IContentHashResolverFilterer{contract: contract}, nil
}

// bindIContentHashResolver binds a generic wrapper to an already deployed contract.
func bindIContentHashResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IContentHashResolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IContentHashResolver *IContentHashResolverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IContentHashResolver.Contract.IContentHashResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IContentHashResolver *IContentHashResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IContentHashResolver.Contract.IContentHashResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IContentHashResolver *IContentHashResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IContentHashResolver.Contract.IContentHashResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IContentHashResolver *IContentHashResolverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IContentHashResolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IContentHashResolver *IContentHashResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IContentHashResolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IContentHashResolver *IContentHashResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IContentHashResolver.Contract.contract.Transact(opts, method, params...)
}

// Contenthash is a free data retrieval call binding the contract method 0xbc1c58d1.
//
// Solidity: function contenthash(bytes32 node) view returns(bytes)
func (_IContentHashResolver *IContentHashResolverCaller) Contenthash(opts *bind.CallOpts, node [32]byte) ([]byte, error) {
	var out []interface{}
	err := _IContentHashResolver.contract.Call(opts, &out, "contenthash", node)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Contenthash is a free data retrieval call binding the contract method 0xbc1c58d1.
//
// Solidity: function contenthash(bytes32 node) view returns(bytes)
func (_IContentHashResolver *IContentHashResolverSession) Contenthash(node [32]byte) ([]byte, error) {
	return _IContentHashResolver.Contract.Contenthash(&_IContentHashResolver.CallOpts, node)
}

// Contenthash is a free data retrieval call binding the contract method 0xbc1c58d1.
//
// Solidity: function contenthash(bytes32 node) view returns(bytes)
func (_IContentHashResolver *IContentHashResolverCallerSession) Contenthash(node [32]byte) ([]byte, error) {
	return _IContentHashResolver.Contract.Contenthash(&_IContentHashResolver.CallOpts, node)
}

// IContentHashResolverContenthashChangedIterator is returned from FilterContenthashChanged and is used to iterate over the raw logs and unpacked data for ContenthashChanged events raised by the IContentHashResolver contract.
type IContentHashResolverContenthashChangedIterator struct {
	Event *IContentHashResolverContenthashChanged // Event containing the contract specifics and raw log

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
func (it *IContentHashResolverContenthashChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IContentHashResolverContenthashChanged)
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
		it.Event = new(IContentHashResolverContenthashChanged)
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
func (it *IContentHashResolverContenthashChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IContentHashResolverContenthashChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IContentHashResolverContenthashChanged represents a ContenthashChanged event raised by the IContentHashResolver contract.
type IContentHashResolverContenthashChanged struct {
	Node [32]byte
	Hash []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterContenthashChanged is a free log retrieval operation binding the contract event 0xe379c1624ed7e714cc0937528a32359d69d5281337765313dba4e081b72d7578.
//
// Solidity: event ContenthashChanged(bytes32 indexed node, bytes hash)
func (_IContentHashResolver *IContentHashResolverFilterer) FilterContenthashChanged(opts *bind.FilterOpts, node [][32]byte) (*IContentHashResolverContenthashChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _IContentHashResolver.contract.FilterLogs(opts, "ContenthashChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &IContentHashResolverContenthashChangedIterator{contract: _IContentHashResolver.contract, event: "ContenthashChanged", logs: logs, sub: sub}, nil
}

// WatchContenthashChanged is a free log subscription operation binding the contract event 0xe379c1624ed7e714cc0937528a32359d69d5281337765313dba4e081b72d7578.
//
// Solidity: event ContenthashChanged(bytes32 indexed node, bytes hash)
func (_IContentHashResolver *IContentHashResolverFilterer) WatchContenthashChanged(opts *bind.WatchOpts, sink chan<- *IContentHashResolverContenthashChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _IContentHashResolver.contract.WatchLogs(opts, "ContenthashChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IContentHashResolverContenthashChanged)
				if err := _IContentHashResolver.contract.UnpackLog(event, "ContenthashChanged", log); err != nil {
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

// ParseContenthashChanged is a log parse operation binding the contract event 0xe379c1624ed7e714cc0937528a32359d69d5281337765313dba4e081b72d7578.
//
// Solidity: event ContenthashChanged(bytes32 indexed node, bytes hash)
func (_IContentHashResolver *IContentHashResolverFilterer) ParseContenthashChanged(log types.Log) (*IContentHashResolverContenthashChanged, error) {
	event := new(IContentHashResolverContenthashChanged)
	if err := _IContentHashResolver.contract.UnpackLog(event, "ContenthashChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDNSRecordResolverMetaData contains all meta data concerning the IDNSRecordResolver contract.
var IDNSRecordResolverMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"resource\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"record\",\"type\":\"bytes\"}],\"name\":\"DNSRecordChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"resource\",\"type\":\"uint16\"}],\"name\":\"DNSRecordDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"DNSZoneCleared\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"resource\",\"type\":\"uint16\"}],\"name\":\"dnsRecord\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a8fa5682": "dnsRecord(bytes32,bytes32,uint16)",
	},
}

// IDNSRecordResolverABI is the input ABI used to generate the binding from.
// Deprecated: Use IDNSRecordResolverMetaData.ABI instead.
var IDNSRecordResolverABI = IDNSRecordResolverMetaData.ABI

// Deprecated: Use IDNSRecordResolverMetaData.Sigs instead.
// IDNSRecordResolverFuncSigs maps the 4-byte function signature to its string representation.
var IDNSRecordResolverFuncSigs = IDNSRecordResolverMetaData.Sigs

// IDNSRecordResolver is an auto generated Go binding around an Ethereum contract.
type IDNSRecordResolver struct {
	IDNSRecordResolverCaller     // Read-only binding to the contract
	IDNSRecordResolverTransactor // Write-only binding to the contract
	IDNSRecordResolverFilterer   // Log filterer for contract events
}

// IDNSRecordResolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDNSRecordResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDNSRecordResolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDNSRecordResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDNSRecordResolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDNSRecordResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDNSRecordResolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDNSRecordResolverSession struct {
	Contract     *IDNSRecordResolver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IDNSRecordResolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDNSRecordResolverCallerSession struct {
	Contract *IDNSRecordResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IDNSRecordResolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDNSRecordResolverTransactorSession struct {
	Contract     *IDNSRecordResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IDNSRecordResolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDNSRecordResolverRaw struct {
	Contract *IDNSRecordResolver // Generic contract binding to access the raw methods on
}

// IDNSRecordResolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDNSRecordResolverCallerRaw struct {
	Contract *IDNSRecordResolverCaller // Generic read-only contract binding to access the raw methods on
}

// IDNSRecordResolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDNSRecordResolverTransactorRaw struct {
	Contract *IDNSRecordResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDNSRecordResolver creates a new instance of IDNSRecordResolver, bound to a specific deployed contract.
func NewIDNSRecordResolver(address common.Address, backend bind.ContractBackend) (*IDNSRecordResolver, error) {
	contract, err := bindIDNSRecordResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDNSRecordResolver{IDNSRecordResolverCaller: IDNSRecordResolverCaller{contract: contract}, IDNSRecordResolverTransactor: IDNSRecordResolverTransactor{contract: contract}, IDNSRecordResolverFilterer: IDNSRecordResolverFilterer{contract: contract}}, nil
}

// NewIDNSRecordResolverCaller creates a new read-only instance of IDNSRecordResolver, bound to a specific deployed contract.
func NewIDNSRecordResolverCaller(address common.Address, caller bind.ContractCaller) (*IDNSRecordResolverCaller, error) {
	contract, err := bindIDNSRecordResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDNSRecordResolverCaller{contract: contract}, nil
}

// NewIDNSRecordResolverTransactor creates a new write-only instance of IDNSRecordResolver, bound to a specific deployed contract.
func NewIDNSRecordResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*IDNSRecordResolverTransactor, error) {
	contract, err := bindIDNSRecordResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDNSRecordResolverTransactor{contract: contract}, nil
}

// NewIDNSRecordResolverFilterer creates a new log filterer instance of IDNSRecordResolver, bound to a specific deployed contract.
func NewIDNSRecordResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*IDNSRecordResolverFilterer, error) {
	contract, err := bindIDNSRecordResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDNSRecordResolverFilterer{contract: contract}, nil
}

// bindIDNSRecordResolver binds a generic wrapper to an already deployed contract.
func bindIDNSRecordResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDNSRecordResolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDNSRecordResolver *IDNSRecordResolverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDNSRecordResolver.Contract.IDNSRecordResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDNSRecordResolver *IDNSRecordResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDNSRecordResolver.Contract.IDNSRecordResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDNSRecordResolver *IDNSRecordResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDNSRecordResolver.Contract.IDNSRecordResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDNSRecordResolver *IDNSRecordResolverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDNSRecordResolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDNSRecordResolver *IDNSRecordResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDNSRecordResolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDNSRecordResolver *IDNSRecordResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDNSRecordResolver.Contract.contract.Transact(opts, method, params...)
}

// DnsRecord is a free data retrieval call binding the contract method 0xa8fa5682.
//
// Solidity: function dnsRecord(bytes32 node, bytes32 name, uint16 resource) view returns(bytes)
func (_IDNSRecordResolver *IDNSRecordResolverCaller) DnsRecord(opts *bind.CallOpts, node [32]byte, name [32]byte, resource uint16) ([]byte, error) {
	var out []interface{}
	err := _IDNSRecordResolver.contract.Call(opts, &out, "dnsRecord", node, name, resource)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DnsRecord is a free data retrieval call binding the contract method 0xa8fa5682.
//
// Solidity: function dnsRecord(bytes32 node, bytes32 name, uint16 resource) view returns(bytes)
func (_IDNSRecordResolver *IDNSRecordResolverSession) DnsRecord(node [32]byte, name [32]byte, resource uint16) ([]byte, error) {
	return _IDNSRecordResolver.Contract.DnsRecord(&_IDNSRecordResolver.CallOpts, node, name, resource)
}

// DnsRecord is a free data retrieval call binding the contract method 0xa8fa5682.
//
// Solidity: function dnsRecord(bytes32 node, bytes32 name, uint16 resource) view returns(bytes)
func (_IDNSRecordResolver *IDNSRecordResolverCallerSession) DnsRecord(node [32]byte, name [32]byte, resource uint16) ([]byte, error) {
	return _IDNSRecordResolver.Contract.DnsRecord(&_IDNSRecordResolver.CallOpts, node, name, resource)
}

// IDNSRecordResolverDNSRecordChangedIterator is returned from FilterDNSRecordChanged and is used to iterate over the raw logs and unpacked data for DNSRecordChanged events raised by the IDNSRecordResolver contract.
type IDNSRecordResolverDNSRecordChangedIterator struct {
	Event *IDNSRecordResolverDNSRecordChanged // Event containing the contract specifics and raw log

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
func (it *IDNSRecordResolverDNSRecordChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDNSRecordResolverDNSRecordChanged)
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
		it.Event = new(IDNSRecordResolverDNSRecordChanged)
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
func (it *IDNSRecordResolverDNSRecordChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDNSRecordResolverDNSRecordChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDNSRecordResolverDNSRecordChanged represents a DNSRecordChanged event raised by the IDNSRecordResolver contract.
type IDNSRecordResolverDNSRecordChanged struct {
	Node     [32]byte
	Name     []byte
	Resource uint16
	Record   []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDNSRecordChanged is a free log retrieval operation binding the contract event 0x52a608b3303a48862d07a73d82fa221318c0027fbbcfb1b2329bface3f19ff2b.
//
// Solidity: event DNSRecordChanged(bytes32 indexed node, bytes name, uint16 resource, bytes record)
func (_IDNSRecordResolver *IDNSRecordResolverFilterer) FilterDNSRecordChanged(opts *bind.FilterOpts, node [][32]byte) (*IDNSRecordResolverDNSRecordChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _IDNSRecordResolver.contract.FilterLogs(opts, "DNSRecordChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &IDNSRecordResolverDNSRecordChangedIterator{contract: _IDNSRecordResolver.contract, event: "DNSRecordChanged", logs: logs, sub: sub}, nil
}

// WatchDNSRecordChanged is a free log subscription operation binding the contract event 0x52a608b3303a48862d07a73d82fa221318c0027fbbcfb1b2329bface3f19ff2b.
//
// Solidity: event DNSRecordChanged(bytes32 indexed node, bytes name, uint16 resource, bytes record)
func (_IDNSRecordResolver *IDNSRecordResolverFilterer) WatchDNSRecordChanged(opts *bind.WatchOpts, sink chan<- *IDNSRecordResolverDNSRecordChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _IDNSRecordResolver.contract.WatchLogs(opts, "DNSRecordChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDNSRecordResolverDNSRecordChanged)
				if err := _IDNSRecordResolver.contract.UnpackLog(event, "DNSRecordChanged", log); err != nil {
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

// ParseDNSRecordChanged is a log parse operation binding the contract event 0x52a608b3303a48862d07a73d82fa221318c0027fbbcfb1b2329bface3f19ff2b.
//
// Solidity: event DNSRecordChanged(bytes32 indexed node, bytes name, uint16 resource, bytes record)
func (_IDNSRecordResolver *IDNSRecordResolverFilterer) ParseDNSRecordChanged(log types.Log) (*IDNSRecordResolverDNSRecordChanged, error) {
	event := new(IDNSRecordResolverDNSRecordChanged)
	if err := _IDNSRecordResolver.contract.UnpackLog(event, "DNSRecordChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDNSRecordResolverDNSRecordDeletedIterator is returned from FilterDNSRecordDeleted and is used to iterate over the raw logs and unpacked data for DNSRecordDeleted events raised by the IDNSRecordResolver contract.
type IDNSRecordResolverDNSRecordDeletedIterator struct {
	Event *IDNSRecordResolverDNSRecordDeleted // Event containing the contract specifics and raw log

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
func (it *IDNSRecordResolverDNSRecordDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDNSRecordResolverDNSRecordDeleted)
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
		it.Event = new(IDNSRecordResolverDNSRecordDeleted)
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
func (it *IDNSRecordResolverDNSRecordDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDNSRecordResolverDNSRecordDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDNSRecordResolverDNSRecordDeleted represents a DNSRecordDeleted event raised by the IDNSRecordResolver contract.
type IDNSRecordResolverDNSRecordDeleted struct {
	Node     [32]byte
	Name     []byte
	Resource uint16
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDNSRecordDeleted is a free log retrieval operation binding the contract event 0x03528ed0c2a3ebc993b12ce3c16bb382f9c7d88ef7d8a1bf290eaf35955a1207.
//
// Solidity: event DNSRecordDeleted(bytes32 indexed node, bytes name, uint16 resource)
func (_IDNSRecordResolver *IDNSRecordResolverFilterer) FilterDNSRecordDeleted(opts *bind.FilterOpts, node [][32]byte) (*IDNSRecordResolverDNSRecordDeletedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _IDNSRecordResolver.contract.FilterLogs(opts, "DNSRecordDeleted", nodeRule)
	if err != nil {
		return nil, err
	}
	return &IDNSRecordResolverDNSRecordDeletedIterator{contract: _IDNSRecordResolver.contract, event: "DNSRecordDeleted", logs: logs, sub: sub}, nil
}

// WatchDNSRecordDeleted is a free log subscription operation binding the contract event 0x03528ed0c2a3ebc993b12ce3c16bb382f9c7d88ef7d8a1bf290eaf35955a1207.
//
// Solidity: event DNSRecordDeleted(bytes32 indexed node, bytes name, uint16 resource)
func (_IDNSRecordResolver *IDNSRecordResolverFilterer) WatchDNSRecordDeleted(opts *bind.WatchOpts, sink chan<- *IDNSRecordResolverDNSRecordDeleted, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _IDNSRecordResolver.contract.WatchLogs(opts, "DNSRecordDeleted", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDNSRecordResolverDNSRecordDeleted)
				if err := _IDNSRecordResolver.contract.UnpackLog(event, "DNSRecordDeleted", log); err != nil {
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

// ParseDNSRecordDeleted is a log parse operation binding the contract event 0x03528ed0c2a3ebc993b12ce3c16bb382f9c7d88ef7d8a1bf290eaf35955a1207.
//
// Solidity: event DNSRecordDeleted(bytes32 indexed node, bytes name, uint16 resource)
func (_IDNSRecordResolver *IDNSRecordResolverFilterer) ParseDNSRecordDeleted(log types.Log) (*IDNSRecordResolverDNSRecordDeleted, error) {
	event := new(IDNSRecordResolverDNSRecordDeleted)
	if err := _IDNSRecordResolver.contract.UnpackLog(event, "DNSRecordDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDNSRecordResolverDNSZoneClearedIterator is returned from FilterDNSZoneCleared and is used to iterate over the raw logs and unpacked data for DNSZoneCleared events raised by the IDNSRecordResolver contract.
type IDNSRecordResolverDNSZoneClearedIterator struct {
	Event *IDNSRecordResolverDNSZoneCleared // Event containing the contract specifics and raw log

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
func (it *IDNSRecordResolverDNSZoneClearedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDNSRecordResolverDNSZoneCleared)
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
		it.Event = new(IDNSRecordResolverDNSZoneCleared)
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
func (it *IDNSRecordResolverDNSZoneClearedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDNSRecordResolverDNSZoneClearedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDNSRecordResolverDNSZoneCleared represents a DNSZoneCleared event raised by the IDNSRecordResolver contract.
type IDNSRecordResolverDNSZoneCleared struct {
	Node [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDNSZoneCleared is a free log retrieval operation binding the contract event 0xb757169b8492ca2f1c6619d9d76ce22803035c3b1d5f6930dffe7b127c1a1983.
//
// Solidity: event DNSZoneCleared(bytes32 indexed node)
func (_IDNSRecordResolver *IDNSRecordResolverFilterer) FilterDNSZoneCleared(opts *bind.FilterOpts, node [][32]byte) (*IDNSRecordResolverDNSZoneClearedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _IDNSRecordResolver.contract.FilterLogs(opts, "DNSZoneCleared", nodeRule)
	if err != nil {
		return nil, err
	}
	return &IDNSRecordResolverDNSZoneClearedIterator{contract: _IDNSRecordResolver.contract, event: "DNSZoneCleared", logs: logs, sub: sub}, nil
}

// WatchDNSZoneCleared is a free log subscription operation binding the contract event 0xb757169b8492ca2f1c6619d9d76ce22803035c3b1d5f6930dffe7b127c1a1983.
//
// Solidity: event DNSZoneCleared(bytes32 indexed node)
func (_IDNSRecordResolver *IDNSRecordResolverFilterer) WatchDNSZoneCleared(opts *bind.WatchOpts, sink chan<- *IDNSRecordResolverDNSZoneCleared, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _IDNSRecordResolver.contract.WatchLogs(opts, "DNSZoneCleared", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDNSRecordResolverDNSZoneCleared)
				if err := _IDNSRecordResolver.contract.UnpackLog(event, "DNSZoneCleared", log); err != nil {
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

// ParseDNSZoneCleared is a log parse operation binding the contract event 0xb757169b8492ca2f1c6619d9d76ce22803035c3b1d5f6930dffe7b127c1a1983.
//
// Solidity: event DNSZoneCleared(bytes32 indexed node)
func (_IDNSRecordResolver *IDNSRecordResolverFilterer) ParseDNSZoneCleared(log types.Log) (*IDNSRecordResolverDNSZoneCleared, error) {
	event := new(IDNSRecordResolverDNSZoneCleared)
	if err := _IDNSRecordResolver.contract.UnpackLog(event, "DNSZoneCleared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDNSZoneResolverMetaData contains all meta data concerning the IDNSZoneResolver contract.
var IDNSZoneResolverMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"lastzonehash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"zonehash\",\"type\":\"bytes\"}],\"name\":\"DNSZonehashChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"zonehash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"5c98042b": "zonehash(bytes32)",
	},
}

// IDNSZoneResolverABI is the input ABI used to generate the binding from.
// Deprecated: Use IDNSZoneResolverMetaData.ABI instead.
var IDNSZoneResolverABI = IDNSZoneResolverMetaData.ABI

// Deprecated: Use IDNSZoneResolverMetaData.Sigs instead.
// IDNSZoneResolverFuncSigs maps the 4-byte function signature to its string representation.
var IDNSZoneResolverFuncSigs = IDNSZoneResolverMetaData.Sigs

// IDNSZoneResolver is an auto generated Go binding around an Ethereum contract.
type IDNSZoneResolver struct {
	IDNSZoneResolverCaller     // Read-only binding to the contract
	IDNSZoneResolverTransactor // Write-only binding to the contract
	IDNSZoneResolverFilterer   // Log filterer for contract events
}

// IDNSZoneResolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDNSZoneResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDNSZoneResolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDNSZoneResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDNSZoneResolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDNSZoneResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDNSZoneResolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDNSZoneResolverSession struct {
	Contract     *IDNSZoneResolver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDNSZoneResolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDNSZoneResolverCallerSession struct {
	Contract *IDNSZoneResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IDNSZoneResolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDNSZoneResolverTransactorSession struct {
	Contract     *IDNSZoneResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IDNSZoneResolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDNSZoneResolverRaw struct {
	Contract *IDNSZoneResolver // Generic contract binding to access the raw methods on
}

// IDNSZoneResolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDNSZoneResolverCallerRaw struct {
	Contract *IDNSZoneResolverCaller // Generic read-only contract binding to access the raw methods on
}

// IDNSZoneResolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDNSZoneResolverTransactorRaw struct {
	Contract *IDNSZoneResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDNSZoneResolver creates a new instance of IDNSZoneResolver, bound to a specific deployed contract.
func NewIDNSZoneResolver(address common.Address, backend bind.ContractBackend) (*IDNSZoneResolver, error) {
	contract, err := bindIDNSZoneResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDNSZoneResolver{IDNSZoneResolverCaller: IDNSZoneResolverCaller{contract: contract}, IDNSZoneResolverTransactor: IDNSZoneResolverTransactor{contract: contract}, IDNSZoneResolverFilterer: IDNSZoneResolverFilterer{contract: contract}}, nil
}

// NewIDNSZoneResolverCaller creates a new read-only instance of IDNSZoneResolver, bound to a specific deployed contract.
func NewIDNSZoneResolverCaller(address common.Address, caller bind.ContractCaller) (*IDNSZoneResolverCaller, error) {
	contract, err := bindIDNSZoneResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDNSZoneResolverCaller{contract: contract}, nil
}

// NewIDNSZoneResolverTransactor creates a new write-only instance of IDNSZoneResolver, bound to a specific deployed contract.
func NewIDNSZoneResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*IDNSZoneResolverTransactor, error) {
	contract, err := bindIDNSZoneResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDNSZoneResolverTransactor{contract: contract}, nil
}

// NewIDNSZoneResolverFilterer creates a new log filterer instance of IDNSZoneResolver, bound to a specific deployed contract.
func NewIDNSZoneResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*IDNSZoneResolverFilterer, error) {
	contract, err := bindIDNSZoneResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDNSZoneResolverFilterer{contract: contract}, nil
}

// bindIDNSZoneResolver binds a generic wrapper to an already deployed contract.
func bindIDNSZoneResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDNSZoneResolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDNSZoneResolver *IDNSZoneResolverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDNSZoneResolver.Contract.IDNSZoneResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDNSZoneResolver *IDNSZoneResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDNSZoneResolver.Contract.IDNSZoneResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDNSZoneResolver *IDNSZoneResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDNSZoneResolver.Contract.IDNSZoneResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDNSZoneResolver *IDNSZoneResolverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDNSZoneResolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDNSZoneResolver *IDNSZoneResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDNSZoneResolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDNSZoneResolver *IDNSZoneResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDNSZoneResolver.Contract.contract.Transact(opts, method, params...)
}

// Zonehash is a free data retrieval call binding the contract method 0x5c98042b.
//
// Solidity: function zonehash(bytes32 node) view returns(bytes)
func (_IDNSZoneResolver *IDNSZoneResolverCaller) Zonehash(opts *bind.CallOpts, node [32]byte) ([]byte, error) {
	var out []interface{}
	err := _IDNSZoneResolver.contract.Call(opts, &out, "zonehash", node)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Zonehash is a free data retrieval call binding the contract method 0x5c98042b.
//
// Solidity: function zonehash(bytes32 node) view returns(bytes)
func (_IDNSZoneResolver *IDNSZoneResolverSession) Zonehash(node [32]byte) ([]byte, error) {
	return _IDNSZoneResolver.Contract.Zonehash(&_IDNSZoneResolver.CallOpts, node)
}

// Zonehash is a free data retrieval call binding the contract method 0x5c98042b.
//
// Solidity: function zonehash(bytes32 node) view returns(bytes)
func (_IDNSZoneResolver *IDNSZoneResolverCallerSession) Zonehash(node [32]byte) ([]byte, error) {
	return _IDNSZoneResolver.Contract.Zonehash(&_IDNSZoneResolver.CallOpts, node)
}

// IDNSZoneResolverDNSZonehashChangedIterator is returned from FilterDNSZonehashChanged and is used to iterate over the raw logs and unpacked data for DNSZonehashChanged events raised by the IDNSZoneResolver contract.
type IDNSZoneResolverDNSZonehashChangedIterator struct {
	Event *IDNSZoneResolverDNSZonehashChanged // Event containing the contract specifics and raw log

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
func (it *IDNSZoneResolverDNSZonehashChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDNSZoneResolverDNSZonehashChanged)
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
		it.Event = new(IDNSZoneResolverDNSZonehashChanged)
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
func (it *IDNSZoneResolverDNSZonehashChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDNSZoneResolverDNSZonehashChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDNSZoneResolverDNSZonehashChanged represents a DNSZonehashChanged event raised by the IDNSZoneResolver contract.
type IDNSZoneResolverDNSZonehashChanged struct {
	Node         [32]byte
	Lastzonehash []byte
	Zonehash     []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDNSZonehashChanged is a free log retrieval operation binding the contract event 0x8f15ed4b723ef428f250961da8315675b507046737e19319fc1a4d81bfe87f85.
//
// Solidity: event DNSZonehashChanged(bytes32 indexed node, bytes lastzonehash, bytes zonehash)
func (_IDNSZoneResolver *IDNSZoneResolverFilterer) FilterDNSZonehashChanged(opts *bind.FilterOpts, node [][32]byte) (*IDNSZoneResolverDNSZonehashChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _IDNSZoneResolver.contract.FilterLogs(opts, "DNSZonehashChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &IDNSZoneResolverDNSZonehashChangedIterator{contract: _IDNSZoneResolver.contract, event: "DNSZonehashChanged", logs: logs, sub: sub}, nil
}

// WatchDNSZonehashChanged is a free log subscription operation binding the contract event 0x8f15ed4b723ef428f250961da8315675b507046737e19319fc1a4d81bfe87f85.
//
// Solidity: event DNSZonehashChanged(bytes32 indexed node, bytes lastzonehash, bytes zonehash)
func (_IDNSZoneResolver *IDNSZoneResolverFilterer) WatchDNSZonehashChanged(opts *bind.WatchOpts, sink chan<- *IDNSZoneResolverDNSZonehashChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _IDNSZoneResolver.contract.WatchLogs(opts, "DNSZonehashChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDNSZoneResolverDNSZonehashChanged)
				if err := _IDNSZoneResolver.contract.UnpackLog(event, "DNSZonehashChanged", log); err != nil {
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

// ParseDNSZonehashChanged is a log parse operation binding the contract event 0x8f15ed4b723ef428f250961da8315675b507046737e19319fc1a4d81bfe87f85.
//
// Solidity: event DNSZonehashChanged(bytes32 indexed node, bytes lastzonehash, bytes zonehash)
func (_IDNSZoneResolver *IDNSZoneResolverFilterer) ParseDNSZonehashChanged(log types.Log) (*IDNSZoneResolverDNSZonehashChanged, error) {
	event := new(IDNSZoneResolverDNSZonehashChanged)
	if err := _IDNSZoneResolver.contract.UnpackLog(event, "DNSZonehashChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC165MetaData contains all meta data concerning the IERC165 contract.
var IERC165MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"01ffc9a7": "supportsInterface(bytes4)",
	},
}

// IERC165ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC165MetaData.ABI instead.
var IERC165ABI = IERC165MetaData.ABI

// Deprecated: Use IERC165MetaData.Sigs instead.
// IERC165FuncSigs maps the 4-byte function signature to its string representation.
var IERC165FuncSigs = IERC165MetaData.Sigs

// IERC165 is an auto generated Go binding around an Ethereum contract.
type IERC165 struct {
	IERC165Caller     // Read-only binding to the contract
	IERC165Transactor // Write-only binding to the contract
	IERC165Filterer   // Log filterer for contract events
}

// IERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC165Session struct {
	Contract     *IERC165          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC165CallerSession struct {
	Contract *IERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC165TransactorSession struct {
	Contract     *IERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC165Raw struct {
	Contract *IERC165 // Generic contract binding to access the raw methods on
}

// IERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC165CallerRaw struct {
	Contract *IERC165Caller // Generic read-only contract binding to access the raw methods on
}

// IERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC165TransactorRaw struct {
	Contract *IERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC165 creates a new instance of IERC165, bound to a specific deployed contract.
func NewIERC165(address common.Address, backend bind.ContractBackend) (*IERC165, error) {
	contract, err := bindIERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC165{IERC165Caller: IERC165Caller{contract: contract}, IERC165Transactor: IERC165Transactor{contract: contract}, IERC165Filterer: IERC165Filterer{contract: contract}}, nil
}

// NewIERC165Caller creates a new read-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Caller(address common.Address, caller bind.ContractCaller) (*IERC165Caller, error) {
	contract, err := bindIERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Caller{contract: contract}, nil
}

// NewIERC165Transactor creates a new write-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC165Transactor, error) {
	contract, err := bindIERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Transactor{contract: contract}, nil
}

// NewIERC165Filterer creates a new log filterer instance of IERC165, bound to a specific deployed contract.
func NewIERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC165Filterer, error) {
	contract, err := bindIERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC165Filterer{contract: contract}, nil
}

// bindIERC165 binds a generic wrapper to an already deployed contract.
func bindIERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.IERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
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

// IInterfaceResolverMetaData contains all meta data concerning the IInterfaceResolver contract.
var IInterfaceResolverMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementer\",\"type\":\"address\"}],\"name\":\"InterfaceChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"interfaceImplementer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"124a319c": "interfaceImplementer(bytes32,bytes4)",
	},
}

// IInterfaceResolverABI is the input ABI used to generate the binding from.
// Deprecated: Use IInterfaceResolverMetaData.ABI instead.
var IInterfaceResolverABI = IInterfaceResolverMetaData.ABI

// Deprecated: Use IInterfaceResolverMetaData.Sigs instead.
// IInterfaceResolverFuncSigs maps the 4-byte function signature to its string representation.
var IInterfaceResolverFuncSigs = IInterfaceResolverMetaData.Sigs

// IInterfaceResolver is an auto generated Go binding around an Ethereum contract.
type IInterfaceResolver struct {
	IInterfaceResolverCaller     // Read-only binding to the contract
	IInterfaceResolverTransactor // Write-only binding to the contract
	IInterfaceResolverFilterer   // Log filterer for contract events
}

// IInterfaceResolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IInterfaceResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterfaceResolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IInterfaceResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterfaceResolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IInterfaceResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterfaceResolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IInterfaceResolverSession struct {
	Contract     *IInterfaceResolver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IInterfaceResolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IInterfaceResolverCallerSession struct {
	Contract *IInterfaceResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IInterfaceResolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IInterfaceResolverTransactorSession struct {
	Contract     *IInterfaceResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IInterfaceResolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IInterfaceResolverRaw struct {
	Contract *IInterfaceResolver // Generic contract binding to access the raw methods on
}

// IInterfaceResolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IInterfaceResolverCallerRaw struct {
	Contract *IInterfaceResolverCaller // Generic read-only contract binding to access the raw methods on
}

// IInterfaceResolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IInterfaceResolverTransactorRaw struct {
	Contract *IInterfaceResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIInterfaceResolver creates a new instance of IInterfaceResolver, bound to a specific deployed contract.
func NewIInterfaceResolver(address common.Address, backend bind.ContractBackend) (*IInterfaceResolver, error) {
	contract, err := bindIInterfaceResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IInterfaceResolver{IInterfaceResolverCaller: IInterfaceResolverCaller{contract: contract}, IInterfaceResolverTransactor: IInterfaceResolverTransactor{contract: contract}, IInterfaceResolverFilterer: IInterfaceResolverFilterer{contract: contract}}, nil
}

// NewIInterfaceResolverCaller creates a new read-only instance of IInterfaceResolver, bound to a specific deployed contract.
func NewIInterfaceResolverCaller(address common.Address, caller bind.ContractCaller) (*IInterfaceResolverCaller, error) {
	contract, err := bindIInterfaceResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IInterfaceResolverCaller{contract: contract}, nil
}

// NewIInterfaceResolverTransactor creates a new write-only instance of IInterfaceResolver, bound to a specific deployed contract.
func NewIInterfaceResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*IInterfaceResolverTransactor, error) {
	contract, err := bindIInterfaceResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IInterfaceResolverTransactor{contract: contract}, nil
}

// NewIInterfaceResolverFilterer creates a new log filterer instance of IInterfaceResolver, bound to a specific deployed contract.
func NewIInterfaceResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*IInterfaceResolverFilterer, error) {
	contract, err := bindIInterfaceResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IInterfaceResolverFilterer{contract: contract}, nil
}

// bindIInterfaceResolver binds a generic wrapper to an already deployed contract.
func bindIInterfaceResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IInterfaceResolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInterfaceResolver *IInterfaceResolverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInterfaceResolver.Contract.IInterfaceResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInterfaceResolver *IInterfaceResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInterfaceResolver.Contract.IInterfaceResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInterfaceResolver *IInterfaceResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInterfaceResolver.Contract.IInterfaceResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInterfaceResolver *IInterfaceResolverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInterfaceResolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInterfaceResolver *IInterfaceResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInterfaceResolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInterfaceResolver *IInterfaceResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInterfaceResolver.Contract.contract.Transact(opts, method, params...)
}

// InterfaceImplementer is a free data retrieval call binding the contract method 0x124a319c.
//
// Solidity: function interfaceImplementer(bytes32 node, bytes4 interfaceID) view returns(address)
func (_IInterfaceResolver *IInterfaceResolverCaller) InterfaceImplementer(opts *bind.CallOpts, node [32]byte, interfaceID [4]byte) (common.Address, error) {
	var out []interface{}
	err := _IInterfaceResolver.contract.Call(opts, &out, "interfaceImplementer", node, interfaceID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterfaceImplementer is a free data retrieval call binding the contract method 0x124a319c.
//
// Solidity: function interfaceImplementer(bytes32 node, bytes4 interfaceID) view returns(address)
func (_IInterfaceResolver *IInterfaceResolverSession) InterfaceImplementer(node [32]byte, interfaceID [4]byte) (common.Address, error) {
	return _IInterfaceResolver.Contract.InterfaceImplementer(&_IInterfaceResolver.CallOpts, node, interfaceID)
}

// InterfaceImplementer is a free data retrieval call binding the contract method 0x124a319c.
//
// Solidity: function interfaceImplementer(bytes32 node, bytes4 interfaceID) view returns(address)
func (_IInterfaceResolver *IInterfaceResolverCallerSession) InterfaceImplementer(node [32]byte, interfaceID [4]byte) (common.Address, error) {
	return _IInterfaceResolver.Contract.InterfaceImplementer(&_IInterfaceResolver.CallOpts, node, interfaceID)
}

// IInterfaceResolverInterfaceChangedIterator is returned from FilterInterfaceChanged and is used to iterate over the raw logs and unpacked data for InterfaceChanged events raised by the IInterfaceResolver contract.
type IInterfaceResolverInterfaceChangedIterator struct {
	Event *IInterfaceResolverInterfaceChanged // Event containing the contract specifics and raw log

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
func (it *IInterfaceResolverInterfaceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IInterfaceResolverInterfaceChanged)
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
		it.Event = new(IInterfaceResolverInterfaceChanged)
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
func (it *IInterfaceResolverInterfaceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IInterfaceResolverInterfaceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IInterfaceResolverInterfaceChanged represents a InterfaceChanged event raised by the IInterfaceResolver contract.
type IInterfaceResolverInterfaceChanged struct {
	Node        [32]byte
	InterfaceID [4]byte
	Implementer common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInterfaceChanged is a free log retrieval operation binding the contract event 0x7c69f06bea0bdef565b709e93a147836b0063ba2dd89f02d0b7e8d931e6a6daa.
//
// Solidity: event InterfaceChanged(bytes32 indexed node, bytes4 indexed interfaceID, address implementer)
func (_IInterfaceResolver *IInterfaceResolverFilterer) FilterInterfaceChanged(opts *bind.FilterOpts, node [][32]byte, interfaceID [][4]byte) (*IInterfaceResolverInterfaceChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var interfaceIDRule []interface{}
	for _, interfaceIDItem := range interfaceID {
		interfaceIDRule = append(interfaceIDRule, interfaceIDItem)
	}

	logs, sub, err := _IInterfaceResolver.contract.FilterLogs(opts, "InterfaceChanged", nodeRule, interfaceIDRule)
	if err != nil {
		return nil, err
	}
	return &IInterfaceResolverInterfaceChangedIterator{contract: _IInterfaceResolver.contract, event: "InterfaceChanged", logs: logs, sub: sub}, nil
}

// WatchInterfaceChanged is a free log subscription operation binding the contract event 0x7c69f06bea0bdef565b709e93a147836b0063ba2dd89f02d0b7e8d931e6a6daa.
//
// Solidity: event InterfaceChanged(bytes32 indexed node, bytes4 indexed interfaceID, address implementer)
func (_IInterfaceResolver *IInterfaceResolverFilterer) WatchInterfaceChanged(opts *bind.WatchOpts, sink chan<- *IInterfaceResolverInterfaceChanged, node [][32]byte, interfaceID [][4]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var interfaceIDRule []interface{}
	for _, interfaceIDItem := range interfaceID {
		interfaceIDRule = append(interfaceIDRule, interfaceIDItem)
	}

	logs, sub, err := _IInterfaceResolver.contract.WatchLogs(opts, "InterfaceChanged", nodeRule, interfaceIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IInterfaceResolverInterfaceChanged)
				if err := _IInterfaceResolver.contract.UnpackLog(event, "InterfaceChanged", log); err != nil {
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

// ParseInterfaceChanged is a log parse operation binding the contract event 0x7c69f06bea0bdef565b709e93a147836b0063ba2dd89f02d0b7e8d931e6a6daa.
//
// Solidity: event InterfaceChanged(bytes32 indexed node, bytes4 indexed interfaceID, address implementer)
func (_IInterfaceResolver *IInterfaceResolverFilterer) ParseInterfaceChanged(log types.Log) (*IInterfaceResolverInterfaceChanged, error) {
	event := new(IInterfaceResolverInterfaceChanged)
	if err := _IInterfaceResolver.contract.UnpackLog(event, "InterfaceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INameResolverMetaData contains all meta data concerning the INameResolver contract.
var INameResolverMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"NameChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"691f3431": "name(bytes32)",
	},
}

// INameResolverABI is the input ABI used to generate the binding from.
// Deprecated: Use INameResolverMetaData.ABI instead.
var INameResolverABI = INameResolverMetaData.ABI

// Deprecated: Use INameResolverMetaData.Sigs instead.
// INameResolverFuncSigs maps the 4-byte function signature to its string representation.
var INameResolverFuncSigs = INameResolverMetaData.Sigs

// INameResolver is an auto generated Go binding around an Ethereum contract.
type INameResolver struct {
	INameResolverCaller     // Read-only binding to the contract
	INameResolverTransactor // Write-only binding to the contract
	INameResolverFilterer   // Log filterer for contract events
}

// INameResolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type INameResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INameResolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type INameResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INameResolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type INameResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INameResolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type INameResolverSession struct {
	Contract     *INameResolver    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// INameResolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type INameResolverCallerSession struct {
	Contract *INameResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// INameResolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type INameResolverTransactorSession struct {
	Contract     *INameResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// INameResolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type INameResolverRaw struct {
	Contract *INameResolver // Generic contract binding to access the raw methods on
}

// INameResolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type INameResolverCallerRaw struct {
	Contract *INameResolverCaller // Generic read-only contract binding to access the raw methods on
}

// INameResolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type INameResolverTransactorRaw struct {
	Contract *INameResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewINameResolver creates a new instance of INameResolver, bound to a specific deployed contract.
func NewINameResolver(address common.Address, backend bind.ContractBackend) (*INameResolver, error) {
	contract, err := bindINameResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &INameResolver{INameResolverCaller: INameResolverCaller{contract: contract}, INameResolverTransactor: INameResolverTransactor{contract: contract}, INameResolverFilterer: INameResolverFilterer{contract: contract}}, nil
}

// NewINameResolverCaller creates a new read-only instance of INameResolver, bound to a specific deployed contract.
func NewINameResolverCaller(address common.Address, caller bind.ContractCaller) (*INameResolverCaller, error) {
	contract, err := bindINameResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &INameResolverCaller{contract: contract}, nil
}

// NewINameResolverTransactor creates a new write-only instance of INameResolver, bound to a specific deployed contract.
func NewINameResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*INameResolverTransactor, error) {
	contract, err := bindINameResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &INameResolverTransactor{contract: contract}, nil
}

// NewINameResolverFilterer creates a new log filterer instance of INameResolver, bound to a specific deployed contract.
func NewINameResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*INameResolverFilterer, error) {
	contract, err := bindINameResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &INameResolverFilterer{contract: contract}, nil
}

// bindINameResolver binds a generic wrapper to an already deployed contract.
func bindINameResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(INameResolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INameResolver *INameResolverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INameResolver.Contract.INameResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INameResolver *INameResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INameResolver.Contract.INameResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INameResolver *INameResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INameResolver.Contract.INameResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INameResolver *INameResolverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INameResolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INameResolver *INameResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INameResolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INameResolver *INameResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INameResolver.Contract.contract.Transact(opts, method, params...)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 node) view returns(string)
func (_INameResolver *INameResolverCaller) Name(opts *bind.CallOpts, node [32]byte) (string, error) {
	var out []interface{}
	err := _INameResolver.contract.Call(opts, &out, "name", node)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 node) view returns(string)
func (_INameResolver *INameResolverSession) Name(node [32]byte) (string, error) {
	return _INameResolver.Contract.Name(&_INameResolver.CallOpts, node)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 node) view returns(string)
func (_INameResolver *INameResolverCallerSession) Name(node [32]byte) (string, error) {
	return _INameResolver.Contract.Name(&_INameResolver.CallOpts, node)
}

// INameResolverNameChangedIterator is returned from FilterNameChanged and is used to iterate over the raw logs and unpacked data for NameChanged events raised by the INameResolver contract.
type INameResolverNameChangedIterator struct {
	Event *INameResolverNameChanged // Event containing the contract specifics and raw log

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
func (it *INameResolverNameChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INameResolverNameChanged)
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
		it.Event = new(INameResolverNameChanged)
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
func (it *INameResolverNameChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INameResolverNameChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INameResolverNameChanged represents a NameChanged event raised by the INameResolver contract.
type INameResolverNameChanged struct {
	Node [32]byte
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNameChanged is a free log retrieval operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: event NameChanged(bytes32 indexed node, string name)
func (_INameResolver *INameResolverFilterer) FilterNameChanged(opts *bind.FilterOpts, node [][32]byte) (*INameResolverNameChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _INameResolver.contract.FilterLogs(opts, "NameChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &INameResolverNameChangedIterator{contract: _INameResolver.contract, event: "NameChanged", logs: logs, sub: sub}, nil
}

// WatchNameChanged is a free log subscription operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: event NameChanged(bytes32 indexed node, string name)
func (_INameResolver *INameResolverFilterer) WatchNameChanged(opts *bind.WatchOpts, sink chan<- *INameResolverNameChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _INameResolver.contract.WatchLogs(opts, "NameChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INameResolverNameChanged)
				if err := _INameResolver.contract.UnpackLog(event, "NameChanged", log); err != nil {
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

// ParseNameChanged is a log parse operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: event NameChanged(bytes32 indexed node, string name)
func (_INameResolver *INameResolverFilterer) ParseNameChanged(log types.Log) (*INameResolverNameChanged, error) {
	event := new(INameResolverNameChanged)
	if err := _INameResolver.contract.UnpackLog(event, "NameChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPubkeyResolverMetaData contains all meta data concerning the IPubkeyResolver contract.
var IPubkeyResolverMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"name\":\"PubkeyChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"pubkey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"c8690233": "pubkey(bytes32)",
	},
}

// IPubkeyResolverABI is the input ABI used to generate the binding from.
// Deprecated: Use IPubkeyResolverMetaData.ABI instead.
var IPubkeyResolverABI = IPubkeyResolverMetaData.ABI

// Deprecated: Use IPubkeyResolverMetaData.Sigs instead.
// IPubkeyResolverFuncSigs maps the 4-byte function signature to its string representation.
var IPubkeyResolverFuncSigs = IPubkeyResolverMetaData.Sigs

// IPubkeyResolver is an auto generated Go binding around an Ethereum contract.
type IPubkeyResolver struct {
	IPubkeyResolverCaller     // Read-only binding to the contract
	IPubkeyResolverTransactor // Write-only binding to the contract
	IPubkeyResolverFilterer   // Log filterer for contract events
}

// IPubkeyResolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPubkeyResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPubkeyResolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPubkeyResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPubkeyResolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPubkeyResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPubkeyResolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPubkeyResolverSession struct {
	Contract     *IPubkeyResolver  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPubkeyResolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPubkeyResolverCallerSession struct {
	Contract *IPubkeyResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IPubkeyResolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPubkeyResolverTransactorSession struct {
	Contract     *IPubkeyResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IPubkeyResolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPubkeyResolverRaw struct {
	Contract *IPubkeyResolver // Generic contract binding to access the raw methods on
}

// IPubkeyResolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPubkeyResolverCallerRaw struct {
	Contract *IPubkeyResolverCaller // Generic read-only contract binding to access the raw methods on
}

// IPubkeyResolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPubkeyResolverTransactorRaw struct {
	Contract *IPubkeyResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPubkeyResolver creates a new instance of IPubkeyResolver, bound to a specific deployed contract.
func NewIPubkeyResolver(address common.Address, backend bind.ContractBackend) (*IPubkeyResolver, error) {
	contract, err := bindIPubkeyResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPubkeyResolver{IPubkeyResolverCaller: IPubkeyResolverCaller{contract: contract}, IPubkeyResolverTransactor: IPubkeyResolverTransactor{contract: contract}, IPubkeyResolverFilterer: IPubkeyResolverFilterer{contract: contract}}, nil
}

// NewIPubkeyResolverCaller creates a new read-only instance of IPubkeyResolver, bound to a specific deployed contract.
func NewIPubkeyResolverCaller(address common.Address, caller bind.ContractCaller) (*IPubkeyResolverCaller, error) {
	contract, err := bindIPubkeyResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPubkeyResolverCaller{contract: contract}, nil
}

// NewIPubkeyResolverTransactor creates a new write-only instance of IPubkeyResolver, bound to a specific deployed contract.
func NewIPubkeyResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*IPubkeyResolverTransactor, error) {
	contract, err := bindIPubkeyResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPubkeyResolverTransactor{contract: contract}, nil
}

// NewIPubkeyResolverFilterer creates a new log filterer instance of IPubkeyResolver, bound to a specific deployed contract.
func NewIPubkeyResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*IPubkeyResolverFilterer, error) {
	contract, err := bindIPubkeyResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPubkeyResolverFilterer{contract: contract}, nil
}

// bindIPubkeyResolver binds a generic wrapper to an already deployed contract.
func bindIPubkeyResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPubkeyResolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPubkeyResolver *IPubkeyResolverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPubkeyResolver.Contract.IPubkeyResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPubkeyResolver *IPubkeyResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPubkeyResolver.Contract.IPubkeyResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPubkeyResolver *IPubkeyResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPubkeyResolver.Contract.IPubkeyResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPubkeyResolver *IPubkeyResolverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPubkeyResolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPubkeyResolver *IPubkeyResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPubkeyResolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPubkeyResolver *IPubkeyResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPubkeyResolver.Contract.contract.Transact(opts, method, params...)
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(bytes32 node) view returns(bytes32 x, bytes32 y)
func (_IPubkeyResolver *IPubkeyResolverCaller) Pubkey(opts *bind.CallOpts, node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	var out []interface{}
	err := _IPubkeyResolver.contract.Call(opts, &out, "pubkey", node)

	outstruct := new(struct {
		X [32]byte
		Y [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Y = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(bytes32 node) view returns(bytes32 x, bytes32 y)
func (_IPubkeyResolver *IPubkeyResolverSession) Pubkey(node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	return _IPubkeyResolver.Contract.Pubkey(&_IPubkeyResolver.CallOpts, node)
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(bytes32 node) view returns(bytes32 x, bytes32 y)
func (_IPubkeyResolver *IPubkeyResolverCallerSession) Pubkey(node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	return _IPubkeyResolver.Contract.Pubkey(&_IPubkeyResolver.CallOpts, node)
}

// IPubkeyResolverPubkeyChangedIterator is returned from FilterPubkeyChanged and is used to iterate over the raw logs and unpacked data for PubkeyChanged events raised by the IPubkeyResolver contract.
type IPubkeyResolverPubkeyChangedIterator struct {
	Event *IPubkeyResolverPubkeyChanged // Event containing the contract specifics and raw log

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
func (it *IPubkeyResolverPubkeyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPubkeyResolverPubkeyChanged)
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
		it.Event = new(IPubkeyResolverPubkeyChanged)
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
func (it *IPubkeyResolverPubkeyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPubkeyResolverPubkeyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPubkeyResolverPubkeyChanged represents a PubkeyChanged event raised by the IPubkeyResolver contract.
type IPubkeyResolverPubkeyChanged struct {
	Node [32]byte
	X    [32]byte
	Y    [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPubkeyChanged is a free log retrieval operation binding the contract event 0x1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46.
//
// Solidity: event PubkeyChanged(bytes32 indexed node, bytes32 x, bytes32 y)
func (_IPubkeyResolver *IPubkeyResolverFilterer) FilterPubkeyChanged(opts *bind.FilterOpts, node [][32]byte) (*IPubkeyResolverPubkeyChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _IPubkeyResolver.contract.FilterLogs(opts, "PubkeyChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &IPubkeyResolverPubkeyChangedIterator{contract: _IPubkeyResolver.contract, event: "PubkeyChanged", logs: logs, sub: sub}, nil
}

// WatchPubkeyChanged is a free log subscription operation binding the contract event 0x1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46.
//
// Solidity: event PubkeyChanged(bytes32 indexed node, bytes32 x, bytes32 y)
func (_IPubkeyResolver *IPubkeyResolverFilterer) WatchPubkeyChanged(opts *bind.WatchOpts, sink chan<- *IPubkeyResolverPubkeyChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _IPubkeyResolver.contract.WatchLogs(opts, "PubkeyChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPubkeyResolverPubkeyChanged)
				if err := _IPubkeyResolver.contract.UnpackLog(event, "PubkeyChanged", log); err != nil {
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

// ParsePubkeyChanged is a log parse operation binding the contract event 0x1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46.
//
// Solidity: event PubkeyChanged(bytes32 indexed node, bytes32 x, bytes32 y)
func (_IPubkeyResolver *IPubkeyResolverFilterer) ParsePubkeyChanged(log types.Log) (*IPubkeyResolverPubkeyChanged, error) {
	event := new(IPubkeyResolverPubkeyChanged)
	if err := _IPubkeyResolver.contract.UnpackLog(event, "PubkeyChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITextResolverMetaData contains all meta data concerning the ITextResolver contract.
var ITextResolverMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"indexedKey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"TextChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"text\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"59d1d43c": "text(bytes32,string)",
	},
}

// ITextResolverABI is the input ABI used to generate the binding from.
// Deprecated: Use ITextResolverMetaData.ABI instead.
var ITextResolverABI = ITextResolverMetaData.ABI

// Deprecated: Use ITextResolverMetaData.Sigs instead.
// ITextResolverFuncSigs maps the 4-byte function signature to its string representation.
var ITextResolverFuncSigs = ITextResolverMetaData.Sigs

// ITextResolver is an auto generated Go binding around an Ethereum contract.
type ITextResolver struct {
	ITextResolverCaller     // Read-only binding to the contract
	ITextResolverTransactor // Write-only binding to the contract
	ITextResolverFilterer   // Log filterer for contract events
}

// ITextResolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITextResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITextResolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITextResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITextResolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITextResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITextResolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITextResolverSession struct {
	Contract     *ITextResolver    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITextResolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITextResolverCallerSession struct {
	Contract *ITextResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ITextResolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITextResolverTransactorSession struct {
	Contract     *ITextResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ITextResolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITextResolverRaw struct {
	Contract *ITextResolver // Generic contract binding to access the raw methods on
}

// ITextResolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITextResolverCallerRaw struct {
	Contract *ITextResolverCaller // Generic read-only contract binding to access the raw methods on
}

// ITextResolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITextResolverTransactorRaw struct {
	Contract *ITextResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITextResolver creates a new instance of ITextResolver, bound to a specific deployed contract.
func NewITextResolver(address common.Address, backend bind.ContractBackend) (*ITextResolver, error) {
	contract, err := bindITextResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITextResolver{ITextResolverCaller: ITextResolverCaller{contract: contract}, ITextResolverTransactor: ITextResolverTransactor{contract: contract}, ITextResolverFilterer: ITextResolverFilterer{contract: contract}}, nil
}

// NewITextResolverCaller creates a new read-only instance of ITextResolver, bound to a specific deployed contract.
func NewITextResolverCaller(address common.Address, caller bind.ContractCaller) (*ITextResolverCaller, error) {
	contract, err := bindITextResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITextResolverCaller{contract: contract}, nil
}

// NewITextResolverTransactor creates a new write-only instance of ITextResolver, bound to a specific deployed contract.
func NewITextResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*ITextResolverTransactor, error) {
	contract, err := bindITextResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITextResolverTransactor{contract: contract}, nil
}

// NewITextResolverFilterer creates a new log filterer instance of ITextResolver, bound to a specific deployed contract.
func NewITextResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*ITextResolverFilterer, error) {
	contract, err := bindITextResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITextResolverFilterer{contract: contract}, nil
}

// bindITextResolver binds a generic wrapper to an already deployed contract.
func bindITextResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITextResolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITextResolver *ITextResolverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITextResolver.Contract.ITextResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITextResolver *ITextResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITextResolver.Contract.ITextResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITextResolver *ITextResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITextResolver.Contract.ITextResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITextResolver *ITextResolverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITextResolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITextResolver *ITextResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITextResolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITextResolver *ITextResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITextResolver.Contract.contract.Transact(opts, method, params...)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(bytes32 node, string key) view returns(string)
func (_ITextResolver *ITextResolverCaller) Text(opts *bind.CallOpts, node [32]byte, key string) (string, error) {
	var out []interface{}
	err := _ITextResolver.contract.Call(opts, &out, "text", node, key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(bytes32 node, string key) view returns(string)
func (_ITextResolver *ITextResolverSession) Text(node [32]byte, key string) (string, error) {
	return _ITextResolver.Contract.Text(&_ITextResolver.CallOpts, node, key)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(bytes32 node, string key) view returns(string)
func (_ITextResolver *ITextResolverCallerSession) Text(node [32]byte, key string) (string, error) {
	return _ITextResolver.Contract.Text(&_ITextResolver.CallOpts, node, key)
}

// ITextResolverTextChangedIterator is returned from FilterTextChanged and is used to iterate over the raw logs and unpacked data for TextChanged events raised by the ITextResolver contract.
type ITextResolverTextChangedIterator struct {
	Event *ITextResolverTextChanged // Event containing the contract specifics and raw log

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
func (it *ITextResolverTextChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITextResolverTextChanged)
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
		it.Event = new(ITextResolverTextChanged)
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
func (it *ITextResolverTextChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITextResolverTextChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITextResolverTextChanged represents a TextChanged event raised by the ITextResolver contract.
type ITextResolverTextChanged struct {
	Node       [32]byte
	IndexedKey common.Hash
	Key        string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTextChanged is a free log retrieval operation binding the contract event 0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550.
//
// Solidity: event TextChanged(bytes32 indexed node, string indexed indexedKey, string key)
func (_ITextResolver *ITextResolverFilterer) FilterTextChanged(opts *bind.FilterOpts, node [][32]byte, indexedKey []string) (*ITextResolverTextChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _ITextResolver.contract.FilterLogs(opts, "TextChanged", nodeRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return &ITextResolverTextChangedIterator{contract: _ITextResolver.contract, event: "TextChanged", logs: logs, sub: sub}, nil
}

// WatchTextChanged is a free log subscription operation binding the contract event 0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550.
//
// Solidity: event TextChanged(bytes32 indexed node, string indexed indexedKey, string key)
func (_ITextResolver *ITextResolverFilterer) WatchTextChanged(opts *bind.WatchOpts, sink chan<- *ITextResolverTextChanged, node [][32]byte, indexedKey []string) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _ITextResolver.contract.WatchLogs(opts, "TextChanged", nodeRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITextResolverTextChanged)
				if err := _ITextResolver.contract.UnpackLog(event, "TextChanged", log); err != nil {
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

// ParseTextChanged is a log parse operation binding the contract event 0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550.
//
// Solidity: event TextChanged(bytes32 indexed node, string indexed indexedKey, string key)
func (_ITextResolver *ITextResolverFilterer) ParseTextChanged(log types.Log) (*ITextResolverTextChanged, error) {
	event := new(ITextResolverTextChanged)
	if err := _ITextResolver.contract.UnpackLog(event, "TextChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverMetaData contains all meta data concerning the Resolver contract.
var ResolverMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"indexedKey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"TextChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"text\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"01ffc9a7": "supportsInterface(bytes4)",
		"59d1d43c": "text(bytes32,string)",
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

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Resolver *ResolverCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Resolver.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Resolver *ResolverSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Resolver.Contract.SupportsInterface(&_Resolver.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Resolver *ResolverCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Resolver.Contract.SupportsInterface(&_Resolver.CallOpts, interfaceId)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(bytes32 node, string key) view returns(string)
func (_Resolver *ResolverCaller) Text(opts *bind.CallOpts, node [32]byte, key string) (string, error) {
	var out []interface{}
	err := _Resolver.contract.Call(opts, &out, "text", node, key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(bytes32 node, string key) view returns(string)
func (_Resolver *ResolverSession) Text(node [32]byte, key string) (string, error) {
	return _Resolver.Contract.Text(&_Resolver.CallOpts, node, key)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(bytes32 node, string key) view returns(string)
func (_Resolver *ResolverCallerSession) Text(node [32]byte, key string) (string, error) {
	return _Resolver.Contract.Text(&_Resolver.CallOpts, node, key)
}

// ResolverTextChangedIterator is returned from FilterTextChanged and is used to iterate over the raw logs and unpacked data for TextChanged events raised by the Resolver contract.
type ResolverTextChangedIterator struct {
	Event *ResolverTextChanged // Event containing the contract specifics and raw log

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
func (it *ResolverTextChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverTextChanged)
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
		it.Event = new(ResolverTextChanged)
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
func (it *ResolverTextChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverTextChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverTextChanged represents a TextChanged event raised by the Resolver contract.
type ResolverTextChanged struct {
	Node       [32]byte
	IndexedKey common.Hash
	Key        string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTextChanged is a free log retrieval operation binding the contract event 0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550.
//
// Solidity: event TextChanged(bytes32 indexed node, string indexed indexedKey, string key)
func (_Resolver *ResolverFilterer) FilterTextChanged(opts *bind.FilterOpts, node [][32]byte, indexedKey []string) (*ResolverTextChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _Resolver.contract.FilterLogs(opts, "TextChanged", nodeRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return &ResolverTextChangedIterator{contract: _Resolver.contract, event: "TextChanged", logs: logs, sub: sub}, nil
}

// WatchTextChanged is a free log subscription operation binding the contract event 0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550.
//
// Solidity: event TextChanged(bytes32 indexed node, string indexed indexedKey, string key)
func (_Resolver *ResolverFilterer) WatchTextChanged(opts *bind.WatchOpts, sink chan<- *ResolverTextChanged, node [][32]byte, indexedKey []string) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _Resolver.contract.WatchLogs(opts, "TextChanged", nodeRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverTextChanged)
				if err := _Resolver.contract.UnpackLog(event, "TextChanged", log); err != nil {
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

// ParseTextChanged is a log parse operation binding the contract event 0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550.
//
// Solidity: event TextChanged(bytes32 indexed node, string indexed indexedKey, string key)
func (_Resolver *ResolverFilterer) ParseTextChanged(log types.Log) (*ResolverTextChanged, error) {
	event := new(ResolverTextChanged)
	if err := _Resolver.contract.UnpackLog(event, "TextChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
