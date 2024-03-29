package main

// import "go-learning/neuralnetwork"
// import ds "go-learning/datastructure"
// import "go-learning/models"
import (
	"log"
	"go-learning/utilities"
	"os"
	"path/filepath"
)
// func init() {
// 	fmt.Println("init in main invoked")
// }

// type rec struct {
// 	word string
// 	count int
// }

func main() {
	// fmt.Println("main invoked")
	// fmt.Println(neuralnetwork.ReLU)
	// a := make(ds.Array, 1200)
	// m := ds.NewMatrixFromArray(a, 100, 12)
	// fmt.Println(m)
	// m.ReShape(24, 50)
	// fmt.Println(m)
	// m.ReShape(5, 240)
	// fmt.Println(m)
	// a := ds.Array{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	// m := a.ToMatrix(6, 2)
	// fmt.Println(*m)
	// fmt.Println(m.Shuffle())
	// fmt.Println("")
	// x := m.SplitN(4)
	// fmt.Println(*x[0])
	// fmt.Println(*x[1])
	// fmt.Println(*x[2])
	// fmt.Println(*x[3])
	// fmt.Println(*x[4])


	// fmt.Println((m.T()))
	// d := ds.NewDataCube(a, 2, 2, 3)

	// fmt.Println(d)
	// a := ds.Array{10, 90, 70, 80, 10, 20}
	// b := ds.Array{3, 2, 1, 5, 5, 4, 7, 3, 6, 9, 6, 8}
	// m1 := a.ToMatrix(2, 3)
	// m2 := b.ToMatrix(3, 4)
	// m3 := m1.Multiply(m2)
	// fmt.Println(*m3)

	// m := models.ScanFile("./datasets/pg10.txt")
	// var counter, total int
	// var max, sec rec
	// for k, v := range m {
	// 	// fmt.Printf("code: %-30s counts: %-10d\n", k, v)
	// 	if v > max.count {
	// 		sec.word = max.word
	// 		sec.count = max.count
	// 		max.count = v
	// 		max.word = k
	// 	}
	// 	total += v
	// 	counter++
	// }
	// fmt.Printf("Total words: %d, \n", counter)
	// fmt.Printf("The number of unique words: %d,\n", total)
	// fmt.Printf("The most frequent: %s, with counts: %d,\n", max.word, max.count)
	// fmt.Printf("The second most frequent: %s, with counts: %d,\n", sec.word, sec.count)
	
	p := os.Args[1]
	p, err := filepath.Abs(p)
	if err != nil {
		log.Fatal("invalid path")
	}
	utilities.CSV2JSON2(p)

		// ll := ds.NewLinkedList()
		// fmt.Println(ll)
		// for i := 0; i < 10; i++ {
		// 	ll.Append(&ds.Node{i, nil})
		// }
		// ll.InsertAtIndex(&ds.Node{100, nil}, 4)
		// fmt.Println(ll)
}
