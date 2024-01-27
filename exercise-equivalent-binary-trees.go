package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}
// Closing the channel allows to terminate the loop
func Walking(t *tree.Tree, ch chan int) {
 Walk(t, ch)
 defer close(ch)  // close channel after Walk() finishes
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
 ch1 := make(chan int)   
 ch2 := make(chan int)
 
 go Walking(t1, ch1)      // call the Walking func for both trees
 go Walking(t2, ch2) 
 
 for {
  v1, ok1 := <-ch1      // receive from both channels    
  v2, ok2 := <-ch2     // the ok param tells us if the channel is closed
  
  if ok1 != ok2 || v1 != v2 {
          return false
        }
        if !ok1 {
          break;
        }
 }
  return true
}

func main() {
	if Same(tree.New(1), tree.New(2)) {
  fmt.Print("Same trees!")
 } else {
  fmt.Print("Not the same...")
 }
}
