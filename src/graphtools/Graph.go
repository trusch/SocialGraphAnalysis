package graphtools

import (
	"log"
	"os"
	"encoding/gob"

)

type Edge struct {
	A,B	uint32
}

type Edges struct {
	Items	[]*Edge
}

func (e *Edges)Save(filename string){
	f,err := os.Create(filename)
	if err!=nil {
		log.Print(err)
		return
	}
	defer f.Close()
	encoder := gob.NewEncoder(f)
	encoder.Encode(e)
}

func (e *Edges)Load(filename string){
	f,err := os.Open(filename)
	if err!=nil {
		log.Print(err)
		return
	}
	defer f.Close()
	decoder := gob.NewDecoder(f)
	decoder.Decode(&e)
}

func ParseEdges(e *Edges) map[uint32][]uint32 {
	result := make(map[uint32][]uint32)
	for _,edge := range e.Items {
		result[edge.A] = append(result[edge.A],edge.B)
		result[edge.B] = append(result[edge.B],edge.A)
	}
	return result
}


type Graph struct {
	AdjMap	map[uint32][]uint32
	Keys	[]uint32	
}

func NewGraph(adj map[uint32][]uint32) *Graph {
	g := new(Graph)
	g.AdjMap = adj
	for key,_ := range adj {
		g.Keys = append(g.Keys,key)
	}
	return g
}

func NewGraphFromFile(filename string) *Graph {
	g := new(Graph)
	g.Load(filename)
	return g
}

func NewGraphFromEdges(filename string) *Graph {
	edges := &Edges{nil}
	edges.Load(filename)
	adjMap := ParseEdges(edges)
	graph := NewGraph(adjMap)
	return graph
}

func (g *Graph)CalcAverageEdges() float64 {
	erg := 0.0
	for _,list := range g.AdjMap {
		erg += float64(len(list))
	}
	return erg/float64(len(g.Keys))
}

func (g *Graph)DeleteNode(nodeid uint32){
	for _,friend := range g.AdjMap[nodeid] {
		fadj := g.AdjMap[friend]
		for idx,val := range fadj {
			if val==nodeid {
				g.AdjMap[friend] = append(fadj[:idx],fadj[idx+1:]...)
				break
			}
		}
	}
	delete(g.AdjMap,nodeid)
	for idx,val := range g.Keys {
		if val==nodeid {
			g.Keys = append(g.Keys[:idx],g.Keys[idx+1:]...)
			break
		}
	}
}

func (g *Graph)Save(filename string){
	f,err := os.Create(filename)
	if err!=nil {
		log.Print(err)
		return
	}
	defer f.Close()
	encoder := gob.NewEncoder(f)
	encoder.Encode(g)
}

func (g *Graph)Load(filename string){
	f,err := os.Open(filename)
	if err!=nil {
		log.Print(err)
		return
	}
	defer f.Close()
	decoder := gob.NewDecoder(f)
	decoder.Decode(g)
}

func (g *Graph) TestIfAdjacent(a,b uint32) bool {
	adjs := g.AdjMap[a]
	for _,val := range adjs {
		if b==val {
			return true
		}
	}
	return false
}

func (g *Graph) GetEdges() []*Edge {
	return nil
} 
