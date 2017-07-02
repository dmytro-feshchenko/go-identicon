package main

// Identicon - contains the structure of each identicon
// such as name, hash and color
type Identicon struct {
	name  string   // received name for identicon generation
	hash  [16]byte // contains hashed value of name
	color [3]byte  // contains red, green and blue colors
	grid  []byte   // contains grid for identicon
}

// BuildGrid - build the symmetric grid for drawing identicon
func BuildGrid(identicon Identicon) Identicon {
	grid := []byte{}

	// go through the hash array with step as 3 items
	// to prevent going beyond the array bounds
	// as the result we will have grid with 5 lines with 3 items in each line
	for i := 0; i < len(identicon.hash) && i+3 < len(identicon.hash)-1; i += 3 {
		chunk := make([]byte, 5)
		// copy items to the new array
		copy(chunk, identicon.hash[i:i+3])
		// mirror the first 2 items to the last items
		// to receive the completed symmetric array with 5 items
		chunk[3] = chunk[1]
		chunk[4] = chunk[0]
		// append new line to the grid array
		grid = append(grid, chunk...)
	}

	identicon.grid = grid
	// return modified identicon
	return identicon
}
