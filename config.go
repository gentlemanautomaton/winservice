package winservice

import "errors"

// Config holds configuration for a service.
//
// NOTE: This structure is subject to change in future revisions.
type Config struct {
	Path        string
	Args        []string
	Type        Type
	Startup     Startup
	Importance  Importance
	Account     string
	Password    string // Write Only
	DisplayName string
	Description string
}

// Apply copies the values from conf to c2. It allows a configuration struct
// to be passed as an option.
func (conf Config) Apply(c2 *Config) {
	c2.Path = conf.Path
	c2.Args = conf.Args
	c2.Account = conf.Account
	c2.Password = conf.Password
	c2.Type = conf.Type
	c2.Startup = conf.Startup
}

// validate returns an error if the config is invalid
func (conf Config) validate() error {
	if conf.Path == "" {
		return errors.New("windows service configuration is missing a path")
	}
	return nil
}

// Option is a service configuration option.
type Option interface {
	Apply(conf *Config)
}

// OptionFunc is a service configuration function.
type OptionFunc func(conf *Config)

// Apply applies the configuration option to conf.
func (f OptionFunc) Apply(conf *Config) {
	f(conf)
}

// Args returns a service configuration option for a set of service arguments.
func Args(args ...string) Option {
	return OptionFunc(func(conf *Config) {
		conf.Args = args
	})
}

// Account returns a service configuration option for the given account
// credentials.
func Account(account, password string) Option {
	return OptionFunc(func(conf *Config) {
		conf.Account = account
		conf.Password = password
	})
}
