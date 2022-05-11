package yc

import (
	"github.com/ydb-platform/ydb-go-yc-metadata/internal/auth"
)

func NewInstanceServiceAccount(opts ...auth.InstanceServiceAccountCredentialsOption) *auth.InstanceServiceAccountCredentials {
	return auth.InstanceServiceAccount(
		append(
			[]auth.InstanceServiceAccountCredentialsOption{
				auth.WithInstanceServiceAccountCredentialsSourceInfo("yc.NewInstanceServiceAccount(ctx)"),
			},
			opts...,
		)...,
	)
}
