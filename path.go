package winservice

// Path is a path to a windows service executable.
type Path string

// Apply applies the service path to conf.
func (p Path) Apply(conf *Config) {
	conf.Path = string(p)
}
