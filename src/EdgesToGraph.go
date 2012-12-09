package main

import (
	"fmt"
	"flag"
	. "./graphtools"
	)

var Filename *string = flag.String("f","edges.gob","filename")
var UseBounds *bool = flag.Bool("b",false,"use bounds?")
var LowerBound *int = flag.Int("l",1,"lower bound") 
var UpperBound *int = flag.Int("u",500,"upper bound")
var Output *string = flag.String("o","output.gob","outputfile")

func main(){
	flag.Parse()
	g := NewGraphFromEdges(*Filename)
	if *UseBounds {
		for key,list := range g.AdjMap {
			if len(list)<*LowerBound || len(list)>*UpperBound {
				g.DeleteNode(key)
			}
		}
	}
	fmt.Println("Save ",len(g.AdjMap)," nodes.")
	g.Save(*Output)
}
