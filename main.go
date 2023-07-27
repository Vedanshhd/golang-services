package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/voices"
)

type TextToSpeechRequest struct {
	Text string `json:"text"`
}

func textToSpeechHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req TextToSpeechRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	// Extract the "text" value from the request
	text := req.Text
	fmt.Println("Received text:", text)
	if text == "" {
		http.Error(w, "Text not provided", http.StatusBadRequest)
		return
	}

	// Create the speech object
	speech := htgotts.Speech{
		Folder:   "audioFile",
		Language: voices.English,
	}

	// Generate the speech audio
	err := speech.Speak(text)
	if err != nil {
		fmt.Println("Error generating speech audio:", err)
		// http.Error(w, "Failed to generate speech audio", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Speech audio generated successfully.")
}

func main() {
	http.HandleFunc("/text-to-speech", textToSpeechHandler)

	port := "8080"
	fmt.Printf("Starting Text-to-Speech API service on port %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Failed to start API service: %s\n", err)
	}
}
