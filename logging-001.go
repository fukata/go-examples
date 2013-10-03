package main

import (
    "github.com/op/go-logging"
)

func main() {
    logging.SetLevel(logging.INFO, "package.example")
    log := logging.MustGetLogger("package.example")
    log.Debug("hello %s", "golang")
    log.Info("hello %s", "golang")
}
