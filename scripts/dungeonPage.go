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
		gridValues := GridTemplate{template.JS(arrayToJavascript(dungeon.getGrid())), template.JS(strconv.Itoa(dungeon.startX)), template.JS(strconv.Itoa(dungeon.startY)), styleSheet}
		err := gridTemplate.Execute(w, gridValues)
		fmt.Fprint(w, err)
}

// Converts a 2D array to a javascript 2D array as a string
func dungeonToJavascript(dungeon Dungeon) string{
	return "{width:"+strconv.Itoa(dungeon.width)+", height:"+strconv.Itoa(dungeon.height)+", grid:"+arrayToJavascript(dungeon.getGrid())+"}"
}

func arrayToJavascript(array [][]int) string{
	javascript := "["
	for _, row := range array {
		javascript += "["
		for _, element := range row {
			javascript += strconv.Itoa(element)+","
		}
		javascript = javascript[:len(javascript)-1]+"],"
	}
	return javascript[:len(javascript)-1]+"]"
}