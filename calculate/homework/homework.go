package main

import (
	"strings"
)

func main() {
	// var str string = "s"
	// changeInput(str)
	//checkError()
}

type homework interface {
	changeInput(str string) string
	MyError() string
	//checkError() string
}

type Input struct {
	Name string
}

func (i Input) changeInput(str string) string {
	str = strings.ToLower(str)
	firstCharacter := str[0:1]
	newFirst := strings.ToUpper(firstCharacter)
	res1 := strings.Replace(str, firstCharacter, newFirst, 1)
	return res1
}

//func (i Input) checkError(err) {
//	//var x string
//	//_, err := fmt.Scanf("%s", &x)
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(0)
//	}
//	fmt.Println("error yok")
//}

type MyError struct{}

func (myerr *MyError) Error() string {
	return "Something is wrong: "
}
