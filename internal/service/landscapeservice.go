package service

import (
	"context"

	pb "landscape/api/landscape/v1"
)

type LandscapeServiceService struct {
	pb.UnimplementedLandscapeServiceServer
}

func NewLandscapeServiceService() *LandscapeServiceService {
	return &LandscapeServiceService{}
}

func (s *LandscapeServiceService) QueryEssayList(ctx context.Context, req *pb.QueryEssayListReq) (*pb.QueryEssayListResp, error) {
	return &pb.QueryEssayListResp{}, nil
}
func (s *LandscapeServiceService) QueryRelatedEssayList(ctx context.Context, req *pb.QueryRelatedEssayListReq) (*pb.QueryRelatedEssayListResp, error) {
	return &pb.QueryRelatedEssayListResp{}, nil
}
func (s *LandscapeServiceService) QueryOneEssay(ctx context.Context, req *pb.QueryOneEssayReq) (*pb.QueryOneEssayResp, error) {
	return &pb.QueryOneEssayResp{}, nil
}
func (s *LandscapeServiceService) QueryCityList(ctx context.Context, req *pb.QueryCityListReq) (*pb.QueryCityListResp, error) {
	return &pb.QueryCityListResp{}, nil
}
func (s *LandscapeServiceService) QuerySwiperImageList(ctx context.Context, req *pb.QuerySwiperImageListReq) (*pb.QuerySwiperImageListResp, error) {
	return &pb.QuerySwiperImageListResp{}, nil
}
func (s *LandscapeServiceService) QueryRankList(ctx context.Context, req *pb.QueryRankListReq) (*pb.QueryRankListResp, error) {
	return &pb.QueryRankListResp{}, nil
}
func (s *LandscapeServiceService) QueryCityImage(ctx context.Context, req *pb.QueryCityImageReq) (*pb.QueryCityImageResp, error) {
	return &pb.QueryCityImageResp{}, nil
}
