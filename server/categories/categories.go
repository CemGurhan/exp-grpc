package categories

import (
	"context"
	pb "exp-grpc/gen/go"
	"fmt"
)

type CategoryServer struct {
	pb.UnimplementedCategoryServer
}

func (c *CategoryServer) SearchCategory(ctx context.Context, seachCategoryRequest *pb.SearchCategoryRequest) (*pb.SearchCategoryResponse, error) {
	return &pb.SearchCategoryResponse{
		Value: fmt.Sprintf("This has no path parameters. You're request was: %s",
			seachCategoryRequest.Value),
	}, nil
}

func (c *CategoryServer) GetCategory(ctx context.Context, getCategoryRequest *pb.GetCategoryRequest) (*pb.GetCategoryResponse, error) {
	return &pb.GetCategoryResponse{
		Value: fmt.Sprintf("Path parameter provided: %s", getCategoryRequest.Id),
	}, nil
}
