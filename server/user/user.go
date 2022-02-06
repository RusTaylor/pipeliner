package user

import "pipeliner/server/rules"

type User struct {
	Name  string
	Login string
	Rules rules.Rules
}
