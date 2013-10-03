package main

// todo: PrintMan(): read readme.md 作為預設的 man 輸出
// test case

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

var version = "0.9"

type SLinksCreateInfo struct {
	fileNames []string
	fromPath  string
	toPath    string
}

func main() {
	slinksJsonPath := GetSymbolicLinkDefineJsonPath()
	if len(slinksJsonPath) == 0 {
		PrintMan()
		return
	}

	sLinksCreateInfo, createInfoErr := NewSLinksCreateInfoFromJsonPath(slinksJsonPath)
	if createInfoErr != nil {
		fmt.Println(createInfoErr)
		return
	}

	fmt.Println("use ", slinksJsonPath, " to create symbolic links")
	fmt.Println("---------------------")
	fmt.Println("files:", sLinksCreateInfo.fileNames)
	fmt.Println("to path: ", sLinksCreateInfo.toPath)
	fmt.Println("---------------------")

	result, makeLinksErr := makeSymbolicLinksWithInfo(sLinksCreateInfo)
	if makeLinksErr != nil {
		fmt.Println(makeLinksErr)
		return
	}
	fmt.Println(result)
	fmt.Println(" ... make links success")
}

func GetSymbolicLinkDefineJsonPath() string {
	if len(os.Args) < 2 {
		return ""
	}

	filePath := os.Args[1]
	fileDir := filepath.Dir(filePath)

	if len(fileDir) > 0 && fileDir != "." {
		return filePath
	}

	fileName := filepath.Base(filePath)
	currentWorkingDir, err := os.Getwd()
	if err != nil {
		panic("can not get current working directory")
		return ""
	}

	fmt.Println(currentWorkingDir)

	return currentWorkingDir + "/" + fileName
}

func PrintMan() {
	fmt.Println("version: ", version)
	fmt.Println("Usage: mkslink <define_file_path>")
}

func NewSLinksCreateInfoFromJsonPath(jsonPath string) (*SLinksCreateInfo, error) {
	if isPathExist(jsonPath) == false {
		return nil, errors.New("can not find file " + jsonPath)
	}

	jsonByte, _ := ioutil.ReadFile(jsonPath)
	var mapFromJson map[string]interface{}
	json.Unmarshal(jsonByte, &mapFromJson)

	var fileNames []string
	for _, fnameInterface := range mapFromJson["files"].([]interface{}) {
		fileNames = append(fileNames, fnameInterface.(string))
	}

	fromPath := filepath.Dir(jsonPath)
	toPath := mapFromJson["toPath"].(string)

	sLinksCreateInfo := SLinksCreateInfo{fileNames, fromPath, toPath}
	return &sLinksCreateInfo, nil
}

func makeSymbolicLinksWithInfo(sLinksInfo *SLinksCreateInfo) (string, error) {
	if sLinksInfo == nil {
		return "error", errors.New("SLinksCreateInfo is nil")
	}

	var issueMessage string = ""
	for _, fileName := range sLinksInfo.fileNames {
		fromFullPath := sLinksInfo.fromPath + "/" + fileName
		toFullPath := sLinksInfo.toPath + "/" + fileName

		if isPathExist(toFullPath) {
			issueMessage += fmt.Sprintln("target path: '", toFullPath, "' is existed")
			continue
		}
		if isPathExist(fromFullPath) == false {
			issueMessage += fmt.Sprintln("source path: '", fromFullPath, "' is not existed")
			continue
		}

		cmd := exec.Command("ln", "-s", fromFullPath, toFullPath)
		output, err := cmd.Output()
		if err == nil {
			if output != nil && len(output) > 0 {
				issueMessage += string(output)
				issueMessage += "\n"
			}
		} else {
			return "error", err
		}
	}

	return issueMessage, nil
}

func isPathExist(path string) bool {
	_, pathErr := os.Stat(path)

	if pathErr != nil {
		return false
	}

	return true
}
