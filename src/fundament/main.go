package main

import ( "fmt")

/**
使用 go install 时，需要注意 go install 的目标文件不能直接在 当前
GOPATH 下面的 src 下面，必须在 src 的子文件夹 下面才可以（至少一层子文件夹）

当 报如下提示错误时：go install: no install location for .go files listed on command line (GOBIN not set)
说明你没有设置 GOBIN ，你只需要将 GOBIN 设置成你当前 GOPATH 的 bin
目录下，即可
*/

func main(){
	a := 3
	laberl1:
	for ; ;  {
		for i := 0; i < 10; i++  {
			if i > 3 {
				fmt.Println("*")
				//goto laberl1 // 如果是 goto,则当执行到该行时，就跳到标签行，继续执行
				break laberl1 // 如果是 breack,则是直接跳出与 该标签 同级的循环;如果后面有代码，则继续执行
			}
		}
	}
	fmt.Println("ok");
	for a >= 3 {
		fmt.Println("a > ", a)
		a++
		if (a > 10) { // if 这里条件表达式允许用 括号 包括，但是并不建议这样使用(应该支持java所有的格式)
			break
		}
	}
	//label2:
}
