package service

import (
	"errors"
	"strings"
	"sync"
	"time"
)

/**
 * user: ZY
 * Date: 2019/9/27 9:40
 */

var (
	sensitiveAll  = []string{}
	sensitiveLock = sync.Mutex{}
	flagModel      = false
)


func sensitive(word string)(err error){
	if flagModel {
		sensitiveLock.Lock()
		defer sensitiveLock.Unlock()
	}
	if len(sensitiveAll)!=0{
		for _,v := range sensitiveAll{
			if strings.Contains(word,v){
				err = errors.New("your some words are sensitive")
				return
			}
		}
	}
	return
}

func AddSensitive(word string){
	flagModel = true
	time.Sleep(100 * time.Millisecond)

	sensitiveLock.Lock()
	defer func(){
		sensitiveLock.Unlock()
		flagModel=false
	}()
	if len(sensitiveAll)!=0{
		for _,v:=range sensitiveAll{
			if v==word{
				return
			}
		}
	}
	sensitiveAll=append(sensitiveAll,word)
}

func RemoveSensitive(word string){
	flagModel = true
	time.Sleep(100 * time.Millisecond)

	sensitiveLock.Lock()
	defer func(){
		sensitiveLock.Unlock()
		flagModel=false
	}()

	for k,v:=range sensitiveAll{
		if v==word {
			sensitiveAll=append(sensitiveAll[:k],sensitiveAll[k+1:]...)
		}
	}
}
