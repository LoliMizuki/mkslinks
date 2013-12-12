package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

type PathsPairAction func(srcPath, dstPath string) (output string, err error)

func makeSymbolicLinkToPathsPair(srcPath, dstPath string) (output string, err error) {
	cmd := exec.Command("ln", "-s", srcPath, dstPath)
	outputBytes, err := cmd.Output()
	if err != nil {
		return "error", err
	}

	return string(outputBytes), nil
}

func makeCopyToPathsPair(srcPath, dstPath string) (output string, err error) {
	dstPathIncludeFileName := dstPath + "/" + filepath.Base(srcPath)

	reader, err := ioutil.ReadFile(srcPath)
	if err != nil {
		return "error", err
	}

	if err = ioutil.WriteFile(dstPathIncludeFileName, reader, os.ModePerm); err != nil {
		return "error", err
	}

	return "", nil
}
