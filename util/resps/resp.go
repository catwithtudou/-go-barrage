package resps

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * user: ZY
 * Date: 2019/9/28 11:30
 */

type RespError struct{
	Code string `json:"code"`
	Msg string `json:"msg"`
}

var (
	paramError = RespError{
		Code:"10001",
		Msg:"param error",
	}

	success = RespError{
		Code:"10000",
		Msg:"success",
	}
	powerError = RespError{
		Code:"10002",
		Msg:"power error",
	}
	innerError = RespError{
		Code: "10003",
		Msg:  "server error",
	}
)

func HandleError(c *gin.Context,err RespError){
	c.JSON(200,err)
}

func OK(c *gin.Context){
	c.JSON(200,success)
}

func ParamError(c *gin.Context){
	c.JSON(http.StatusBadRequest,paramError)
}

func PowerError(c *gin.Context){
	c.JSON(401,powerError)
}

func InnerError(c *gin.Context){
	c.JSON(500,innerError)
}
