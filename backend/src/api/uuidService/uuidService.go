package uuidService

import (
	"github.com/google/uuid"
)

func newUuid() (uuid.UUID, error) {
	return uuid.NewV7()
}

func NewClientUuid() (uuid.UUID, error) {
	return newUuid()
}

func NewRequestUuid() (uuid.UUID, error) {
	return newUuid()
}
