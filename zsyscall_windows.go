// Code generated by 'go generate'; DO NOT EDIT.

package winfirewall

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modfwpuclnt = windows.NewLazySystemDLL("fwpuclnt.dll")

	procFwpmEngineClose0               = modfwpuclnt.NewProc("FwpmEngineClose0")
	procFwpmEngineOpen0                = modfwpuclnt.NewProc("FwpmEngineOpen0")
	procFwpmFilterCreateEnumHandle0    = modfwpuclnt.NewProc("FwpmFilterCreateEnumHandle0")
	procFwpmFilterDestroyEnumHandle0   = modfwpuclnt.NewProc("FwpmFilterDestroyEnumHandle0")
	procFwpmFilterEnum0                = modfwpuclnt.NewProc("FwpmFilterEnum0")
	procFwpmFreeMemory0                = modfwpuclnt.NewProc("FwpmFreeMemory0")
	procFwpmLayerCreateEnumHandle0     = modfwpuclnt.NewProc("FwpmLayerCreateEnumHandle0")
	procFwpmLayerDestroyEnumHandle0    = modfwpuclnt.NewProc("FwpmLayerDestroyEnumHandle0")
	procFwpmLayerEnum0                 = modfwpuclnt.NewProc("FwpmLayerEnum0")
	procFwpmProviderAdd0               = modfwpuclnt.NewProc("FwpmProviderAdd0")
	procFwpmProviderCreateEnumHandle0  = modfwpuclnt.NewProc("FwpmProviderCreateEnumHandle0")
	procFwpmProviderDeleteByKey0       = modfwpuclnt.NewProc("FwpmProviderDeleteByKey0")
	procFwpmProviderDestroyEnumHandle0 = modfwpuclnt.NewProc("FwpmProviderDestroyEnumHandle0")
	procFwpmProviderEnum0              = modfwpuclnt.NewProc("FwpmProviderEnum0")
	procFwpmSubLayerAdd0               = modfwpuclnt.NewProc("FwpmSubLayerAdd0")
	procFwpmSubLayerCreateEnumHandle0  = modfwpuclnt.NewProc("FwpmSubLayerCreateEnumHandle0")
	procFwpmSubLayerDeleteByKey0       = modfwpuclnt.NewProc("FwpmSubLayerDeleteByKey0")
	procFwpmSubLayerDestroyEnumHandle0 = modfwpuclnt.NewProc("FwpmSubLayerDestroyEnumHandle0")
	procFwpmSubLayerEnum0              = modfwpuclnt.NewProc("FwpmSubLayerEnum0")
	procFwpmTransactionAbort0          = modfwpuclnt.NewProc("FwpmTransactionAbort0")
	procFwpmTransactionBegin0          = modfwpuclnt.NewProc("FwpmTransactionBegin0")
	procFwpmTransactionCommit0         = modfwpuclnt.NewProc("FwpmTransactionCommit0")
)

