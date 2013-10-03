package main

import (
    "./hoge"
    "./hoge/foo"
    "fmt"
)

func main() {
    bars := make([]hoge.Hoge, 1)
    bars[0] = foo.New()

    for i, b := range bars {
        fmt.Printf("%d, Id: %d\n", i, b.GetId())
        name, name_err := b.Name()
        if name_err == nil {
            fmt.Printf("%d, Name: %s\n", i, name)
        }
    }
}
