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

// Generates and returns the 2D char array representing the dungeon (X = room, # = hallway)
func (dungeon *Dungeon) getWalls() [][]int {

	// Create empty grid
	grid := make([][]bool, dungeon.height)
	for row := range grid {
		grid[row] = make([]bool, dungeon.width)
		for col := range grid[row] {
			grid[row][col] = false
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
				grid[y][x] = true
			}
		}
	}

	// Add rooms to grid
	for _, room := range dungeon.rooms {
		for x := room.X; x <= room.Width + room.X; x++ {
			for y := room.Y; y <= room.Height+room.Y; y++ {
				grid[y][x] = true
			}
		}
	}

	// Build walls
	walls := make([][]int, 0)
	for y := range grid {
		for x := range grid[y] {
			if !grid[y][x] && ((x+1<dungeon.width && grid[y][x+1]) || (x-1>=0 && grid[y][x-1]) || (y+1<dungeon.height && grid[y+1][x]) || (y-1>=0 && grid[y-1][x])) {
				walls = append(walls, []int{x, y});
			}
		}
	}

	return walls
}