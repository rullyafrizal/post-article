package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/rullyafrizal/post-article-api/data/model"
	"gorm.io/gorm"
)

type PostRepository interface {
	GetAll(ctx context.Context, limit int, offset int) ([]model.Article, error)
	FindById(ctx context.Context, id int) (*model.Article, error)
	Insert(ctx context.Context, article *model.Article) error
	UpdateById(ctx context.Context, article *model.Article) error
	DeleteById(ctx context.Context, id int) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{
		db: db,
	}
}

func (r *postRepository) GetAll(ctx context.Context, limit int, offset int) ([]model.Article, error) {
	var articles []model.Article
	result := r.db.Limit(limit).Offset(offset).Find(&articles)

	if result.Error != nil {
		return nil, result.Error
	}

	return articles, nil
}

func (r *postRepository) FindById(ctx context.Context, id int) (*model.Article, error) {
	article := new(model.Article)
	article.Id = int64(id)

	res := r.db.WithContext(ctx).Unscoped().First(article)
	if res.RowsAffected == 0 {
		return nil, errors.New(fmt.Sprintf("article id %v not found", id))
	}
	return article, nil
}

func (r *postRepository) Insert(ctx context.Context, article *model.Article) error {
	err := r.db.WithContext(ctx).Create(&article).Error
	return err
}

func (r *postRepository) UpdateById(ctx context.Context, article *model.Article) error {
	res := r.db.WithContext(ctx).Where("id = ?", article.Id).Updates(&article)

	if res.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("article id %v not found", article.Id))
	}

	return nil
}

func (r *postRepository) DeleteById(ctx context.Context, id int) error {
	res := r.db.WithContext(ctx).Delete(&model.Article{}, id)

	if res.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("article id %v not found", id))
	}

	return nil
}
