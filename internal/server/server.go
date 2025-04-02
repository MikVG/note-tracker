package server

import (
	"fmt"
	"net/http"

	"github.com/MikVG/note-tracker/internal/config"
	"github.com/MikVG/note-tracker/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ServerApi struct {
	server   *http.Server
	valid    *validator.Validate
	uService *service.UserService
	tService *service.TaskService
}

func New(cfg config.Config, uService *service.UserService, tService *service.TaskService) *ServerApi {
	server := http.Server{
		Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
	}

	return &ServerApi{
		server:   &server,
		valid:    validator.New(),
		uService: uService,
		tService: tService,
	}
}

func (s *ServerApi) configRoutes() {
	router := gin.Default()
	router.GET("/tasks", s.getTasks)
	router.POST("/tasks", s.createTask)
	task := router.Group("/tasks")
	{
		task.PUT("/:id", func(ctx *gin.Context) {})
		task.DELETE("/:id", func(ctx *gin.Context) {})
		task.GET("/:id", func(ctx *gin.Context) {})
	}
	users := router.Group("/users")
	{
		users.POST("/register", s.registerUser)
		users.POST("/login", s.loginUser)
	}
	s.server.Handler = router
}

func (s *ServerApi) Start() error {
	s.configRoutes()
	return s.server.ListenAndServe()
}
