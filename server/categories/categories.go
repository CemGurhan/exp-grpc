package categories

import (
	"context"

	pb "github.com/cemgurhan/exp-grpc/gen/go"
)

type CategoryServer struct {
	pb.UnimplementedCategoryServer
}

func (c *CategoryServer) SearchCategory(ctx context.Context, seachCategoryRequest *pb.SearchCategoryRequest) (*pb.SearchCategoryResponse, error) {
	return &pb.SearchCategoryResponse{
		Value: "/v1/categories/search was selected",
	}, nil
}

func (c *CategoryServer) GetCategory(ctx context.Context, getCategoryRequest *pb.GetCategoryRequest) (*pb.GetCategoryResponse, error) {
	return &pb.GetCategoryResponse{
		Value: "v1/categories/{id} has been selected",
	}, nil
}
