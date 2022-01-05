package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

var reader *bufio.Reader

//this is not a variable
//it is a type for a variable
type User struct {
	UserName string
	Age int
	FavoriteNumber float64
	OwnsADog bool
}

func main(){
	reader=bufio.NewReader(os.Stdin)

	var user User
	user.UserName= readString("What is your name?")
	user.Age=readInt("How old are you?")
	user.FavoriteNumber = readFloat("What is your favorite number?")
	user.OwnsADog = readBool("Do you own a dog (y/n)?")

	s:=fmt.Sprintf("Your name is %s, your age is %d years old, your favorite number is %.2f, and you own a dog: %t",
	user.UserName,
	user.Age,
	user.FavoriteNumber, user.OwnsADog)
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

		// get user input and replace ENTER with an empty string
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput,"\n","",-1)
		
		// is the user input a valid float number?
		num, err:= strconv.ParseFloat(userInput,64)

		//if there is an error, message and repeat
		// due to endless loop
		// otherwise return num
		if err != nil {
			fmt.Println("Please enter a number")
		} else {
			return num	
		}
	}
}

func readBool(s string) bool{
	err:=keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	// from here and downwards there is no error
	// from Open() function
	defer func(){
		_=keyboard.Close()
	}()

	for {
		fmt.Println(s)
		char,_,err:=keyboard.GetSingleKey()

		if err != nil {
			log.Fatal(err)
		}
		// from here and downwards there is no error
		// from GetSingleKey()
		if strings.ToLower(string(char))!="y" && strings.ToLower(string(char))!="n" {
			fmt.Println("Please type y or n")
		} else if char == 'n' || char == 'N'{
			return false
		} else if char == 'y' || char == 'Y'{
			return true
		}
	}
}