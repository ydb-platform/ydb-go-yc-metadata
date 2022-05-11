package yc

import (
	"github.com/ydb-platform/ydb-go-sdk/v3"

	"github.com/ydb-platform/ydb-go-yc-metadata/internal/auth"
	"github.com/ydb-platform/ydb-go-yc-metadata/internal/pem"
	"github.com/ydb-platform/ydb-go-yc-metadata/trace"
)

func WithURL(url string) auth.InstanceServiceAccountCredentialsOption {
	return auth.WithInstanceServiceAccountURL(url)
}

func WithCredentials(opts ...auth.InstanceServiceAccountCredentialsOption) ydb.Option {
	return ydb.WithCredentials(
		NewInstanceServiceAccount(opts...),
	)
}

func WithInternalCA() ydb.Option {
	return ydb.WithCertificatesFromPem(pem.YcPEM)
}

func WithTrace(t trace.Trace) auth.InstanceServiceAccountCredentialsOption {
	return auth.WithTrace(t)
}
