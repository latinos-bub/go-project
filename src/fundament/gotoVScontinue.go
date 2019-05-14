package main

import "fmt"

/**
如下两种 continue 和 goto 的输出结果是一样
 */
func main()  {
	LABEL:
		for i := 0; i < 10; i++{
			fmt.Println("i")
			for {
				fmt.Println(i)
				continue LABEL  // continue 并不会中断循环，只是当前continue所在行之后的代码不再执行
				// 当continue配合行 label 时，会每次都跳到该 label 所在的行，但是之前的循环记录并不会中断(即 i 还是保持上一次的值)
				fmt.Println("ok")
			}
		}




	//LABEL2:
	//	for i := 0; i < 10; i++{
	//		for{
	//			fmt.Println(i)
	//			goto LABEL2
	//		}
	//	}
}
