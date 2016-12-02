package main
// Advent of code day 1 problem 2
import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type coordinate struct {
	x int
	y int
}

const (
	north = iota
	south = iota
	west  = iota
	east  = iota
)

var currentDirection int
var visitedCords []coordinate
var currentCord coordinate

func main() {
	currentDirection = north
	myMap := getMap()
	stringInstructions := strings.Split(myMap, ", ")
	for _, e := range stringInstructions {
		if strings.HasPrefix(e, "R") {
			updateDirectionRight()
			nbrMoves, _ := strconv.Atoi(e[1:])
			if move(nbrMoves) {
				fmt.Println(intAbs(currentCord.x) + intAbs(currentCord.y))
				break
			}

		} else {
			updateDirectionLeft()
			nbrMoves, _ := strconv.Atoi(e[1:])
			if move(nbrMoves) {
				fmt.Println(intAbs(currentCord.x) + intAbs(currentCord.y))
				break
			}
		}
	}
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func move(steps int) bool {
	for i := 0; i < steps; i++ {
		switch currentDirection {
		case north:
			currentCord.y++
		case south:
			currentCord.y--
		case west:
			currentCord.x--
		case east:
			currentCord.x++
		}
		if contains(visitedCords, currentCord) {
			return true
		}
		//fmt.Println(currentCord.x)
		//fmt.Println(currentCord.y)
		visitedCords = append(visitedCords, currentCord)
	}
	return false
}

func contains(list []coordinate, elem coordinate) bool {
	for _, e := range list {
		if e == elem {
			return true
		}
	}
	return false
}

func updateDirectionRight() {
	switch currentDirection {
	case north:
		currentDirection = east
	case south:
		currentDirection = west
	case west:
		currentDirection = north
	case east:
		currentDirection = south
	}
}

func updateDirectionLeft() {
	switch currentDirection {
	case north:
		currentDirection = west
	case south:
		currentDirection = east
	case west:
		currentDirection = south
	case east:
		currentDirection = north
	}
}

func getMap() string {
	b, err := ioutil.ReadFile("instructions")
	if err != nil {
		return "FUCK OFF"
	}
	return string(b)
}
