package necrotheif

import (
	"net/http"
	"regexp"
	"fmt"
)

var err error
var pathRegex = regexp.MustCompile("(.*)/(.*)")
var console = ""

func init() {
	http.HandleFunc("/", main)
	err = loadFiles()
}

func main(w http.ResponseWriter, r *http.Request) {
	if(err==nil){
		path := pathRegex.FindStringSubmatch(r.URL.Path)
		loadPage(w, path[1], path[2])
		fmt.Fprint(w, console)
	} else {
		loadError(w, err)
	}
}

func abs(value int) int {
	if (value<0){
		return value*-1
	} else {
		return value
	}
}

func min(value1 int, value2 int) int {
	if(value1<=value2){
		return value1
	} else {
		return value2
	}
}

func max(value1 int, value2 int) int {
	if(value1>value2){
		return value1
	} else {
		return value2
	}
}

func consoleLog(thing interface{}) {
	console += fmt.Sprint(thing)
}