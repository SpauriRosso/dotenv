package main

import (
	"os"
	"strings"
)

var Env string
var File []byte

//func main() {
//	Define(".env")
//	fmt.Println(GetEnv("SUPER_SECRET_KEY"))
//}

func Define(path string) any {
	_, err := os.Getwd()
	if err != nil {
		return err
	}
	Env = path
	return Env
}

func GetEnv(key string) any {
	var value string
	File, _ = os.ReadFile(Env)

	for _, v := range strings.Split(string(File), "\n") {
		if strings.HasPrefix(v, key+"=") {
			value = strings.TrimPrefix(v, key+"=")
			value = strings.Trim(value, "\n")
			break
		}
	}
	return value
}
