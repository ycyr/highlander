package main

import (
	"flag"
	"time"
)

type Cmd struct {
	bind         string
	remote       string
	preferredIp  string
	healthTick   time.Duration
	healthExpiry time.Duration
}

func parseCmd() Cmd {
	var cmd Cmd
	flag.DurationVar(&cmd.healthTick, "t", 1*time.Second, "health check every n second")
	flag.DurationVar(&cmd.healthExpiry, "e", 5*time.Second, "disable stream after n seconds of inactivity")
	flag.StringVar(&cmd.bind, "l", "0.0.0.0:9091", "listen on ip:port")
	flag.StringVar(&cmd.remote, "r", "http://localhost:9090", "reverse proxy addr")
	flag.StringVar(&cmd.preferredIp, "i", "", "preferred source ip address")
	flag.Parse()
	return cmd
}
