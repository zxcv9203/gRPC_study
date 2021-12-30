package service

import (
	"context"

	"github.com/gofrs/uuid"
	pb "github.com/zxcv9203/grpc_study/ecommerce"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server는 ecommerce/product_info를 구현하는데 사용합니다.
type Server struct {
	productMap map[string]*pb.Product
}

// AddProduct는 ecommerce.AddProduct를 구현합니다.
func (s *Server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while generating Product ID", err)
	}
	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}
	s.productMap[in.Id] = in
	return &pb.ProductID{Value: in.Id}, status.New(codes.OK, "").Err()
}

// GetProduct는 ecommerce.GetProduct를 구현합니다.
func (s *Server) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	value, exists := s.productMap[in.Value]
	if exists {
		return value, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Product does not exist.", in.Value)
}
