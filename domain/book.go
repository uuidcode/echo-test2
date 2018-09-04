package domain

import "time"

type Book struct {
	BookId      int64
	UserId      int64
	Name        string
	RegDatetime time.Time
	ModDatetime time.Time
}
