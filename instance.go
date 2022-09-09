package winservice

import (
	"strings"

	"golang.org/x/sys/windows/svc/mgr"
)

// Instances returns the names of per-user instances for the given service
// template name.
//
// TODO: Use lower level system calls that don't require elevated rights.
//
// FIXME: Right now this returns all services that are prefixed with the
// service template name and an underscore. In addition, we should verify
// that the services returned actually have the correct service type flags.
func Instances(template string) ([]string, error) {
	m, err := mgr.Connect()
	if err != nil {
		return nil, OpError{Op: "instances of", Service: template, Err: err}
	}
	defer m.Disconnect()

	services, err := m.ListServices()
	if err != nil {
		return nil, err
	}

	prefix := template + "_"

	var instances []string
	for _, service := range services {
		if strings.HasPrefix(service, prefix) {
			instances = append(instances, service)
		}
	}

	return instances, nil
}
