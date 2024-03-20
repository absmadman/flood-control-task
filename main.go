package main

import (
	"context"
	"fmt"
	"task/floodcontrol"
)

func main() {
	fc := floodcontrol.NewFC(context.Background())
	for i := 0; i < 100; i++ {
		for j := 0; j < 1000; j++ {
			res, err := fc.Check(context.Background(), int64(i))
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("UserID = ", i, " ", res)
			}
			//time.Sleep(time.Millisecond )
		}
		fmt.Println()
	}

}
