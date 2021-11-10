package tools

import "os"

func SetConfig() {
	os.Setenv("token_duration" ,"8")
	os.Setenv("token_secret", "aduashduhwuahduhd")
	os.Setenv("port","8080")
}