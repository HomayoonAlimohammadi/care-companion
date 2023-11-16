package core

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/homayoonalimohammadi/care-companion/autogen/care_companion"
	"github.com/homayoonalimohammadi/care-companion/internal/app/providers"
)

type serviceImplementation struct {
	care_companion.UnimplementedCareCompanionServer
	careSeekerProvider providers.CareSeeker
	careSeekProvider   providers.CareSeek
}

func New(
	careSeekerProvider providers.CareSeeker,
	careSeekProvider providers.CareSeek,
) *serviceImplementation {
	return &serviceImplementation{
		careSeekerProvider: careSeekerProvider,
		careSeekProvider:   careSeekProvider,
	}
}

func (s *serviceImplementation) GetCareSeeker(ctx context.Context,
	request *care_companion.GetCareSeekerRequest) (*care_companion.GetCareSeekerResponse, error) {
	seeker, err := s.careSeekerProvider.Get(request.GetId())
	if err != nil {
		return nil, status.Error(
			codes.Internal,
			fmt.Errorf("failed to get careSeeker: %w", err).Error(),
		)
	}

	return &care_companion.GetCareSeekerResponse{
		CareSeeker: &care_companion.CareSeeker{
			Id:   seeker.ID,
			Name: seeker.Name,
		},
	}, nil
}

func (s *serviceImplementation) CreateCareSeeker(ctx context.Context,
	request *care_companion.CreateCareSeekerRequest) (*care_companion.CreateCareSeekerResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *serviceImplementation) GetCareSeek(ctx context.Context,
	request *care_companion.GetCareSeekRequest) (*care_companion.GetCareSeekResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *serviceImplementation) CreateCareSeek(ctx context.Context,
	request *care_companion.CreateCareSeekRequest) (*care_companion.CreateCareSeekResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *serviceImplementation) Ping(ctx context.Context, request *care_companion.Empty) (*care_companion.PingResponse, error) {
	return &care_companion.PingResponse{
		Message: "I am alive!",
	}, nil
}
