package main

import (
	"GGblog/internal/model"
	"GGblog/internal/routers"
	"GGblog/internal/setting"
)

func main() {
	setting.Init("./config/config.yaml")
	model.InitDatabase()
	routers.InitRouter()
}
