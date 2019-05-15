package main

import "fmt"

/**
 冒泡排序: 后一个与前一个相比较，如果后一个比前一个大，则往后移动,一次重复该过程
 */


func main(){
	a := [...]int{5,2,3,1,7,9,0}
	fmt.Println(a)

	num := len(a)

	for i := 0; i < num ; i++ {
		for j := i+1; j < num ; j++  {
			if a[i] < a[j]{
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}

	fmt.Println(a)
}
