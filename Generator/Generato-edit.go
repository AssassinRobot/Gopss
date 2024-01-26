package generator

import (
	"fmt"
	"os"
	"sync"
)

var (
	mu = &sync.Mutex{}
	wg = &sync.WaitGroup{}
)

func createFile() *os.File {
	var fileName string
	fmt.Print("Enter Name of your password list file:")
	fmt.Scanln(&fileName)
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return file
}

func Do() {
	infos := getInfos()
	file := createFile()
	defer file.Close()
	wg.Add(1)
	go generator(file, infos)
	wg.Wait()
}

func getInfos() []string {
	infs := []string{}
	var (
		fName       string
		lName       string
		phoneNumber string
		year        string
		month       string
		day         string
	)
	fmt.Print("Enter first name:")
	fmt.Scanln(&fName)
	infs = append(infs, fName)
	fmt.Print("Enter last name:")
	fmt.Scanln(&lName)
	infs = append(infs, lName)
	fmt.Print("Enter phone number:")
	fmt.Scanln(&phoneNumber)
	infs = append(infs, phoneNumber)
	fmt.Print("Enter year of Birth:")
	fmt.Scanln(&year)
	infs = append(infs, year)
	fmt.Print("Enter month of Birth:")
	fmt.Scanln(&month)
	infs = append(infs, month)
	fmt.Print("Enter day of Birth:")
	fmt.Scanln(&day)
	infs = append(infs, day)
	return infs
}
func generator(file *os.File, infos []string) {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()
	for i := 0; i <= len(infos)-1; i++ {
		for j := 0; j <= len(infos)-1; j++ {
			for k := 0; k <= len(infos)-1; k++ {
				for l := 0; l <= len(infos)-1; l++ {
					for m := 0; m <= len(infos)-1; m++ {
						for n := 0; n <= len(infos)-1; n++ {
							password := fmt.Sprintf("%s%s%s%s%s%s\n", infos[i], infos[j], infos[k], infos[l], infos[m], infos[n])
							fmt.Fprint(file, password)
						}
					}
				}
			}
		}
	}
}
