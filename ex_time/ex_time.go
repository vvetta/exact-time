package ex_time

import (
	"time"

	"github.com/beevik/ntp"
)

type NTPServer struct {
	host string
}

func NewNTPServer(remoteNTPHost string) *NTPServer {
	return &NTPServer{
		host: remoteNTPHost,
	}
}

func (n *NTPServer) GetExactTime() (time.Time, error) {
	exactTime, err := ntp.Time(n.host)
	if err != nil {
		return time.Time{}, err
	}

	return exactTime, nil
}
