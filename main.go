package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

var direction string
var xAxis int
var yAxis int
var command []string

func main() {
	direction, xAxis, yAxis = initializePosition()
	fmt.Print("Input command (A, L, R): ")
	scanner.Scan()
	command = strings.Split(scanner.Text(), "")
	fmt.Println(command)

}

func initializePosition() (direction string, x int, y int) {
	fmt.Println("===Initialize robot position===")
	fmt.Print("Facing Direction (N / E / S / W): ")
	scanner.Scan()
	direction = scanner.Text()
	fmt.Print("X axis (number): ")
	scanner.Scan()
	x, _ = strconv.Atoi(scanner.Text())
	fmt.Print("Y axis (number): ")
	scanner.Scan()
	y, _ = strconv.Atoi(scanner.Text())
	return direction, x, y
}
