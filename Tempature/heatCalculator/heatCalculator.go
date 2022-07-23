package main

import (
	"errors"
	"fmt"
	heatFunction "fundamentals/ch07/Tempature/heatMath"
	"os"
	"strings"
)

var flag bool = true

type Calculator struct {
	functions []heatFunction.HeatFunction
}

func (c *Calculator) addHeatFunction(m heatFunction.HeatFunction) {
	c.functions = append(c.functions, m)
}

func (c *Calculator) doCalculation(name string, arg float64, toChange string) (float64, error) {
	var result float64
	for _, f := range c.functions {
		if strings.ToLower(name) == strings.ToLower(f.GetName()) {
			result = f.Calculate(arg, toChange)
			return result, nil
		}
	}
	return 0, errors.New("no such function exists:" + name)
}

func main() {
	startCalculator()
}

func startCalculator() {
	myCalculator := Calculator{}
	//myCalculator.functions[0] = Sin{"Sinus"}
	//myCalculator.functions[1] = Cos{"Cosinus"}
	//myCalculator.functions[2] = Log{"Log"}
	myCalculator.addHeatFunction(heatFunction.Celsius{"Celsius"})
	myCalculator.addHeatFunction(heatFunction.Fahrenheit{"Fahrenheit"})
	myCalculator.addHeatFunction(heatFunction.Kelvin{"Kelvin"})

	fmt.Println("\nCalculator started.")
	fmt.Println("You can calculate using following functions")
	for _, f := range myCalculator.functions {
		fmt.Println(f.GetName())
	}

	for flag {
		var fName string
		var fNameTwo string
		var arg float64
		fmt.Println("> Enter name temperature systems that you want to change then enter the tempature system to be changed  or enter x to exit:")
		_, err := fmt.Scanf("%s\n", &fName)
		_, err2 := fmt.Scanf("%s\n", &fNameTwo)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		if err2 != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		//
		if fName == "x" {
			//os.Exit(0)
			flag = false
		} else {
			fmt.Println("> Enter a value for the calculation:")
			_, err := fmt.Scanf("%f", &arg)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			value, err := myCalculator.doCalculation(fName, arg, fNameTwo)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Result %s changed value is : %f\n", fName, value)
			}
		}
	}
	println("Bye!")
}
