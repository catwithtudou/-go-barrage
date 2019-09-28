package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go-barrage/service"
	"log"
	"net/http"
)

/**
 * user: ZY
 * Date: 2019/9/28 14:29
 */

var(
	upgrade = websocket.Upgrader{
		ReadBufferSize:0,
		WriteBufferSize:0,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func Screen(c *gin.Context){
	conn,err:=upgrade.Upgrade(c.Writer,c.Request,nil)
	if err!=nil{
		log.Println("websocket failed",err)
		return
	}
	service.ConnPoolWaitClose(conn)
}
