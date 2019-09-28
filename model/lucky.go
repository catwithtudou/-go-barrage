package model

import (
	"github.com/jinzhu/gorm"
	"log"
)

/**
 * user: ZY
 * Date: 2019/9/27 8:28
 */


type Lucky struct{
	Redid string
	Num int
}

func (l *Lucky)Add(){
	var err error
	tmp:=Lucky{}
	err=DB.Where("redid = ?",l.Redid).First(&tmp).Error
	if tmp.Redid!=""&&tmp.Redid==l.Redid{
		err=DB.Where("redid = ?",l.Redid).Update("num",gorm.Expr("num + 1")).Error
		if err!=nil{
			log.Println("update the lucky in db failed :",err)
		}
		return
	}
	err=DB.Create(l).Error
	if err!=nil{
		log.Println("create the lucky in db failed :",err)
	}

}


