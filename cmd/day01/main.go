package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	numbers := make(map[int]bool)

	scanner := bufio.NewScanner(file)
	var contents [200]int
	data := make(map[int]bool)
	
	index := 0
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if numbers[2020-i] {
			fmt.Println("Two sum:")
			fmt.Println(i * (2020 - i))
		}
		numbers[i] = true
		data[i] = true
		contents[index] = i
		index +=1
	}



	for i, value := range contents {
		goal := 2020 - value
		for j := i + 1; j < 200; j++ {
			target := goal-contents[j]
			/*if target >= 0 {
				fmt.Printf("Value: %d , Current: %d Target: %d \n", value, contents[j], target)
			}*/
			if data[target] {
				fmt.Printf("Three sum: \n %d, %d, %d \n", value, contents[j], goal-contents[j])
				fmt.Println(value * contents[j] * (goal-contents[j]))
				break
			}
		}
	}

}
