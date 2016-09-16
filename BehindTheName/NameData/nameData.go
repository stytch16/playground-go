// Package nameData stores data of user's name
package NameData

const (
	URL = "http://www.behindthename.com/name/"
)

type name struct {
	theName string

	gender  string
	usage   string
	meaning string
}

var Name name

const (
	nFields = 3 /* gender, usage, and meaning */
)
