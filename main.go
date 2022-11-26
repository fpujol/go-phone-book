package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Entry struct {
	Name       string
	Surname    string
	Number     string
	LastAccess string
}

var myData = []Entry{}
var index map[string]int
var CSVFILE = "data.csv"

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: insert|delete|search|list <arguments>")
		return
	}

	//if the CSVFILE does not exist, create an empty one
	_, err := os.Stat(CSVFILE)
	if err != nil {
		fmt.Println("Creating", CSVFILE)
		f, err := os.Create(CSVFILE)
		if err != nil {
			f.Close()
			fmt.Println(err)
			return
		}
		f.Close()
	}

	fileInfo, err := os.Stat(CSVFILE)
	//Is it a regular file ?
	mode := fileInfo.Mode()
	if !mode.IsRegular() {
		fmt.Println(CSVFILE, "not a regular file!")
		return
	}

	err = readCsvFile(CSVFILE)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = createIndex()
	if err != nil {
		fmt.Println("Cannot create index.")
		return
	}

}

func createIndex() error {

	for i, v := range myData {
		index[v.Number] = i
	}
	return nil
}

func readCsvFile(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}
	for _, line := range lines {
		temp := Entry{
			Name:       line[0],
			Surname:    line[1],
			Number:     line[2],
			LastAccess: line[3],
		}
		myData = append(myData, temp)
	}
	return nil
}

func saveCSVFile(filepath string) error {
	csvFile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvFile.Close()
	csvWriter := csv.NewWriter(csvFile)
	csvWriter.Comma = '\t'
	for _, row := range myData {
		temp := []string{row.Name, row.Surname, row.Number}
		_ = csvWriter.Write(temp)
	}
	csvWriter.Flush()
	return nil
}
