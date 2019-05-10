package winservice

// Importance is the importance of a windows service to the operating system.
// It determines the action taken when a service fails to start.
type Importance uint32

// Apply applies the service importance to conf.
func (i Importance) Apply(conf *Config) {
	conf.Importance = i
}

// Windows service startup types.
//
// https://docs.microsoft.com/en-us/windows/desktop/api/winsvc/nf-winsvc-createservicew
const (
	// Ignored services will be ignored if they fail to start.
	Ignored Importance = 0x00000000 // SERVICE_ERROR_IGNORE

	// Logged services will have startup failures logged.
	Logged Importance = 0x00000001 // SERVICE_ERROR_NORMAL

	// Essential services cause the operating system to revert to a last known
	// good configuration when the service fails to start, if possible.
	//
	// If the system is already is a last known good configuration it boots
	// normally.
	Essential Importance = 0x00000002 // SERVICE_ERROR_SEVERE

	// Critical services cause the operating system to revert to a last known
	// good configuration when the service fails to start, if possible.
	//
	// If the system is already is a last known good configuration the system
	// halts.
	Critical Importance = 0x00000003 // SERVICE_ERROR_CRITICAL
)
