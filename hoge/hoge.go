package hoge

import (
    "errors"
)

var UNIMPLEMENTED_ERROR = errors.New("UNIMPLEMENTED")

type Hoge interface {
    GetId() int
    Name() (string, error)
}

func DefaultName() (string, error) {
    return "", nil
}
