package necrotheif

import (
	"math/rand"
)

// The number of units per room
const UNITS_PER_ROOM = 300

// A dungeon with randomly placed rooms and paths connecting all the rooms
type Dungeon struct{
	width    int
	height   int
	rooms    []Room
	paths    []Path
}

// Generates the dungeon with rooms and paths connecting them
func (dungeon *Dungeon) generate(possibleRooms []Room, numPossibleRooms int){

	// Create the rooms
	dungeon.rooms = make([]Room, dungeon.width*dungeon.height / UNITS_PER_ROOM)
	for i := 0; i < len(dungeon.rooms); i++ {
		curRoom := rand.Intn(numPossibleRooms)
		dungeon.rooms[i] = Room{X:possibleRooms[curRoom].X, Y:possibleRooms[curRoom].Y, Width:possibleRooms[curRoom].Width, Height:possibleRooms[curRoom].Height}
		intersects := true
		for ;intersects; {
			intersects = false
			dungeon.rooms[i].X = rand.Intn(dungeon.width)
			dungeon.rooms[i].Y = rand.Intn(dungeon.height)
			for j := 0; j < len(dungeon.rooms); j++ {
				if j!=i && dungeon.rooms[j].isIntersectingRoom(dungeon.rooms[i]) {
					intersects = true
				}
			}
		}
	}
	
	// Create the paths
	dungeon.paths = make([]Path, 0)
	emptyRoom := true
	for ;emptyRoom; {
		room1 := rand.Intn(len(dungeon.rooms))
		room2 := room1
		for ;dungeon.rooms[room1].NumPaths >= 2; {
			room1 = rand.Intn(len(dungeon.rooms))
		}
		for ;room1 == room2 || dungeon.rooms[room2].NumPaths >= 3; {
			room2 = rand.Intn(len(dungeon.rooms))
		}
		dungeon.paths = append(dungeon.paths, dungeon.rooms[room1].createPathTo(dungeon.rooms[room2]))

		emptyRoom = false
		for i := 0; i < len(dungeon.rooms) && !emptyRoom; i++ {
			if dungeon.rooms[i].NumPaths == 0 {
				emptyRoom = true
			}
		}
	}
}

// Generates and returns the 2D char array representing the dungeon (X = room, # = hallway)
func (dungeon *Dungeon) charArray() [][]string {

	characters := make([][]string, dungeon.width)

	for y := 0; y<dungeon.height; y++ {
		characters[y] = make([]string, dungeon.width)
		for x := 0; x<dungeon.width; x++ {
			contains := false
			for i := 0; i<len(dungeon.rooms) && !contains; i++  {
				if dungeon.rooms[i].contains(x,y) {
					characters[y][x] = "X"
					contains = true
				}
			}
			for i := 0; i<len(dungeon.paths) && !contains; i++ {
				if dungeon.paths[i].contains(x,y) {
					characters[y][x] = "#"
					contains = true
				}
			}
			if !contains {
				characters[y][x] = "."
			}
		}
	}

	return characters
}