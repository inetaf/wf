package winfirewall

import (
	"golang.org/x/sys/windows"
)

type fwpmDisplayData0 struct {
	Name        *uint16
	Description *uint16
}

type fwpmSession0Flags uint32

const fwpmSession0FlagDynamic = 1

type fwpmSession0 struct {
	SessionKey           windows.GUID
	DisplayData          fwpmDisplayData0
	Flags                fwpmSession0Flags
	TxnWaitTimeoutMillis uint32
	ProcessID            uint32
	SID                  *windows.SID
	Username             *uint16
	KernelMode           uint8
}

type authnService uint32

const (
	authnServiceWinNT   authnService = 0xa
	authnServiceDefault authnService = 0xffffffff
)

type fwpmLayerEnumTemplate0 struct {
	reserved uint64
}

// LayerFlags are flags associated with a layer.
type LayerFlags uint32

const (
	LayerFlagsKernel             = 1 // classification occurs in kernel mode
	fwpmLayerFlagsBuiltin        = 2 // built-in layer, cannot be deleted
	fwpmLayerFlagsClassifyMostly = 4 // optimized for classifying not enumerating
	fwpmLayerFlagsBuffered       = 8 // buffered?
)

type fwpmLayer0 struct {
	LayerKey           windows.GUID
	DisplayData        fwpmDisplayData0
	Flags              LayerFlags
	NumFields          uint32
	Field              *fwpmField0
	DefaultSublayerKey windows.GUID
	LayerID            uint16
}

type fwpmField0 struct {
	FieldKey *windows.GUID
	Type     fwpmFieldType
	DataType fwpmDataType
}

type fwpmFieldType uint32

const (
	fwpmFieldTypeRawData fwpmFieldType = iota
	fwpmFieldTypeIPAddress
	fwpmFieldTypeFlags
)

type fwpmDataType uint32

const (
	fwpmDataTypeEmpty fwpmDataType = iota
	fwpmDataTypeUint8
	fwpmDataTypeUint6
	fwpmDataTypeUint2
	fwpmDataTypeUint64
	fwpmDataTypeInt8
	fwpmDataTypeInt16
	fwpmDataTypeInt32
	fwpmDataTypeInt64
	fwpmDataTypeFloat
	fwpmDataTypeDouble
	fwpmDataTypeByteArray16
	fwpmDataTypeByteBlob
	fwpmDataTypeSID
	fwpmDataTypeSecurityDescriptor
	fwpmDataTypeTokenInformation
	fwpmDataTypeTokenAccessInformation
	fwpmDataTypeUnicodeString
	fwpmDataTypeV4AddrMask = 0x100 + iota
	fwpmDataTypeV6AddrMask
	fwpmDataTypeRange
)
