package session

import (
	"github.com/satori/go.uuid"
	"sync"
	"time"
	"vido_server/api/dbops"
	"vido_server/api/defs"
)

var sessionMap *sync.Map

func init(){
	sessionMap = &sync.Map{}
}

func nowInMilli() int64{
	return time.Now().UnixNano() / 1000000
}

func deleteExpiredSession(sid string){
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

func LoadSessionsFromDB(){
	r, err := dbops.RetrieveAllSessions()
	if err != nil{
		return
	}
	r.Range(func(k, v interface{}) bool {
		ss := v.(defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})
}

func GenerateNewSessionId(un string) string{
	u1 := uuid.Must(uuid.NewV4())
	var id string
	for i := 0; i < len(u1); i++ {
		if u1[i] == 0 {
			id = string(u1[0:i])
		}
	}

	ct := nowInMilli()
	ttl := ct + 30 * 60 * 1000

	ss := &defs.SimpleSession{Username:un, TTL:ttl}
	sessionMap.Store(id, ss)
	dbops.InserSession(id, ttl, un)

	return id
}

func IsSessionExpired(sid string) (string, bool){
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := nowInMilli()
		if ss.(*defs.SimpleSession).TTL < ct{
			deleteExpiredSession(sid)
			return "", true
		}
		return ss.(*defs.SimpleSession).Username, false
	}

	return "", true
}