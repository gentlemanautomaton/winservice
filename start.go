package winservice

import (
	"context"
	"time"

	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
)

// Start issues a start command to a service and waits for it to start. It returns
// an error if it fails or the context is cancelled.
//
// Start returns without error if the service is already started.
func Start(ctx context.Context, name string) error {
	m, err := mgr.Connect()
	if err != nil {
		return OpError{Op: "start", Service: name, Err: err}
	}
	defer m.Disconnect()

	if err := startService(ctx, m, name); err != nil {
		return OpError{Op: "start", Service: name, Err: err}
	}

	return nil
}

func startService(ctx context.Context, m *mgr.Mgr, name string) error {
	s, err := m.OpenService(name)
	if err != nil {
		return err
	}
	defer s.Close()

	ticker := time.NewTicker(PollingInterval)
	defer ticker.Stop()

	for {
		status, err := s.Query()
		if err != nil {
			return err
		}
		if status.State == svc.Stopped {
			// The service has stopped
			break
		} else if status.State != svc.StopPending {
			// The service is already running
			return nil
		} else {
			// The service is stopping (wait for it)
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-ticker.C:
			}
		}
	}

	if err := s.Start(); err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			status, err := s.Query()
			if err != nil {
				return err
			}
			if status.State != svc.StartPending {
				return nil
			}
		}
	}
}
