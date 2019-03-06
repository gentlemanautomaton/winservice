package winservice

import "syscall"

const (
	// ErrAccessDenied is returned when the calling process has insufficient
	// permissions.
	ErrAccessDenied = syscall.ERROR_ACCESS_DENIED // ERROR_ACCESS_DENIED

	// ErrInvalidHandle is returned when an invalid service handle has
	// been provided to a system call.
	ErrInvalidHandle = syscall.Errno(0x00000006) // ERROR_INVALID_HANDLE

	// ErrInvalidParameter is returned when an invalid argument has been
	// provided to a system call.
	ErrInvalidParameter = syscall.Errno(0x00000057) // ERROR_INVALID_PARAMETER

	// ErrDependentServicesRunning is returned when a service cannot be
	// stopped because it has other services depending on it.
	ErrDependentServicesRunning = syscall.Errno(0x0000041B) // ERROR_DEPENDENT_SERVICES_RUNNING

	// ErrInvalidServiceControl is returned when an invalid service control
	// code has been sent.
	ErrInvalidServiceControl = syscall.Errno(0x0000041C) // ERROR_INVALID_SERVICE_CONTROL

	// ErrServiceRequestTimeout is returned when a service does not respond
	// to a control code within the windows service manager's timeout.
	ErrServiceRequestTimeout = syscall.Errno(0x0000041D) // ERROR_SERVICE_REQUEST_TIMEOUT

	// ErrServiceDoesNotExist is returned when a requested service does not
	// exist.
	ErrServiceDoesNotExist = syscall.Errno(0x00000424) // ERROR_SERVICE_DOES_NOT_EXIST

	// ErrServiceCannotAcceptControl is returned when a service is not in a condition to
	// accept a particular control code.
	ErrServiceCannotAcceptControl = syscall.Errno(0x00000425) // ERROR_SERVICE_CANNOT_ACCEPT_CTRL

	// ErrServiceNotActive is returned when a service is not running.
	ErrServiceNotActive = syscall.Errno(0x00000426) // ERROR_SERVICE_NOT_ACTIVE

	// ErrShutdownInProgress is returned when an action cannot be taken on a
	// service because the system is shutting down.
	ErrShutdownInProgress = syscall.Errno(0x0000045B) // ERROR_SHUTDOWN_IN_PROGRESS
)
