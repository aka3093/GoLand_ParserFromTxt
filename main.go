package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readFile(fileName string) map[string][]string {
	//	To read and provide a map of obtained data from a file.

	// Create a Hash table (dict)
	data := make(map[string][]string)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	// Create a scanner to read a file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		data[fields[0]] = append(data[fields[0]], fields[1])
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return data
}

func meanSlice(slice []string) float64 {
	res := 0.0
	for _, num := range slice {
		num, _ := strconv.ParseFloat(num, 64)
		res += num
	}
	return res / float64(len(slice))
}

func main() {
	data := readFile("./students.txt")
	// Creating a list of keys and sort it
	lstKeys := make([]string, 0, len(data))
	for key := range data {
		lstKeys = append(lstKeys, key)
	}
	sort.Strings(lstKeys)

	for _, key := range lstKeys {
		fmt.Println(key)
		mean := meanSlice(data[key])
		fmt.Printf("Scores: %v\n", strings.Join(data[key], " "))
		fmt.Printf("Average: %.2f\n\n", mean)
	}
}
