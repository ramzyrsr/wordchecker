package main

import "fmt"

func countEdit(str1, str2 string) bool {
	var sum = 0
	var indexA, indexB = 0, 0

	for i := 0; i < len(str1); i++ {
		if indexA != len(str1)-1 {
			if str1[indexA] != str2[indexB] {
				if str1[indexA+1] != str2[indexB] {
					sum++
				}
			}
		}
		if sum >= 2 {
			return false
		}
		indexA++
		indexB++
	}

	return true
}

func main() {
	fmt.Println(countEdit("telkom", "telecom"))
	fmt.Println(countEdit("telkom", "tlkom"))
}
