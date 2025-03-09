package main

import (
	"context"
	"database/sql"
	sso "github.com/alexzin1331/Web_With_Articles/protos/proto/gen"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"mod1/internal/repository"
	serv "mod1/internal/server"
	"mod1/internal/services"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

/*
для генерации по proto-файлам из корневой директории проекта:
protoc -I=C:\Users\DELL\GolandProjects\Web_With_Artices\protos\sso --go-grpc_out=./protos/proto/gen --go-grpc_opt=paths=source_relative C:\Users\DELL\GolandProjects\Web_With_Artices\protos\sso\sso.proto
protoc -I=C:\Users\DELL\GolandProjects\Web_With_Artices\protos\sso --go_out=./protos/proto/gen --go_opt=paths=source_relative C:\Users\DELL\GolandProjects\Web_With_Artices\protos\sso\sso.proto

*/

type Server struct {
	httpserver *http.Server
}

func (s *Server) Run(port string) error {
	s.httpserver = &http.Server{
		Addr:           port,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpserver.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) {
	//return s.httpserver.Shutdown()
}

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:alex1234@127.0.0.1:5436/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	repo := repository.New(db)
	service := services.NewService(repo)
	server := serv.NewWebServer(service)
	grpcServer := grpc.NewServer()
	sso.RegisterWebServiceServer(grpcServer, server)

	list, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("Server is running on port 8080")
	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Shutting down server...")

}

/*
вроде +- исправил косяки. осталось допилить след три туду и можно писать фронт.
TODO: разобраться с конфигами
TODO: доделать изящную остановку
TODO: доделать delete articles
TODO: разобраться с jwt-токенами (все ли ок с ними?)
TODO: попросить у гпт написать простой фронт для тестирования

TODO: как доп задание написать тесты (юнит и функциональные)
*/

/*
middleware для валидации токена:

import (
    "context"
    "fmt"
    "github.com/golang-jwt/jwt/v5"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/metadata"
    "google.golang.org/grpc/status"
)

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
    // Пропускаем аутентификацию для методов Register и Login
    if info.FullMethod == "/auth.WebService/Register" || info.FullMethod == "/auth.WebService/Login" {
        return handler(ctx, req)
    }

    // Извлекаем токен из метаданных
    md, ok := metadata.FromIncomingContext(ctx)
    if !ok {
        return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
    }

    token := md["authorization"]
    if len(token) == 0 {
        return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
    }

    // Парсим и валидируем токен
    claims, err := parseToken(token[0])
    if err != nil {
        return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
    }

    // Добавляем claims в контекст
    ctx = context.WithValue(ctx, "claims", claims)
    return handler(ctx, req)
}

func parseToken(tokenString string) (jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(os.Getenv("JWT_SECRET")), nil
    })
    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    }
    return nil, fmt.Errorf("invalid token")
}*/
