package main

import (
	"fmt"
	"time"

	"github.com/xupin/astar/astar"
)

func main() {
	a := &astar.AStar{
		Rows:      5,
		Cols:      8,
		Heuristic: astar.Diagonal,
	}
	// 5x8地图
	// 0是可移动的网格
	// 1是障碍网格
	mapData := [][]int{
		0: {0, 0, 1, 1, 0, 0, 0, 0},
		1: {0, 0, 0, 0, 1, 0, 0, 0},
		2: {0, 0, 0, 1, 1, 0, 0, 0},
		3: {0, 0, 0, 0, 1, 0, 0, 0},
		4: {0, 0, 0, 0, 0, 0, 0, 0},
	}
	a.Init(mapData)
	fmt.Println("开始时间", time.Now().UnixNano())
	node := a.FindPath(
		&astar.Node{X: 0, Y: 0},
		&astar.Node{X: 5, Y: 0},
	)
	fmt.Println("结束时间", time.Now().UnixNano())
	a.Print(node, mapData)
}
