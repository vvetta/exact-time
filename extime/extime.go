package extime

import (
	"time"

	"github.com/beevik/ntp"
)

// NTPServer wraps a remote NTP host and provides time queries.
type NTPServer struct {
	host string
}

// NewNTPServer returns a new NTPServer for the given host.
func NewNTPServer(remoteNTPHost string) *NTPServer {
	return &NTPServer{
		host: remoteNTPHost,
	}
}

// GetExactTime returns current time from the configured NTP host.
func (n *NTPServer) GetExactTime() (time.Time, error) {
	exactTime, err := ntp.Time(n.host)
	if err != nil {
		return time.Time{}, err
	}

	return exactTime, nil
}
