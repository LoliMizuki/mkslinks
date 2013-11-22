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

type LinksSetInfo struct {
	name             string
	pathType         PathType
	fromPathParent   string
	fromPathChildren []string
	destPath         string
}

func newLinksSetInfoFromInterface(key, absDir string, rawInterface map[string]interface{}) (linksSetInfo LinksSetInfo) {
	linksSetInfo.pathType = func() PathType {
		pathTypeStr := rawInterface["path_type"].(string)
		switch strings.ToLower(pathTypeStr) {
		case "absolute", "a", "abs":
			return kPathType_Absolute

		case "relative", "r", "rel":
			return kPathType_Relative
		}
		return kPathType_Unknow
	}()

	linksSetInfo.name = key

	linksSetInfo.destPath = rawInterface["dest_path"].(string)

	linksSetInfo.fromPathParent = rawInterface["from_path_parent"].(string)

	if linksSetInfo.pathType == kPathType_Relative {
		linksSetInfo.destPath = absDir + "/" + linksSetInfo.destPath
		linksSetInfo.fromPathParent = absDir + "/" + linksSetInfo.fromPathParent
	}

	linksSetInfo.fromPathChildren = func() (fromPathChildren []string) {
		for _, pathInterface := range rawInterface["from_path_children"].([]interface{}) {
			fromPathChildren = append(fromPathChildren, pathInterface.(string))
		}

		return
	}()

	return linksSetInfo
}
