package server

import (
	"fmt"
	"net/http"

	"github.com/MikVG/note-tracker/internal/config"
	"github.com/MikVG/note-tracker/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ServerApi struct {
	server *http.Server
	valid  *validator.Validate
	repo   any
}

func New(cfg config.Config, repo any) *ServerApi {
	server := http.Server{
		Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
	}
	router := gin.Default()
	router.GET("/tasks", func(ctx *gin.Context) {})
	router.POST("/tasks", func(ctx *gin.Context) {})
	task := router.Group("/tasks")
	{
		task.PUT("/:id", func(ctx *gin.Context) {})
		task.DELETE("/:id", func(ctx *gin.Context) {})
		task.GET("/:id", func(ctx *gin.Context) {})
	}
	user := router.Group("/user")
	{
		user.POST("/login", func(ctx *gin.Context) {})
		user.POST("/register", func(ctx *gin.Context) {})
		user.GET("/profile", func(ctx *gin.Context) {})
	}

	server.Handler = router
	return &ServerApi{
		server: &server,
		valid:  validator.New(),
		repo:   repo,
	}
}

func (s *ServerApi) Start() error {
	log := logger.Get()
	log.Info().Str("server address", s.server.Addr).Msg("server was started")
	// if err := s.server.ListenAndServe(); err != nil {
	// 	return err
	// }
	// return nil
	return s.server.ListenAndServe()
}
