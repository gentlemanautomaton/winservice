package winservice

import (
	"errors"

	"golang.org/x/sys/windows/svc/mgr"
)

// Install installs a service with the given name and options.
//
// NOTE: This function signature is subject to change in future revisions.
//
// TODO: Consider making name an option instead of a required parameter.
func Install(name string, options ...Option) error {
	// Validate the name
	if name == "" {
		return OpError{Op: "install", Err: errors.New("empty service name")}
	}
	if len(name) > 255 {
		// FIXME: Count utf16 characters, not bytes
		return OpError{Op: "install", Service: name, Err: errors.New("service name exceeds 255 bytes")}
	}

	// Build up a composite of the configuration options
	var conf Config
	for _, option := range options {
		option.Apply(&conf)
	}

	// Validate the configuration
	if err := conf.validate(); err != nil {
		return OpError{Op: "install", Service: name, Err: err}
	}

	// Copy options to a mgr.Config struct.
	var mconf mgr.Config
	mconf.ServiceType = uint32(conf.Type)
	mconf.StartType = uint32(conf.Startup)
	mconf.ErrorControl = uint32(conf.Importance)
	mconf.ServiceStartName = conf.Account
	mconf.Password = conf.Password
	mconf.DisplayName = conf.DisplayName
	mconf.Description = conf.Description

	// Connect to the service control manager
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()

	// Make sure the service doesn't already exist
	// TODO: Add options to replace an existing service
	exists, err := serviceExists(m, name)
	if err != nil {
		return err
	}
	if exists {
		return OpError{Op: "install", Service: name, Err: errors.New("a service with that name already exists")}
	}

	// Create the service
	s, err := m.CreateService(name, conf.Path, mconf, conf.Args...)
	if err != nil {
		return OpError{Op: "install", Service: name, Err: err}
	}
	defer s.Close()

	return nil
}
