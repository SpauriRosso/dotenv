# dotenv
 Simple program to retrieve .env key and more 

## How to use ?
Usage is pretty simple, first you have to "Define" the path where your .env file is located and then bind a variable to the value your are trying to resolve

````go
package main

import "github.com/SpauriRosso/dotenv"

func main() {
	dotenv.Define(".env")
	key := dotenv.GetEnv("MY_SUPER_SECRET_KEY")
	println(key)
}
````