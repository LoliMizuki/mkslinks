package main

import (
	"strings"
)

type PathType uint

const (
	kPathType_Unknow = iota
	kPathType_Relative
	kPathType_Absolute
)

type PathsPairSetInfo struct {
	name            string
	pathType        PathType
	srcPathParent   string
	srcPathChildren []string
	destPath        string
}

func newPathsPairSetInfoFromInterface(
	key, absDir string, rawInterface map[string]interface{}) (pathsPairSetInfo PathsPairSetInfo) {
	pathsPairSetInfo.pathType = func() PathType {
		pathTypeStr := rawInterface["path_type"].(string)
		switch strings.ToLower(pathTypeStr) {
		case "absolute", "a", "abs":
			return kPathType_Absolute

		case "relative", "r", "rel":
			return kPathType_Relative
		}
		return kPathType_Unknow
	}()

	pathsPairSetInfo.name = key

	pathsPairSetInfo.destPath = rawInterface["dest_path"].(string)

	pathsPairSetInfo.srcPathParent = rawInterface["src_path_parent"].(string)

	if pathsPairSetInfo.pathType == kPathType_Relative {
		pathsPairSetInfo.destPath = absDir + "/" + pathsPairSetInfo.destPath
		pathsPairSetInfo.srcPathParent = absDir + "/" + pathsPairSetInfo.srcPathParent
	}

	pathsPairSetInfo.srcPathChildren = func() (srcPathChildren []string) {
		for _, pathInterface := range rawInterface["src_path_children"].([]interface{}) {
			srcPathChildren = append(srcPathChildren, pathInterface.(string))
		}

		return
	}()

	return
}
