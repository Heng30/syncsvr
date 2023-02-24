package main

import (
    "log"
)

func setupLogger() {
    log.SetFlags(log.Lshortfile | log.LstdFlags)
}
