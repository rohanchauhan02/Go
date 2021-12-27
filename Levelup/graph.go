package main

import (
	"fmt"
)
//Graph represents an adjacency list graph

type Graph struct{
	vertices []*Vertex
}
//Vertex represents a graph vertex

type Vertex struct{
	key int
	adjacent []* Vertex
}

//Add Vertex
//Add Edge

//getVertex
//contains