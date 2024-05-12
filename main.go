package main

import (
	"GGblog/routers"
	"GGblog/utils"
)

func main() {
	utils.Init("./config/config.yaml")
	routers.InitRouter()
}
