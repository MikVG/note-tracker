package main

import (
	"github.com/MikVG/note-tracker/internal/app"
	"github.com/MikVG/note-tracker/internal/config"
	"github.com/MikVG/note-tracker/internal/repo/memstorage"
	"github.com/MikVG/note-tracker/internal/server"
	"github.com/MikVG/note-tracker/internal/service"
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

	repo := memstorage.New()
	userService := service.NewUserService(repo)
	taskService := service.NewTaskService(repo)
	server := server.New(*cfg, userService, taskService)

	app := app.NewApp(*cfg, server, repo)

	if err := app.StartApp(); err != nil {
		panic(err)
	}
}
