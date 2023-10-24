// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teleporterregistry

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = interfaces.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ProtocolRegistryEntry is an auto generated low-level Go binding around an user-defined struct.
type ProtocolRegistryEntry struct {
	Version         *big.Int
	ProtocolAddress common.Address
}

// TeleporterregistryMetaData contains all meta data concerning the Teleporterregistry contract.
var TeleporterregistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"protocolAddress\",\"type\":\"address\"}],\"internalType\":\"structProtocolRegistryEntry[]\",\"name\":\"initialEntries\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidDestinationAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDestinationChainID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOriginSenderAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProtocolAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProtocolVersion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSourceChainID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWarpMessage\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"protocolAddress\",\"type\":\"address\"}],\"name\":\"AddProtocolVersion\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VALIDATORS_SOURCE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WARP_MESSENGER\",\"outputs\":[{\"internalType\":\"contractWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"addProtocolVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"getAddressFromVersion\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestTeleporter\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"getTeleporterFromVersion\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"protocolAddress\",\"type\":\"address\"}],\"name\":\"getVersionFromAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TeleporterregistryABI is the input ABI used to generate the binding from.
// Deprecated: Use TeleporterregistryMetaData.ABI instead.
var TeleporterregistryABI = TeleporterregistryMetaData.ABI

// Teleporterregistry is an auto generated Go binding around an Ethereum contract.
type Teleporterregistry struct {
	TeleporterregistryCaller     // Read-only binding to the contract
	TeleporterregistryTransactor // Write-only binding to the contract
	TeleporterregistryFilterer   // Log filterer for contract events
}

// TeleporterregistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeleporterregistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterregistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeleporterregistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterregistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeleporterregistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterregistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeleporterregistrySession struct {
	Contract     *Teleporterregistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TeleporterregistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeleporterregistryCallerSession struct {
	Contract *TeleporterregistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// TeleporterregistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeleporterregistryTransactorSession struct {
	Contract     *TeleporterregistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// TeleporterregistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeleporterregistryRaw struct {
	Contract *Teleporterregistry // Generic contract binding to access the raw methods on
}

// TeleporterregistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeleporterregistryCallerRaw struct {
	Contract *TeleporterregistryCaller // Generic read-only contract binding to access the raw methods on
}

// TeleporterregistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeleporterregistryTransactorRaw struct {
	Contract *TeleporterregistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeleporterregistry creates a new instance of Teleporterregistry, bound to a specific deployed contract.
func NewTeleporterregistry(address common.Address, backend bind.ContractBackend) (*Teleporterregistry, error) {
	contract, err := bindTeleporterregistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Teleporterregistry{TeleporterregistryCaller: TeleporterregistryCaller{contract: contract}, TeleporterregistryTransactor: TeleporterregistryTransactor{contract: contract}, TeleporterregistryFilterer: TeleporterregistryFilterer{contract: contract}}, nil
}

// NewTeleporterregistryCaller creates a new read-only instance of Teleporterregistry, bound to a specific deployed contract.
func NewTeleporterregistryCaller(address common.Address, caller bind.ContractCaller) (*TeleporterregistryCaller, error) {
	contract, err := bindTeleporterregistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeleporterregistryCaller{contract: contract}, nil
}

// NewTeleporterregistryTransactor creates a new write-only instance of Teleporterregistry, bound to a specific deployed contract.
func NewTeleporterregistryTransactor(address common.Address, transactor bind.ContractTransactor) (*TeleporterregistryTransactor, error) {
	contract, err := bindTeleporterregistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeleporterregistryTransactor{contract: contract}, nil
}

// NewTeleporterregistryFilterer creates a new log filterer instance of Teleporterregistry, bound to a specific deployed contract.
func NewTeleporterregistryFilterer(address common.Address, filterer bind.ContractFilterer) (*TeleporterregistryFilterer, error) {
	contract, err := bindTeleporterregistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeleporterregistryFilterer{contract: contract}, nil
}

