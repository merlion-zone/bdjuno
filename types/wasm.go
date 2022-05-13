package types

import (
	"time"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// WasmParams represents the CosmWasm code in x/wasm module
type WasmParams struct {
	CodeUploadAccess             *wasmtypes.AccessConfig
	InstantiateDefaultPermission int32
	MaxWasmCodeSize              uint64
	Height                       int64
}

// NewWasmParams allows to build a new x/wasm params instance
func NewWasmParams(
	codeUploadAccess *wasmtypes.AccessConfig, instantiateDefaultPermission int32, maxWasmCodeSize uint64, height int64,
) WasmParams {
	return WasmParams{
		CodeUploadAccess:             codeUploadAccess,
		InstantiateDefaultPermission: instantiateDefaultPermission,
		MaxWasmCodeSize:              maxWasmCodeSize,
		Height:                       height,
	}
}

// WasmCode represents the CosmWasm code in x/wasm module
type WasmCode struct {
	Sender                string
	WasmByteCode          []byte
	InstantiatePermission *wasmtypes.AccessConfig
	CodeID                uint64
	Height                int64
}

// NewWasmCode allows to build a new x/wasm code instance
func NewWasmCode(
	sender string, wasmByteCode []byte, initPermission *wasmtypes.AccessConfig, codeID uint64, height int64,
) WasmCode {
	return WasmCode{
		Sender:                sender,
		WasmByteCode:          wasmByteCode,
		InstantiatePermission: initPermission,
		CodeID:                codeID,
		Height:                height,
	}
}

// WasmContract represents the CosmWasm contract in x/wasm module
type WasmContract struct {
	Sender                string
	Creator               string
	Admin                 string
	CodeID                uint64
	Label                 string
	RawContractMsg        wasmtypes.RawContractMessage
	Funds                 sdk.Coins
	ContractAddress       string
	Data                  string
	InstantiatedAt        time.Time
	ContractInfoExtension string
	Height                int64
}

// NewWasmCode allows to build a new x/wasm contract instance
func NewWasmContract(
	sender string, admin string, codeID uint64, label string, rawMsg wasmtypes.RawContractMessage, funds sdk.Coins, contractAddress string, data string,
	instantiatedAt time.Time, creator string, contractInfoExtension string, height int64,
) WasmContract {
	rawContractMsg, _ := rawMsg.MarshalJSON()

	return WasmContract{
		Sender:                sender,
		Creator:               creator,
		Admin:                 admin,
		CodeID:                codeID,
		Label:                 label,
		RawContractMsg:        rawContractMsg,
		Funds:                 funds,
		ContractAddress:       contractAddress,
		Data:                  data,
		InstantiatedAt:        instantiatedAt,
		ContractInfoExtension: contractInfoExtension,
		Height:                height,
	}
}

// WasmExecuteContract represents the CosmWasm execute contract in x/wasm module
type WasmExecuteContract struct {
	Sender          string
	ContractAddress string
	RawContractMsg  []byte
	Funds           sdk.Coins
	Data            string
	ExecutedAt      time.Time
	Height          int64
}

// NewWasmExecuteContract allows to build a new x/wasm execute contract instance
func NewWasmExecuteContract(
	sender string, contractAddress string, rawMsg wasmtypes.RawContractMessage,
	funds sdk.Coins, data string, executedAt time.Time, height int64,
) WasmExecuteContract {
	rawContractMsg, _ := rawMsg.MarshalJSON()

	return WasmExecuteContract{
		Sender:          sender,
		ContractAddress: contractAddress,
		RawContractMsg:  rawContractMsg,
		Funds:           funds,
		Data:            data,
		ExecutedAt:      executedAt,
		Height:          height,
	}
}