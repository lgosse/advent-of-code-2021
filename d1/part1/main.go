package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide input file: go run main.go <input,file>")
	}

	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("error opening file %v", err)
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("error reading file %v", err)
	}

	depths := strings.Split(string(content), "\n")

	increasingCount := 0
	for i, depth := range depths {
		if i == 0 || depth == "" {
			continue
		}
		previousDepth, err := strconv.Atoi(depths[i-1])
		if err != nil {
			log.Fatalf("invalid depth #%d:%s [%v]", i, depth, err)
		}

		actualDepth, err := strconv.Atoi(depth)
		if err != nil {
			log.Fatalf("invalid depth #%d:%s [%v]", i, depth, err)
		}

		if actualDepth > previousDepth {
			increasingCount++
		}
	}

	fmt.Printf("Depth increased %d times", increasingCount)
}
