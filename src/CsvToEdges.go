package main

import (
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
	"flag"
	"fmt"
	. "./graphtools"
)

func ReadCSV(file string) []*Edge {
	f,err := os.Open(file)
	if err!=nil {
		log.Print(err)
		return nil
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	edges := make([]*Edge,0,400000)
	for{
		l,_,err := reader.ReadLine()
		if err!=nil||len(l)==0 {
			break
		}
		if l[0]=='#' {
			continue
		}
		line := strings.Trim(string(l)," \r\n")
		words := strings.Split(line,"\t")
		if len(words)<2 {
			log.Print(words)
			break
		}
		from,err1 := strconv.ParseInt(words[0],10,32)
		to,err2   := strconv.ParseInt(words[1],10,32)
		if err1!=nil || err2!=nil {
			log.Print("Failed parsing line.")
			break
		}
		edges = append(edges,&Edge{uint32(from),uint32(to)})
	}
	return edges
}

func Delete(list []*Edge,idx int) []*Edge {
	return append(list[:idx],list[idx+1:]...)
}

func CalcBidirectionalGraph(edges []*Edge) []*Edge {
	result := make([]*Edge,0)
	for idx1:=0;idx1<len(edges);idx1++ {
		for idx2:=idx1+1;idx2<len(edges);idx2++ {
			e1,e2 := edges[idx1],edges[idx2]
			if e1.A==e2.B && e1.B==e2.A {
				result = append(result,e1)
				edges = Delete(edges,idx2)
				break
			}
		}
		if idx1%500==0 {
			percent := float64(idx1)/float64(381217)
			log.Print("finished ",percent*100,"% : found ",len(result))	
		}
	}
	return result
}


var Filename *string = flag.String("f","trust_data.txt","filename") 
var Output *string = flag.String("o","output.gob","outputfile")

func main(){
	flag.Parse()
	edges := ReadCSV(*Filename)
	fmt.Println("Read ",len(edges)," edges.")
	bidirect := CalcBidirectionalGraph(edges)
	fmt.Println("found ",len(bidirect)," bidirectional edges.")
	e1 := &Edges{bidirect}
	e1.Save(*Output)
	log.Print("wrote ",len(e1.Items)," edges (",e1.Items[0].A,")")

}
