package service

import (
	"context"
	"database/sql"
	"github.com/rullyafrizal/post-article-api/data/entity"
	"github.com/rullyafrizal/post-article-api/data/model"
	"github.com/rullyafrizal/post-article-api/repository"
	"github.com/sirupsen/logrus"
	"time"
)

type ArticleService interface {
	Create(ctx context.Context, request *entity.CreateArticleRequest) error
	FindAll(ctx context.Context, request *entity.GetArticlesRequest) (*entity.GetArticlesResponse, error)
	FindById(ctx context.Context, id int) (*entity.ArticleResponse, error)
	Update(ctx context.Context, request *entity.UpdateArticleRequest) error
	Delete(ctx context.Context, id int) error
}

type articleService struct {
	articleRepo repository.PostRepository
}

func NewArticleService(articleRepo repository.PostRepository) ArticleService {
	return &articleService{
		articleRepo: articleRepo,
	}
}

func (svc *articleService) Create(ctx context.Context, request *entity.CreateArticleRequest) error {
	article := new(model.Article)
	article.Title = request.Title
	article.Content = request.Content
	article.Status = request.Status
	article.Category = request.Category
	article.CreatedDate = sql.NullTime{
		Valid: true,
		Time:  time.Now(),
	}
	article.UpdatedDate = sql.NullTime{
		Valid: true,
		Time:  time.Now(),
	}

	err := svc.articleRepo.Insert(ctx, article)
	if err != nil {
		logrus.Errorf("[article-service] error svc.articleRepo.Insert, cause: %v\n", err)
	}
	return err
}

func (svc *articleService) FindAll(ctx context.Context, request *entity.GetArticlesRequest) (*entity.GetArticlesResponse, error) {
	articles, err := svc.articleRepo.GetAll(ctx, request.Limit, request.Offset)
	if err != nil {
		logrus.Errorf("[article-service] error svc.articleRepo.GetAll, cause: %v\n", err)
		return nil, err
	}

	resp := new(entity.GetArticlesResponse)
	for _, article := range articles {
		resp.Articles = append(resp.Articles, entity.ArticleResponse{
			Id:       int(article.Id),
			Title:    article.Title,
			Content:  article.Content,
			Category: article.Category,
			Status:   article.Status,
		})
	}

	return resp, nil
}

func (svc *articleService) FindById(ctx context.Context, id int) (*entity.ArticleResponse, error) {
	article, err := svc.articleRepo.FindById(ctx, id)
	if err != nil {
		logrus.Errorf("[article-service] error svc.articleRepo.FindById, cause: %v\n", err)
		return nil, err
	}

	articleResponse := new(entity.ArticleResponse)
	articleResponse.Id = int(article.Id)
	articleResponse.Title = article.Title
	articleResponse.Category = article.Category
	articleResponse.Status = article.Status
	articleResponse.Content = article.Content

	return articleResponse, nil
}

func (svc *articleService) Update(ctx context.Context, request *entity.UpdateArticleRequest) error {
	article, err := svc.articleRepo.FindById(ctx, request.Id)
	if err != nil {
		logrus.Errorf("[article-service] error svc.articleRepo.FindById, cause: %v\n", err)
		return err
	}
	article.Title = request.Title
	article.Content = request.Content
	article.Category = request.Category
	article.Status = request.Status
	article.UpdatedDate = sql.NullTime{
		Valid: true,
		Time:  time.Now(),
	}

	err = svc.articleRepo.UpdateById(ctx, article)
	if err != nil {
		logrus.Errorf("[article-service] error svc.articleRepo.UpdateById, cause: %v\n", err)
	}
	return err
}

func (svc *articleService) Delete(ctx context.Context, id int) error {
	err := svc.articleRepo.DeleteById(ctx, id)
	if err != nil {
		logrus.Errorf("[article-service] error delete article, cause: %v\n", err)
	}
	return err
}
