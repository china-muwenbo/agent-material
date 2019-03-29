package ope_env

import (
	"testing"
	"fmt"
)

func TestEnv(t *testing.T){
	systemVersion,err:=GetVersion()
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println("current operation system :",systemVersion)
}
