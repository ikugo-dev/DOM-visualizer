package main

import (
	"io"
	"log"
	"net/http"

	"github.com/m1gwings/treedrawer/tree"
)

func parseDOM(_ string) string {
	t := tree.NewTree(tree.NodeInt64(5))

	t.AddChild(tree.NodeString("adding a string"))
	t.AddChild(tree.NodeInt64(42))
	t.AddChild(tree.NodeInt64(3))

	return t.String()
}

func parseDOMHandler(w http.ResponseWriter, r *http.Request) {
	// Read the body of the request (plain text)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	// Convert the body to a string
	inputText := string(body)

	// Call your Go function to process the text
	// parsedText := parseDOM(inputText)
	parsedText := inputText

	// Send the result back as plain text
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(parsedText))
}

func main() {
	const portName = ":6969"
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/parse-text", parseDOMHandler)

	log.Printf("Starting server on :%s", portName)
	http.ListenAndServe(portName, nil)
}
