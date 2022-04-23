package session_buntdb

import (
	"github.com/chefsgo/session"
	sb "github.com/chefsgo/session-buntdb"
)

func init() {
	session.Register("file", sb.Driver("store/session.db"))
}
