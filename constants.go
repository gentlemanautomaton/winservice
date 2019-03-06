package winservice

import "time"

const (
	// PollingInterval is the service polling interval used by some winservice
	// functions.
	PollingInterval = time.Millisecond * 50
)
