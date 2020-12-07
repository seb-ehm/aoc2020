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
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	counter_part_1 := 0
	counter_part_2 := 0
	re := regexp.MustCompile(`(?P<start>\d+)-(?P<end>\d+) (?P<char>\S): (?P<content>\S+)`)
	for scanner.Scan() {
		line := scanner.Text()
		contents := re.FindStringSubmatch(line)
		start, _ := strconv.Atoi(contents[1])
		ende, _ := strconv.Atoi(contents[2])
		char := []rune(contents[3])[0]
		content := contents[4]
		i := 0
		for _, c := range content {
			if c == char {
				i++
			}
		}
		if i >= start && i <= ende {
			counter_part_1 += 1
		}

		c := 0
		if content[start-1] == byte(char) {
			c += 1
		}

		if content[ende-1] == byte(char) {
			c += 1
		}
		if c == 1 {
			counter_part_2 += 1
		}

	}
	fmt.Printf("Task 1: %d \n", counter_part_1)
	fmt.Printf("Task 2: %d \n", counter_part_2)
}
