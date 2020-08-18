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
	var bar model.ProgressBar
	bar.NewOption(0,1000, 50)
	for i := 0; i<=1000; i++{
		time.Sleep(time.Millisecond * 10)
		bar.Play(i)
	}
	bar.Finish()
}