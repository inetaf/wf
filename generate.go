package wf

//go:generate go run generators/gen_guids.go includes/fwpmu.h zguids.go
//go:generate go run golang.org/x/sys/windows/mkwinsyscall -output zsyscall_windows.go syscall.go

//go:generate stringer -output=zlayerflags_strings.go -type=fwpmLayerFlags -trimprefix=fwpmLayerFlags
//go:generate stringer -output=zfieldtype_strings.go -type=fwpmFieldType -trimprefix=fwpmFieldtype
//go:generate stringer -output=zsublayerflags_strings.go -type=fwpmSublayerFlags -trimprefix=fwpmSublayerFlags
//go:generate stringer -output=zfilterenumtype_strings.go -type=filterEnumType -trimprefix=filterEnumType
//go:generate stringer -output=zfilterenumflags_strings.go -type=filterEnumFlags -trimprefix=filterEnumFlags
//go:generate stringer -output=zactiontype_strings.go -type=actionType -trimprefix=actionType
//go:generate stringer -output=zmatchtype_strings.go -type=matchType -trimprefix=matchType
//go:generate stringer -output=zfilterflags_strings.go -type=filterFlags -trimprefix=filterFlags
//go:generate stringer -output=zproviderflags_strings.go -type=fwpmProviderFlags -trimprefix=fwpmProviderFlags
//go:generate stringer -output=zdatatype_strings.go -type=dataType -trimprefix=dataType
