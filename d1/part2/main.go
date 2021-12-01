package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func getSum(depths []string, index int) int {
	d1, err := strconv.Atoi(depths[index])
	if err != nil {
		log.Fatalf("invalid depth #%d:%s [%v]", index, depths[index], err)
	}
	d2, err := strconv.Atoi(depths[index+1])
	if err != nil {
		log.Fatalf("invalid depth #%d:%s [%v]", index+1, depths[index+1], err)
	}
	d3, err := strconv.Atoi(depths[index+2])
	if err != nil {
		log.Fatalf("invalid depth #%d:%s [%v]", index+2, depths[index+2], err)
	}

	return d1 + d2 + d3
}

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
	for i := range depths {
		if i == 0 || i+3 > len(depths) || depths[i+2] == "" {
			continue
		}
		sum1 := getSum(depths, i-1)

		sum2 := getSum(depths, i)

		if sum2 > sum1 {
			increasingCount++
		}
	}

	fmt.Printf("Depth increased %d times", increasingCount)
}

// if len(depths) < index + 1 {

// }
