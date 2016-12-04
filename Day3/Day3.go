package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := getFileLines()

	counter := 0
	for i := 0 ; i < len(lines);i +=3 {

		a1, a2, a3 := sides(lines[i])
		b1, b2, b3 := sides(lines[i+1])
		c1, c2, c3 := sides(lines[i+2])
		if isATriangle(a1, b1, c1) {
			counter++
		}
		if isATriangle(a2, b2, c2) {
			counter++
		}
		if isATriangle(a3, b3, c3) {
			counter++
		}
	}
	fmt.Println("Total Counter: ", counter)
	fmt.Println(len(lines))
	fmt.Println(lines[0]," ", lines[1634])
}

func sides(s string) (int, int, int) {
	// Prints were used for debugging
	fmt.Println("Start of sides")
	abc := strings.TrimSpace(s)
	fmt.Println("abc:", abc)
	// Take first number
	abcArray1 := strings.SplitAfterN(abc, " ", 2)
	fmt.Println("abcArray1:", abcArray1)
	fmt.Println("abcArray1[0]:", abcArray1[0])
	a, _ := strconv.Atoi(strings.TrimSpace(abcArray1[0]))
	fmt.Println("a:", a)
	// Take second number
	abc = strings.TrimSpace(abcArray1[1])
	fmt.Println("abc:", abc)
	abcArray2 := strings.SplitAfterN(abc, " ", 2)
	fmt.Println("abcArray2:", abcArray2)
	b, _ := strconv.Atoi(strings.TrimSpace(abcArray2[0]))
	fmt.Println("b", b)

	// Take last number
	abc = strings.TrimSpace(abcArray2[1])
	fmt.Println("abc", abc)
	c, _ := strconv.Atoi(strings.TrimSpace(abc))
	fmt.Println("c:", c)
	fmt.Println(a, b, c)
	fmt.Println("End of sides")
	return a, b, c
}

func getFileLines() []string {
	fileLines := make([]string, 0)
	f, err := os.Open("input")
	if err != nil {
		fmt.Println("error opening file= ", err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	s, _, e := r.ReadLine()
	for e == nil {
		newLine := string(s)
		fileLines = append(fileLines, newLine)
		s, _, e = r.ReadLine()
	}
	return fileLines
}

func isATriangle(a int, b int, c int) bool {
	/* In a valid triangle, the sum of any two sides must be larger than the remaining side.
	For example, the "triangle" given above is impossible, because 5 + 10 is not larger than 25.

	In your puzzle input, how many of the listed triangles are possible?*/
	return a+b>c && a+c>b && b+c>a && a>0 && b>0 && c>0
}
