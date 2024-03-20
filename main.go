package main

import (
	"context"
	"fmt"
	"task/floodcontrol"
	"time"
)

func main() {
	fc := floodcontrol.NewFC(context.Background())
	for i := 0; i < 100; i++ {
		for j := 0; j < 10; j++ {
			res, err := fc.Check(context.Background(), int64(i))
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("UserID = ", i, "  Check res: ", res)
			}
			time.Sleep(time.Millisecond * 200)
		}
		fmt.Println()
	}

}
