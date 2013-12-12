package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var version = "1.5"

func main() {
	jsonPath, pathsPairAction, showMan := jsonPathAndActionFromArgs()
	if showMan {
		printMan()
		return
	}

	if message, err := applyPathsPairActionToJsonPath(pathsPairAction, jsonPath); err != nil {
		fmt.Println(err)
	} else {
		if len(message) > 0 {
			fmt.Println(message)
		}
	}
}

func jsonPathAndActionFromArgs() (jsonPath string, pathsPairAction PathsPairAction, showMan bool) {
	if len(os.Args) < 3 {
		return "", nil, true
	}

	pathsPairAction = pathsPairActionFromString(os.Args[1])
	jsonPath = os.Args[2]
	showMan = false

	return
}

func printMan() {
	fmt.Println("version: ", version)
	fmt.Println("Usage: path2pathAction [-sl|-c] <define_file_path>")
}

func pathsPairActionFromString(arg string) PathsPairAction {
	switch strings.ToLower(arg) {
	case "-sl":
		return makeSymbolicLinkToPathsPair
	case "-c":
		return makeCopyToPathsPair
	}

	fmt.Println("unknow arg to make action, must be \"-sl(symbolic link)\" or \"-c(copy)\", but you are", arg)
	os.Exit(1)
	return nil
}

func applyPathsPairActionToJsonPath(pathsPairAction PathsPairAction, jsonPath string) (message string, err error) {
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
		pathsPairSetInfo := newPathsPairSetInfoFromInterface(key, absDir, value.(map[string]interface{}))
		issueMessage, err := applyPathsPairActionToSetInfo(pathsPairAction, &pathsPairSetInfo)
		if err != nil {
			message += err.Error()
		}

		message += issueMessage
	}

	return message, nil
}

func applyPathsPairActionToSetInfo(
	pathsPairAction PathsPairAction,
	pathsPairSetInfo *PathsPairSetInfo) (issueMessage string, err error) {
	if pathsPairSetInfo == nil {
		return "error", errors.New("pathsPairSetInfo is nil")
	}

	for _, srcPathChild := range pathsPairSetInfo.srcPathChildren {
		srcFullPath := pathsPairSetInfo.srcPathParent + "/" + srcPathChild
		if isPathExist(srcFullPath) == false {
			issueMessage += fmt.Sprintf("source path: '%s' is not existed\n", srcFullPath)
			continue
		}

		if isPathExist(pathsPairSetInfo.destPath) == false {
			if err := os.Mkdir(pathsPairSetInfo.destPath, os.ModeDir|os.ModePerm); err != nil {
				return "error", errors.New("can not make dest dir " + pathsPairSetInfo.destPath)
			}
		}

		output, err := pathsPairAction(pathsPairSetInfo.srcPathParent, srcPathChild, pathsPairSetInfo.destPath)
		if err != nil {
			return "error", err
		}

		if len(output) > 0 {
			issueMessage += string(output) + "\n"
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
