package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reFwd, _ := regexp.Compile("^forward ([0-9]+)$")
	reUp, _ := regexp.Compile("^up ([0-9]+)$")
	reDown, _ := regexp.Compile("^down ([0-9]+)$")

	scanner := bufio.NewScanner(file)
	var intFwdTotal, intDepthSimple, intDepth, intAim = 0, 0, 0, 0
	for scanner.Scan() {
		strLine := scanner.Text()
		if reFwd.MatchString(strLine) {
			m, err := strconv.Atoi(reFwd.FindStringSubmatch(strLine)[1])
			if err != nil {
				log.Fatal(err)
			}
			intFwdTotal += m
			if intAim != 0 {
				intDepth += (m * intAim)
			}
		} else if reUp.MatchString(strLine) {
			m, err := strconv.Atoi(reUp.FindStringSubmatch(strLine)[1])
			if err != nil {
				log.Fatal(err)
			}
			intDepthSimple -= m
			intAim -= m
		} else if reDown.MatchString(strLine) {
			m, err := strconv.Atoi(reDown.FindStringSubmatch(strLine)[1])
			if err != nil {
				log.Fatal(err)
			}
			intDepthSimple += m
			intAim += m
		} else {
			fmt.Println("Invalid line: ", strLine)
		}
	}
	fmt.Println("Forward total: ", intFwdTotal)
	fmt.Println("Depth (simple): ", intDepthSimple)
	fmt.Println("Depth (with aim): ", intDepth)
	fmt.Println("Simple multiplier: ", intFwdTotal*intDepthSimple)
	fmt.Println("Multiplier: ", intFwdTotal*intDepth)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