func fwpmEngineClose0(engineHandle windows.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procFwpmEngineClose0.Addr(), 1, uintptr(engineHandle), 0, 0)
	if r1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func fwpmEngineOpen0(mustBeNil *uint16, authnService authnService, authIdentity *uintptr, session *fwpmSession0, engineHandle *windows.Handle) (err error) {
	r1, _, e1 := syscall.Syscall6(procFwpmEngineOpen0.Addr(), 5, uintptr(unsafe.Pointer(mustBeNil)), uintptr(authnService), uintptr(unsafe.Pointer(authIdentity)), uintptr(unsafe.Pointer(session)), uintptr(unsafe.Pointer(engineHandle)), 0)
	if r1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func fwpmFilterCreateEnumHandle0(engineHandle windows.Handle, enumTemplate *fwpmFilterEnumTemplate0, handle *windows.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procFwpmFilterCreateEnumHandle0.Addr(), 3, uintptr(engineHandle), uintptr(unsafe.Pointer(enumTemplate)), uintptr(unsafe.Pointer(handle)))
	if r1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func fwpmFilterDestroyEnumHandle0(engineHandle windows.Handle, enumHandle windows.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procFwpmFilterDestroyEnumHandle0.Addr(), 2, uintptr(engineHandle), uintptr(enumHandle), 0)
	if r1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func fwpmFilterEnum0(engineHandle windows.Handle, enumHandle windows.Handle, numEntriesRequested uint32, entries ***fwpmFilter0, numEntriesReturned *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procFwpmFilterEnum0.Addr(), 5, uintptr(engineHandle), uintptr(enumHandle), uintptr(numEntriesRequested), uintptr(unsafe.Pointer(entries)), uintptr(unsafe.Pointer(numEntriesReturned)), 0)
	if r1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func fwpmFreeMemory0(p uintptr) {
	syscall.Syscall(procFwpmFreeMemory0.Addr(), 1, uintptr(p), 0, 0)
	return
}

func fwpmLayerCreateEnumHandle0(engineHandle windows.Handle, enumTemplate *fwpmLayerEnumTemplate0, handle *windows.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procFwpmLayerCreateEnumHandle0.Addr(), 3, uintptr(engineHandle), uintptr(unsafe.Pointer(enumTemplate)), uintptr(unsafe.Pointer(handle)))
	if r1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func fwpmLayerDestroyEnumHandle0(engineHandle windows.Handle, enumHandle windows.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procFwpmLayerDestroyEnumHandle0.Addr(), 2, uintptr(engineHandle), uintptr(enumHandle), 0)
	if r1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func fwpmLayerEnum0(engineHandle windows.Handle, enumHandle windows.Handle, numEntriesRequested uint32, entries ***fwpmLayer0, numEntriesReturned *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procFwpmLayerEnum0.Addr(), 5, uintptr(engineHandle), uintptr(enumHandle), uintptr(numEntriesRequested), uintptr(unsafe.Pointer(entries)), uintptr(unsafe.Pointer(numEntriesReturned)), 0)
	if r1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func fwpmProviderAdd0(engineHandle windows.Handle, provider *fwpmProvider0, nilForNow *uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procFwpmProviderAdd0.Addr(), 3, uintptr(engineHandle), uintptr(unsafe.Pointer(provider)), uintptr(unsafe.Pointer(nilForNow)))
	if r1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func fwpmProviderCreateEnumHandle0(engineHandle windows.Handle, enumTemplate *fwpmProviderEnumTemplate0, handle *windows.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procFwpmProviderCreateEnumHandle0.Addr(), 3, uintptr(engineHandle), uintptr(unsafe.Pointer(enumTemplate)), uintptr(unsafe.Pointer(handle)))
	if r1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func fwpmProviderDeleteByKey0(engineHandle windows.Handle, guid *windows.GUID) (err error) {
	r1, _, e1 := syscall.Syscall(procFwpmProviderDeleteByKey0.Addr(), 2, uintptr(engineHandle), uintptr(unsafe.Pointer(guid)), 0)
	if r1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func fwpmProviderDestroyEnumHandle0(engineHandle windows.Handle, enumHandle windows.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procFwpmProviderDestroyEnumHandle0.Addr(), 2, uintptr(engineHandle), uintptr(enumHandle), 0)
	if r1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func fwpmProviderEnum0(engineHandle windows.Handle, enumHandle windows.Handle, numEntriesRequested uint32, entries ***fwpmProvider0, numEntriesReturned *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procFwpmProviderEnum0.Addr(), 5, uintptr(engineHandle), uintptr(enumHandle), uintptr(numEntriesRequested), uintptr(unsafe.Pointer(entries)), uintptr(unsafe.Pointer(numEntriesReturned)), 0)
	if r1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func fwpmSubLayerAdd0(engineHandle windows.Handle, sublayer *fwpmSublayer0, nilForNow *uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procFwpmSubLayerAdd0.Addr(), 3, uintptr(engineHandle), uintptr(unsafe.Pointer(sublayer)), uintptr(unsafe.Pointer(nilForNow)))
	if r1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func fwpmSubLayerCreateEnumHandle0(engineHandle windows.Handle, enumTemplate *fwpmSublayerEnumTemplate0, handle *windows.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procFwpmSubLayerCreateEnumHandle0.Addr(), 3, uintptr(engineHandle), uintptr(unsafe.Pointer(enumTemplate)), uintptr(unsafe.Pointer(handle)))
	if r1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func fwpmSubLayerDeleteByKey0(engineHandle windows.Handle, guid *windows.GUID) (err error) {
	r1, _, e1 := syscall.Syscall(procFwpmSubLayerDeleteByKey0.Addr(), 2, uintptr(engineHandle), uintptr(unsafe.Pointer(guid)), 0)
	if r1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func fwpmSubLayerDestroyEnumHandle0(engineHandle windows.Handle, enumHandle windows.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procFwpmSubLayerDestroyEnumHandle0.Addr(), 2, uintptr(engineHandle), uintptr(enumHandle), 0)
	if r1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func fwpmSubLayerEnum0(engineHandle windows.Handle, enumHandle windows.Handle, numEntriesRequested uint32, entries ***fwpmSublayer0, numEntriesReturned *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procFwpmSubLayerEnum0.Addr(), 5, uintptr(engineHandle), uintptr(enumHandle), uintptr(numEntriesRequested), uintptr(unsafe.Pointer(entries)), uintptr(unsafe.Pointer(numEntriesReturned)), 0)
	if r1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func fwpmTransactionAbort0(engineHandle windows.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procFwpmTransactionAbort0.Addr(), 1, uintptr(engineHandle), 0, 0)
	if r1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func fwpmTransactionBegin0(engineHandle windows.Handle, flags uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procFwpmTransactionBegin0.Addr(), 2, uintptr(engineHandle), uintptr(flags), 0)
	if r1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func fwpmTransactionCommit0(engineHandle windows.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procFwpmTransactionCommit0.Addr(), 1, uintptr(engineHandle), 0, 0)
	if r1 != 0 {
		err = errnoErr(e1)
	}
	return
}
