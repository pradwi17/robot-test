package entity

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

type Robot struct {
	direction string
	xAxis     int
	yAxis     int
}

func InitializeRobotPosition() Robot {
	fmt.Println("===Initialize robot position===")
	fmt.Print("Facing Direction (N / E / S / W): ")
	scanner.Scan()
	direction = scanner.Text()

	fmt.Print("X axis (number): ")
	scanner.Scan()
	xAxis, _ = strconv.Atoi(scanner.Text())
	fmt.Print("Y axis (number): ")
	scanner.Scan()
	yAxis, _ = strconv.Atoi(scanner.Text())

	return Robot{direction: direction, xAxis: xAxis, yAxis: yAxis}

}

func (robot *Robot) turnLeft() {
	switch robot.direction {
	case "N":
		robot.direction = "W"
	case "E":
		robot.direction = "N"
	case "S":
		robot.direction = "E"
	case "W":
		robot.direction = "S"

	}
}

func (robot *Robot) turnRight() {
	switch robot.direction {
	case "N":
		robot.direction = "E"
	case "E":
		robot.direction = "S"
	case "S":
		robot.direction = "W"
	case "W":
		robot.direction = "N"

	}
}

func (robot *Robot) getMovingCoordinate() {
	switch robot.direction {
	case "N":
		robot.yAxis = robot.yAxis + 1
	case "E":
		robot.xAxis = robot.xAxis + 1
	case "S":
		robot.yAxis = robot.yAxis - 1
	case "W":
		robot.xAxis = robot.xAxis - 1

	}
}

func (robot Robot) Move(command []string) {
	for i := 0; i < len(command); i++ {
		switch command[i] {
		case "L":
			robot.turnLeft()
			fmt.Printf("%s -> (%d,%d)\n", robot.direction, robot.xAxis, robot.yAxis)
		case "R":
			robot.turnRight()
			fmt.Printf("%s -> (%d,%d)\n", robot.direction, robot.xAxis, robot.yAxis)
		case "A":
			robot.getMovingCoordinate()
			fmt.Printf("%s -> (%d,%d)\n", robot.direction, robot.xAxis, robot.yAxis)
		default:
			fmt.Printf("%s -> (%d,%d)\n", robot.direction, robot.xAxis, robot.yAxis)
		}
	}
}
