package main

import (
	// "github.com/MikVG/note-tracker/internal/app"
	"github.com/MikVG/note-tracker/internal/config"
	"github.com/MikVG/note-tracker/internal/repo/memstorage"
	"github.com/MikVG/note-tracker/internal/server"
	// "github.com/MikVG/note-tracker/pkg/logger"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}

	repo := memstorage.New()

	server := server.New(*cfg, repo)

	// log := logger.Get(cfg.Debug)
	// log.Debug().Msg("logger was initialized")
	// log.Debug().Str("host", cfg.Host).Int("port", cfg.Port).Send()
	// app := app.NewApp(*cfg, server, nil)

	if err := server.Start(); err != nil {
		panic(err)
	}
}
