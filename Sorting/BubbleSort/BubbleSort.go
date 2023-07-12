package main

import "fmt"

func main() {
	fmt.Println("Buuble Sort in Golang")
	arr:=[5]int{}
	fmt.Println("Enter the number")
    for i:=0;i<5;i++ {
		fmt.Scanln(&arr[i])
	}

	for i:=0;i<len(arr)-1;i++ {
		for j:=0;j<len(arr)-1;j++ {
            if arr[j] > arr[j+1] {
				// temp:=arr[j+1]
				// arr[j+1]= arr[j]
				// arr[j] = temp
				swap(&arr[j],&arr[j+1])
			}
		}
	}
	print(arr);
}

func print(arr [5]int){
    fmt.Println(arr);
}

func swap(num1,num2 *int){
   temp:=*num1
   *num1=*num2
   *num2=temp
}