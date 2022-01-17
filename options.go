package yc

import (
	"context"
	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-yc-metadata/internal/auth"
	"github.com/ydb-platform/ydb-go-yc-metadata/internal/pem"
)

func WithURL(url string) auth.InstanceServiceAccountCredentialsOption {
	return auth.WithInstanceServiceAccountURL(url)
}

func WithCredentials(ctx context.Context, opts ...auth.InstanceServiceAccountCredentialsOption) ydb.Option {
	return ydb.WithCredentials(
		NewInstanceServiceAccount(ctx, opts...),
	)
}

func WithInternalCA() ydb.Option {
	return ydb.WithCertificatesFromPem(pem.YcPEM)
}
