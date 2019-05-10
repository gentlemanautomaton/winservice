package winservice

// DisplayName is the user-friendly name of a windows service.
type DisplayName string

// Apply applies the display name to conf.
func (dn DisplayName) Apply(conf *Config) {
	conf.DisplayName = string(dn)
}
