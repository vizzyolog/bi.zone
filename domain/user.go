package domain

import "time"

type User struct {
	UserID  int
	Created time.Time
	Role    string
	Confirm bool
}
