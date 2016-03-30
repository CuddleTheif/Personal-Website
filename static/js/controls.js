var dirKeyStack = [], playerMovements = [], curMovement;
const DIRECTIONS = ['Down', 'Left', 'Right', 'Up'], CONTROLS = {'38': 3, '87': 3, '40': 0, '83': 0, '37': 1, '65': 1, '39': 2, '68': 2};

document.addEventListener('DOMContentLoaded', function() {
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
});


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