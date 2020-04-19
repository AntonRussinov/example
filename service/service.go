package service

//DirReduc deletes unnecessary directions
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
