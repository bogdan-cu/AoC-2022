package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	sort.Slice(loads, func(i, j int) bool {
		return loads[i] > loads[j]
	})

	maxLoads := 0
	for i := 0; i < 3 && i < len(loads); i++ {
		maxLoads += loads[i]
	}
	fmt.Println("the largest three loads weigh ", maxLoads)
}
