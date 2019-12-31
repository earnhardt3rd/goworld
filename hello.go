package main

import (
   "fmt"
   "github.com/earnhardt3rd/goworld/myString"
)
// To Run: go run hello.go
func main() {
   fmt.Println("============ PRINT VARIABLES ==================")
   fmt.Print("hello, world\n")
   // Using new Reverse from string package
   fmt.Println("--NOW REVERSE...")
   fmt.Println(myString.Reverse("hello, world!"))

   fmt.Println("============ DECLARE VARIABLES ==================")
   // Declare variable as: var <var> type = (val)
   var x int
   x = 5
   fmt.Println("declare var x int and set x =",x)
   var y int = 7
   fmt.Println("declare var y int =",y)
   var sum int = x + y
   fmt.Println("declare sum int = x + y result:",x,"+",y,"=",sum)
   if x < 7 {
      fmt.Println("if x < 7")
   } else if y < 10 {
      fmt.Println("else if y < 10")
   } else {
      fmt.Println("else")
   }

   fmt.Println("============ ARRAY VARIABLES ==================")
   // ARRAY Example:
   var a [5]int // Fixed ARRAY
   fmt.Println("ARRAYS ARE FIXED SIZE(var a [5]int) and default to zero:",a)
   //Update ARRAY[i] Value
   a[2] = 7
   fmt.Println("ARRAYS ARE ZERO BASED INDEX: Update a[2]=7  :=",a)
   // Declare ARRAY with initial values
   b := [5]int{9,8,7,6,5}
   fmt.Println("Declare ARRAY with initial values: (b := [5]int{9,8,7,6,5}) b = ",b)

   fmt.Println("-- NON-FIXED LENGTH ARRAY ARE CALLED SLICES - USED WITH append")
   c := []int{1,3,5,7,9}
   fmt.Println("-- c := []int{1,3,5,7,9} c=",c)
   c = append(c,13)
   fmt.Println("-- c = append(c,13) c = ",c)

   fmt.Println("============ HASH VARIABLES CALLED MAP ==================")
   // HASH Example: make(map[keyType]valType)
   vertices := make(map[string]int)  // Use make() to initialize an empty hash map

   vertices["triangle"] = 2
   vertices["square"] = 3
   vertices["dodecagon"] = 12
   for key, value := range vertices {
      fmt.Printf("  Key:%-15v Value:%-v\n",key,value) // Notice the Left Alignment
   }
   // print single hash by reference
   fmt.Printf("  --Single ByVal... triangle:=%3v\n",vertices["triangle"])  // Notice the Right Alignment

   type Node struct {
      Next  *Node
      Value interface{}
   }
   var first *Node

   visited := make(map[*Node]bool)
   for n := first; n != nil; n = n.Next {
      if visited[n] {
          fmt.Println("cycle detected")
          break
      }
      visited[n] = true
      fmt.Println(n.Value)
   }
}