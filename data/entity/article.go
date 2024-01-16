package entity

type CreateArticleRequest struct {
	Title    string `json:"title" validate:"required,min=20"`
	Content  string `json:"content" validate:"required,min=200"`
	Category string `json:"category" validate:"required,min=3"`
	Status   string `json:"status" validate:"required,oneof=Publish Draft Thrash"`
}

type GetArticlesRequest struct {
	Limit  int
	Offset int
}

type GetArticlesResponse struct {
	Articles []ArticleResponse `json:"articles"`
}

type ArticleResponse struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Status   string `json:"status"`
}

type UpdateArticleRequest struct {
	Id       int
	Title    string `json:"title" validate:"required,min=20"`
	Content  string `json:"content" validate:"required,min=200"`
	Category string `json:"category" validate:"required,min=3"`
	Status   string `json:"status" validate:"required,oneof=Publish Draft Thrash"`
}
