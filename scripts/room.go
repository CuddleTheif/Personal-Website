package necrotheif

import (
	"math/rand"
)

// A room in dungeon generation
type Room struct {
	X        int
	Y        int
	Width    int
	Height   int
	NumPaths int
}

// Checks if the given room intersects the room
func (room1 *Room) isIntersectingRoom(room2 Room) bool {
	return room1.X <= room2.X+room2.Width && room1.X+room1.Width >= room2.X && room1.Y <= room2.Y+room2.Height && room1.Y+room1.Height >= room2.Y
}

// Checks if the given position is in the room
func (room *Room) contains(x int, y int) bool {
	return x >= room.X && x <= room.X+room.Width && y>=room.Y && y<=room.Y+room.Height
}

// Creates a path to the given room from this room
func (room1 *Room) createPathTo(room2 Room) Path {
	var startX, startY, endX, endY int
	if (rand.Intn(2) == 0){
		startX = room1.X + rand.Intn(2)*room1.Width
		startY = room1.Y + rand.Intn(room1.Height-2) + 1
	} else {
		startX = room1.X + rand.Intn(room1.Width-2) + 1
		startY = room1.Y + rand.Intn(2)*room1.Height
	}
	if (rand.Intn(2) == 0){
		endX = room2.X + rand.Intn(2)*room2.Width
		endY = room2.Y + rand.Intn(room2.Height-2) + 1
	} else {
		endX = room2.X + rand.Intn(room2.Width-2) + 1
		endY = room2.Y + rand.Intn(2)*room2.Height
	}
	room1.NumPaths++
	room2.NumPaths++
	path := Path{startX:startX, startY:startY, endX:endX, endY:endY}
	path.generate()
	return path
}