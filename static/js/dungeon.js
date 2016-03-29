var dungeonStage, dungeonLayer, characterLayer, player, playerMovements = [], curMovement, dirKeyStack = [];
const DIRECTIONS = ['Down', 'Left', 'Right', 'Up'], CONTROLS = {'38': 3, '87': 3, '40': 0, '83': 0, '37': 1, '65': 1, '39': 2, '68': 2};

// Gets a sprite's current width
Konva.Sprite.prototype.getWidth = function(){
	return this.animations()[this.animation()][this.frameIndex()*4+2]*this.scaleX();
}

// Gets a sprite's current height
Konva.Sprite.prototype.getHeight = function(){
	return this.animations()[this.animation()][this.frameIndex()*4+3]*this.scaleY();
}

// Create the basic grid with the player after the page has loaded
document.addEventListener('DOMContentLoaded', function() {
	
	// Set all keys to false
	keys = [];
	for(var i=0;i<=222;i++)
		keys[i] = false;
	
	// Create the stage and layers
	dungeonStage = new Konva.Stage({
	  container: 'grid',
	  width: dungeon.width*scale,
	  height: dungeon.height*scale
	});
	dungeonLayer = new Konva.Layer();
	characterLayer = new Konva.Layer();
	
	// Create all the paths first since rooms should be over them
	for (var i in dungeon.paths)
		dungeonLayer.add(new Konva.Rect({
						      x: dungeon.paths[i].x*scale,
						      y: dungeon.paths[i].y*scale,
						      width: scale,
						      height: scale,
						      fill: 'green',
						      strokeWidth: 0
						    }));
	
	// Create all the rooms next
	for (var i in dungeon.rooms)
		dungeonLayer.add(new Konva.Rect({
						      x: dungeon.rooms[i].X*scale,
						      y: dungeon.rooms[i].Y*scale,
						      width: dungeon.rooms[i].Width*scale,
						      height: dungeon.rooms[i].Height*scale,
						      fill: 'red',
						      strokeWidth: 0
						    }));
	
	// Create all the walls last
	for (var i in dungeon.walls)
		dungeonLayer.add(new Konva.Rect({
						      x: dungeon.walls[i].x*scale,
						      y: dungeon.walls[i].y*scale,
						      width: scale,
						      height: scale,
						      fill: 'yellow',
						      strokeWidth: 0
						    }));
						      
	
	// Add the layer with all the background to the stage (rooms, paths, walls, etc.)
	dungeonStage.add(dungeonLayer);
						    
	// Create the sprite for the player
	var playerImg = new Image();
    playerImg.onload = function() {

	  // Actually create the player's sprite with animations
      player = new Konva.Sprite({
        x: xPos,
        y: yPos,
        image: playerImg,
        animation: 'idleDown',
        animations: {
				      idleDown: [
				        32, 0, 32, 32
				      ],
				      moveDown: [
				      	0, 0, 32, 32,
				        32, 0, 32, 32,
				        64, 0, 32, 32,
				        32, 0, 32, 32
				      ],
				      idleLeft: [
				      	32, 32, 32, 32
				      ],
				      moveLeft: [
				      	0, 32, 32, 32,
				        32, 32, 32, 32,
				        64, 32, 32, 32,
				        32, 32, 32, 32
				      ],
				      idleRight: [
				      	32, 64, 32, 32
				      ],
				      moveRight: [
				      	0, 64, 32, 32,
				        32, 64, 32, 32,
				        64, 64, 32, 32,
				        32, 64, 32, 32
				      ],
				      idleUp: [
				      	32, 96, 32, 32
				      ],
				      moveUp: [
				      	0, 96, 32, 32,
				        32, 96, 32, 32,
				        64, 96, 32, 32,
				        32, 96, 32, 32
				      ]
				    },
        frameRate: 10,
        frameIndex: 0,
		          scale: { x:scale/64, y:scale/64 }
      });
      
      // add the shape to the layer
	  characterLayer.add(player);

      // add the layer to the stage
      dungeonStage.add(characterLayer);
      
	  // Update the viewport
	  updateViewport();
    };
    playerImg.src = 'static/images/$placeholder_sprite.png';
    
    // Create the player movement animations
    playerMovements.push(new Konva.Animation(function(frame) {
	        move(player, 0, frame.timeDiff*playerSpeed);
	        updateViewport();
	    }, characterLayer));
	playerMovements.push(new Konva.Animation(function(frame) {
	        move(player, 1, frame.timeDiff*playerSpeed);
	        updateViewport();
	    }, characterLayer));
	playerMovements.push(new Konva.Animation(function(frame) {
	        move(player, 2, frame.timeDiff*playerSpeed);
	        updateViewport();
	    }, characterLayer));
	playerMovements.push(new Konva.Animation(function(frame) {
	        move(player, 3, frame.timeDiff*playerSpeed);
	        updateViewport();
	    }, characterLayer));
	
}, false);

