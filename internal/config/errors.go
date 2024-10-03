package config

import "errors"

var (
	ErrNoKey            = errors.New("no such key in config")
	ErrInvalidUnmarshal = errors.New("123")
)
