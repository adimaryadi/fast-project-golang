package tools

import (
	"os"
)

func SetConfig() {
	os.Setenv("token_duration" ,"1")
	os.Setenv("token_secret", "adi")
	os.Setenv("port","8080")
	os.Setenv("host","localhost")
	os.Setenv("userMysql","root")
	os.Setenv("passMysql","")
	os.Setenv("portMysql","3306")
	os.Setenv("dbMysql","tester")
}