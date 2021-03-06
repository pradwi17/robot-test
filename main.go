package main

import (
	"bufio"
	"fmt"
	"os"
	"robot-curiosity/entity"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

var command []string

func main() {

	robot := entity.InitializeRobotPosition()
	fmt.Print("Input command (A, L, R): ")
	scanner.Scan()
	command = strings.Split(scanner.Text(), "")
	robot.Move(command)

}
