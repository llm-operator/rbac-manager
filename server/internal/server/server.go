package server

import (
	"context"
	"fmt"
	"net"

	v1 "github.com/llm-operator/rbac-manager/api/v1"
	"github.com/llm-operator/rbac-manager/server/internal/cache"
	"github.com/llm-operator/rbac-manager/server/internal/dex"
	"google.golang.org/grpc"
)

type cacheGetter interface {
	GetAPIKeyBySecret(secret string) (*cache.K, bool)

	GetClusterByRegistrationKey(key string) (*cache.C, bool)
	GetClustersByTenantID(tenantID string) []cache.C

	GetOrganizationByID(organizationID string) (*cache.O, bool)
	GetOrganizationsByUserID(userID string) []cache.OU

	GetProjectsByOrganizationID(organizationID string) []cache.P
	GetProjectByID(projectID string) (*cache.P, bool)
	GetProjectsByUserID(userID string) []cache.PU

	GetUserByID(id string) (*cache.U, bool)
}

type tokenIntrospector interface {
	TokenIntrospect(token string) (*dex.Introspection, error)
}

// New returns a new Server.
func New(issuerURL string, cache cacheGetter, roleScopes map[string][]string) *Server {
	return &Server{
		tokenIntrospector: dex.NewDefaultClient(issuerURL),

		cache: cache,

		roleScopesMapper: roleScopes,
	}
}

// Server implementes the gRPC interface.
type Server struct {
	v1.UnimplementedRbacInternalServiceServer

	tokenIntrospector tokenIntrospector

	cache cacheGetter

	roleScopesMapper map[string][]string
}

// Run starts the gRPC server.
func (s *Server) Run(ctx context.Context, port int) error {
	serv := grpc.NewServer()
	v1.RegisterRbacInternalServiceServer(serv, s)
	return listenAndServe(serv, port)
}

// listenAndServe is a helper function for starting a gRPC server.
func listenAndServe(grpcServer *grpc.Server, port int) error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	if err := grpcServer.Serve(l); err != nil {
		return fmt.Errorf("failed to start gRPC server: %v", err)
	}
	return nil
}
