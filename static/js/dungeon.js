var dungeonStage, dungeonLayer, characterLayer, player;

// Move the player a certain distance with checking for invaild spaces
function move(x, y) {
	player.x(player.x()+x);
	player.y(player.y()+y);
	var found = false;
	for (i=0;i<dungeon.walls.length && !found;i++){
		if((dungeon.walls[i].x==Math.trunc((player.x())/scale) || dungeon.walls[i].x==Math.trunc((player.x()+player.getWidth())/scale)) && 
		(dungeon.walls[i].y==Math.trunc((player.y())/scale) || dungeon.walls[i].y==Math.trunc((player.y()+player.getHeight())/scale))){
			found = true;
			player.x(player.x()-x);
			player.y(player.y()-y);
		}
	}
	if(!found){
		characterLayer.draw();
		var viewport = document.getElementById("viewport");
		viewport.scrollLeft = player.getAbsolutePosition().x-viewport.clientWidth/2;
		viewport.scrollTop = player.getAbsolutePosition().y-viewport.clientHeight/2;
	}
}

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
	dungeonStage = new Konva.Stage({
	  container: 'grid',
	  width: dungeon.width*scale,
	  height: dungeon.height*scale
	});
	dungeonLayer = new Konva.Layer();
	characterLayer = new Konva.Layer();
	
	for (var i in dungeon.paths)
		dungeonLayer.add(new Konva.Rect({
						      x: dungeon.paths[i].x*scale,
						      y: dungeon.paths[i].y*scale,
						      width: scale,
						      height: scale,
						      fill: 'green',
						      strokeWidth: 0
						    }));
	
	for (var i in dungeon.rooms)
		dungeonLayer.add(new Konva.Rect({
						      x: dungeon.rooms[i].X*scale,
						      y: dungeon.rooms[i].Y*scale,
						      width: dungeon.rooms[i].Width*scale,
						      height: dungeon.rooms[i].Height*scale,
						      fill: 'red',
						      strokeWidth: 0
						    }));
	
	for (var i in dungeon.walls)
		dungeonLayer.add(new Konva.Rect({
						      x: dungeon.walls[i].x*scale,
						      y: dungeon.walls[i].y*scale,
						      width: scale,
						      height: scale,
						      fill: 'yellow',
						      strokeWidth: 0
						    }));
						    
	var playerImg = new Image();
    playerImg.onload = function() {

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
        frameRate: 7,
        frameIndex: 0,
		          scale: { x:scale/64, y:scale/64 }
      });
      
      // add the shape to the layer
	  characterLayer.add(player);

      // add the layer to the stage
      dungeonStage.add(characterLayer);

      // start sprite animation
      player.start();
      
	  var viewport = document.getElementById("viewport");
	  viewport.scrollLeft = player.getAbsolutePosition().x-viewport.clientWidth/2;
	  viewport.scrollTop = player.getAbsolutePosition().y-viewport.clientHeight/2;
    };
    
    playerImg.src = 'static/images/$placeholder_sprite.png';
	
	dungeonStage.add(dungeonLayer);
}, false);

// Add Key input
document.addEventListener('keydown', function(event) {
    if(event.keyCode == 38 || event.keyCode == 87) {
    	event.preventDefault();
        move(0, -1);
    }
    else if(event.keyCode == 40 || event.keyCode == 83) {
    	event.preventDefault();
        move(0, 1);
    }
    else if(event.keyCode == 37 || event.keyCode == 65) {
    	event.preventDefault();
        move(-1, 0);
    }
    else if(event.keyCode == 39 || event.keyCode == 68) {
    	event.preventDefault();
        move(1, 0);
    }
});