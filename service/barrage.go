package service

/**
 * user: ZY
 * Date: 2019/9/27 9:00
 */
import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"go-barrage/model"
	"log"
	"sync"
	"time"
)

var (
	bigScreen   = make(map[string]*websocket.Conn, 5)  //websocket连接池
	screenLock  = sync.Mutex{} //连接屏幕锁
	sendCh      = make(chan []byte, 15) //弹幕通道
	intervalMap = sync.Map{} //检查间隔锁
	lock        = sync.Mutex{}
	syncModel   = false //不通过通道写进数据标志
)



//websocket连接若失效则关闭连接
func ConnPoolWaitClose(conn *websocket.Conn) {
	screenLock.Lock()
	bigScreen[conn.RemoteAddr().String()] = conn
	screenLock.Unlock()
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println(conn.RemoteAddr(), "连接断开")
			delete(bigScreen, conn.RemoteAddr().String())
			break
		}
	}
}


type BarrageEntity struct {
	Time time.Time `json:"time"`
	Text string `json:"text"`
	Color string `json:"color"`
	Redid string `json:"redid"`
}


func SendBarrage(b BarrageEntity,u model.User)(err error){

	//检查间隔
	//每个弹幕的间隔为15秒,根据map存入时间判断
	if ts,ok:=intervalMap.Load(u.Redid);ok&&time.Now().Unix()<(ts.(int64)+15){
		return errors.New("please wait a moment")
	}

	//检查敏感词
	if err := sensitive(b.Text); err!=nil{
		return err
	}

	//检查黑名单
	if flag:= blackPerson(u.Redid);flag{
		return errors.New("you are in blacklist")
	}

	//将用户发送弹幕的时间存入
	timeNowUnix:=time.Now().Unix()
	timeNow:=time.Now()
	intervalMap.Store(u.Redid,timeNowUnix)
	res,_ :=json.Marshal(BarrageEntity{
		Time:timeNow,
		Text:b.Text,
		Color:b.Color,
		Redid:b.Redid,
	})

	//检查是否开启弹幕抽奖模式
	//为避免json反序列化会降低效率,即在发送时判断
	//若开启则判断取出的弹幕内容是否含有关键词
	if syncLucky{
		if existKeyWord(b.Text){
			sendLucky<-b.Redid
		}
	}

	//向管道中输入弹幕json序列
	sendCh<- res
	return
}

func SendSync(data []byte)(err error){
	syncModel = true
	defer func() {
		syncModel = false
	}()
	time.Sleep(1 * time.Second)

	lock.Lock()
	for _, conn := range bigScreen {
		if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
			log.Println("websocket write failed:",err)
		}
	}
	lock.Unlock()
	return
}

func work(){
	if syncModel {
		lock.Lock()
		defer lock.Unlock()
	}

	//从管道取出弹幕json序列
	b := <-sendCh

	for _, conn := range bigScreen {
		if err := conn.WriteMessage(websocket.TextMessage, b); err != nil {
			log.Println(err)
		}
	}
}

func GoSend(){
	for{
		work()
	}
}


