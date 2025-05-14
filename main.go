package main

import (
	"io"
	"log"
	"net/http"
	"regexp"

	"github.com/m1gwings/treedrawer/tree"
)

func main() {
	const portName = ":8080"
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/parse-text", parseDOMHandler)

	log.Printf("Starting server on %s", portName)
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
	root := tree.NewTree(tree.NodeString("document"))
	tags := regexp.MustCompile("<[^!](.*?)>").FindAllString(inputText, -1)

	for _, line := range tags {
		line = line[1 : len(line)-1]

		if line[0] == '/' { // Go up the tree
			newRoot, _ := root.Parent()
			root = newRoot
		} else { // Add child and traverse the tree
			tagName := regexp.MustCompile("^\\w+").FindString(line)
			root.AddChild(tree.NodeString(tagName))

			newRoot, err := root.Child(len(root.Children()) - 1)
			if err != nil {
				log.Fatalf("Failed to read child node of root: %s", root)
			}
			root = newRoot
		}
		if line[len(line)-1] == '/' { // Go up the tree
			newRoot, _ := root.Parent()
			root = newRoot
		}
	}

	return root.String()
}
