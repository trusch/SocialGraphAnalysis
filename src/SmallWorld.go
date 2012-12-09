package main

import (
	. "./graphtools"
	"fmt"
	"flag"
	"math/rand"
	"log"
	"time"
)

var Filename *string = flag.String("file","graph.gob","input graph")
var Deep *int = flag.Int("deep",6,"searchdeep")
var NumNodes *int = flag.Int("nodes",1,"number of nodes per fail-scene")
var NumScenes *int = flag.Int("scenes",1,"number of fail-scenarios")
var Random *bool = flag.Bool("r",false,"run randomly with seed time.Now().Unix()?")
var Verbose *bool = flag.Bool("v",false,"verbose")
var FailureProb *float64 = flag.Float64("fail",0.0,"failure probability")

func main(){
	flag.Parse()

	//Set seed
	if *Random {
		rand.Seed(time.Now().Unix())	
	}else{
		rand.Seed(0)
	}
	
	//Load graph
	g := NewGraphFromFile(*Filename)
	
	if *Verbose {
		log.Print(fmt.Sprintf("loaded graph from %v with %v nodes",*Filename,len(g.Keys)))
		log.Print(fmt.Sprintf("run test(s) with a search-deep of %v and failure prob of %6.2f",*Deep,*FailureProb))
	}
	
	//Generate fail-scenarios
	result := 0.0
	for i:=0;i<*NumScenes;i++ {
	
		//Kill some nodes
		onlineMap := make(map[uint32]bool)
		online := 0
		for _,key := range g.Keys {
			if rand.Intn(1000)<int((1000.*(*FailureProb))){
				onlineMap[key]=false
			}else{
				onlineMap[key]=true
				online++
			}
		}
		if *Verbose {
			log.Print("Killed ",len(g.Keys)-online," nodes")
		}
		
		//Run tests on scenario
		for j:=0;j<*NumNodes;j++ {
			//Find a starting point which is online
			var start uint32
			for{
				start = g.Keys[rand.Intn(len(g.Keys))]
				if onlineMap[start]==true {
					break
				}
			}
		
			//Compute how many nodes are reachable
			reach := ReachableIn(start,g,uint8(*Deep),onlineMap)

			//Accumulate		
			r := float64(reach)/float64(online)
			result += r

			if *Verbose {
				log.Print(fmt.Sprintf("test started at %v. reached %v nodes",start,reach))
				log.Print(fmt.Sprintf("reached %6.2f%% of all online nodes",100.*r))
			}
		}		
	}
	result /= float64((*NumNodes)*(*NumScenes))
	result*=100.
	fmt.Printf("%5.3f\n",result)
}
