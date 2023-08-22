package main

import (
	"bufio"
	"fmt"
	"geektrust/expense"
	"os"
	"strings"
)

func main() {
	cliArgs := os.Args[1:]

	if len(cliArgs) == 0 {
		fmt.Println("Please provide the input file path")
		return
	}

	filePath := cliArgs[0]
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening the input file")
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	executeOp := expense.NewExpenseHandler()
	for scanner.Scan() {
		args := scanner.Text()
		argList := strings.Fields(args)
		executeOp.ExecuteOperation(argList)
	}
}
