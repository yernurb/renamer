package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	dirPath := "./PhotoBase/"
	var fileNames []string
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	curNameID := ""
	nameID := ""
	newDir := ""
	newName := ""

	uniqueID := 0
	numOfFiles := len(files)
	numOfFemales := 0
	numOfMales := 0
	var histNationality [500]int
	var histAge [500]int

	for i, file := range files {
		for j := 0; j < 100000000; j++ {
		}
		fileNames = append(fileNames, file.Name())
		nameSplit := strings.Split(file.Name(), "_")
		nameID = nameSplit[0]
		newName = strings.Join(nameSplit[1:4], "_")
		id, _ := strconv.Atoi(nameSplit[2])
		histNationality[id]++

		id, _ = strconv.Atoi(strings.Split(nameSplit[3], ".")[0])
		histAge[id]++
		if curNameID != nameID {
			if nameSplit[1] == "F" {
				numOfFemales++
			} else {
				numOfMales++
			}
			curNameID = nameID
			uniqueID++
			newDir = dirPath + "Person_" + strconv.Itoa(uniqueID)
			os.Mkdir(newDir, 0777)
			os.Rename(dirPath+file.Name(), newDir+"/"+newName)
		} else {
			os.Rename(dirPath+file.Name(), newDir+"/"+newName)
			//fmt.Println(uniqueID)
		}
		fmt.Printf("\r%d%%", 100*i/numOfFiles)
	}
	fmt.Println(" ")
	fmt.Println(numOfFiles)
	fmt.Println(numOfFemales)
	fmt.Println(numOfMales)
	fmt.Println(uniqueID)
	fmt.Println(histAge)
	fmt.Println(histNationality)

	/*
		currentFP := ""
		items := 0
		for _, fp := range firstPart {
			if currentFP != fp {
				currentFP = fp
				items++
				os.Mkdir("./test/"+string(currentFP), 0777)
				os.Rename()
			} else {
				fmt.Println(fp)
			}
		}
		fmt.Println(items)
		os.Mkdir("."+string(filepath.Separator)+"testDir", 0777)
	*/
}
