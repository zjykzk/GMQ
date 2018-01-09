package store

import "errors"

var (
	ErrBadPath = errors.New("bad path")
	ErrBadSize = errors.New("bad size")
)
