package gopherjson

import "time"

func (cd CustomDate) Serialize() interface{} {
	return cd.Format(time.RFC3339)
}

func (cr CustomRegex) Serialize() interface{} {
	return cr.String()
}
