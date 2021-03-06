// Code generated by gtrace. DO NOT EDIT.

package trace

import (
	"context"
	"time"
)

// Compose returns a new Trace which has functional fields composed
// both from t and x.
func (t Trace) Compose(x Trace) (ret Trace) {
	switch {
	case t.OnRefreshToken == nil:
		ret.OnRefreshToken = x.OnRefreshToken
	case x.OnRefreshToken == nil:
		ret.OnRefreshToken = t.OnRefreshToken
	default:
		h1 := t.OnRefreshToken
		h2 := x.OnRefreshToken
		ret.OnRefreshToken = func(r RefreshTokenStartInfo) func(RefreshTokenDoneInfo) {
			r1 := h1(r)
			r2 := h2(r)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(r RefreshTokenDoneInfo) {
					r1(r)
					r2(r)
				}
			}
		}
	}
	switch {
	case t.OnGetToken == nil:
		ret.OnGetToken = x.OnGetToken
	case x.OnGetToken == nil:
		ret.OnGetToken = t.OnGetToken
	default:
		h1 := t.OnGetToken
		h2 := x.OnGetToken
		ret.OnGetToken = func(g GetTokenStartInfo) func(GetTokenDoneInfo) {
			r1 := h1(g)
			r2 := h2(g)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(g GetTokenDoneInfo) {
					r1(g)
					r2(g)
				}
			}
		}
	}
	return ret
}
func (t Trace) onRefreshToken(r RefreshTokenStartInfo) func(RefreshTokenDoneInfo) {
	fn := t.OnRefreshToken
	if fn == nil {
		return func(RefreshTokenDoneInfo) {
			return
		}
	}
	res := fn(r)
	if res == nil {
		return func(RefreshTokenDoneInfo) {
			return
		}
	}
	return res
}
func (t Trace) onGetToken(g GetTokenStartInfo) func(GetTokenDoneInfo) {
	fn := t.OnGetToken
	if fn == nil {
		return func(GetTokenDoneInfo) {
			return
		}
	}
	res := fn(g)
	if res == nil {
		return func(GetTokenDoneInfo) {
			return
		}
	}
	return res
}
func TraceOnRefreshToken(t Trace, c *context.Context) func(token string, expiresIn time.Duration, _ error) {
	var p RefreshTokenStartInfo
	p.Context = c
	res := t.onRefreshToken(p)
	return func(token string, expiresIn time.Duration, e error) {
		var p RefreshTokenDoneInfo
		p.Token = token
		p.ExpiresIn = expiresIn
		p.Error = e
		res(p)
	}
}
func TraceOnGetToken(t Trace, c *context.Context) func(token string, _ error) {
	var p GetTokenStartInfo
	p.Context = c
	res := t.onGetToken(p)
	return func(token string, e error) {
		var p GetTokenDoneInfo
		p.Token = token
		p.Error = e
		res(p)
	}
}