// bindTeleporterregistry binds a generic wrapper to an already deployed contract.
func bindTeleporterregistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeleporterregistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Teleporterregistry *TeleporterregistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Teleporterregistry.Contract.TeleporterregistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Teleporterregistry *TeleporterregistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Teleporterregistry.Contract.TeleporterregistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Teleporterregistry *TeleporterregistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Teleporterregistry.Contract.TeleporterregistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Teleporterregistry *TeleporterregistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Teleporterregistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Teleporterregistry *TeleporterregistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Teleporterregistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Teleporterregistry *TeleporterregistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Teleporterregistry.Contract.contract.Transact(opts, method, params...)
}

// VALIDATORSSOURCEADDRESS is a free data retrieval call binding the contract method 0x0731775d.
//
// Solidity: function VALIDATORS_SOURCE_ADDRESS() view returns(address)
func (_Teleporterregistry *TeleporterregistryCaller) VALIDATORSSOURCEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Teleporterregistry.contract.Call(opts, &out, "VALIDATORS_SOURCE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VALIDATORSSOURCEADDRESS is a free data retrieval call binding the contract method 0x0731775d.
//
// Solidity: function VALIDATORS_SOURCE_ADDRESS() view returns(address)
func (_Teleporterregistry *TeleporterregistrySession) VALIDATORSSOURCEADDRESS() (common.Address, error) {
	return _Teleporterregistry.Contract.VALIDATORSSOURCEADDRESS(&_Teleporterregistry.CallOpts)
}

// VALIDATORSSOURCEADDRESS is a free data retrieval call binding the contract method 0x0731775d.
//
// Solidity: function VALIDATORS_SOURCE_ADDRESS() view returns(address)
func (_Teleporterregistry *TeleporterregistryCallerSession) VALIDATORSSOURCEADDRESS() (common.Address, error) {
	return _Teleporterregistry.Contract.VALIDATORSSOURCEADDRESS(&_Teleporterregistry.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_Teleporterregistry *TeleporterregistryCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Teleporterregistry.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_Teleporterregistry *TeleporterregistrySession) WARPMESSENGER() (common.Address, error) {
	return _Teleporterregistry.Contract.WARPMESSENGER(&_Teleporterregistry.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_Teleporterregistry *TeleporterregistryCallerSession) WARPMESSENGER() (common.Address, error) {
	return _Teleporterregistry.Contract.WARPMESSENGER(&_Teleporterregistry.CallOpts)
}

// GetAddressFromVersion is a free data retrieval call binding the contract method 0x46f9ef49.
//
// Solidity: function getAddressFromVersion(uint256 version) view returns(address)
func (_Teleporterregistry *TeleporterregistryCaller) GetAddressFromVersion(opts *bind.CallOpts, version *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Teleporterregistry.contract.Call(opts, &out, "getAddressFromVersion", version)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressFromVersion is a free data retrieval call binding the contract method 0x46f9ef49.
//
// Solidity: function getAddressFromVersion(uint256 version) view returns(address)
func (_Teleporterregistry *TeleporterregistrySession) GetAddressFromVersion(version *big.Int) (common.Address, error) {
	return _Teleporterregistry.Contract.GetAddressFromVersion(&_Teleporterregistry.CallOpts, version)
}

// GetAddressFromVersion is a free data retrieval call binding the contract method 0x46f9ef49.
//
// Solidity: function getAddressFromVersion(uint256 version) view returns(address)
func (_Teleporterregistry *TeleporterregistryCallerSession) GetAddressFromVersion(version *big.Int) (common.Address, error) {
	return _Teleporterregistry.Contract.GetAddressFromVersion(&_Teleporterregistry.CallOpts, version)
}

// GetLatestTeleporter is a free data retrieval call binding the contract method 0xd820e64f.
//
// Solidity: function getLatestTeleporter() view returns(address)
func (_Teleporterregistry *TeleporterregistryCaller) GetLatestTeleporter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Teleporterregistry.contract.Call(opts, &out, "getLatestTeleporter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLatestTeleporter is a free data retrieval call binding the contract method 0xd820e64f.
//
// Solidity: function getLatestTeleporter() view returns(address)
func (_Teleporterregistry *TeleporterregistrySession) GetLatestTeleporter() (common.Address, error) {
	return _Teleporterregistry.Contract.GetLatestTeleporter(&_Teleporterregistry.CallOpts)
}

// GetLatestTeleporter is a free data retrieval call binding the contract method 0xd820e64f.
//
// Solidity: function getLatestTeleporter() view returns(address)
func (_Teleporterregistry *TeleporterregistryCallerSession) GetLatestTeleporter() (common.Address, error) {
	return _Teleporterregistry.Contract.GetLatestTeleporter(&_Teleporterregistry.CallOpts)
}

// GetLatestVersion is a free data retrieval call binding the contract method 0x0e6d1de9.
//
// Solidity: function getLatestVersion() view returns(uint256)
func (_Teleporterregistry *TeleporterregistryCaller) GetLatestVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Teleporterregistry.contract.Call(opts, &out, "getLatestVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestVersion is a free data retrieval call binding the contract method 0x0e6d1de9.
//
// Solidity: function getLatestVersion() view returns(uint256)
func (_Teleporterregistry *TeleporterregistrySession) GetLatestVersion() (*big.Int, error) {
	return _Teleporterregistry.Contract.GetLatestVersion(&_Teleporterregistry.CallOpts)
}

// GetLatestVersion is a free data retrieval call binding the contract method 0x0e6d1de9.
//
// Solidity: function getLatestVersion() view returns(uint256)
func (_Teleporterregistry *TeleporterregistryCallerSession) GetLatestVersion() (*big.Int, error) {
	return _Teleporterregistry.Contract.GetLatestVersion(&_Teleporterregistry.CallOpts)
}

// GetTeleporterFromVersion is a free data retrieval call binding the contract method 0x215abce9.
//
// Solidity: function getTeleporterFromVersion(uint256 version) view returns(address)
func (_Teleporterregistry *TeleporterregistryCaller) GetTeleporterFromVersion(opts *bind.CallOpts, version *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Teleporterregistry.contract.Call(opts, &out, "getTeleporterFromVersion", version)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTeleporterFromVersion is a free data retrieval call binding the contract method 0x215abce9.
//
// Solidity: function getTeleporterFromVersion(uint256 version) view returns(address)
func (_Teleporterregistry *TeleporterregistrySession) GetTeleporterFromVersion(version *big.Int) (common.Address, error) {
	return _Teleporterregistry.Contract.GetTeleporterFromVersion(&_Teleporterregistry.CallOpts, version)
}

// GetTeleporterFromVersion is a free data retrieval call binding the contract method 0x215abce9.
//
// Solidity: function getTeleporterFromVersion(uint256 version) view returns(address)
func (_Teleporterregistry *TeleporterregistryCallerSession) GetTeleporterFromVersion(version *big.Int) (common.Address, error) {
	return _Teleporterregistry.Contract.GetTeleporterFromVersion(&_Teleporterregistry.CallOpts, version)
}

// GetVersionFromAddress is a free data retrieval call binding the contract method 0x4c1f08ce.
//
// Solidity: function getVersionFromAddress(address protocolAddress) view returns(uint256)
func (_Teleporterregistry *TeleporterregistryCaller) GetVersionFromAddress(opts *bind.CallOpts, protocolAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Teleporterregistry.contract.Call(opts, &out, "getVersionFromAddress", protocolAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVersionFromAddress is a free data retrieval call binding the contract method 0x4c1f08ce.
//
// Solidity: function getVersionFromAddress(address protocolAddress) view returns(uint256)
func (_Teleporterregistry *TeleporterregistrySession) GetVersionFromAddress(protocolAddress common.Address) (*big.Int, error) {
	return _Teleporterregistry.Contract.GetVersionFromAddress(&_Teleporterregistry.CallOpts, protocolAddress)
}

// GetVersionFromAddress is a free data retrieval call binding the contract method 0x4c1f08ce.
//
// Solidity: function getVersionFromAddress(address protocolAddress) view returns(uint256)
func (_Teleporterregistry *TeleporterregistryCallerSession) GetVersionFromAddress(protocolAddress common.Address) (*big.Int, error) {
	return _Teleporterregistry.Contract.GetVersionFromAddress(&_Teleporterregistry.CallOpts, protocolAddress)
}

// AddProtocolVersion is a paid mutator transaction binding the contract method 0x41f34ed9.
//
// Solidity: function addProtocolVersion(uint32 messageIndex) returns()
func (_Teleporterregistry *TeleporterregistryTransactor) AddProtocolVersion(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _Teleporterregistry.contract.Transact(opts, "addProtocolVersion", messageIndex)
}

// AddProtocolVersion is a paid mutator transaction binding the contract method 0x41f34ed9.
//
// Solidity: function addProtocolVersion(uint32 messageIndex) returns()
func (_Teleporterregistry *TeleporterregistrySession) AddProtocolVersion(messageIndex uint32) (*types.Transaction, error) {
	return _Teleporterregistry.Contract.AddProtocolVersion(&_Teleporterregistry.TransactOpts, messageIndex)
}

// AddProtocolVersion is a paid mutator transaction binding the contract method 0x41f34ed9.
//
// Solidity: function addProtocolVersion(uint32 messageIndex) returns()
func (_Teleporterregistry *TeleporterregistryTransactorSession) AddProtocolVersion(messageIndex uint32) (*types.Transaction, error) {
	return _Teleporterregistry.Contract.AddProtocolVersion(&_Teleporterregistry.TransactOpts, messageIndex)
}

// TeleporterregistryAddProtocolVersionIterator is returned from FilterAddProtocolVersion and is used to iterate over the raw logs and unpacked data for AddProtocolVersion events raised by the Teleporterregistry contract.
type TeleporterregistryAddProtocolVersionIterator struct {
	Event *TeleporterregistryAddProtocolVersion // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeleporterregistryAddProtocolVersionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterregistryAddProtocolVersion)
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
		it.Event = new(TeleporterregistryAddProtocolVersion)
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
func (it *TeleporterregistryAddProtocolVersionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterregistryAddProtocolVersionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterregistryAddProtocolVersion represents a AddProtocolVersion event raised by the Teleporterregistry contract.
type TeleporterregistryAddProtocolVersion struct {
	Version         *big.Int
	ProtocolAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAddProtocolVersion is a free log retrieval operation binding the contract event 0xa5eed93d951a9603d5f7c0a57de79a299dd3dbd5e51429be209d8053a42ab43a.
//
// Solidity: event AddProtocolVersion(uint256 indexed version, address indexed protocolAddress)
func (_Teleporterregistry *TeleporterregistryFilterer) FilterAddProtocolVersion(opts *bind.FilterOpts, version []*big.Int, protocolAddress []common.Address) (*TeleporterregistryAddProtocolVersionIterator, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var protocolAddressRule []interface{}
	for _, protocolAddressItem := range protocolAddress {
		protocolAddressRule = append(protocolAddressRule, protocolAddressItem)
	}

	logs, sub, err := _Teleporterregistry.contract.FilterLogs(opts, "AddProtocolVersion", versionRule, protocolAddressRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterregistryAddProtocolVersionIterator{contract: _Teleporterregistry.contract, event: "AddProtocolVersion", logs: logs, sub: sub}, nil
}

// WatchAddProtocolVersion is a free log subscription operation binding the contract event 0xa5eed93d951a9603d5f7c0a57de79a299dd3dbd5e51429be209d8053a42ab43a.
//
// Solidity: event AddProtocolVersion(uint256 indexed version, address indexed protocolAddress)
func (_Teleporterregistry *TeleporterregistryFilterer) WatchAddProtocolVersion(opts *bind.WatchOpts, sink chan<- *TeleporterregistryAddProtocolVersion, version []*big.Int, protocolAddress []common.Address) (event.Subscription, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var protocolAddressRule []interface{}
	for _, protocolAddressItem := range protocolAddress {
		protocolAddressRule = append(protocolAddressRule, protocolAddressItem)
	}

	logs, sub, err := _Teleporterregistry.contract.WatchLogs(opts, "AddProtocolVersion", versionRule, protocolAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterregistryAddProtocolVersion)
				if err := _Teleporterregistry.contract.UnpackLog(event, "AddProtocolVersion", log); err != nil {
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

// ParseAddProtocolVersion is a log parse operation binding the contract event 0xa5eed93d951a9603d5f7c0a57de79a299dd3dbd5e51429be209d8053a42ab43a.
//
// Solidity: event AddProtocolVersion(uint256 indexed version, address indexed protocolAddress)
func (_Teleporterregistry *TeleporterregistryFilterer) ParseAddProtocolVersion(log types.Log) (*TeleporterregistryAddProtocolVersion, error) {
	event := new(TeleporterregistryAddProtocolVersion)
	if err := _Teleporterregistry.contract.UnpackLog(event, "AddProtocolVersion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
