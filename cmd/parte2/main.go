package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type Column struct {
	Organizacion string `json:"organizacion"`
	Usuario      string `json:"usuario"`
	Rol          string `json:"rol"`
}

type FileCsv1 struct {
	column []*Column
}

type Answers struct {
	Organizacion string `json:"organizacion"`
	Users        []User `json:"users"`
}
type User struct {
	UserName string   `json:"username"`
	Roles    []string `json:"roles"`
}

func Search(length int, f func(index int) bool) int {
	for index := 0; index < length; index++ {
		if f(index) {
			return index
		}
	}
	return -1
}

func main() {
	fileName := "./archivo.csv"
	var fileCsv1 FileCsv1
	fileCsv1 = readCsvFile(fileName)

	answers := []Answers{}
	for _, col := range fileCsv1.column {

		indexOrganizacion := len(answers)
		index := Search(len(answers), func(i int) bool { return answers[i].Organizacion == col.Organizacion })
		if index != -1 {
			indexOrganizacion = index
		} else {
			answers = append(answers, Answers{Organizacion: col.Organizacion})
		}

		indexUser := len(answers[indexOrganizacion].Users)
		index = Search(len(answers[indexOrganizacion].Users), func(i int) bool { return answers[indexOrganizacion].Users[i].UserName == col.Usuario })
		if index != -1 {
			indexUser = index
		} else {
			answers[indexOrganizacion].Users = append(answers[indexOrganizacion].Users, User{UserName: col.Usuario})
		}

		indexRole := len(answers[indexOrganizacion].Users[indexUser].Roles)
		index = Search(indexRole, func(i int) bool { return answers[indexOrganizacion].Users[indexUser].Roles[i] == col.Rol })
		if index == -1 {
			answers[indexOrganizacion].Users[indexUser].Roles = append(answers[indexOrganizacion].Users[indexUser].Roles, col.Rol)
		}

	}

	out, err := json.MarshalIndent(answers, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)

	}
	fmt.Println(string(out))
}

func readCsvFile(filePath string) FileCsv1 {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()
	var fileCsv1 FileCsv1 = FileCsv1{}
	var wg sync.WaitGroup
	wg.Add(1)
	go func(file io.Reader) {

		defer wg.Done()
		records, _ := csv.NewReader(file).ReadAll()
		for _, row := range records {
			col := &Column{
				Organizacion: string(row[0]),
				Usuario:      string(row[1]),
				Rol:          string(row[2]),
			}
			fileCsv1.column = append(fileCsv1.column, col)
		}

	}(f)

	wg.Wait()
	return fileCsv1
}
