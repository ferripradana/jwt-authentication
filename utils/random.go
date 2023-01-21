package utils

import "github.com/google/uuid"

func Uuid() string {
	uid := uuid.New()
	return uid.String()
}
