package necrotheif

import (
	"net/http"
	"fmt"
	"math/rand"
)

func init() {
	http.HandleFunc("/dungeon", dungeon)
}

func dungeon(w http.ResponseWriter, r *http.Request) {
	rooms := make([]Room, 10)
	for i := range rooms {
		rooms[i] = Room{X:0, Y:0, Width:rand.Intn(5)+3, Height:rand.Intn(5)+3}
	}

	dungeon := Dungeon{width:100, height:30}
	dungeon.generate(rooms, 10)
	for _, charArray := range dungeon.charArray() {
		line := ""
		for _, character := range charArray {
			line += character
		}
		fmt.Fprintln(w, line)
	}
}