// Code generated by "stringer -output=zproviderflags_strings.go -type=ProviderFlags -trimprefix=ProviderFlags"; DO NOT EDIT.

package wf

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ProviderFlagsPersistent-1]
	_ = x[ProviderFlagsDisabled-16]
}

const (
	_ProviderFlags_name_0 = "Persistent"
	_ProviderFlags_name_1 = "Disabled"
)

func (i ProviderFlags) String() string {
	switch {
	case i == 1:
		return _ProviderFlags_name_0
	case i == 16:
		return _ProviderFlags_name_1
	default:
		return "ProviderFlags(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
