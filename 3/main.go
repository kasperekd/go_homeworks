package main

import (
	"flag"
	"fmt"
	"my_mod/calculations"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Please provide a number as argument.")
		return
	}

	n, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid input. Please provide a valid number.")
		return
	}

	var useLogging bool
	flagSet := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flagSet.BoolVar(&useLogging, "log", false, "Enable logging")
	flagSet.Parse(args[1:])

	result := calculations.Calculate(int64(n), useLogging)

	log := logrus.New()
	log.Infof("Factorial of %d is %d", n, result)
}
