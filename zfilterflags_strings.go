// Code generated by "stringer -output=zfilterflags_strings.go -type=fwpmFilterFlags -trimprefix=fwpmFilterFlags"; DO NOT EDIT.

package wf

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[fwpmFilterFlagsPersistent-1]
	_ = x[fwpmFilterFlagsBootTime-2]
	_ = x[fwpmFilterFlagsHasProviderContext-4]
	_ = x[fwpmFilterFlagsClearActionRight-8]
	_ = x[fwpmFilterFlagsPermitIfCalloutUnregistered-16]
	_ = x[fwpmFilterFlagsDisabled-32]
	_ = x[fwpmFilterFlagsIndexed-64]
}

const (
	_fwpmFilterFlags_name_0 = "PersistentBootTime"
	_fwpmFilterFlags_name_1 = "HasProviderContext"
	_fwpmFilterFlags_name_2 = "ClearActionRight"
	_fwpmFilterFlags_name_3 = "PermitIfCalloutUnregistered"
	_fwpmFilterFlags_name_4 = "Disabled"
	_fwpmFilterFlags_name_5 = "Indexed"
)

var (
	_fwpmFilterFlags_index_0 = [...]uint8{0, 10, 18}
)

func (i fwpmFilterFlags) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _fwpmFilterFlags_name_0[_fwpmFilterFlags_index_0[i]:_fwpmFilterFlags_index_0[i+1]]
	case i == 4:
		return _fwpmFilterFlags_name_1
	case i == 8:
		return _fwpmFilterFlags_name_2
	case i == 16:
		return _fwpmFilterFlags_name_3
	case i == 32:
		return _fwpmFilterFlags_name_4
	case i == 64:
		return _fwpmFilterFlags_name_5
	default:
		return "fwpmFilterFlags(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
