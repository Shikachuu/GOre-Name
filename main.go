package main

import (
	"log"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"encoding/csv"

	"github.com/rainycape/unidecode"
)

func main() {
	var csvExport [][]string
	workFolder := "."
	files, err := ioutil.ReadDir(workFolder)
	check("Cannot read files from folder, maybe permission problems?",err)
	for i, file := range files {
		newName := slugify(file.Name())
		var dummyToCsv []string
		dummyToCsv = append(dummyToCsv, newName)
		csvExport = append(csvExport,dummyToCsv)
		go os.Rename(file.Name(), workFolder+"/"+newName)
		log.Print("Renamed files:", i+1)
	}
	createCsvFromFileNames(csvExport)
	fmt.Println("Press Enter to exit")
	fmt.Scanln()
}

func check(message string,err error) {
	if err != nil {
		log.Fatal(message,err)
	}
}

func slugify(stringToSlugify string) string {
	changeSpacesToDashes := strings.NewReplacer(" ", "_")
	finalString := changeSpacesToDashes.Replace(strings.ToLower(stringToSlugify))
	return unidecode.Unidecode(finalString)
}

func createCsvFromFileNames(fileNames [][]string) {
	file,err := os.Create("rename-result.csv")
	check("Cannot create file, maybe permission problems?",err)
	defer file.Close()
	writer := csv.NewWriter(file)
	check("Cannot write to file, maybe permission problems?",writer.WriteAll(fileNames))
}