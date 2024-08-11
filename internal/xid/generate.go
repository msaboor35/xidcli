package xid

import "github.com/rs/xid"

func GenerateXID() string {
	xid := xid.New()
	return xid.String()
}
