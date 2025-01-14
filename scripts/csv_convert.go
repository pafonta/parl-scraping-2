package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	input := os.Args[1]
	output := os.Args[2]

	outFile, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}
	w := csv.NewWriter(outFile)

	data, err := ioutil.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	var documents []map[string]interface{}
	err = json.Unmarshal(data, &documents)
	if err != nil {
		log.Fatal(err)
	}

	nFields := len(documents[0])

	// Get list of map keys from first map in slice
	headers := make([]string, 0, nFields)
	for header := range documents[0] {
		headers = append(headers, header)
	}

	w.Write(headers)

	row := make([]string, nFields)
	for _, document := range documents {
		for i, header := range headers {
			if document[header] == nil {
				row[i] = ""
			} else {
				row[i] = fmt.Sprintf("%v", document[header])
			}
		}
		w.Write(row)
	}
	w.Flush()

	log.Println("Number of documents", len(documents))
}
