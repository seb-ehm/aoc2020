package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	hill := make([]string, 0, 0)

	scanner := bufio.NewScanner(file)
	lines := 0
	for scanner.Scan() {
		line := scanner.Text()
		hill = append(hill, line)
		lines += 1
	}

	counter := 0
	width := len(hill[0])
	incy := 2
	incx := 1

	posx := 0
	posy := 0

	for posy < lines {
		if hill[posy][posx] == '#' {
			counter += 1
		}
		posx += incx
		posy += incy
		posx = posx % width
	}

	fmt.Println(counter)

}
