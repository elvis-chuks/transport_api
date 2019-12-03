package handlers

import (
	"log"
	"os"
	"fmt"
)

// improve logic to accept request status and print accordingly

func Logger(w string){
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    log.SetOutput(file)
    log.Print(fmt.Sprintf("%s",w))
}