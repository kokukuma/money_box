// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package money_box

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

// MoneyBoxMetaData contains all meta data concerning the MoneyBox contract.
var MoneyBoxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"god\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"godCollect\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pray\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60806040526000341161001157600080fd5b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550426001819055506103d8806100676000396000f3fe60806040526004361061003f5760003560e01c8063419905e5146100445780634b1d73e11461004e578063b69ef8a814610065578063ba56f6ee14610090575b600080fd5b61004c6100bb565b005b34801561005a57600080fd5b5061006361014f565b005b34801561007157600080fd5b5061007a6101e0565b6040516100879190610225565b60405180910390f35b34801561009c57600080fd5b506100a56101e8565b6040516100b29190610281565b60405180910390f35b6000600a42336001546040516020016100d693929190610305565b6040516020818303038152906040528051906020012060001c6100f99190610371565b9050600281101561014c573373ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f1935050505015801561014a573d6000803e3d6000fd5b505b50565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101a757600080fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b600047905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000819050919050565b61021f8161020c565b82525050565b600060208201905061023a6000830184610216565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061026b82610240565b9050919050565b61027b81610260565b82525050565b60006020820190506102966000830184610272565b92915050565b6000819050919050565b6102b76102b28261020c565b61029c565b82525050565b60008160601b9050919050565b60006102d5826102bd565b9050919050565b60006102e7826102ca565b9050919050565b6102ff6102fa82610260565b6102dc565b82525050565b600061031182866102a6565b60208201915061032182856102ee565b60148201915061033182846102a6565b602082019150819050949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061037c8261020c565b91506103878361020c565b92508261039757610396610342565b5b82820690509291505056fea2646970667358221220e6c3bb58bf206f89184ac592e0b46e7cacc9ed57ffbd2ffad3faa402f21b939264736f6c634300080d0033",
}

// MoneyBoxABI is the input ABI used to generate the binding from.
// Deprecated: Use MoneyBoxMetaData.ABI instead.
var MoneyBoxABI = MoneyBoxMetaData.ABI

// MoneyBoxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MoneyBoxMetaData.Bin instead.
var MoneyBoxBin = MoneyBoxMetaData.Bin

// DeployMoneyBox deploys a new Ethereum contract, binding an instance of MoneyBox to it.
func DeployMoneyBox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MoneyBox, error) {
	parsed, err := MoneyBoxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MoneyBoxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MoneyBox{MoneyBoxCaller: MoneyBoxCaller{contract: contract}, MoneyBoxTransactor: MoneyBoxTransactor{contract: contract}, MoneyBoxFilterer: MoneyBoxFilterer{contract: contract}}, nil
}

// MoneyBox is an auto generated Go binding around an Ethereum contract.
type MoneyBox struct {
	MoneyBoxCaller     // Read-only binding to the contract
	MoneyBoxTransactor // Write-only binding to the contract
	MoneyBoxFilterer   // Log filterer for contract events
}

