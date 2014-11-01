package specifics

import "github.com/paolobueno/euler/helpers"

/*
Finds biggest predecessor on triangle-shaped tree/DA graphs
Used in p18 and p67
*/
func BiggestTrianglePredecessor(depth int, pos int, graph [][]int) int{
	prevLevel := graph[depth-1]
	if pos == 0 {
		return prevLevel[pos]
	} else if pos >= len(prevLevel) {
		return prevLevel[pos-1]
	} else {
		return helpers.Max(prevLevel[pos-1], graph[depth-1][pos])
	}
}