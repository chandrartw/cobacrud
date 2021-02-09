package main

import (
	"github.com/chandrartw/cobacrud/util"
	"github.com/chandrartw/cobacrud/server"
)
func main (){
	util.LoadEnv()
	util.SetupLoggerOutput()
	server.start()
}