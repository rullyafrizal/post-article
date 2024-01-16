package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rullyafrizal/post-article-api/data/entity"
	"github.com/rullyafrizal/post-article-api/pkg/http"
	"github.com/rullyafrizal/post-article-api/pkg/validation"
	"github.com/rullyafrizal/post-article-api/service"
	"strconv"
)

type ArticleHandler struct {
	validator      validation.Validator
	articleService service.ArticleService
}

func NewArticleHandler(articleService service.ArticleService, validator validation.Validator) *ArticleHandler {
	return &ArticleHandler{
		articleService: articleService,
		validator:      validator,
	}
}

func (h *ArticleHandler) CreateArticle(ctx *fiber.Ctx) error {
	req := new(entity.CreateArticleRequest)
	err := ctx.BodyParser(req)
	if err != nil {
		return http.ResponseBadRequest(ctx, []error{err})
	}

	errs := h.validator.Validate(req)
	if errs != nil {
		return http.ResponseStatusUnprocessableEntityError(ctx, errs)
	}

	err = h.articleService.Create(ctx.Context(), req)
	if err != nil {
		return http.ResponseError(ctx, []error{err}, http.TypeInternalServerError)
	}
	return http.ResponseSuccess(ctx, nil)
}

func (h *ArticleHandler) GetArticles(ctx *fiber.Ctx) error {
	limit, _ := strconv.Atoi(ctx.Params("limit"))
	offset, _ := strconv.Atoi(ctx.Params("offset"))
	req := &entity.GetArticlesRequest{
		Limit:  limit,
		Offset: offset,
	}

	resp, err := h.articleService.FindAll(ctx.Context(), req)
	if err != nil {
		return http.ResponseError(ctx, []error{err}, http.TypeInternalServerError)
	}
	return http.ResponseSuccess(ctx, resp.Articles)
}

func (h *ArticleHandler) FindArticle(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	resp, err := h.articleService.FindById(ctx.Context(), id)
	if err != nil {
		return http.ResponseError(ctx, []error{err}, http.TypeInternalServerError)
	}
	return http.ResponseSuccess(ctx, resp)
}

func (h *ArticleHandler) UpdateArticle(ctx *fiber.Ctx) error {
	req := new(entity.UpdateArticleRequest)
	err := ctx.BodyParser(req)
	if err != nil {
		return http.ResponseBadRequest(ctx, []error{err})
	}

	errs := h.validator.Validate(req)
	if errs != nil {
		return http.ResponseStatusUnprocessableEntityError(ctx, errs)
	}

	req.Id, _ = strconv.Atoi(ctx.Params("id"))

	err = h.articleService.Update(ctx.Context(), req)
	if err != nil {
		return http.ResponseError(ctx, []error{err}, http.TypeInternalServerError)
	}
	return http.ResponseSuccess(ctx, nil)
}

func (h *ArticleHandler) DeleteArticle(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	err := h.articleService.Delete(ctx.Context(), id)
	if err != nil {
		return http.ResponseError(ctx, []error{err}, http.TypeInternalServerError)
	}
	return http.ResponseSuccess(ctx, nil)
}
