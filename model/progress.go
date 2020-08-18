package model

import (
	"fmt"
	"math"
	"strconv"
)

/**
* @Author: super
* @Date: 2020-08-18 10:37
* @Description:
**/

type ProgressBar struct {
	percent       int    //百分比
	cur           int    //当前进度位置
	total         int    //总进度
	totalLen      int    //total的长度
	count         int    //图案显示的个数
	rate          string //进度条
	showGraphRate int    //控制多少个字节显示一个图案
	graph         string //显示的图案
}

func (bar *ProgressBar) NewOption(start, total, count int) {
	bar.count = count
	bar.cur = start
	bar.total = total
	s := strconv.Itoa(total)
	bar.totalLen = len(s)
	if bar.graph == "" {
		bar.graph = "#"
	}
	bar.showGraphRate = int(math.Ceil(float64(bar.total) / float64(bar.count)))
	bar.percent = 0
}

func (bar *ProgressBar) NewOptionWithGraph(start, total, count int, graph string) {
	bar.graph = graph
	bar.NewOption(start, total, count)
}

func (bar *ProgressBar) Play(cur int) {
	bar.cur = cur
	bar.percent = bar.getPercent()


	if cur%(bar.showGraphRate) == 0 && cur <= bar.total && cur != 0 {
		bar.rate += bar.graph
	}

	fmt.Printf("\r[%-"+strconv.Itoa(bar.count)+"s]%3d%% %" + strconv.Itoa(bar.totalLen) +"d/%d", bar.rate, bar.percent, bar.cur, bar.total)
}

func (bar *ProgressBar) Finish() {
	fmt.Println()
}

func (bar *ProgressBar) getPercent() int {
	return int(float32(bar.cur) / float32(bar.total) * 100)
}
