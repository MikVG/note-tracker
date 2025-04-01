package main

import (
	"github.com/MikVG/note-tracker/internal/app"
	"github.com/MikVG/note-tracker/internal/config"
	"github.com/MikVG/note-tracker/internal/server"
	"github.com/MikVG/note-tracker/pkg/logger"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}

	log := logger.Get(cfg.Debug)
	log.Debug().Msg("logger was initialized")
	log.Debug().Str("host", cfg.Host).Int("port", cfg.Port).Send()
	server := server.New(*cfg, nil)
	app := app.NewApp(*cfg, server, nil)

	if err := app.StartApp(); err != nil {
		panic(err)
	}
}
