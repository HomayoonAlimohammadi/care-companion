package core

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/homayoonalimohammadi/care-companion/autogen/care_companion"
	"github.com/homayoonalimohammadi/care-companion/internal/app/providers"
)

type serviceImplementation struct {
	care_companion.UnimplementedCareCompanionServer
	careNeedProvider providers.CareNeedProvider
}

func New(
	careNeedProvider providers.CareNeedProvider,
) *serviceImplementation {
	return &serviceImplementation{
		careNeedProvider: careNeedProvider,
	}
}

func (s *serviceImplementation) GetCareSeeker(ctx context.Context,
	request *care_companion.GetCareSeekerRequest) (*care_companion.GetCareSeekerResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *serviceImplementation) CreateCareSeeker(ctx context.Context,
	request *care_companion.CreateCareSeekerRequest) (*care_companion.CreateCareSeekerResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *serviceImplementation) GetCareNeed(ctx context.Context,
	request *care_companion.GetCareNeedRequest) (*care_companion.GetCareNeedResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *serviceImplementation) CreateCareNeed(ctx context.Context,
	request *care_companion.CreateCareNeedRequest) (*care_companion.CreateCareNeedResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *serviceImplementation) Ping(ctx context.Context, request *care_companion.Empty) (*care_companion.PingResponse, error) {
	return &care_companion.PingResponse{
		Message: "I am alive!",
	}, nil
}
