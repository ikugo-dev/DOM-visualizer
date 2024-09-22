package main

import (
	"fmt"
	"io"
	"log"
	"strings"

	"net/http"

	"github.com/m1gwings/treedrawer/tree"
)

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

func parseDOM(_ string) string {
	t := tree.NewTree(tree.NodeInt64(5))

	t.AddChild(tree.NodeString("adding a string"))
	t.AddChild(tree.NodeInt64(42))
	t.AddChild(tree.NodeInt64(3))

	return t.String()
}

func splitXML(r rune) bool {
	return r == ' ' || r == '>'
}

func main() {
	inputText := strings.Split(testXML, "<")

	root := tree.NewTree(tree.NodeString("HTML"))

	for _, text := range inputText {
		if text == " " {
			continue
		}
		textParts := strings.FieldsFunc(text, splitXML)
		if len(textParts) == 0 {
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

	fmt.Printf("%v\n", root)
}

// const portName = ":6969"
// http.Handle("/", http.FileServer(http.Dir("./")))
// http.HandleFunc("/parse-text", parseDOMHandler)
//
// log.Printf("Starting server on :%s", portName)
// http.ListenAndServe(portName, nil)
const testXML = "<Tests xmlns=\"http://www.adatum.com\"> <Test TestId=\"0001\" TestType=\"CMD\"> <Name>Convert number to string</Name> <CommandLine>Examp1.EXE</CommandLine> <Input>1</Input> <Output>One</Output> </Test> <Test TestId=\"0002\" TestType=\"CMD\"> <Name>Find succeeding characters</Name> <CommandLine>Examp2.EXE</CommandLine> <Input>abc</Input> <Output>def</Output> </Test> <Test TestId=\"0003\" TestType=\"GUI\"> <Name>Convert multiple numbers to strings</Name> <CommandLine>Examp2.EXE /Verbose</CommandLine> <Input>123</Input> <Output>One Two Three</Output> </Test> <Test TestId=\"0004\" TestType=\"GUI\"> <Name>Find correlated key</Name> <CommandLine>Examp3.EXE</CommandLine> <Input>a1</Input> <Output>b1</Output> </Test> <Test TestId=\"0005\" TestType=\"GUI\"> <Name>Count characters</Name> <CommandLine>FinalExamp.EXE</CommandLine> <Input>This is a test</Input> <Output>14</Output> </Test> <Test TestId=\"0006\" TestType=\"GUI\"> <Name>Another Test</Name> <CommandLine>Examp2.EXE</CommandLine> <Input>Test Input</Input> <Output>10</Output> </Test> </Tests>"
