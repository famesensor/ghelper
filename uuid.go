package ghelper

import "github.com/google/uuid"

type Uuider interface {
	New() string
}

type uuids struct{}

func NewUUID() Uuider {
	return &uuids{}
}

func (u *uuids) New() string {
	return uuid.NewString()
}
