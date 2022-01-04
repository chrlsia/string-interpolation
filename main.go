package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

//this is not a variable
//it is a type for a variable
type User struct {
	UserName string
	Age int
	FavoriteNumber float64
}

func main(){
	reader=bufio.NewReader(os.Stdin)

	var user User
	user.UserName= readString("What is your name?")
	user.Age=readInt("How old are you?")
	user.FavoriteNumber = readFloat("What is your favorite number?")

	s:=fmt.Sprintf("Your name is %s, your age is %d years old and your favorite number is %.2f",
	user.UserName,
	user.Age,
	user.FavoriteNumber)
	fmt.Println(s)
}

func prompt() {
	fmt.Print("-> ")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput,"\n","",-1)

		if userInput != "" {
			return userInput
		} else{
			fmt.Println("Please enter a value")
		}

		
	}
}

func readInt(s string) int {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput,"\n","",-1)

		num, err:= strconv.Atoi(userInput)

		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return num	
		}
	}
}

	func readFloat(s string) float64 {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput,"\n","",-1)
		
		num, err:= strconv.ParseFloat(userInput,64)

		if err != nil {
			fmt.Println("Please enter a number")
		} else {
			return num	
		}
	}
}