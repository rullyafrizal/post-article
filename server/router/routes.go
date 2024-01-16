package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rullyafrizal/post-article-api/config"
	"github.com/rullyafrizal/post-article-api/handler"
	"log"
)

type Router struct {
	cfg            *config.Config
	router         *fiber.App
	articleHandler *handler.ArticleHandler
}

func NewRouter(router *fiber.App, cfg *config.Config, articleHandler *handler.ArticleHandler) *Router {
	return &Router{
		router:         router,
		cfg:            cfg,
		articleHandler: articleHandler,
	}
}

func (r *Router) Run(port string) {
	r.setRoutes()
	r.run(port)
}

func (r *Router) setRoutes() {
	rg := r.router.Group("")
	rg.Use(recover.New())
	rg.Use(cors.New())
	rg.Use(logger.New())
	r.setArticleRoutes(rg)
}

func (r *Router) setArticleRoutes(rg fiber.Router) {
	rg.Post("/article", r.articleHandler.CreateArticle)
	rg.Get("/article/:limit/:offset", r.articleHandler.GetArticles)
	rg.Get("/article/:id", r.articleHandler.FindArticle)
	rg.Put("/article/:id", r.articleHandler.UpdateArticle)
	rg.Delete("/article/:id", r.articleHandler.DeleteArticle)
}

func (r *Router) run(port string) {
	if port == "" {
		port = "8080"
	}
	log.Printf("running on port ::[:%v]", port)
	if err := r.router.Listen(":" + port); err != nil {
		log.Fatalf("failed to run on port [::%v]", port)
	}
}
