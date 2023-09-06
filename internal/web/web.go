package web

import (
	"auth/internal/config"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type WebServer struct {
	log    *zap.Logger
	cfg    *config.WebServerConfig
	client *fiber.App
}

func NewWebServer(logger *zap.Logger, cfg *config.WebServerConfig) *WebServer {
	return &WebServer{
		log:    logger,
		cfg:    cfg,
		client: fiber.New(fiber.Config{}),
	}
}

func (w *WebServer) registerRoutes() {
	w.client.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
}

func (w *WebServer) Run() error {
	w.registerRoutes()
	w.log.Info("Starting web server", zap.Int("port", w.cfg.Port))
	return w.client.Listen(fmt.Sprintf("0.0.0.0:%d", w.cfg.Port))
}
