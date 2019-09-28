package service

import (
	"math/rand"
	"strings"
	"time"
)

/**
 * user: ZY
 * Date: 2019/9/27 17:42
 */


var(
	syncLucky   = false //弹幕红包标志
	sendLucky   = make(chan string) //弹幕红包人数通道
	guys =[]LuckGuy{}
)

type BarrageLucky struct{
	KeyWord string `json:"key_word"`
	Num   int `json:"num"`
	Duration time.Duration `json:"duration"`
}

type LuckGuy struct{
	Redid string `json:"redid"`
}

var barrageLucky  = BarrageLucky{}

//开启弹幕红包模式
//当选择相关关键词和持续时间后,用户在发送弹幕的时候通过是否存在关键词
//若存在时则将用户Redid存入弹幕红包通道中,在持续时间结束后关闭该通道
//最后从通道中取出相关用户Redid即用Slice存储,并根据Slice长度和红包个数进行随机数生成
///用协程开启避免堵塞主线程
func (lucky *BarrageLucky)openBarrageLucky(){
	barrageLucky=*lucky
	go func(){
		select{
		case v:=<-sendLucky:
			guys=append(guys,LuckGuy{Redid:v})
		}
	}()
	syncLucky=true

}

func (lucky *BarrageLucky)closeBarrageLucky(){
	timeTicker:=time.NewTicker(barrageLucky.Duration)
	go func(){
		select {
		case <-timeTicker.C:
			syncLucky=false
			return
		}
	}()
}

//检测是否有关键词
func existKeyWord(word string)bool{
	if strings.Contains(word,barrageLucky.KeyWord){
		return true
	}
	return false
}

//从通道取出后生成随机幸运儿,并初始化guys变量
func getLuckyGuys()(luckyGuys []LuckGuy){
	for{
		if !syncLucky{
			rand.Seed(time.Now().UnixNano())
			length:=len(guys)
			luckyGuys=make([]LuckGuy,barrageLucky.Num)
			Hash:=make(map[string]string,barrageLucky.Num)
			j:=0
			for i:=0;i<barrageLucky.Num;i++{
				n:=rand.Intn(length-1)
				redid:=guys[n].Redid
				if Hash[redid]=="1"{
					i--
					continue
				}
				Hash[redid]="1"
				luckyGuys[j]=LuckGuy{
					Redid:redid,
				}
				j++
			}
			return luckyGuys
		}
	}
}

//弹幕红包模式封装
func (lucky *BarrageLucky)BarrageLuckyWork()(result []LuckGuy){
	lucky.openBarrageLucky()
	lucky.closeBarrageLucky()
	result=getLuckyGuys()
	return
}