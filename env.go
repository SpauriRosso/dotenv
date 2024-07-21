package dotenv

import (
	"log"
	"os"
	"strings"
)

var Env string
var File []byte

//func main() {
//	Define(".env")
//	fmt.Println(GetEnv("SUPER_SECRET_KEY"))
//}

func Define(path string) string {
	_, err := os.Getwd()
	CheckErr(err)
	Env = path
	return Env
}

func GetEnv(key string) string {
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

func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
