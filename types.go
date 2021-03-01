package wf

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

type fwpmLayerFlags uint32

const (
	fwpmLayerFlagsKernel fwpmLayerFlags = 1 << iota
	fwpmLayerFlagsBuiltin
	fwpmLayerFlagsClassifyMostly
	fwpmLayerFlagsBuffered
)

type fwpmLayer0 struct {
	LayerKey           windows.GUID
	DisplayData        fwpmDisplayData0
	Flags              fwpmLayerFlags
	NumFields          uint32
	Fields             *fwpmField0
	DefaultSublayerKey windows.GUID
	LayerID            uint16
}

type fwpmFieldType uint32

const (
	fwpmFieldTypeRawData   fwpmFieldType = iota // no special semantics
	fwpmFieldTypeIPAddress                      // data is an IP address
	fwpmFieldTypeFlags                          // data is a flag bitfield
)

type dataType uint32

const (
	dataTypeEmpty dataType = iota
	dataTypeUint8
	dataTypeUint16
	dataTypeUint32
	dataTypeUint64
	dataTypeInt8
	dataTypeInt16
	dataTypeInt32
	dataTypeInt64
	dataTypeFloat
	dataTypeDouble
	dataTypeByteArray16
	dataTypeByteBlob
	dataTypeSID
	dataTypeSecurityDescriptor
	dataTypeTokenInformation
	dataTypeTokenAccessInformation
	dataTypeUnicodeString
	dataTypeArray6
	dataTypeBitmapIndex
	dataTypeBitmapArray64

	dataTypeV4AddrMask = 256
	dataTypeV6AddrMask = 257
	dataTypeRange      = 258
)

type fwpmField0 struct {
	FieldKey *windows.GUID
	Type     fwpmFieldType
	DataType dataType
}

type fwpmSublayerEnumTemplate0 struct {
	ProviderKey *windows.GUID
}

type fwpByteBlob struct {
	Size uint32
	Data *uint8
}

type fwpmSublayerFlags uint32

const fwpmSublayerFlagsPersistent fwpmSublayerFlags = 1

type fwpmSublayer0 struct {
	SublayerKey  windows.GUID
	DisplayData  fwpmDisplayData0
	Flags        fwpmSublayerFlags
	ProviderKey  *windows.GUID
	ProviderData fwpByteBlob
	Weight       uint16
}

type fwpmProviderEnumTemplate0 struct {
	Reserved uint64
}

type fwpmProviderFlags uint32

const (
	fwpmProviderFlagsPersistent fwpmProviderFlags = 0x01
	fwpmProviderFlagsDisabled   fwpmProviderFlags = 0x10
)

type fwpmProvider0 struct {
	ProviderKey  windows.GUID
	DisplayData  fwpmDisplayData0
	Flags        fwpmProviderFlags
	ProviderData fwpByteBlob
	ServiceName  *uint16
}

type fwpValue0 struct {
	Type  dataType
	Value uintptr // unioned value
}

type fwpmFilterFlags uint32

const (
	fwpmFilterFlagsPersistent fwpmFilterFlags = 1 << iota
	fwpmFilterFlagsBootTime
	fwpmFilterFlagsHasProviderContext
	fwpmFilterFlagsClearActionRight
	fwpmFilterFlagsPermitIfCalloutUnregistered
	fwpmFilterFlagsDisabled
	fwpmFilterFlagsIndexed
)

type fwpmAction0 struct {
	Type Action
	GUID windows.GUID
}

type fwpmFilter0 struct {
	FilterKey           windows.GUID
	DisplayData         fwpmDisplayData0
	Flags               fwpmFilterFlags
	ProviderKey         *windows.GUID
	ProviderData        fwpByteBlob
	LayerKey            windows.GUID
	SublayerKey         windows.GUID
	Weight              fwpValue0
	NumFilterConditions uint32
	FilterConditions    *fwpmFilterCondition0
	Action              fwpmAction0
	ProviderContextKey  windows.GUID
	Reserved            *windows.GUID
	FilterID            uint64
	EffectiveWeight     fwpValue0
}

type fwpConditionValue0 struct {
	Type  dataType
	Value uintptr
}

type fwpmFilterCondition0 struct {
	FieldKey  windows.GUID
	MatchType MatchType
	Value     fwpConditionValue0
}

type fwpV4AddrAndMask struct {
	Addr, Mask uint32
}

type fwpV6AddrAndMask struct {
	Addr         [16]byte
	PrefixLength uint8
}

type fwpmProviderContextEnumTemplate0 struct {
	ProviderKey         *windows.GUID
	ProviderContextType uint32
}

type fwpmFilterEnumTemplate0 struct {
	ProviderKey             *windows.GUID
	LayerKey                windows.GUID
	EnumType                filterEnumType
	Flags                   filterEnumFlags
	ProviderContextTemplate *fwpmProviderContextEnumTemplate0 // TODO: wtf?
	NumConditions           uint32
	Conditions              *fwpmFilterCondition0
	ActionMask              uint32
	CalloutKey              *windows.GUID
}

type fwpRange0 struct {
	From, To fwpValue0
}
