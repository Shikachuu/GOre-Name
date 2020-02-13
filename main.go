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
	path := getPath()
	files, dirReadErr := ioutil.ReadDir(path)
	check("Cannot read files from folder, maybe permission problems?",dirReadErr)
	fileNames := renameFolderContent(files,path)
	go createCsvFromFileNames(fileNames)
	fmt.Println("Press Enter to exit")
	fmt.Scanln()
}

func check(message string,err error) {
	if err != nil {
		log.Fatal(message,err)
	}
}

func slugify(stringToSlugify string, c chan string) {
	changeSpacesToDashes := strings.NewReplacer(" ", "_")
	finalString := changeSpacesToDashes.Replace(strings.ToLower(stringToSlugify))
	c <- unidecode.Unidecode(finalString)
}

func createCsvFromFileNames(fileNames [][]string) {
	file,err := os.Create("rename-result.csv")
	check("Cannot create file, maybe permission problems?",err)
	writer := csv.NewWriter(file)
	check("Cannot write to file, maybe permission problems?",writer.WriteAll(fileNames))
	file.Close()
}

func renameFolderContent(files []os.FileInfo,workFolder string) (csvExport [][]string) {
	slugChan := make(chan string)
	for _, file := range files {
		go slugify(file.Name(),slugChan)
		newName := <- slugChan
		csvRow := []string{newName}
		csvExport = append(csvExport,csvRow)
		log.Print(workFolder+file.Name())
		go os.Rename(workFolder+file.Name(), workFolder+"/"+newName)
	}
	return
}

func getPath() string {
	workFolder := flag.String("path", "./", "define workdir")
	flag.Parse()
	workFolderDeref := *workFolder
	return workFolderDeref
}