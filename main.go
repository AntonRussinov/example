package main

//added comment dev
//one more
import (
	"fmt"

	"github.com/example/service"
)

func main() {
	result := service.DirReduc([]string{"SOUT", "WEST", "WEST", "EAST", "NORTH", "SOUTH", "SOUTH", "WEST"})
	fmt.Printf("result of work is %v \n", result)
}
