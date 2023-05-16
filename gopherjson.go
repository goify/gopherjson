package gopherjson

import "time"

func (cd CustomDate) Serialize() interface{} {
	return cd.Format(time.RFC3339)
}
