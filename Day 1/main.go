package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var intLastLine, intLastSum, intLineIncreased, intSumIncreased = -1, -1, 0, 0
	var queue []int
	for scanner.Scan() {
		intLine, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}
		if intLine > intLastLine && intLastLine != -1 {
			intLineIncreased++
		}
		intLastLine = intLine
		queue = append(queue, intLine)
		if len(queue) == 3 {
			intSum := 0
			for _, v := range queue {
				intSum += v
			}
			if intSum > intLastSum && intLastSum != -1 {
				intSumIncreased++
			}
			intLastSum = intSum
			queue = queue[1:]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Lines increased: ", intLineIncreased)
	fmt.Println("Sums increased: ", intSumIncreased)
}
