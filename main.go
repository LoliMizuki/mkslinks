// test flow
// 1. call mkslinks with test json files
// 2. check file is exist?

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

var version = "1.00"

func main() {
	jsonDir, jsonFileName, hasFile := jsonDirAndFileNameFromArgs()
	if hasFile == false {
		printMan()
		return
	}

	if message, err := makeLinksFromJsonPath(jsonDir + "/" + jsonFileName); err != nil {
		fmt.Println(err)
	} else {
		if len(message) > 0 {
			fmt.Println(message)
		}
	}
}

func jsonDirAndFileNameFromArgs() (dir, fileName string, hasFile bool) {
	if len(os.Args) < 2 {
		return "", "", false
	}

	filePath := os.Args[1]
	return filepath.Dir(filePath), filepath.Base(filePath), true
}

func printMan() {
	fmt.Println("version: ", version)
	fmt.Println("Usage: mkslink <define_file_path>")
}

func makeLinksFromJsonPath(jsonPath string) (message string, err error) {
	if isPathExist(jsonPath) == false {
		return "error", errors.New("can not found json file: " + jsonPath)
	}

	linksInfoSetMap := func() map[string]interface{} {
		returnMap := make(map[string]interface{})
		jsonByte, _ := ioutil.ReadFile(jsonPath)
		if err := json.Unmarshal(jsonByte, &returnMap); err != nil {
			fmt.Println("unmarshal fail:", err)
			os.Exit(1)
		}

		return returnMap
	}()

	absDir, _ := filepath.Abs(jsonPath)
	absDir = filepath.Dir(absDir)

	for key, value := range linksInfoSetMap {
		linksSetInfo := newLinksSetInfoFromInterface(key, absDir, value.(map[string]interface{}))
		issueMessage, err := makeSymbolicLinksWithSetInfo(&linksSetInfo)
		if err != nil {
			message += err.Error()
		}

		message += issueMessage
	}

	return message, nil
}

func makeSymbolicLinksWithSetInfo(linksSetInfo *LinksSetInfo) (issueMessage string, err error) {
	if linksSetInfo == nil {
		return "error", errors.New("linksSetInfo is nil")
	}

	for _, fromPathChild := range linksSetInfo.fromPathChildren {
		fromFullPath := linksSetInfo.fromPathParent + "/" + fromPathChild

		if isPathExist(fromFullPath) == false {
			issueMessage += fmt.Sprintf("source path: '%s' is not existed\n", fromFullPath)
			continue
		}

		if isPathExist(linksSetInfo.destPath) == false {
			if err := os.Mkdir(linksSetInfo.destPath, os.ModeDir|os.ModePerm); err != nil {
				return "error", errors.New("cna not make dest dir " + linksSetInfo.destPath)
			}
		}

		cmd := exec.Command("ln", "-s", fromFullPath, linksSetInfo.destPath)
		output, err := cmd.Output()
		if err != nil {
			return "error", err
		} else {
			if output != nil && len(output) > 0 {
				issueMessage += string(output) + "\n"
			}
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
