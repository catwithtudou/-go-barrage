package service

import "go-barrage/model"

/**
 * user: ZY
 * Date: 2019/9/27 17:07
 */

func blackPerson(redid string)bool{
	user:=model.FindByRedid(redid)
	if user==(model.User{})||user.Black!=1{
		return false
	}
	return true
}

func AddBlack(redid string){
	model.BlackByRedid(redid,1)
}

func RemoveBlack(redid string){
	model.BlackByRedid(redid,0)
}