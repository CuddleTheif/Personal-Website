package necrotheif

import (
	"net/http"
	"math/rand"
	"strconv"
	"html/template"
"fmt"
)

var gridTemplate *template.Template

type GridTemplate struct {
	Dungeon    		template.JS
	PlayerX			template.JS
	PlayerY			template.JS
	StyleSheet		template.CSS
}

func init() {
	http.HandleFunc("/dungeon", dungeon)
	gridTemplate, err = template.ParseFiles("templates/dungeon.html")
}

// Display a new dungeon or a given one
func dungeon(w http.ResponseWriter, r *http.Request) {

	// Create the rooms for the new dungeon
		rooms := make([]Room, 10)
		for i := range rooms {
			rooms[i] = Room{X:0, Y:0, Width:rand.Intn(4)+4, Height:rand.Intn(4)+4}
		}

	// Create the dungeon
		dungeon := Dungeon{width:200, height:100}
		dungeon.generate(rooms, 20)

	// Add the grid add buttons to the actual page (and draw the player in the grid)
		gridValues := GridTemplate{template.JS(dungeonToJavascript(dungeon)), template.JS(strconv.Itoa(dungeon.startX)), template.JS(strconv.Itoa(dungeon.startY)), styleSheet}
		err := gridTemplate.Execute(w, gridValues)
		fmt.Fprint(w, err)
}

// Converts a 2D array to a javascript 2D array as a string
func dungeonToJavascript(dungeon Dungeon) string{
	javascript := "{width:"+strconv.Itoa(dungeon.width)+", height:"+strconv.Itoa(dungeon.height)+", walls:"+arrayToJavascript(dungeon.getWalls())+",paths:"+arrayToJavascript(dungeon.getPaths())+",rooms:["
	for _, room := range dungeon.rooms {
		javascript += roomToJavascript(room)+","
	}
	return javascript[:len(javascript)-1]+"]}"
}

func roomToJavascript(room Room) string{
	return "{X:"+strconv.Itoa(room.X)+", Y:"+strconv.Itoa(room.Y)+", Width:"+strconv.Itoa(room.Width)+", Height:"+strconv.Itoa(room.Height)+"}"
}

func pathToJavascript(path Path) string{
	javascript := "{startX:"+strconv.Itoa(path.startX)+", startY:"+strconv.Itoa(path.startY)+", endX:"+strconv.Itoa(path.endX)+", endY:"+strconv.Itoa(path.endY)+", segments:["
	for _, segment := range path.segments {
		javascript += segmentToJavascript(segment)+","
	}
	return javascript[:len(javascript)-1]+"]}"
}

func segmentToJavascript(segment Segment) string{
	return "{startX:"+strconv.Itoa(segment.startX)+", startY:"+strconv.Itoa(segment.startY)+", distance:"+strconv.Itoa(segment.distance)+", direction:"+strconv.FormatBool(segment.direction)+"}"
}

func arrayToJavascript(array []Point) string{
	javascript := "["
	for _, element := range array {
		javascript += "{x:"+strconv.Itoa(element.x)+",y:"+strconv.Itoa(element.y)+"},"
	}
	return javascript[:len(javascript)-1]+"]"
}