package services

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	article "mod1/internal/models"
	rep "mod1/internal/repository"
	"time"
)

/*type WebService interface {
	sso.UnimplementedWebServiceServer
	Register(ctx context.Context, req *sso.RegisterRequest) (*sso.RegisterResponse, error)
	Login(ctx context.Context, req *sso.LoginRequest) (*sso.LoginResponse, error)
	CreateArticle(ctx context.Context, req sso.Article) (*sso.CreateArticleResponse, error)
	AddComment(ctx context.Context, req *sso.Comment) (*sso.AddCommentResponse, error)
	GetArticles(ctx context.Context, req *sso.Empty) (*[]sso.Article, error)
}*/

type Service struct {
	repo *rep.Repository
	//w WebService
}

func NewService(repo *rep.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Register(ctx context.Context, username, email, password string) (int64, error) {
	return s.repo.Register(ctx, username, email, password)
}

func (s *Service) Login(ctx context.Context, email string, password string, appid int) (string, error) {
	check, err := s.repo.Login(ctx, email, password)
	if err != nil {
		return "", err
	}
	if !check {
		return "", errors.New("Invalid credentials")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte("secret"))
}

func (s *Service) CreateArticle(ctx context.Context, title string, content string, username string) error {
	return s.repo.CreateArticle(ctx, title, content, username)
}

func (s *Service) GetArticles() ([]article.Article, error) {
	return s.repo.GetArticles()
}

func (s *Service) AddComment(ctx context.Context, articleID int64, content, username string) error {
	return s.repo.AddComment(ctx, articleID, content, username)
}
