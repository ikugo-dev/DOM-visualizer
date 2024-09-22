package main

import (
	"github.com/m1gwings/treedrawer/tree"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	const portName = ":6969"
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/parse-text", parseDOMHandler)

	log.Printf("Starting server on :%s", portName)
	http.ListenAndServe(portName, nil)
}

func parseDOMHandler(w http.ResponseWriter, r *http.Request) {
	// Read the body of the request (plain text as bytes)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	inputText := string(body)
	parsedText := parseDOM(inputText)

	// Send the result back as plain text
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(parsedText))
}

func parseDOM(inputText string) string {
	root := tree.NewTree(tree.NodeString("HTML"))

	text := strings.Split(inputText, "<")

	for _, line := range text {
		if line == " " { // Stupid edge case 1
			continue
		}
		textParts := strings.FieldsFunc(line, splitXML)

		if len(textParts) == 0 { // Stupid edge case 2
			continue
		}

		if string(textParts[0][0]) == "/" { // Go up the tree
			newRoot, _ := root.Parent()
			root = newRoot
		} else { // Add child and traverse the tree
			root.AddChild(tree.NodeString(textParts[0]))

			newRoot, err := root.Child(len(root.Children()) - 1)
			if err != nil {
				log.Fatalf("Failed to read child node of root: %s", root)
			}
			root = newRoot
		}

	}
	return root.String()
}

func splitXML(r rune) bool {
	return r == ' ' || r == '>'
}
