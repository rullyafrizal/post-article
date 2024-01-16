package dependency

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rullyafrizal/post-article-api/config"
	"github.com/rullyafrizal/post-article-api/handler"
	"github.com/rullyafrizal/post-article-api/infra/database"
	"github.com/rullyafrizal/post-article-api/pkg/validation"
	"github.com/rullyafrizal/post-article-api/repository"
	"github.com/rullyafrizal/post-article-api/server/router"
	"github.com/rullyafrizal/post-article-api/service"
)

func InitRouter(r *fiber.App) *router.Router {
	cfg := config.Get()
	db := database.New(cfg)
	validator := validation.NewValidator()
	articleRepo := repository.NewPostRepository(db)
	articleService := service.NewArticleService(articleRepo)
	articleHandler := handler.NewArticleHandler(articleService, validator)
	return router.NewRouter(r, cfg, articleHandler)
}
