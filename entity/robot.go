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
var turn *string = &direction

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

func (robot *Robot) getMovingCoordinate(movingRobot Robot, turn *string) {
	switch *turn {
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
			turnLeft(turn)
			fmt.Printf("%s -> (%d,%d)\n", direction, xAxis, yAxis)
		case "R":
			turnRight(turn)
			fmt.Printf("%s -> (%d,%d)\n", direction, xAxis, yAxis)
		case "A":
			robot.getMovingCoordinate(robot, turn)
			fmt.Printf("%s -> (%d,%d)\n", robot.direction, robot.xAxis, robot.yAxis)
		default:
			fmt.Printf("%s -> (%d,%d)\n", direction, xAxis, yAxis)
		}
	}
}
