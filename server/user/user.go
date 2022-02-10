package user

import "pipeliner/server/rules"

type User struct {
	Id    int
	Login string
	Name  string
	Rules rules.Rules
}
