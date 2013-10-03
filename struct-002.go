package main

import "fmt"

type Dummy interface {
    GetName() string
}

type Dummy1 struct {}
type Dummy2 struct {}

func (d *Dummy1) GetName() string {
    return "dummy1"
}
func (d *Dummy2) GetName() string {
    return "dummy2"
}

func main() {
    d1 := &Dummy1{}
    d2 := &Dummy2{}

    fmt.Println(d1.GetName())
    fmt.Println(d2.GetName())
}
