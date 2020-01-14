package main

import (
	"log"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"encoding/csv"
	"flag"

	"github.com/rainycape/unidecode"
)

func main() {
	pathVal := getPath()
	files, err := ioutil.ReadDir(pathVal)
	check("Cannot read files from folder, maybe permission problems?",err)
	fileNames := rename(files,pathVal)
	createCsvFromFileNames(fileNames)
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
	writer := csv.NewWriter(file)
	check("Cannot write to file, maybe permission problems?",writer.WriteAll(fileNames))
	file.Close()
}

func rename(files []os.FileInfo,workFolder string) (csvExport [][]string) {
	for _, file := range files {
		newName := slugify(file.Name())
		var dummyToCsv []string
		dummyToCsv = append(dummyToCsv, newName)
		csvExport = append(csvExport,dummyToCsv)
		log.Print(workFolder+file.Name())
		go os.Rename(workFolder+file.Name(), workFolder+"/"+newName)
		//log.Print("Renamed files:", i+1)
	}
	return
}

func getPath() string {
	workFolder := flag.String("path", "./", "define workdir")
	flag.Parse()
	returnVal := *workFolder
	return returnVal
}