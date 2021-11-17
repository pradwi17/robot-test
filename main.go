package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

var direction string
var xAxis int
var yAxis int

func main() {
	direction, xAxis, yAxis = 	initializePosition()

}

func initializePosition() (direction string, x int, y int) {
	fmt.Println("===Initialize robot position===")
	fmt.Print("Facing Direction (N / E / S / W): ")
	scanner.Scan()
	direction = scanner.Text()
	fmt.Print("X axis (number): ")
	scanner.Scan()
	x, _ = strconv.Atoi(scanner.Text())

	return direction, x, y
}
