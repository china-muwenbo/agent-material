package ope_env

import (
	"syscall"
	"fmt"
	"runtime"
)
//
//func getOperationEnv() string {
//
//}

//操作系统内核版本
func GetVersion() (string, error) {
	if runtime.GOOS=="windows"{
		version, err := syscall.GetVersion()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%d.%d (%d)", byte(version), uint8(version>>8), version>>16),nil
	}
	version, err := syscall.GetVersion()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d.%d (%d)", byte(version), uint8(version>>8), version>>16),nil
}
