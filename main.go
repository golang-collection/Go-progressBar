package main

import (
	"Go-progressBar/model"
	"time"
)

/**
* @Author: super
* @Date: 2020-08-18 10:47
* @Description:
**/

func main() {
	var bar model.Bar
	bar.NewOption(0,100)
	for i := 0; i<=100; i++{
		time.Sleep(time.Millisecond * 100)
		bar.Play(int64(i))
	}
	bar.Finish()
}