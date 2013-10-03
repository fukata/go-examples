package foo

import (
    "../../hoge"
)

const ID = 1

type Foo struct {
}

func New() *Foo {
    return &Foo{}
}

func (d *Foo) GetId() int {
    return ID
}

func (d *Foo) Name() (string, error) {
    return hoge.DefaultName()               // default
    return "FOO", hoge.UNIMPLEMENTED_ERROR  // optional
}
