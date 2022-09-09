package winservice

import (
	"syscall"

	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/svc/mgr"
)

// Exists returns true if a service with the given name exists.
func Exists(name string) (bool, error) {
	m, err := mgr.Connect()
	if err != nil {
		return false, OpError{Op: "check", Service: name, Err: err}
	}
	defer m.Disconnect()

	return serviceExists(m, name)
}

func serviceExists(m *mgr.Mgr, name string) (bool, error) {
	pname, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return false, err
	}

	h, err := windows.OpenService(m.Handle, pname, windows.SERVICE_QUERY_STATUS)
	if err != nil {
		return false, err
	}
	defer windows.CloseServiceHandle(h)

	return true, nil
}
