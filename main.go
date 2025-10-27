package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/vvetta/exact-time/ex_time"
)

const (
	NTPHOST = "0.beevik-ntp.pool.ntp.org"
)

func main() {
	_ = godotenv.Load(".env")

	hostEnv := os.Getenv("NTPHOST")
	var host string

	if hostEnv == "" {
		host = NTPHOST
	} else {
		host = hostEnv
	}

	ntpServer := ex_time.NewNTPServer(host)

	time, err := ntpServer.GetExactTime()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v \n", err)
		os.Exit(1)
	}

	fmt.Printf("Exact time: %v", time)
}
