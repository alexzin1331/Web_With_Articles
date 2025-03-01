package cmd

import (
	"context"
	"database/sql"
	"google.golang.org/grpc"
	"log"
	"mod1/internal/repository"
	serv "mod1/internal/server"
	"mod1/internal/services"
	sso "mod1/protos/proto/gen"
	"net"
	"net/http"
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

	log.Println("Server is running on port 50051")
	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
