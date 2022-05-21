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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"rewarded\",\"type\":\"bool\"}],\"name\":\"Pray\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"god\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"godCollect\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pray\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60806040526000341161001157600080fd5b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055504260018190555061046b806100676000396000f3fe60806040526004361061003f5760003560e01c8063419905e5146100445780634b1d73e11461004e578063b69ef8a814610065578063ba56f6ee14610090575b600080fd5b61004c6100bb565b005b34801561005a57600080fd5b50610063610190565b005b34801561007157600080fd5b5061007a610221565b6040516100879190610266565b60405180910390f35b34801561009c57600080fd5b506100a5610229565b6040516100b291906102c2565b60405180910390f35b6000600a42336001546040516020016100d693929190610346565b6040516020818303038152906040528051906020012060001c6100f991906103b2565b905060006002821090508015610151573373ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f1935050505015801561014f573d6000803e3d6000fd5b505b7fe31d5a2b10d83bbf72c38efb28632d3d815ddd94d4c35793cace01c819300ece333483604051610184939291906103fe565b60405180910390a15050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101e857600080fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b600047905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000819050919050565b6102608161024d565b82525050565b600060208201905061027b6000830184610257565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006102ac82610281565b9050919050565b6102bc816102a1565b82525050565b60006020820190506102d760008301846102b3565b92915050565b6000819050919050565b6102f86102f38261024d565b6102dd565b82525050565b60008160601b9050919050565b6000610316826102fe565b9050919050565b60006103288261030b565b9050919050565b61034061033b826102a1565b61031d565b82525050565b600061035282866102e7565b602082019150610362828561032f565b60148201915061037282846102e7565b602082019150819050949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006103bd8261024d565b91506103c88361024d565b9250826103d8576103d7610383565b5b828206905092915050565b60008115159050919050565b6103f8816103e3565b82525050565b600060608201905061041360008301866102b3565b6104206020830185610257565b61042d60408301846103ef565b94935050505056fea2646970667358221220b979d22559365e5d49c36368260a056a5cc75e0ac48b01d9bb4e5086cc63324a64736f6c634300080d0033",
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

// MoneyBoxPrayIterator is returned from FilterPray and is used to iterate over the raw logs and unpacked data for Pray events raised by the MoneyBox contract.
type MoneyBoxPrayIterator struct {
	Event *MoneyBoxPray // Event containing the contract specifics and raw log

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
func (it *MoneyBoxPrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoneyBoxPray)
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
		it.Event = new(MoneyBoxPray)
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
func (it *MoneyBoxPrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoneyBoxPrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoneyBoxPray represents a Pray event raised by the MoneyBox contract.
type MoneyBoxPray struct {
	Prayer   common.Address
	Amount   *big.Int
	Rewarded bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPray is a free log retrieval operation binding the contract event 0xe31d5a2b10d83bbf72c38efb28632d3d815ddd94d4c35793cace01c819300ece.
//
// Solidity: event Pray(address prayer, uint256 amount, bool rewarded)
func (_MoneyBox *MoneyBoxFilterer) FilterPray(opts *bind.FilterOpts) (*MoneyBoxPrayIterator, error) {

	logs, sub, err := _MoneyBox.contract.FilterLogs(opts, "Pray")
	if err != nil {
		return nil, err
	}
	return &MoneyBoxPrayIterator{contract: _MoneyBox.contract, event: "Pray", logs: logs, sub: sub}, nil
}

// WatchPray is a free log subscription operation binding the contract event 0xe31d5a2b10d83bbf72c38efb28632d3d815ddd94d4c35793cace01c819300ece.
//
// Solidity: event Pray(address prayer, uint256 amount, bool rewarded)
func (_MoneyBox *MoneyBoxFilterer) WatchPray(opts *bind.WatchOpts, sink chan<- *MoneyBoxPray) (event.Subscription, error) {

	logs, sub, err := _MoneyBox.contract.WatchLogs(opts, "Pray")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoneyBoxPray)
				if err := _MoneyBox.contract.UnpackLog(event, "Pray", log); err != nil {
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

// ParsePray is a log parse operation binding the contract event 0xe31d5a2b10d83bbf72c38efb28632d3d815ddd94d4c35793cace01c819300ece.
//
// Solidity: event Pray(address prayer, uint256 amount, bool rewarded)
func (_MoneyBox *MoneyBoxFilterer) ParsePray(log types.Log) (*MoneyBoxPray, error) {
	event := new(MoneyBoxPray)
	if err := _MoneyBox.contract.UnpackLog(event, "Pray", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
