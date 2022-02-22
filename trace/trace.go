package trace

import "time"

// tool gtrace used from repository github.com/asmyasnikov/cmd/gtrace

//go:generate gtrace

//gtrace:gen
//gtrace:set Shortcut
type Trace struct {
	OnRefreshToken func(RefreshTokenStartInfo) func(RefreshTokenDoneInfo)
	OnGetToken     func(GetTokenStartInfo) func(GetTokenDoneInfo)
}

type (
	RefreshTokenStartInfo struct{}
	RefreshTokenDoneInfo  struct {
		Token     string
		ExpiresIn time.Duration
		Error     error
	}
	GetTokenStartInfo struct{}
	GetTokenDoneInfo  struct {
		Token string
		Error error
	}
)
