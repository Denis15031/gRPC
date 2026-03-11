package server

import (
	"context"
	"log"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"grpc-stub/api"
)

// UserServer реализует UserService (stub-логика)
type UserServer struct {
	api.UnimplementedUserServiceServer
	nextID int64
}

// ChatServer реализует ChatService (stub-логика)
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

// user service

func (s *UserServer) CreateUser(ctx context.Context, req *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	log.Printf("[UserService.CreateUser] name=%s, email=%s, role=%s",
		req.GetName(), req.GetEmail(), req.GetRole().String())
	log.Printf("  password_confirm=%s", req.GetPasswordConfirm())

	id := s.nextID
	s.nextID++

	return &api.CreateUserResponse{Id: id}, nil
}

func (s *UserServer) GetUser(ctx context.Context, req *api.GetUserRequest) (*api.GetUserResponse, error) {
	log.Printf("[UserService.GetUser] id=%d", req.GetId())

	// GetUserRequest имеет только Id, поэтому возвращаем заглушку
	return &api.GetUserResponse{
		Id:        req.GetId(),
		Name:      "John Doe",
		Email:     "john@example.com",
		Role:      api.Role_ROLE_USER,
		CreatedAt: timestamppb.New(time.Now().Add(-24 * time.Hour)),
		UpdatedAt: timestamppb.Now(),
	}, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, req *api.UpdateUserRequest) (*emptypb.Empty, error) {
	log.Printf("[UserService.UpdateUser] id=%d", req.GetId())
	if req.GetName() != nil {
		log.Printf("  name -> %s", req.GetName().GetValue())
	}
	if req.GetEmail() != nil {
		log.Printf("  email -> %s", req.GetEmail().GetValue())
	}
	return &emptypb.Empty{}, nil
}

func (s *UserServer) DeleteUser(ctx context.Context, req *api.DeleteUserRequest) (*emptypb.Empty, error) {
	log.Printf("[UserService.DeleteUser] id=%d", req.GetId())
	return &emptypb.Empty{}, nil
}

// chat service

func (s *ChatServer) CreateChat(ctx context.Context, req *api.CreateChatRequest) (*api.CreateChatResponse, error) {
	log.Printf("[ChatService.CreateChat] usernames=%v", req.GetUsernames())

	id := s.nextID
	s.nextID++

	return &api.CreateChatResponse{Id: id}, nil
}

func (s *ChatServer) DeleteChat(ctx context.Context, req *api.DeleteChatRequest) (*emptypb.Empty, error) {
	log.Printf("[ChatService.DeleteChat] id=%d", req.GetId())
	return &emptypb.Empty{}, nil
}

func (s *ChatServer) SendMessage(ctx context.Context, req *api.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("[ChatService.SendMessage] from=%s, text=%q", req.GetFrom(), req.GetText())
	if req.GetTimestamp() != nil {
		log.Printf("  timestamp=%s", req.GetTimestamp().AsTime().Format(time.RFC3339))
	}
	return &emptypb.Empty{}, nil
}
