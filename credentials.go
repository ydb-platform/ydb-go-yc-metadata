package yc

import (
	"context"

	"github.com/ydb-platform/ydb-go-sdk/v3/credentials"
	"github.com/ydb-platform/ydb-go-yc/internal/auth"
)

func NewInstanceServiceAccount(ctx context.Context, opts ...auth.InstanceServiceAccountCredentialsOption) credentials.Credentials {
	return auth.InstanceServiceAccount(
		ctx,
		append(
			[]auth.InstanceServiceAccountCredentialsOption{
				auth.WithInstanceServiceAccountCredentialsSourceInfo("yc.NewInstanceServiceAccount(ctx)"),
			},
			opts...,
		)...,
	)
}
