package uploadservice

import (
	"context"
	"net/http"

	"github.com/ServiceWeaver/weaver"
)

type T interface {
	RegisterRoutes(context.Context) error
}

type impl struct {
	weaver.Implements[T]
}

func (i *impl) RegisterRoutes(ctx context.Context) error {
	log := i.Logger()
	for _, t := range RouteList {
		http.HandleFunc(t.path, t.handler)
		log.Info("Register Path(" + t.path + ") success")
	}

	return nil
}
