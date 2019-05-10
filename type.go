package winservice

// Type is a windows service type.
type Type uint32

// Apply applies the service type to conf. Its bits are OR'd with the current
// value of conf.Type.
func (t Type) Apply(conf *Config) {
	conf.Type |= t
}

// Windows service types.
//
// https://docs.microsoft.com/en-us/windows/desktop/api/winsvc/nf-winsvc-createservicew
const (
	KernelDriver        Type = 0x00000001 // SERVICE_KERNEL_DRIVER
	FileSystemDriver    Type = 0x00000002 // SERVICE_FILE_SYSTEM_DRIVER
	Adapter             Type = 0x00000004 // SERVICE_ADAPTER
	RecognizerDriver    Type = 0x00000008 // SERVICE_RECOGNIZER_DRIVER
	IsolatedProcess     Type = 0x00000010 // SERVICE_WIN32_OWN_PROCESS
	SharedProcess       Type = 0x00000020 // SERVICE_WIN32_SHARE_PROCESS
	IsolatedUserProcess Type = 0x00000050 // SERVICE_USER_OWN_PROCESS
	SharedUserProcess   Type = 0x00000060 // SERVICE_USER_SHARE_PROCESS
)
