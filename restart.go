package winservice

import (
	"context"

	"golang.org/x/sys/windows/svc/mgr"
)

// Restart issues a stop command to a service if it is running and waits
// for it to stop, then issues a start command and waits for it to start.
// It returns an error if either command fails or the context is cancelled.
//
// If the service is not already running this is equivalent to calling Start.
func Retart(ctx context.Context, name string) error {
	m, err := mgr.Connect()
	if err != nil {
		return OpError{Op: "start", Service: name, Err: err}
	}
	defer m.Disconnect()

	if err := stopService(ctx, m, name); err != nil {
		return OpError{Op: "stop", Service: name, Err: err}
	}

	if err := startService(ctx, m, name); err != nil {
		return OpError{Op: "start", Service: name, Err: err}
	}

	return nil
}
