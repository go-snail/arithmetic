package dpr

import (
	"fmt"
	"time"
	"os"
)

func Recover()  {
	defer func() {
		fmt.Println("defer...")
		if r := recover();r != nil {
			fmt.Println("Recoverd in f", r)
		}
	}()
	go func() {
		fmt.Println(os.Getppid())
		time.Sleep(time.Second * 2)
	}()

	fmt.Println(os.Getegid())
	//for ;; {
		fmt.Println("sleeping...")
		time.Sleep(time.Second * 4)
	//}

}
