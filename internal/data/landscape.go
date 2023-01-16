package data

import (
	"context"

	"landscape/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type essayRepo struct {
	data *Data
	log  *log.Helper
}

// NewEssayRepo .
func NewEssayRepo(data *Data, logger log.Logger) biz.EssayRepo {
	return &essayRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *essayRepo) Save(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

func (r *essayRepo) Update(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

func (r *essayRepo) FindByID(context.Context, int64) (*biz.Greeter, error) {
	return nil, nil
}

func (r *essayRepo) ListByHello(context.Context, string) ([]*biz.Greeter, error) {
	return nil, nil
}

func (r *essayRepo) ListAll(context.Context) ([]*biz.Greeter, error) {
	return nil, nil
}
