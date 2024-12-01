package decorator

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

type commandHttpServeDecorator[C any] struct {
	base   CommandHandler[C]
	logger *logrus.Entry
}

func (d commandHttpServeDecorator[C]) Handle(ctx context.Context, cmd C) (err error) {

	return d.base.Handle(ctx, cmd)
}

type queryHttpServeDecorator[C any, R any] struct {
	base   QueryHandler[C, R]
	logger *logrus.Entry
}

func (d queryHttpServeDecorator[C, R]) Handle(ctx context.Context, cmd C) (result R, err error) {

	return d.base.Handle(ctx, cmd)
}
