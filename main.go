package main

import "fmt"

func main() {

	result := DirReduc([]string{"SOUT", "WEST", "WEST", "EAST", "NORTH", "SOUTH", "SOUTH", "WEST"})

	fmt.Printf("result of work is %v \n", result)
}

//DirReduc delete unnecessary directions
func DirReduc(arr []string) []string {

	for i := 1; i < len(arr); i++ {
		if arr[i-1] == "NORTH" && arr[i] == "SOUTH" || arr[i-1] == "SOUTH" && arr[i] == "NORTH" {
			arr = append(arr[:i-1], arr[i+1:]...)

			return DirReduc(arr)
		} else if arr[i-1] == "EAST" && arr[i] == "WEST" || arr[i-1] == "WEST" && arr[i] == "EAST" {
			arr = append(arr[:i-1], arr[i+1:]...)

			return DirReduc(arr)
		}
	}

	if arr == nil {
		arr = make([]string, 0, 0)
	}

	return arr
}
