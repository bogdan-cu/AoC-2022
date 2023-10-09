package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	filePath := wd + "/day1.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("failed to open file: ", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	loads := make([]int, 0)
	currentLoad := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if currentLoad != 0 {
				loads = append(loads, currentLoad)
			}
			currentLoad = 0
			continue
		}
		weight, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("invalid value encountered: ", err)
		}
		currentLoad += weight
	}

	maxLoad := 0
	for _, load := range loads {
		if load >= maxLoad {
			maxLoad = load
		}
	}
	fmt.Println("the largest load weighs ", maxLoad)
}
