package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/vvetta/exact-time/extime"
)

const (
	defautlNTPHOST = "0.beevik-ntp.pool.ntp.org"
)

func main() {
	_ = godotenv.Load(".env")

	hostEnv := os.Getenv("NTPHOST")
	var host string

	if hostEnv == "" {
		host = defautlNTPHOST
	} else {
		host = hostEnv
	}

	ntpServer := extime.NewNTPServer(host)

	exactTime, err := ntpServer.GetExactTime()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v \n", err)
		os.Exit(1)
	}

	fmt.Printf("Exact time: %v", exactTime.Format(time.RFC3339Nano))
}
