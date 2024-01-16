package http

type ErrorType uint8

const (
	Default ErrorType = iota
	TypeBadRequest
	TypeNotFound
	TypeConflict
	TypeInternalServerError
	TypeUnprocessableEntity
)
