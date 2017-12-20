package main

import (
        "flag"
        "fmt"
        "encoding/json"
)

const (
	str = "version"
)

type V interface {
        getVersion() string
}

type Version struct {
	major, minor, patch int
}

func (v Version) getVersion() string {
	return fmt.Sprintf("%d.%d.%d", v.major, v.minor, v.patch)
}

func printJson(aMap map[string]string) {
        mapB, _ := json.Marshal(aMap)
        fmt.Println(string(mapB))
}

func main() {
        var v V
	var x = flag.Int("major", 0, "Major Version")
	var y = flag.Int("minor", 0, "Minor Version")
	var z = flag.Int("patch", 1, "Patch Version")
        flag.Parse()
        v = Version{*x,*y,*z}
	var myVersion map[string]string
        myVersion = make(map[string]string)
        myVersion[str] = v.getVersion()
	printJson(myVersion)
}
