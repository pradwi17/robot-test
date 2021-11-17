package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	initializePosition()
}

func initializePosition() (direction string, x int, y int) {
	fmt.Println("===Initialize robot position===")
	fmt.Print("Facing Direction (N / E / S / W): ")
	scanner.Scan()
	direction = scanner.Text()

	return direction, 0, 0
}
