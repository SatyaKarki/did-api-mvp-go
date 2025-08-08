// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

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

// EthMetaData contains all meta data concerning the Eth contract.
var EthMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"document\",\"type\":\"string\"}],\"name\":\"DIDDocumentSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getDIDDocument\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"document\",\"type\":\"string\"}],\"name\":\"setDIDDocument\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506107b38061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c80635f87e23c14610038578063784afc4e14610054575b5f5ffd5b610052600480360381019061004d91906102e6565b610084565b005b61006e6004803603810190610069919061035c565b6100ec565b60405161007b9190610403565b60405180910390f35b805f83604051610094919061045d565b908152602001604051809103902090816100ae9190610679565b507fea0c4f36b3b964cb11bef6a9a3a2cb19ccb99c71d51afa83ca8d5f4b0fb0271282826040516100e0929190610748565b60405180910390a15050565b60605f826040516100fd919061045d565b90815260200160405180910390208054610116906104a0565b80601f0160208091040260200160405190810160405280929190818152602001828054610142906104a0565b801561018d5780601f106101645761010080835404028352916020019161018d565b820191905f5260205f20905b81548152906001019060200180831161017057829003601f168201915b50505050509050919050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6101f8826101b2565b810181811067ffffffffffffffff82111715610217576102166101c2565b5b80604052505050565b5f610229610199565b905061023582826101ef565b919050565b5f67ffffffffffffffff821115610254576102536101c2565b5b61025d826101b2565b9050602081019050919050565b828183375f83830152505050565b5f61028a6102858461023a565b610220565b9050828152602081018484840111156102a6576102a56101ae565b5b6102b184828561026a565b509392505050565b5f82601f8301126102cd576102cc6101aa565b5b81356102dd848260208601610278565b91505092915050565b5f5f604083850312156102fc576102fb6101a2565b5b5f83013567ffffffffffffffff811115610319576103186101a6565b5b610325858286016102b9565b925050602083013567ffffffffffffffff811115610346576103456101a6565b5b610352858286016102b9565b9150509250929050565b5f60208284031215610371576103706101a2565b5b5f82013567ffffffffffffffff81111561038e5761038d6101a6565b5b61039a848285016102b9565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f6103d5826103a3565b6103df81856103ad565b93506103ef8185602086016103bd565b6103f8816101b2565b840191505092915050565b5f6020820190508181035f83015261041b81846103cb565b905092915050565b5f81905092915050565b5f610437826103a3565b6104418185610423565b93506104518185602086016103bd565b80840191505092915050565b5f610468828461042d565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806104b757607f821691505b6020821081036104ca576104c9610473565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261052c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826104f1565b61053686836104f1565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61057a6105756105708461054e565b610557565b61054e565b9050919050565b5f819050919050565b61059383610560565b6105a761059f82610581565b8484546104fd565b825550505050565b5f5f905090565b6105be6105af565b6105c981848461058a565b505050565b5b818110156105ec576105e15f826105b6565b6001810190506105cf565b5050565b601f82111561063157610602816104d0565b61060b846104e2565b8101602085101561061a578190505b61062e610626856104e2565b8301826105ce565b50505b505050565b5f82821c905092915050565b5f6106515f1984600802610636565b1980831691505092915050565b5f6106698383610642565b9150826002028217905092915050565b610682826103a3565b67ffffffffffffffff81111561069b5761069a6101c2565b5b6106a582546104a0565b6106b08282856105f0565b5f60209050601f8311600181146106e1575f84156106cf578287015190505b6106d9858261065e565b865550610740565b601f1984166106ef866104d0565b5f5b82811015610716578489015182556001820191506020850194506020810190506106f1565b86831015610733578489015161072f601f891682610642565b8355505b6001600288020188555050505b505050505050565b5f6040820190508181035f83015261076081856103cb565b9050818103602083015261077481846103cb565b9050939250505056fea26469706673582212207b74ce29717d3354e0a3ebe64405e7e28a8a1ed8d9073bb5c0cdecb886be18ce64736f6c634300081e0033",
}

// EthABI is the input ABI used to generate the binding from.
// Deprecated: Use EthMetaData.ABI instead.
var EthABI = EthMetaData.ABI

// EthBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EthMetaData.Bin instead.
var EthBin = EthMetaData.Bin

// DeployEth deploys a new Ethereum contract, binding an instance of Eth to it.
func DeployEth(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Eth, error) {
	parsed, err := EthMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EthBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Eth{EthCaller: EthCaller{contract: contract}, EthTransactor: EthTransactor{contract: contract}, EthFilterer: EthFilterer{contract: contract}}, nil
}

// Eth is an auto generated Go binding around an Ethereum contract.
type Eth struct {
	EthCaller     // Read-only binding to the contract
	EthTransactor // Write-only binding to the contract
	EthFilterer   // Log filterer for contract events
}

// EthCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthSession struct {
	Contract     *Eth              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthCallerSession struct {
	Contract *EthCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthTransactorSession struct {
	Contract     *EthTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthRaw struct {
	Contract *Eth // Generic contract binding to access the raw methods on
}

// EthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthCallerRaw struct {
	Contract *EthCaller // Generic read-only contract binding to access the raw methods on
}

// EthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthTransactorRaw struct {
	Contract *EthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEth creates a new instance of Eth, bound to a specific deployed contract.
func NewEth(address common.Address, backend bind.ContractBackend) (*Eth, error) {
	contract, err := bindEth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Eth{EthCaller: EthCaller{contract: contract}, EthTransactor: EthTransactor{contract: contract}, EthFilterer: EthFilterer{contract: contract}}, nil
}

// NewEthCaller creates a new read-only instance of Eth, bound to a specific deployed contract.
func NewEthCaller(address common.Address, caller bind.ContractCaller) (*EthCaller, error) {
	contract, err := bindEth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthCaller{contract: contract}, nil
}

// NewEthTransactor creates a new write-only instance of Eth, bound to a specific deployed contract.
func NewEthTransactor(address common.Address, transactor bind.ContractTransactor) (*EthTransactor, error) {
	contract, err := bindEth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthTransactor{contract: contract}, nil
}

// NewEthFilterer creates a new log filterer instance of Eth, bound to a specific deployed contract.
func NewEthFilterer(address common.Address, filterer bind.ContractFilterer) (*EthFilterer, error) {
	contract, err := bindEth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthFilterer{contract: contract}, nil
}

// bindEth binds a generic wrapper to an already deployed contract.
func bindEth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EthMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eth *EthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Eth.Contract.EthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eth *EthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eth.Contract.EthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eth *EthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eth.Contract.EthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eth *EthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Eth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eth *EthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eth *EthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eth.Contract.contract.Transact(opts, method, params...)
}

// GetDIDDocument is a free data retrieval call binding the contract method 0x784afc4e.
//
// Solidity: function getDIDDocument(string did) view returns(string)
func (_Eth *EthCaller) GetDIDDocument(opts *bind.CallOpts, did string) (string, error) {
	var out []interface{}
	err := _Eth.contract.Call(opts, &out, "getDIDDocument", did)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetDIDDocument is a free data retrieval call binding the contract method 0x784afc4e.
//
// Solidity: function getDIDDocument(string did) view returns(string)
func (_Eth *EthSession) GetDIDDocument(did string) (string, error) {
	return _Eth.Contract.GetDIDDocument(&_Eth.CallOpts, did)
}

// GetDIDDocument is a free data retrieval call binding the contract method 0x784afc4e.
//
// Solidity: function getDIDDocument(string did) view returns(string)
func (_Eth *EthCallerSession) GetDIDDocument(did string) (string, error) {
	return _Eth.Contract.GetDIDDocument(&_Eth.CallOpts, did)
}

// SetDIDDocument is a paid mutator transaction binding the contract method 0x5f87e23c.
//
// Solidity: function setDIDDocument(string did, string document) returns()
func (_Eth *EthTransactor) SetDIDDocument(opts *bind.TransactOpts, did string, document string) (*types.Transaction, error) {
	return _Eth.contract.Transact(opts, "setDIDDocument", did, document)
}

// SetDIDDocument is a paid mutator transaction binding the contract method 0x5f87e23c.
//
// Solidity: function setDIDDocument(string did, string document) returns()
func (_Eth *EthSession) SetDIDDocument(did string, document string) (*types.Transaction, error) {
	return _Eth.Contract.SetDIDDocument(&_Eth.TransactOpts, did, document)
}

// SetDIDDocument is a paid mutator transaction binding the contract method 0x5f87e23c.
//
// Solidity: function setDIDDocument(string did, string document) returns()
func (_Eth *EthTransactorSession) SetDIDDocument(did string, document string) (*types.Transaction, error) {
	return _Eth.Contract.SetDIDDocument(&_Eth.TransactOpts, did, document)
}

// EthDIDDocumentSetIterator is returned from FilterDIDDocumentSet and is used to iterate over the raw logs and unpacked data for DIDDocumentSet events raised by the Eth contract.
type EthDIDDocumentSetIterator struct {
	Event *EthDIDDocumentSet // Event containing the contract specifics and raw log

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
func (it *EthDIDDocumentSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthDIDDocumentSet)
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
		it.Event = new(EthDIDDocumentSet)
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
func (it *EthDIDDocumentSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthDIDDocumentSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthDIDDocumentSet represents a DIDDocumentSet event raised by the Eth contract.
type EthDIDDocumentSet struct {
	Did      string
	Document string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDIDDocumentSet is a free log retrieval operation binding the contract event 0xea0c4f36b3b964cb11bef6a9a3a2cb19ccb99c71d51afa83ca8d5f4b0fb02712.
//
// Solidity: event DIDDocumentSet(string did, string document)
func (_Eth *EthFilterer) FilterDIDDocumentSet(opts *bind.FilterOpts) (*EthDIDDocumentSetIterator, error) {

	logs, sub, err := _Eth.contract.FilterLogs(opts, "DIDDocumentSet")
	if err != nil {
		return nil, err
	}
	return &EthDIDDocumentSetIterator{contract: _Eth.contract, event: "DIDDocumentSet", logs: logs, sub: sub}, nil
}

// WatchDIDDocumentSet is a free log subscription operation binding the contract event 0xea0c4f36b3b964cb11bef6a9a3a2cb19ccb99c71d51afa83ca8d5f4b0fb02712.
//
// Solidity: event DIDDocumentSet(string did, string document)
func (_Eth *EthFilterer) WatchDIDDocumentSet(opts *bind.WatchOpts, sink chan<- *EthDIDDocumentSet) (event.Subscription, error) {

	logs, sub, err := _Eth.contract.WatchLogs(opts, "DIDDocumentSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthDIDDocumentSet)
				if err := _Eth.contract.UnpackLog(event, "DIDDocumentSet", log); err != nil {
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

// ParseDIDDocumentSet is a log parse operation binding the contract event 0xea0c4f36b3b964cb11bef6a9a3a2cb19ccb99c71d51afa83ca8d5f4b0fb02712.
//
// Solidity: event DIDDocumentSet(string did, string document)
func (_Eth *EthFilterer) ParseDIDDocumentSet(log types.Log) (*EthDIDDocumentSet, error) {
	event := new(EthDIDDocumentSet)
	if err := _Eth.contract.UnpackLog(event, "DIDDocumentSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
