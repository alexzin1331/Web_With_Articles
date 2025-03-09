package repository

import (
	"context"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	article "mod1/internal/models"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Register(ctx context.Context, username, email, password string) (int64, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}
	var id int64
	err = r.db.QueryRow("INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id", username, email, hashPass).Scan(&id)
	if err != nil {
		return -1, status.Errorf(codes.AlreadyExists, "User already exists")
	}
	return id, nil
}

func (r *Repository) Login(ctx context.Context, email, password string) (bool, error) {
	row := r.db.QueryRow("SELECT password FROM users WHERE email = $1", email)
	hashdbPass := ""
	err := row.Scan(&hashdbPass)
	if err != nil {
		return false, status.Errorf(codes.NotFound, "User not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashdbPass), []byte(password))
	if err != nil {
		return false, status.Errorf(codes.Unauthenticated, "Invalid password")
	}
	return true, nil
}

func (r *Repository) CreateArticle(ctx context.Context, title, content, username string) error {
	_, err := r.db.Exec("INSERT INTO artiсles (title , content, username) VALUES ($1, $2, $3)", title, content, username)
	return err
}

func (r *Repository) AddComment(ctx context.Context, articleID int64, content, username string) error {
	_, err := r.db.Exec("INSERT INTO comments (article_id, content, username) VALUES ($1, $2, $3)", articleID, content, username)
	return err
}

func (r *Repository) GetArticles() ([]article.Article, error) {
	rows, err := r.db.Query("SELECT id, title, content, username FROM articles")
	if err != nil {
		return nil, status.Error(codes.NotFound, "Articles not found")
	}
	defer rows.Close()
	var articles []article.Article
	for rows.Next() {
		var article article.Article
		if err := rows.Scan(&article.Id, &article.Title, &article.Content, &article.Username); err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func (r *Repository) DeleteArticle(ctx context.Context, articleID int64) error {
	_, err := r.db.Exec("DELETE FROM articles WHERE id = $1", articleID)
	if err != nil {
		return status.Errorf(codes.NotFound, "Article not found")
	}
	return nil
}
