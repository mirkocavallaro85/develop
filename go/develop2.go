package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Crea un handler che catturi tutte le chiamate al percorso /example
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ottieni l'ora corrente
		now := time.Now()

		// Crea un oggetto JSON con i dati della chiamata
		jsonData := map[string]interface{}{
			"timestamp": now.Format(time.RFC3339),
			"method":    r.Method,
			"url":       r.URL.String(),
			"headers":   r.Header,
			"body":      r.Body,
		}

		// Serializza l'oggetto JSON
		jsonBytes, err := json.Marshal(jsonData)
		if err != nil {
			// Gestisci l'errore
			fmt.Println(err)
			return
		}

		// Scrivi il JSON nella risposta
		fmt.Fprintf(w, "%s", jsonBytes)
	})

	// Avvia il server HTTP
	http.ListenAndServe(":8080", handler)
}
