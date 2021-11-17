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

	var robot = entity.InitializeRobotPosition()
	fmt.Println(robot)
	fmt.Print("Input command (A, L, R): ")
	scanner.Scan()
	command = strings.Split(scanner.Text(), "")
	entity.Move(command)

}
