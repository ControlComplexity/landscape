package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Landscape is a Landscape model.
type Landscape struct {
	Hello string
}

// LandscapeRepo is a Landscape repo.
type LandscapeRepo interface {
	Save(context.Context, *Landscape) (*Landscape, error)
	Update(context.Context, *Landscape) (*Landscape, error)
	FindByID(context.Context, int64) (*Landscape, error)
	ListByHello(context.Context, string) ([]*Landscape, error)
	ListAll(context.Context) ([]*Landscape, error)
}

// LandscapeUsecase is a Greeter usecase.
type LandscapeUsecase struct {
	repo LandscapeRepo
	log  *log.Helper
}

// NewLandscapeUsecase new a Landscape usecase.
func NewLandscapeUsecase(repo LandscapeRepo, logger log.Logger) *LandscapeUsecase {
	return &LandscapeUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateLandscape creates a Landscape, and returns the new Landscape.
func (uc *LandscapeUsecase) CreateLandscape(ctx context.Context, g *Landscape) (*Landscape, error) {
	uc.log.WithContext(ctx).Infof("CreateLandscape: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
