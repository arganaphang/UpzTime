package sendz

import "context"

type Provider interface {
	Send(ctx context.Context, err error) error
}

type Sendz struct {
	provider Provider
}

func New(provider Provider) *Sendz {
	return &Sendz{provider: provider}
}

func (c *Sendz) IsFeatureEnabled(ctx context.Context, err error) error {
	return c.provider.Send(ctx, err)
}
