package necrotheif

import (
	"math/rand"
)

// The number of units per segment of path
const UNITS_PER_SEG = 1

// A segment in a path in dungeon generation
type Segment struct {
	startX    int
	startY    int
	distance  int
	direction bool
}

// A path in dungeon generation that connects two rooms
type Path struct {
	startX   int
	startY   int
	endX     int
	endY     int
	segments []Segment
}

// Generates the segments of a path and stores them
func (path *Path) generate() {

	// Create the variables for number of segments in each direction and the direction of the path
	var xSeg, ySeg, dirX, dirY int
	if(path.startX==path.endX){
		xSeg = 0
	} else {
		xSeg = rand.Intn(abs((path.startX - path.endX) / UNITS_PER_SEG)) + 1
	}
	if(path.startY==path.endY){
		ySeg = 0
	} else {
		ySeg = rand.Intn(abs((path.startY - path.endY) / UNITS_PER_SEG)) + 1
	}
	curX := path.startX
	curY := path.startY
	if(path.startX<path.endX){
		dirX = 1
	} else {
		dirX = -1
	}
	if(path.startY<path.endY){
		dirY = 1
	} else {
		dirY = -1
	}
	path.segments = make([]Segment, xSeg + ySeg)

	// Create the segments with a random order
	for i := 0; i < len(path.segments); i++ {
		path.segments[i].startX = curX
		path.segments[i].startY = curY
		if ((rand.Intn(2) == 0 && xSeg>0) || ySeg<=0) {
			if xSeg == 1{
				path.segments[i].distance = abs(curX - path.endX)*dirX
			} else {
				path.segments[i].distance = (rand.Intn(abs(curX - path.endX) - (xSeg-1) + 1) + 1)*dirX
			}
			xSeg--
			path.segments[i].direction = false
			curX += path.segments[i].distance
		} else {
			if ySeg == 1{
				path.segments[i].distance = abs(curY - path.endY)*dirY
			} else {
				path.segments[i].distance = (rand.Intn(abs(curY - path.endY) - (ySeg-1) + 1) + 1)*dirY
			}
			ySeg--
			path.segments[i].direction = true
			curY += path.segments[i].distance
		}
	}	
}