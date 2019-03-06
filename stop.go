package winservice

import (
	"context"
	"time"

	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
)

// Stop issues a stop command to a service and waits for it to stop. It returns
// an error if it fails or the context is cancelled.
//
// Stop returns without error if the service is already stopped.
func Stop(ctx context.Context, name string) error {
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()

	return stopService(ctx, m, name)
}

func stopService(ctx context.Context, m *mgr.Mgr, name string) error {
	s, err := m.OpenService(name)
	if err != nil {
		return err
	}
	defer s.Close()

	status, err := s.Control(svc.Stop)
	switch err {
	case nil:
		// The stop command was issued
	case ErrServiceNotActive:
		// The service is not running
		return nil
	case ErrServiceCannotAcceptControl:
		// The service is in some state that can't accept a stop command
		switch status.State {
		case svc.StopPending:
			// The stop is already pending, which is what we wanted anyway
		default:
			return err
		}
	default:
		return err
	}

	ticker := time.NewTicker(PollingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			status, err := s.Query()
			if err != nil {
				return err
			}
			if status.State == svc.Stopped {
				return nil
			}
		}
	}
}