// MoneyBoxCaller is an auto generated read-only Go binding around an Ethereum contract.
type MoneyBoxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MoneyBoxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MoneyBoxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MoneyBoxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MoneyBoxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MoneyBoxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MoneyBoxSession struct {
	Contract     *MoneyBox         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MoneyBoxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MoneyBoxCallerSession struct {
	Contract *MoneyBoxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MoneyBoxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MoneyBoxTransactorSession struct {
	Contract     *MoneyBoxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MoneyBoxRaw is an auto generated low-level Go binding around an Ethereum contract.
type MoneyBoxRaw struct {
	Contract *MoneyBox // Generic contract binding to access the raw methods on
}

// MoneyBoxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MoneyBoxCallerRaw struct {
	Contract *MoneyBoxCaller // Generic read-only contract binding to access the raw methods on
}

// MoneyBoxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MoneyBoxTransactorRaw struct {
	Contract *MoneyBoxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMoneyBox creates a new instance of MoneyBox, bound to a specific deployed contract.
func NewMoneyBox(address common.Address, backend bind.ContractBackend) (*MoneyBox, error) {
	contract, err := bindMoneyBox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MoneyBox{MoneyBoxCaller: MoneyBoxCaller{contract: contract}, MoneyBoxTransactor: MoneyBoxTransactor{contract: contract}, MoneyBoxFilterer: MoneyBoxFilterer{contract: contract}}, nil
}

// NewMoneyBoxCaller creates a new read-only instance of MoneyBox, bound to a specific deployed contract.
func NewMoneyBoxCaller(address common.Address, caller bind.ContractCaller) (*MoneyBoxCaller, error) {
	contract, err := bindMoneyBox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MoneyBoxCaller{contract: contract}, nil
}

// NewMoneyBoxTransactor creates a new write-only instance of MoneyBox, bound to a specific deployed contract.
func NewMoneyBoxTransactor(address common.Address, transactor bind.ContractTransactor) (*MoneyBoxTransactor, error) {
	contract, err := bindMoneyBox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MoneyBoxTransactor{contract: contract}, nil
}

// NewMoneyBoxFilterer creates a new log filterer instance of MoneyBox, bound to a specific deployed contract.
func NewMoneyBoxFilterer(address common.Address, filterer bind.ContractFilterer) (*MoneyBoxFilterer, error) {
	contract, err := bindMoneyBox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MoneyBoxFilterer{contract: contract}, nil
}

// bindMoneyBox binds a generic wrapper to an already deployed contract.
func bindMoneyBox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MoneyBoxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MoneyBox *MoneyBoxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MoneyBox.Contract.MoneyBoxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MoneyBox *MoneyBoxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MoneyBox.Contract.MoneyBoxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MoneyBox *MoneyBoxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MoneyBox.Contract.MoneyBoxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MoneyBox *MoneyBoxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MoneyBox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MoneyBox *MoneyBoxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MoneyBox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MoneyBox *MoneyBoxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MoneyBox.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_MoneyBox *MoneyBoxCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MoneyBox.contract.Call(opts, &out, "balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_MoneyBox *MoneyBoxSession) Balance() (*big.Int, error) {
	return _MoneyBox.Contract.Balance(&_MoneyBox.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_MoneyBox *MoneyBoxCallerSession) Balance() (*big.Int, error) {
	return _MoneyBox.Contract.Balance(&_MoneyBox.CallOpts)
}

// God is a free data retrieval call binding the contract method 0xba56f6ee.
//
// Solidity: function god() view returns(address)
func (_MoneyBox *MoneyBoxCaller) God(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MoneyBox.contract.Call(opts, &out, "god")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// God is a free data retrieval call binding the contract method 0xba56f6ee.
//
// Solidity: function god() view returns(address)
func (_MoneyBox *MoneyBoxSession) God() (common.Address, error) {
	return _MoneyBox.Contract.God(&_MoneyBox.CallOpts)
}

// God is a free data retrieval call binding the contract method 0xba56f6ee.
//
// Solidity: function god() view returns(address)
func (_MoneyBox *MoneyBoxCallerSession) God() (common.Address, error) {
	return _MoneyBox.Contract.God(&_MoneyBox.CallOpts)
}

// GodCollect is a paid mutator transaction binding the contract method 0x4b1d73e1.
//
// Solidity: function godCollect() returns()
func (_MoneyBox *MoneyBoxTransactor) GodCollect(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MoneyBox.contract.Transact(opts, "godCollect")
}

// GodCollect is a paid mutator transaction binding the contract method 0x4b1d73e1.
//
// Solidity: function godCollect() returns()
func (_MoneyBox *MoneyBoxSession) GodCollect() (*types.Transaction, error) {
	return _MoneyBox.Contract.GodCollect(&_MoneyBox.TransactOpts)
}

// GodCollect is a paid mutator transaction binding the contract method 0x4b1d73e1.
//
// Solidity: function godCollect() returns()
func (_MoneyBox *MoneyBoxTransactorSession) GodCollect() (*types.Transaction, error) {
	return _MoneyBox.Contract.GodCollect(&_MoneyBox.TransactOpts)
}

// Pray is a paid mutator transaction binding the contract method 0x419905e5.
//
// Solidity: function pray() payable returns()
func (_MoneyBox *MoneyBoxTransactor) Pray(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MoneyBox.contract.Transact(opts, "pray")
}

// Pray is a paid mutator transaction binding the contract method 0x419905e5.
//
// Solidity: function pray() payable returns()
func (_MoneyBox *MoneyBoxSession) Pray() (*types.Transaction, error) {
	return _MoneyBox.Contract.Pray(&_MoneyBox.TransactOpts)
}

// Pray is a paid mutator transaction binding the contract method 0x419905e5.
//
// Solidity: function pray() payable returns()
func (_MoneyBox *MoneyBoxTransactorSession) Pray() (*types.Transaction, error) {
	return _MoneyBox.Contract.Pray(&_MoneyBox.TransactOpts)
}
