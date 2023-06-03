package core

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	care_companion "github.com/homayoonalimohammadi/care-companion/proto"
)

type serviceImplementation struct {
	care_companion.UnimplementedCareCompanionServer
}

func New() *serviceImplementation {
	return &serviceImplementation{}
}

func (s *serviceImplementation) GetCareSeeker(ctx context.Context,
	request *care_companion.GetCareSeekerRequest) (*care_companion.GetCareSeekerResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *serviceImplementation) CreateCareSeeker(ctx context.Context,
	request *care_companion.CreateCareSeekerRequest) (*care_companion.CreateCareSeekerResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *serviceImplementation) Ping(ctx context.Context, request *care_companion.Empty) (*care_companion.PingResponse, error) {
	return &care_companion.PingResponse{
		Message: "I am alive!",
	}, nil
}
