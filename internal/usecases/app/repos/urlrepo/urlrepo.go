package urlrepo

import (
	"context"
	"fmt"
	"github.com/PudgeRo/LinkShortener/internal/entities/urlentity"
	"github.com/thanhpk/randstr"
)

type UrlStore interface {
	Create(ctx context.Context, u urlentity.Url) (string, error)
	Read(ctx context.Context, id string) (*urlentity.Url, error)
}

type Urls struct {
	ustore UrlStore
}

func NewUrl(ustore UrlStore) *Urls {
	return &Urls{
		ustore: ustore,
	}
}

func (url *Urls) Create(ctx context.Context, u urlentity.Url) (*urlentity.Url, error) {
	u.ID = randstr.String(8)
	id, err := url.ustore.Create(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("create url error: %w", err)
	}
	u.ID = id
	return &u, nil
}

func (url *Urls) Read(ctx context.Context, id string) (*urlentity.Url, error) {
	u, err := url.ustore.Read(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("read url error: %w", err)
	}
	return u, nil
}
