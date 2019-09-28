package model

import (
	"log"
)

/**
 * user: ZY
 * Date: 2019/9/26 9:47
 */

type User struct{
	Id uint
	Username string
	Image string
	Redid string
	Black int
	Power int
}


func (u *User)AddUser(){
	var err error
	tmp:=User{}
	err=DB.Table("user").Where("redid = ?",u.Redid).First(&tmp).Error
	if tmp.Redid!=""&&u.Redid==tmp.Redid{
		err=DB.Table("user").Where("redid = ?",u.Redid).Update(u).Error
	}else{
		err=DB.Table("user").Create(u).Error
	}
	if err!=nil{
		log.Println("save the user in db failed:",err)
		return
	}
	return
}

func FindByRedid(redid string)(tmp User){
	if err:=DB.Table("user").Where("redid = ?",redid).First(&tmp).Error;err!=nil{
		log.Println("find the user in db failed:",err)
		return
	}
	return
}

//0为默认,1为拉黑
func BlackByRedid(redid string,choice int){
	if err:=DB.Table("user").Where("redid = ?",redid).Update("black",choice).Error;err!=nil{
		log.Println("black the user in db failed:",err)
		return
	}
}

//0为默认用户,1为管理员
func PowerByRedid(redid string,choice int){
	if err:=DB.Table("user").Where("redid = ?",redid).Update("power",choice).Error;err!=nil{
		log.Println("black the user in db failed:",err)
		return
	}
}





