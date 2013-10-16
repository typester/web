package session

import (
	"github.com/typester/web"
	uuid "github.com/nu7hatch/gouuid"
	"crypto/hmac"
	"crypto/sha256"
	"log"
	"encoding/hex"
	"time"
)

type Session struct {
	state State
	store Store
	secret string
	expires time.Duration
}

type SessionData struct {
	session *Session
	sessionId string
	data map[string]interface{}
	inStorage bool
}

type State interface {
	GetSessionID(c *web.Context) (string, error)
	SaveSessionID(c *web.Context, sd *SessionData) error
}

type Store interface {
	GetData(sessionId string) interface{}
	SetData(sessionId string, data map[string]interface{}, expires time.Duration)
}

var defaultSession *Session

func Setup(state State, store Store, secret string, expires time.Duration) {
	defaultSession = &Session{}
	defaultSession.state = state
	defaultSession.store = store
	defaultSession.secret = secret
	defaultSession.expires = expires
}

func Restore(c *web.Context) (*SessionData, error) {
	session_id, err := defaultSession.state.GetSessionID(c)

	inStorage := true
	if err != nil {
		session_id = newSessionID()
		inStorage = false
	}

	data := defaultSession.store.GetData(session_id)

	if m, ok := data.(map[string]interface{}); ok {
		return &SessionData{ defaultSession, session_id, m, inStorage }, nil
	} else {
		if inStorage {
			// re-generate cookie
			session_id = newSessionID()
		}

		return &SessionData{ defaultSession, session_id, map[string]interface{}{}, false }, nil
	}
}

func newSessionID() string {
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Panicf("Failed to create new session id: %s", err)
	}

	mac := hmac.New(sha256.New, []byte(defaultSession.secret))
	mac.Write(uuid[:])

	digest := mac.Sum(nil)
	
	return hex.EncodeToString(digest)
}

func (sd *SessionData) Save(ctx *web.Context) error {
	sd.session.state.SaveSessionID(ctx, sd)
	sd.session.store.SetData(sd.sessionId, sd.data, sd.session.expires)
	return nil
}

func (sd *SessionData) SessionID() string {
	return sd.sessionId
}

func (sd *SessionData) Expires() time.Duration {
	return sd.session.expires
}

func (sd *SessionData) Get(key string) interface{} {
	return sd.data[key]
}

func (sd *SessionData) Set(key string, value interface{}) {
	sd.data[key] = value
}

















