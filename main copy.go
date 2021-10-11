package main

// import (
// 	"fmt"

// 	"github.com/eapache/queue"
// )

// // Vertex is 邻接表
// type Vertex struct {
// 	Key      string
// 	Parents  []*Vertex
// 	Children []*Vertex
// 	Value    interface{}
// }

// type DAG struct {
// 	Vertexes []*Vertex
// }

// // AddVertex is xx
// func (dag *DAG) AddVertex(v *Vertex) {
// 	dag.Vertexes = append(dag.Vertexes, v)
// }

// // AddEdge is xx
// func (dag *DAG) AddEdge(from, to *Vertex) {
// 	from.Children = append(from.Children, to)

// 	to.Parents = append(from.Parents, from)
// }

// // BFS is xxx
// func (dag *DAG) BFS(root *Vertex) {
// 	q := queue.New()

// 	visitMap := make(map[string]bool)
// 	visitMap[root.Key] = true

// 	q.Add(root)

// 	for {
// 		if q.Length() == 0 {
// 			fmt.Println("done")
// 			break
// 		}
// 		current := q.Remove().(*Vertex)

// 		fmt.Println("bfs key", current.Key)

// 		for _, v := range current.Children {
// 			fmt.Printf("from:%v to:%s\n", current.Key, v.Key)
// 			if v.Key == root.Key {
// 				panic("back root")
// 			}
// 			if _, ok := visitMap[v.Key]; !ok {
// 				visitMap[v.Key] = true
// 				//fmt.Println("add visit", v.Key)

// 				q.Add(v)
// 			}
// 		}
// 	}

// }

// // DFS is xxx
// func (dag *DAG) DFS(root *Vertex) {
// 	stack := []*Vertex{root}

// 	visitMap := make(map[string]bool)
// 	visitMap[root.Key] = true

// 	for i := 0; i < 10; i++ {
// 		if len(stack) == 0 {
// 			fmt.Println("done")
// 			break
// 		}
// 		if len(stack)-1 < 0 {
// 			panic("unexpected")
// 		}
// 		current := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]

// 		fmt.Println("dfs key", current.Key)

// 		for _, v := range current.Children {
// 			fmt.Println("from:%v to:%s", current.Key, v.Key)
// 			if v.Key == root.Key {
// 				panic("back root")
// 			}
// 			if _, ok := visitMap[v.Key]; !ok {
// 				visitMap[v.Key] = true
// 				//fmt.Println("add visit", v.Key)
// 				if v.Key == root.Key {
// 					panic("back root")
// 				}
// 				stack = append(stack, v)
// 			}
// 		}
// 	}
// }

// func main() {

// 	dag := &DAG{}
// 	v1 := &Vertex{Key: "1"}
// 	v2 := &Vertex{Key: "2"}
// 	v3 := &Vertex{Key: "3"}
// 	v4 := &Vertex{Key: "4"}
// 	v5 := &Vertex{Key: "5"}

// 	// dag
// 	//     5
// 	//   >
// 	//  /
// 	// 1----->2
// 	//  \   >   \
// 	//   > /     >
// 	//   3-------->4

// 	dag.AddEdge(v1, v5)
// 	dag.AddEdge(v1, v2)
// 	//dag.AddEdge(v2, v1)
// 	dag.AddEdge(v1, v3)
// 	dag.AddEdge(v3, v4)
// 	dag.AddEdge(v3, v2)
// 	dag.AddEdge(v2, v4)
// 	dag.BFS(v1)
// 	//dag.DFS(v1)
// }
