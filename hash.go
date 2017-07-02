package main

import "crypto/md5"

// GenerateHash - transforms array of bytes to the Identicon
// object using Sum function from crypto/md5 package
func GenerateHash(text []byte) Identicon {
	// generate check sum from text
	checkSum := md5.Sum(text)
	// return the Identicon object
	return Identicon{
		name: string(text),
		hash: checkSum,
	}
}
