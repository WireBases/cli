package Pkg

import (
	"fmt"
	"os/exec"
	"strings"
	"wirebasec/Utils"
)



func isProgramRunning(programName string) (bool, error) {
	cmd := exec.Command("tasklist.exe")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	return strings.Contains(string(output), programName), nil
}



func Mysql_stop() (bool, string) {
	cmd := exec.Command("taskkill.exe", "/F", "/IM", "mysqld.exe")
	err := cmd.Run()
	if err != nil {
		return false, "Mysql Server could not be shut down or is not turned on"
	}
	return true,"Mysql server was shut down successfully"
}

func isRunning() (bool, error) {
	cmd := exec.Command("tasklist.exe")
	output, err := cmd.Output()
	fmt.Println(err)
	if err != nil {
		fmt.Println(err, output, "sa")
		return true, err
	}
	prosname := "mysqld.exe"
	
	return strings.Contains(string(output), prosname), nil
}

func Mysql_start() (bool, string) {
	mystat := Utils.GetStat("client","port")
	programName := "mysqld.exe"
	isRunning, err := isProgramRunning(programName)

	if isRunning {
		return false, "Mysql Server is already running"
	}

	if err != nil {
		return false,"Mysql Server is already running"
	}

	cmd := exec.Command("C:/Users/yusao/OneDrive/Masaüstü/fonksüyonlar/mysql/bin/mysqld", "--defaults-file=C:/Users/yusao/OneDrive/Masaüstü/fonksüyonlar/mysql/bin/my.ini", "--standalone")
	err = cmd.Start()

	if err != nil {
		return false, "Mysql Server failed to start"
	}


	return true, fmt.Sprintf("Starting MySQL Server in %s port!", mystat)

}
