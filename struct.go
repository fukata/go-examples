package main

import (
    "fmt"
)

type Context struct {
    Param *Param
}

type Param struct {
    Id      int
}

func NewContext() *Context {
    ctx := &Context{}
    ctx.Param = &Param{}
    return ctx
}

func main() {
    ctx := NewContext()
    ctx.Param.Id = 100
    fmt.Println(ctx.Param.Id)
}
