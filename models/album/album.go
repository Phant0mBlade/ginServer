// album.go
/*
I am defining this schema in 2 places
1. inside schema
	- all subsequent schemas for album and its relation will be inside this file
2. inside schema/album
	- all subsequent schemas for album and its relation will be in different files inside mentioned folder
*/

package albumschema

// a Album entity
type Album struct {
	ID        string `json:"id"`
	FirstName string `json:"title"`
	LastName  string `json:"artist"`
}
