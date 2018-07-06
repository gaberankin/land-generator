var gen = require('random-seed');

var initTerrain = function(w, h) {
};

var Terrain = function(options) {
	this.options = {
		width: (isNaN(options.width) || options.width <= 0) ? 10 : options.width,
		height: (isNaN(options.height) || options.height <= 0) ? 10 : options.height,
		seed: typeof options.seed === 'undefined' ? Math.random() * 10000 : options.seed,
		hills: (isNaN(options.hills) || options.hills < 0) ? 0 : options.hills,
		min: isNaN(options.min) ? 0 : options.min,
		max: isNaN(options.max) ? 10 : options.max,
		flattening: isNaN(options.flattening) || options.flattening <= 0 ? 0 : options.flattening,
		island: typeof options.island === 'undefined' ? false : options.island
	};

	this.cellValMin = null;
	this.cellValMax = null;

	this.init();
	this.rand = gen.create(this.options.seed);

	//do stuff
};

Terrain.prototype.init = function(){
	//create array. 0 is 'sea level'
	this.terrain = [];
	for(var i = 0; i < this.options.height; i++) {
		this.terrain[i] = [];
		for(var j = 0; j < this.options.width; j++) {
			this.setCellValue(j, i, 0);
		}
	}
};

Terrain.prototype.generate = function(){
	for(var h = 0; h < this.options.hills; h++) {
		this.addHill();
	}
	this.normalize();
	this.flatten();
};

Terrain.prototype.getCellValue = function(x, y) {
	if(x >= 0 && y >= 0 && this.options.height > y && this.options.width > x){
		return this.terrain[y][x];
	}
	return null;
};

Terrain.prototype.setCellValue = function(x, y, val) {
	if(x >= 0 && y >= 0 && this.options.height > y && this.options.width > x){
		if(this.cellValMin === null || this.cellValMin > val) {
			this.cellValMin = val;
		}
		if(this.cellValMax === null || this.cellValMax < val) {
			this.cellValMax = val;
		}
		this.terrain[y][x] = val;
	}
};

Terrain.prototype.clear = function() {
	for(var i = 0; i < this.options.height; i++) {
		for(var j = 0; j < this.options.width; j++) {
			this.setCellValue(j, i, 0);
		}
	}
};

Terrain.prototype.addHill = function(){
	var radius = this.rand.intBetween(this.options.min, this.options.max),
		minDimension = this.options.width < this.options.height ? this.options.width : this.options.height;
	var x, y, distSq, height, xMin, xMax, yMin, yMax,
		radiusSq = radius * radius;
	if(this.options.island) {
		var theta = this.rand.floatBetween(0, 6.28),
			distance = this.rand.floatBetween(radius/2, minDimension/2 - radius);
		x = Math.round(minDimension / 2 + Math.cos(theta) * distance);
		y = Math.round(minDimension / 2 + Math.sin(theta) * distance);
	} else {
		x = Math.round(this.rand.floatBetween(-radius, this.options.width + radius));
		y = Math.round(this.rand.floatBetween(-radius, this.options.height + radius));
	}

	xMin = x - radius - 1;
	xMax = x + radius + 1;
	if(xMin < 0) xMin = 0;
	if(xMax >= this.options.width) xMax = this.options.width - 1;

	yMin = y - radius - 1;
	yMax = y + radius + 1;
	if(yMin < 0) yMin = 0;
	if(yMax >= this.options.height) yMax = this.options.height - 1;

	for(var i = xMin; i <= xMax; ++i) {
		for(var j = yMin; j <= yMax; ++j) {
			distSq = (x - i) * (x - i) + (y - j) * (y - j);
			height = radiusSq - distSq;
			if(height > 0){
				this.setCellValue(i, j, this.getCellValue(i, j) + height);
			}
		}
	}
};

Terrain.prototype.normalize = function(){
	// avoiding divide by zero (unlikely with floats, but just in case)
	if( this.cellValMax != this.cellValMin ) {
		// divide every height by the maximum to normalize to ( 0.0, 1.0 )
		for( var x = 0; x < this.options.width; ++x ) {
			for( var y = 0; y < this.options.height; ++y ) {
				this.setCellValue( x, y, Math.round(10*(( this.getCellValue( x, y ) - this.cellValMin ) / ( this.cellValMax - this.cellValMin )) ));
			}
		}
	}
	else
	{
		// if the min and max are the same, then the terrain has no height, so just clear it
		// to 0.0.
		this.clear();
	}


};

Terrain.prototype.flatten = function(){
	// if flattening is one, then nothing would be changed, so just skip the
	// process altogether.
	if( this.options.flattening > 1 ) {
		var flat, original;
		for( var x = 0; x < this.options.width; ++x ) {
			for( var y = 0; y < this.options.height; ++y ) {
				flat = 1.0;
				original = this.getCellValue( x, y );
				
				// flatten as many times as desired
				for( var i = 0; i < this.options.flattening; ++i ) {
					flat *= original;
				}
				
				// put it back into the cell
				this.setCellValue( x, y, flat );
			}
		}
	}

};

Terrain.prototype.print = function(){
	for(var i = 0, h = this.terrain.length; i < h; i++){
		console.log.apply(console, this.terrain[i]);
	}
};

t = new Terrain({ width:40, height:40, min: -1, max:20, hills: 10});

t.print();

console.log('\n---------------------------\n');
t.generate();

t.print();
