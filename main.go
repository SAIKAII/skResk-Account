package main

import (
	"github.com/SAIKAII/skResk-Infra"
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
)

func main() {
	file := kvs.GetCurrentFilePath("config/config.ini", 1)
	conf := ini.NewIniFileCompositeConfigSource(file)
	app := infra.New(conf)
	app.Start()
}
