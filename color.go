package main

// PickColors - pick color for the identicon
// uses first 3 bytes as red, green and blue values accordingly
func PickColors(identicon Identicon) Identicon {
	rgb := [3]byte{}

	// copy first 3 bytes to the rbg array
	copy(rgb[:], identicon.hash[:3])

	identicon.color = rgb
	// return changed identicon
	return identicon
}
