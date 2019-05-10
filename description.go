package winservice

// Description is the description of a windows service.
type Description string

// Apply applies the description to conf.
func (d Description) Apply(conf *Config) {
	conf.Description = string(d)
}