// Move the given sprite a certain direction with checking for invaild spaces
function move(sprite, dir, distance) {
	var x = dir==1 ? -distance : (dir==2 ? distance : 0);
	var y = dir==0 ? distance : (dir==3 ? -distance : 0);
	sprite.x(sprite.x()+x);
	sprite.y(sprite.y()+y);
	for (var i=0, found=false;i<dungeon.walls.length && !found;i++){
		if((dungeon.walls[i].x==Math.trunc((sprite.x())/scale) || dungeon.walls[i].x==Math.trunc((sprite.x()+sprite.getWidth())/scale)) && 
		(dungeon.walls[i].y==Math.trunc((sprite.y()+sprite.getHeight()/2)/scale) || dungeon.walls[i].y==Math.trunc((sprite.y()+sprite.getHeight())/scale))){
			found = true;
			sprite.x(sprite.x()-x);
			sprite.y(sprite.y()-y);
		}
	}
}

// Udate the player's viewport onto the player
function updateViewport(){
	var viewport = document.getElementById("viewport");
	viewport.scrollLeft = player.getAbsolutePosition().x-viewport.clientWidth/2;
	viewport.scrollTop = player.getAbsolutePosition().y-viewport.clientHeight/2;
}

// Move the player a certain direction with checking for invaild spaces and playing an animation (0 - down, 1 - left, 2 - right, 3 - down)
function movePlayer(dir){
	
	// if not animation set to moving direction set it and start the animation
	if(player.animation()!='move'+DIRECTIONS[dir]){
		var start = player.animation().startsWith('idle');
		player.setAnimation('move'+DIRECTIONS[dir]);
		if(start)
			player.start();
		if(curMovement)
			curMovement.stop();
		curMovement = playerMovements[dir];
		playerMovements[dir].start();
	}
}

// Stop the player moving in the given direction
function stopPlayer(){
	if(curMovement)
		curMovement.stop();
	player.setAnimation('idle'+player.animation().substring('move'.length, player.animation().length-1));
	player.stop();
}

// Add Key input
document.addEventListener('keydown', function(event) {
	if(CONTROLS[event.keyCode.toString()]!=null){
    	event.preventDefault();
		addKeyToStack(event.keyCode);
	}
});
document.addEventListener('keyup', function(event) {
	if(CONTROLS[event.keyCode.toString()]!=null){
    	event.preventDefault();
		removeKeyFromStack(event.keyCode);
	}
});

function addKeyToStack(key){
	if(dirKeyStack.indexOf(key)==-1)
		dirKeyStack.push(key);
	if(CONTROLS[key.toString()]!=CONTROLS[dirKeyStack[dirKeyStack.length-2]])
		movePlayer(CONTROLS[key.toString()]);
}

function removeKeyFromStack(key){
	if(dirKeyStack.indexOf(key)==-1)
		return;
	if(dirKeyStack.indexOf(key)==dirKeyStack.length-1 && dirKeyStack.length>1)
		movePlayer(CONTROLS[dirKeyStack[dirKeyStack.length-2]]);
	dirKeyStack.splice(dirKeyStack.indexOf(key), 1);
	if(dirKeyStack.length==0)
		stopPlayer();
}