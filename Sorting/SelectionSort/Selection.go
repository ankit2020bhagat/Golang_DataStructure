package main

import "fmt"

func main() {
	fmt.Println("Welcom to Sorting in Golang")
	arr:=[5]int{65,78,45,23,45}
	var minIndex int
	for i:=0;i<len(arr);i++ {
		minIndex=i;
		for j:=i+1;j<len(arr);j++ {
           
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
			
		}
		temp:=arr[i]
		arr[i] = arr[minIndex]
		arr[minIndex] = temp
	}
	fmt.Println(arr)
}