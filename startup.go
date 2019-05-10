package winservice

// Startup is a windows service startup mode.
type Startup uint32

// Apply applies the service startup mode to conf.
func (s Startup) Apply(conf *Config) {
	conf.Startup = s
}

// Windows service startup modes.
//
// https://docs.microsoft.com/en-us/windows/desktop/api/winsvc/nf-winsvc-createservicew
const (
	// BootStart services are driver services started by the system loader.
	BootStart Startup = 0x00000000 // SERVICE_BOOT_START

	// SystemStart services are driver services started by IoInitSystem.
	SystemStart Startup = 0x00000001 // SERVICE_SYSTEM_START

	// AutoStart services are started automatically when the system starts.
	AutoStart Startup = 0x00000002 // SERVICE_AUTO_START

	// DemandStart services are started upon request.
	DemandStart Startup = 0x00000003 // SERVICE_DEMAND_START

	// Disabled services cannot be started.
	Disabled Startup = 0x00000004 // SERVICE_DISABLED
)
