package main

import (
	"io/ioutil"
	"os"
	"os/exec"
)

type PathsPairAction func(srcParent, srcChild, dstPath string) (output string, err error)

func makeSymbolicLinkToPathsPair(srcParent, srcChild, dstPath string) (output string, err error) {
	srcPath := srcParent + "/" + srcChild

	cmd := exec.Command("ln", "-s", srcPath, dstPath)
	outputBytes, err := cmd.Output()
	if err != nil {
		return "error", err
	}

	return string(outputBytes), nil
}

func makeCopyToPathsPair(srcParent, srcChild, dstPath string) (output string, err error) {
	srcPath := srcParent + "/" + srcChild
	dstFullPath := dstPath + "/" + srcChild

	reader, err := ioutil.ReadFile(srcPath)
	if err != nil {
		return "error", err
	}

	if err = ioutil.WriteFile(dstFullPath, reader, os.ModePerm); err != nil {
		return "error", err
	}

	return "", nil
}
