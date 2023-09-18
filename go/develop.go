package main

import (
	"fmt"
	"os"
	"time"

	"os/signal"
)

func main() {
	// Crea un canale os.Signal
	signalChan := make(chan os.Signal, 1)

	// Registra un handler per il segnale SIGINT
	signal.Notify(signalChan, os.Interrupt)

	// Inizia una goroutine che rimane in running
	go func() {
		// Fa qualcosa di continuo
		for {
			fmt.Println("In esecuzione...")
			time.Sleep(1 * time.Second)
		}
	}()

	// Attende un segnale
	select {
	case <-signalChan:
		fmt.Println("Fine")
	}
}
