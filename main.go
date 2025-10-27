package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"

	"acad.learn2earn.ng/git/cnnamani/quad/piscine"
)

func captureOutput(fn func(int, int), width, height int) string {
	r, w, _ := os.Pipe()
	stdout := os.Stdout
	os.Stdout = w

	fn(width, height)

	w.Close()
	os.Stdout = stdout
	var builder strings.Builder
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
		builder.WriteString("\n")
	}
	return builder.String()
}

func main() {
	input, _ := bufio.NewReader(os.Stdin).ReadString(0)
	if len(input) == 0 {
		data, _ := io.ReadAll(os.Stdin)
		input = string(data)
	}

	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	if len(lines) == 0 || len(lines[0]) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	height := len(lines)
	width := len(lines[0])

	results := map[string]string{
		"quadA": captureOutput(piscine.QuadA, width, height),
		"quadB": captureOutput(piscine.QuadB, width, height),
		"quadC": captureOutput(piscine.QuadC, width, height),
		"quadD": captureOutput(piscine.QuadD, width, height),
		"quadE": captureOutput(piscine.QuadE, width, height),
	}

	var matches []string
	for name, pattern := range results {
		if input+"\n" == pattern {
			matches = append(matches, fmt.Sprintf("[%s] [%d] [%d]", name, width, height))
		}
	}

	slices.Sort(matches)

	if len(matches) == 0 {
		fmt.Println("Not a quad function")
	} else {
		fmt.Println(strings.Join(matches, " || "))
	}
}

// package main

// import (
// 	"os"
// 	"strconv"

// 	"acad.learn2earn.ng/git/cnnamani/quad/piscine"
// )

// func main() {
// 	args := os.Args

// 	x, _ := strconv.Atoi(args[1])
// 	y, _ := strconv.Atoi(args[2])

// 	if args[0] == "./quadA" {
// 		piscine.QuadA(x, y)
// 	} else if args[0] == "./quadB" {
// 		piscine.QuadB(x, y)
// 	} else if args[0] == "./quadC" {
// 		piscine.QuadC(x, y)
// 	} else if args[0] == "./quadD" {
// 		piscine.QuadD(x, y)
// 	} else if args[0] == "./quadE" {
// 		piscine.QuadE(x, y)
// 	} else {
// 		os.Stdout.WriteString("Not a quad function\n")
// 	}
// }
