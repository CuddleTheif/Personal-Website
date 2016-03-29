package necrotheif

import (
	"math/rand"
)

// A dungeon with randomly placed rooms and paths connecting all the rooms
type Dungeon struct{
	width    int
	height   int
	rooms    []Room
	paths    []Path
	startX   int
	startY   int
}

const MAX_PATHS = 2

// Generates the dungeon with rooms and paths connecting them
func (dungeon *Dungeon) generate(possibleRooms []Room, numRooms int){

	// Create the rooms
		dungeon.rooms = make([]Room, numRooms)
		for i := 0; i < len(dungeon.rooms); i++ {
			curRoom := rand.Intn(len(possibleRooms))
			dungeon.rooms[i] = possibleRooms[curRoom]
			intersects := true
			for ;intersects; {
				intersects = false
				dungeon.rooms[i].X = rand.Intn(dungeon.width-dungeon.rooms[i].Width-2)+1
				dungeon.rooms[i].Y = rand.Intn(dungeon.height-dungeon.rooms[i].Height-2)+1
				for j := 0; j < len(dungeon.rooms); j++ {
					if j!=i && dungeon.rooms[j].isIntersectingRoom(dungeon.rooms[i]) {
						intersects = true
					}
				}
			}
		}

	// Get the start position in the middle of a random room
		startRoom := dungeon.rooms[rand.Intn(len(dungeon.rooms))]
		dungeon.startX = startRoom.X+startRoom.Width/2
		dungeon.startY = startRoom.Y+startRoom.Height/2
	
	// Create the paths
		dungeon.paths = make([]Path, 0)
		emptyRoom := true
		for ;emptyRoom; {
			room1 := rand.Intn(len(dungeon.rooms))
			room2 := room1
			for ;dungeon.rooms[room1].NumPaths >= MAX_PATHS; {
				room1 = rand.Intn(len(dungeon.rooms))
			}
			for ;room1 == room2 || dungeon.rooms[room2].NumPaths >= MAX_PATHS; {
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

type Point struct {
	x	int
	y	int
}

// Generates and returns the 2D char array representing the dungeon (-1 = nothing, 0 = wall, 1 = room, 2 = hallway)
func (dungeon *Dungeon) getGrid() [][]int {

	// Create empty grid
	grid := make([][]int, dungeon.width)
	for row := range grid {
		grid[row] = make([]int, dungeon.height)
		for col := range grid[row] {
			grid[row][col] = -1
		}
	}

	// Add paths to grid
	for _, path := range dungeon.paths {
		for _, segment := range path.segments {
			for j := 0; abs(j) < abs(segment.distance); j+=segment.distance/abs(segment.distance) {
				var x, y int
				if segment.direction {
					x = segment.startX
					y = segment.startY + j
				} else {
					x = segment.startX + j
					y = segment.startY
				}
				grid[x][y] = 2
			}
		}
	}

	// Add rooms to grid
	for _, room := range dungeon.rooms {
		for x := room.X; x < room.Width + room.X; x++ {
			for y := room.Y; y < room.Height+room.Y; y++ {
				grid[x][y] = 1
			}
		}
	}

	// Build walls
	for x := range grid {
		for y := range grid[x] {
			if(y<len(grid[x])-1 && grid[x][y]==-1 && grid[x][y+1]>0){
				grid[x][y] = 0
			}
		}
	}

	return grid
}