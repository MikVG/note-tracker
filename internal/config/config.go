package config

import (
	"cmp"
	"flag"
	"os"
	"strconv"
)

type Config struct {
	Host  string
	Port  int
	Debug bool
}

const (
	defaultHost = "localhost"
	defaultPort = 8080
)

func ReadConfig() (*Config, error) {
	var cfg Config

	flag.StringVar(&cfg.Host, "host", defaultHost, "flag for host")
	flag.IntVar(&cfg.Port, "port", defaultPort, "flag for host")
	flag.BoolVar(&cfg.Debug, "debug", false, "flag for debug")

	flag.Parse()

	if cfg.Host == "localhost" {
		cfg.Host = cmp.Or(os.Getenv("HOST"), cfg.Host)
	}

	if cfg.Port == 8080 {
		defPort := strconv.Itoa(cfg.Port)
		envPort := cmp.Or(os.Getenv("PORT"), defPort)
		port, err := strconv.Atoi(envPort)
		if err != nil {
			return nil, err
		}
		cfg.Port = port
	}

	return &cfg, nil
}
