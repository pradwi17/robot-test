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
var turn *string = &direction

func main() {
	
	direction, xAxis, yAxis = initializePosition()
	fmt.Print("Input command (A, L, R): ")
	scanner.Scan()
	command = strings.Split(scanner.Text(), "")
	fmt.Println(command)
	move(command)

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

func move(command []string) {
	for i := 0; i < len(command); i++ {
		switch command[i] {
		case "L":
			turnLeft(turn)
			fmt.Printf("%s -> (%d,%d)\n", direction, xAxis, yAxis)
		case "R":
			turnRight(turn)
			fmt.Printf("%s -> (%d,%d)\n", direction, xAxis, yAxis)
		case "A":
			getMovingCoordinate(&xAxis, &yAxis, turn)
			fmt.Printf("%s -> (%d,%d)\n", direction, xAxis, yAxis)
		}
	}
}

func turnLeft(turn *string) {
	switch *turn {
	case "N":
		*turn = "W"
	case "E":
		*turn = "N"
	case "S":
		*turn = "E"
	case "W":
		*turn = "S"

	}
}

func turnRight(turn *string) {
	switch *turn {
	case "N":
		*turn = "E"
	case "E":
		*turn = "S"
	case "S":
		*turn = "W"
	case "W":
		*turn = "N"

	}
}

func getMovingCoordinate(x *int, y *int, turn *string) {
	switch *turn {
	case "N":
		*y = *y + 1
	case "E":
		*x = *x + 1
	case "S":
		*y = *y - 1
	case "W":
		*x = *x - 1

	}
}
