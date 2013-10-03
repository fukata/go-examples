package main

import (
    "fmt"
)

func main() {
    id_map := make(map[int]bool, 2)
    id_map[1] = true
    id_map[3] = true

    for i := 0; i < 5; i++ {
        if id_map[i] == false {
            fmt.Printf("%d: false\n", i)
        }
    }
}
