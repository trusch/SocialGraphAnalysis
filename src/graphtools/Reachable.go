package graphtools

func reach(act uint32,g *Graph,deep uint8,resultMap map[uint32]uint8, onlineMap map[uint32]bool) {
	if onlineMap[act]==false {
		return
	}
	if val,ok := resultMap[act]; deep>val||ok==false {
		resultMap[act]=deep
		if deep>0 {
			for _,next := range g.AdjMap[act] {
				reach(next,g,deep-1,resultMap,onlineMap)
			}
		}
	}
}

func ReachableIn(start uint32,g *Graph,deep uint8,onlineMap map[uint32]bool) uint32 {
	result := make(map[uint32]uint8)
	reach(start,g,deep,result,onlineMap)
	return uint32(len(result))
}

