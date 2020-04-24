package main

import (
	"fmt"
	"os"
)

type Dijkstra struct {
	ShortTablePath int

	//已知最短路径
	disList []string

	//候选集
	visitList []visit

	join []map[string]string
}

type visit struct {
	end   string
	start string
	path  int
}

//有向图
var graph map[string]map[string]int

func (d *Dijkstra) dijkstra(start string, end string) {
	if start == end {
		fmt.Printf("最短路径:%d\n", d.ShortTablePath)
		os.Exit(1)
	}

	if len(d.join) == 0 {
		//第一次
		d.ShortTablePath = 0
	}

	//放入候选集
	if len(d.visitList) == 0 {
		for j, weight := range graph[start] {
			visit := visit{}
			visit.end = j
			visit.start = start
			visit.path = visit.path + weight
			d.visitList = append(d.visitList, visit)
		}
	} else {
		var visitList []visit
		for _, list := range d.visitList {
			if list.end == end {
				visit := visit{}
				visit.end = end
				visit.path = list.path
				visitList = append(visitList, visit)
			} else {
				for j, weight := range graph[list.end] {
					visit := visit{}
					visit.end = j
					visit.start = list.end
					visit.path = list.path + weight
					visitList = append(visitList, visit)
				}
			}
		}
		d.visitList = visitList
	}

	var min visit
	//寻找候选集权重最小的
	for _, list := range d.visitList {
		if min.path == 0 {
			min.path = list.path
			min.end = list.end
			min.start = list.start
			continue
		}
		if list.path < min.path {
			min.start = list.start
			min.path = list.path
			min.end = list.end
		}
	}

	d.ShortTablePath = min.path
	d.dijkstra(min.end, end)
}

func main() {
	graph = make(map[string]map[string]int)
	graph["a"] = make(map[string]int)
	graph["a"]["c"] = 3
	graph["a"]["b"] = 6
	graph["b"] = make(map[string]int)
	graph["b"]["d"] = 5
	graph["c"] = make(map[string]int)
	graph["c"]["d"] = 3
	graph["c"]["b"] = 2
	graph["c"]["e"] = 4
	graph["d"] = make(map[string]int)
	graph["d"]["e"] = 2
	graph["d"]["f"] = 3
	graph["e"] = make(map[string]int)
	graph["e"]["f"] = 5

	d := &Dijkstra{}
	d.dijkstra("a", "f")
}
