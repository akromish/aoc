package main

import (
	"os"
	"strings"
)



func main() {
    bytesRead, _ := os.ReadFile("/Users/akramw/Downloads/aoc_day_1_challenge_1.txt")
    fileContent := string(bytesRead)
    lines := strings.Split(fileContent, "\n")
    print(lines[1])
}



