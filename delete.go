package winservice

import (
	"context"
	"time"

	"golang.org/x/sys/windows/svc/mgr"
)

// Delete removes a service and waits for it to dissolve. It returns an error
// if it fails or the context is cancelled.
//
// Delete will attempt to stop the service first if it is running.
//
// Delete returns without error if the service doesn't exist.
func Delete(ctx context.Context, name string) error {
	m, err := mgr.Connect()
	if err != nil {
		return OpError{Op: "delete", Service: name, Err: err}
	}
	defer m.Disconnect()

	// Try to stop the service if it's running
	stopService(ctx, m, name)

	// Try to delete the service
	if err := deleteService(m, name); err != nil {
		switch err {
		case ErrServiceDoesNotExist:
			return nil
		default:
			return OpError{Op: "delete", Service: name, Err: err}
		}
	}

	// Return early if the service has already dissolved
	exists, err := serviceExists(m, name)
	if err != nil {
		return OpError{Op: "delete", Service: name, Err: err}
	}
	if !exists {
		return nil
	}

	// Wait for the service to dissolve
	ticker := time.NewTicker(PollingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			exists, err := serviceExists(m, name)
			if err != nil {
				return OpError{Op: "delete", Service: name, Err: err}
			}
			if !exists {
				return nil
			}
		}
	}

}

func deleteService(m *mgr.Mgr, name string) error {
	s, err := m.OpenService(name)
	if err != nil {
		return err
	}
	defer s.Close()

	// Ask the system to delete the service. The deletion won't take effect
	// until all open handles to the service are closed, including our own.
	return s.Delete()
}
