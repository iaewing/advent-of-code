package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"strconv"
)

func main() {
	var lines []string;
	file, err :=  os.Open("input.txt");
	if err != nil {
		fmt.Println(err);
	}
	defer file.Close();

	scanner := bufio.NewScanner(file);
	scanner.Split(bufio.ScanLines);

	runningTotal := 0;

	for scanner.Scan() {
		runningTotal += parseInput(scanner.Text());
		lines = append(lines, scanner.Text());
	}
	fmt.Println(runningTotal);
}

func parseInput(s string) int {
	regex := regexp.MustCompile("[0-9]+");
	digits := regex.FindAllString(s, -1);
	digitString := strings.Join(digits, "");

	firstAndLast := digitString[0:1] + digitString[len(digitString)-1:];
	int, err := strconv.Atoi(firstAndLast);
	if err != nil {
		fmt.Println(err);
	}
	return int;
}