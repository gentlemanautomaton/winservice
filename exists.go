package winservice

import (
	"strings"

	"golang.org/x/sys/windows/svc/mgr"
)

// Exists returns true if a service with the given name exists.
func Exists(name string) (bool, error) {
	m, err := mgr.Connect()
	if err != nil {
		return false, err
	}
	defer m.Disconnect()

	return serviceExists(m, name)
}

func serviceExists(m *mgr.Mgr, name string) (bool, error) {
	services, err := m.ListServices()
	if err != nil {
		return false, err
	}

	for _, service := range services {
		if strings.EqualFold(name, service) {
			return true, nil
		}
	}

	return false, nil
}
