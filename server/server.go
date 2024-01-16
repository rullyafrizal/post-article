package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rullyafrizal/post-article-api/config"
	"github.com/rullyafrizal/post-article-api/server/dependency"
	"log"
)

type Server interface {
	Run()
}

type HttpServer struct {
	cfg *config.Config
}

func NewHttp() Server {
	if err := config.Init(); err != nil {
		log.Fatalf("error initialize config, cause:%v\n", err)
	}

	return &HttpServer{
		cfg: config.Get(),
	}
}

func (s *HttpServer) Run() {
	r := fiber.New()
	dependency.InitRouter(r).Run(s.cfg.AppPort)
}
