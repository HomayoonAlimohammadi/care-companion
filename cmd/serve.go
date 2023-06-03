package main

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/homayoonalimohammadi/care-companion/internal/app/core"
	care_companion "github.com/homayoonalimohammadi/care-companion/proto"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start server",
	Run:   serve,
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve(cmd *cobra.Command, args []string) {
	config := loadConfigOrPanic(cmd)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	service := core.New()
	care_companion.RegisterCareCompanionServer(grpcServer, service)
	reflection.Register(grpcServer)

	log.Printf("starting gRPC server on port %d", config.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func loadConfigOrPanic(cmd *cobra.Command) *Config {
	config, err := LoadConfig(cmd)
	if err != nil {
		panic("failed to load configurations")
	}
	return config
}
