package main

import (
	"fmt"
	generator "gopps/Generator"
	"os"
	"strings"
)

func main() {
	var fileName string
	fmt.Print("Enter name of your password list:")
	fmt.Scanln(&fileName)

	file, err := createFile(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	info := getInfo()

	generator.Generate(file, info)
}

func createFile(fileName string) (*os.File,error) {
	if !strings.HasSuffix(fileName, ".txt") {
		fileName += ".txt"
	}
	
	file, err := os.Create(fileName)
	if err != nil {
		return nil,err
	}

	return file,nil
}

func getInfo() []string {
	info := []string{}
	var (
		firstName       string
		lastName       string
		phoneNumber string
		year        string
		month       string
		monthName       string
		day         string
	)

	fmt.Print("Enter first name:")
	fmt.Scanln(&firstName)
	info = append(info, firstName)
	info = append(info, strings.ToLower(firstName))

	fmt.Print("Enter last name:")
	fmt.Scanln(&lastName)
	info = append(info, lastName)
	info = append(info, strings.ToLower(lastName))

	fmt.Print("Enter phone number:")
	fmt.Scanln(&phoneNumber)
	info = append(info, phoneNumber)

	fmt.Print("Enter year of Birth(1980):")
	fmt.Scanln(&year)
	info = append(info, year)

	fmt.Print("Enter month of Birth(1-12):")
	fmt.Scanln(&month)
	info = append(info, month)

	fmt.Print("Enter month of Birth(May):")
	fmt.Scanln(&monthName)
	info = append(info, monthName)
	info = append(info, strings.ToLower(monthName))

	fmt.Print("Enter day of Birth:")
	fmt.Scanln(&day)
	info = append(info, day)

	return info
}
