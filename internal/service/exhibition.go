package service

import (
	"context"

	pb "landscape/api/landscape/v1"
)

type ExhibitionService struct {
	pb.UnimplementedExhibitionServer
}

func NewExhibitionService() *ExhibitionService {
	return &ExhibitionService{}
}

func (s *ExhibitionService) CreateExhibition(ctx context.Context, req *pb.CreateExhibitionRequest) (*pb.CreateExhibitionReply, error) {
	return &pb.CreateExhibitionReply{}, nil
}
func (s *ExhibitionService) UpdateExhibition(ctx context.Context, req *pb.UpdateExhibitionRequest) (*pb.UpdateExhibitionReply, error) {
	return &pb.UpdateExhibitionReply{}, nil
}
func (s *ExhibitionService) DeleteExhibition(ctx context.Context, req *pb.DeleteExhibitionRequest) (*pb.DeleteExhibitionReply, error) {
	return &pb.DeleteExhibitionReply{}, nil
}
func (s *ExhibitionService) GetExhibition(ctx context.Context, req *pb.GetExhibitionRequest) (*pb.GetExhibitionReply, error) {
	return &pb.GetExhibitionReply{}, nil
}
func (s *ExhibitionService) ListExhibition(ctx context.Context, req *pb.ListExhibitionRequest) (*pb.ListExhibitionReply, error) {
	return &pb.ListExhibitionReply{}, nil
}
