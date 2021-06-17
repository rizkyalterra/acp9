package main

type User struct {
	Name string
}

// func main() {
// 	fmt.Println(maxDifference([]int32{1, 2, 3, 4}))
// }

func maxDifference(px []int32) int32 {
	// Write your code here
	var max int32 = -1
	var min int32 = px[0]
	var difference int32 = 0

	for step := 1; step < len(px); step++ {
		if px[step] <= min {
			min = px[step]
		} else {
			difference = px[step] - min
			if difference != 0 && difference > max {
				max = difference
			}
		}
	}
	return max

}
