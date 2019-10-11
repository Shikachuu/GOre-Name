package main
import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"github.com/rainycape/unidecode"
)
func main() {
	workFolder := "."
	files, err := ioutil.ReadDir(workFolder)
	if err != nil {
		fmt.Println(err)
	}
	for i, file := range files {
		trueFilename := strings.TrimSuffix(file.Name(), "\r")
		newName := slugify(trueFilename)
		os.Rename(trueFilename, workFolder+"/"+newName)
		fmt.Println("Renamed files:",i+1)
	}
	fmt.Println("Press Enter to exit")
	fmt.Scanln()
}
func slugify(stringToSlugify string) string{
	changeSpacesToDashes := strings.NewReplacer(" ", "_")
	finalString := changeSpacesToDashes.Replace(strings.ToLower(stringToSlugify))
	return unidecode.Unidecode(finalString)
}