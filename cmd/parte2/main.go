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

func main() {
	fileName := "./archivo.csv"
	var fileCsv1 FileCsv1
	fileCsv1 = readCsvFile(fileName)

	//fmt.Println(fileCsv1)

	//emp := &Employee{Name: "Rocky", Number: 5454}
	salida := "["

	for _, col := range fileCsv1.column {

		e, err := json.Marshal(col)
		if err != nil {
			fmt.Println(err)
			return
		}
		salida += string(e) + ","

	}

	salida += "]"
	fmt.Println(salida)
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
