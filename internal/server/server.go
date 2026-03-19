package server

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"grpc-stub/api"
)

type UserServer struct {
	api.UnimplementedUserServiceServer
	nextID int64
}

type ChatServer struct {
	api.UnimplementedChatServiceServer
	nextID int64
}

func NewUserServer() *UserServer {
	return &UserServer{nextID: 1}
}

func NewChatServer() *ChatServer {
	return &ChatServer{nextID: 1}
}

func (s *UserServer) CreateUser(ctx context.Context, req *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	log.Printf("[CreateUser] name=%s, email=%s, role=%s",
		req.Name, req.Email, req.Role.String())

	// Валидация роли
	if req.Role == api.Role_ROLE_UNSPECIFIED {
		return nil, status.Error(codes.InvalidArgument, "role is required")
	}

	id := s.nextID
	s.nextID++
	return &api.CreateUserResponse{Id: id}, nil
}

func (s *UserServer) GetUser(ctx context.Context, req *api.GetUserRequest) (*api.GetUserResponse, error) {
	// Прямой доступ: req.Id
	log.Printf("[GetUser] id=%d", req.Id)

	return &api.GetUserResponse{
		Id:        req.Id,
		Name:      "John Doe",
		Email:     "john@example.com",
		Role:      api.Role_ROLE_USER,
		CreatedAt: timestamppb.New(time.Now().Add(-24 * time.Hour)),
		UpdatedAt: timestamppb.Now(),
	}, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, req *api.UpdateUserRequest) (*emptypb.Empty, error) {
	log.Printf("[UpdateUser] id=%d", req.Id)

	// Если в proto поля — обёртки (StringValue), проверяем != nil
	// Если обычные string — просто req.Name / req.Email
	if req.Name != nil && req.Name.Value != "" {
		log.Printf("  name -> %s", req.Name)
	}
	if req.Email != nil && req.Email.Value != "" {
		log.Printf("  email -> %s", req.Email)
	}
	return &emptypb.Empty{}, nil
}

func (s *UserServer) DeleteUser(ctx context.Context, req *api.DeleteUserRequest) (*emptypb.Empty, error) {
	log.Printf("[DeleteUser] id=%d", req.Id)
	return &emptypb.Empty{}, nil
}

func (s *ChatServer) CreateChat(ctx context.Context, req *api.CreateChatRequest) (*api.CreateChatResponse, error) {
	log.Printf("[CreateChat] usernames=%v", req.Usernames)

	id := s.nextID
	s.nextID++
	return &api.CreateChatResponse{Id: id}, nil
}

func (s *ChatServer) DeleteChat(ctx context.Context, req *api.DeleteChatRequest) (*emptypb.Empty, error) {
	log.Printf("[DeleteChat] id=%d", req.Id)
	return &emptypb.Empty{}, nil
}

func (s *ChatServer) SendMessage(ctx context.Context, req *api.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("[SendMessage] from=%s, text=%q", req.From, req.Text)
	if req.Timestamp != nil {
		log.Printf("  timestamp=%s", req.Timestamp.AsTime().Format(time.RFC3339))
	}
	return &emptypb.Empty{}, nil
}
