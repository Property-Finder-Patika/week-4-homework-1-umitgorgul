package main

import (
	"errors"
	"fmt"
	heatFunction "Tempature/heatMath"
	"os"
	"strings"
)

var flag bool = true

type Calculator struct { // get HeatFunction functions
	functions []heatFunction.HeatFunction
}

func (c *Calculator) addHeatFunction(m heatFunction.HeatFunction) { // append functions
	c.functions = append(c.functions, m)
}

func (c *Calculator) doCalculation(name string, arg float64, toChange string) (float64, error) { // calculation started here
	var result float64
	for _, f := range c.functions {
		if strings.ToLower(name) == strings.ToLower(f.GetName()) { // make lowercase to name and getname beacause it makes checking equality between them is easy 
			result = f.Calculate(arg, toChange) // going inside heatMath
			return result, nil
		}
	}
	return 0, errors.New("no such function exists:" + name)
}

func main() {
	startCalculator() // start code
}

func startCalculator() {
	myCalculator := Calculator{}  // add Celsius, Fahrenheit, Kelvin
	myCalculator.addHeatFunction(heatFunction.Celsius{"Celsius"})
	myCalculator.addHeatFunction(heatFunction.Fahrenheit{"Fahrenheit"})
	myCalculator.addHeatFunction(heatFunction.Kelvin{"Kelvin"})

	fmt.Println("\nCalculator started.")
	fmt.Println("You can calculate using following functions")
	for _, f := range myCalculator.functions {
		fmt.Println(f.GetName())
	}

	for flag {
		var fName string // for temperature systems name which wanted to change
		var fNameTwo string // for temperature systems name which wanted to be changed
		var arg float64 // value of tempature
		fmt.Println("> Enter name temperature systems that you want to change then enter the tempature system to be changed  or enter x to exit:")
		_, err := fmt.Scanf("%s\n", &fName)
		_, err2 := fmt.Scanf("%s\n", &fNameTwo)
		if err != nil { //check error for fName
			fmt.Println(err)
			os.Exit(0)
		}
		if err2 != nil { //check error for fNameTwo
			fmt.Println(err)
			os.Exit(0)
		}

		//
		if fName == "x" {
			flag = false
		} else {
			fmt.Println("> Enter a value for the tempature:")
			_, err := fmt.Scanf("%f", &arg) 
			if err != nil {
				fmt.Println(err) // check error for tempature value
				os.Exit(0)
			}
			value, err := myCalculator.doCalculation(fName, arg, fNameTwo) // start calculation for getting value
			if err != nil {
				fmt.Println(err) 
			} else {
				fmt.Printf("Result %s changed value is : %f\n", fName, value)
			}
		}
	}
	println("Bye!")
}
