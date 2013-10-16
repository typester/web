package cookie

import (
	"github.com/typester/web"
	"github.com/typester/web/session"
	"net/http"
	"time"
)

type CookieState struct {
	Name, Path, Domain string
	HttpOnly, Secure bool
	session.State
}

func (state *CookieState) GetSessionID(ctx *web.Context) (session_id string, err error) {
	cookie, err := ctx.Request.Cookie(state.Name)
	if err != nil {
		return
	}

	session_id = cookie.Value
	return
}

func (state *CookieState) SaveSessionID(ctx *web.Context, sd *session.SessionData) error {
	cookie := &http.Cookie{
		Name: state.Name,
		Value: sd.SessionID(),
	}

	if len(state.Path) > 0 {
		cookie.Path = state.Path
	}
	if len(state.Domain) > 0 {
		cookie.Domain = state.Domain
	}

	cookie.Expires = time.Now().Add( sd.Expires() )

	if state.HttpOnly {
		cookie.HttpOnly = true
	}
	if state.Secure {
		cookie.Secure = true
	}

	http.SetCookie(ctx, cookie)
		
	return nil
}

















