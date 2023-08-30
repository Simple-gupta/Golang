package main
import "fmt"

func main() {
   var i, j int
   var matrixC [3][3]int
   matrixA := [3][3]int{
      {0, 1 ,2},
      {3, 4, 5},
      {6, 7, 8},
   }
   matrixB := [3][3]int{
	{0, 1 ,2},
	{3, 4, 5},
	{6, 7, 8},
   }
   
   fmt.Println()
   fmt.Println("Result")
   for i = 0; i < 3; i++ {
      for j = 0; j < 3; j++ {
         matrixC[i][j] = matrixA[i][j] + matrixB[i][j]
      }
   }
   for i = 0; i < 3; i++ {
      for j = 0; j < 3; j++ {
         fmt.Print(matrixC[i][j], "\t")
      }
      fmt.Println()
   }
}