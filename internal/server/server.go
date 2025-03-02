package server

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	serv "mod1/internal/services"
	//sso "mod1/protos/proto/gen"
	sso "github.com/alexzin1331/Web_With_Articles/protos/proto/gen"
	//sso1 "github"
	"sync"
)

type WebService interface {
	sso.UnimplementedWebServiceServer
	Register(ctx context.Context, req *sso.RegisterRequest) (*sso.RegisterResponse, error)
	Login(ctx context.Context, req *sso.LoginRequest) (*sso.LoginResponse, error)
	CreateArticle(ctx context.Context, req *sso.Article) (*sso.CreateArticleResponse, error)
	AddComment(ctx context.Context, req *sso.Comment) (*sso.AddCommentResponse, error)
	GetArticles(req *sso.Empty, stream grpc.ServerStreamingServer[sso.Article]) error
}

type WebServiceServer struct {
	sso.UnimplementedWebServiceServer
	Service *serv.Service
	mu      sync.Mutex
}

func NewWebServer(service *serv.Service) *WebServiceServer {
	return &WebServiceServer{Service: service}
}

func (s *WebServiceServer) Register(ctx context.Context, req *sso.RegisterRequest) (*sso.RegisterResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if req.Email == "" {
		return nil, status.Error(codes.FailedPrecondition, "Email is empty")
	}
	if req.Password == "" {
		return nil, status.Error(codes.FailedPrecondition, "Password is empty")
	}
	uid, err := s.Service.Register(ctx, req.Username, req.Email, req.Password)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}
	return &sso.RegisterResponse{UserId: uid}, nil
}

func (s *WebServiceServer) Login(ctx context.Context, req *sso.LoginRequest) (*sso.LoginResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}
	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}
	if req.AddId == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	token, err := s.Service.Login(ctx, req.GetEmail(), req.GetPassword(), int(req.GetAddId()))
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, "")
	}
	return &sso.LoginResponse{Token: token}, nil
}

func (s *WebServiceServer) CreateArticle(ctx context.Context, req *sso.Article) (*sso.CreateArticleResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	err := s.Service.CreateArticle(ctx, req.GetTitle(), req.GetContent(), req.GetUsername())
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}
	return &sso.CreateArticleResponse{
		Message: "Article is created: " + req.GetTitle(),
	}, nil
}

func (s *WebServiceServer) AddComment(ctx context.Context, req *sso.Comment) (*sso.AddCommentResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	err := s.Service.AddComment(ctx, req.GetArticleId(), req.GetContent(), req.GetUsername())
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}
	return &sso.AddCommentResponse{
		Message: "Comment is added: " + req.GetContent(),
	}, nil
}

func (s *WebServiceServer) GetArticles(req *sso.Empty, stream grpc.ServerStreamingServer[sso.Article]) error {
	articles, err := s.Service.GetArticles()
	if err != nil {
		log.Printf("Failed to get articles: %v", err)
		return status.Errorf(codes.Internal, "Failed to retrieve articles")
	}
	for _, article := range articles {
		if err := stream.Send(&sso.Article{
			Id:       article.Id,
			Title:    article.Title,
			Content:  article.Content,
			Username: article.Username,
		}); err != nil {
			log.Printf("Failed to send article: %v", err)
			return status.Errorf(codes.Internal, "Failed to send article")
		}
	}
	return nil
}
