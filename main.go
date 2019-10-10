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
		fmt.Println(workFolder)
	}
	for _, file := range files {
		trueFilename := strings.TrimSuffix(file.Name(), "\r")
		newName := slugify(trueFilename)
		fmt.Println(trueFilename)
		fmt.Println(newName)
		os.Rename(trueFilename, workFolder+"/"+newName)
	}
	fmt.Println("Done")
}
func slugify(stringToSlugify string) string{
	changeSpacesToDashes := strings.NewReplacer(" ", "_")
	finalString := changeSpacesToDashes.Replace(strings.ToLower(stringToSlugify))
	return unidecode.Unidecode(finalString)
}