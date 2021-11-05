package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"redEnv_v1/app/redEnv/dbtools"
	"redEnv_v1/app/redEnv/handler"
	"redEnv_v1/filepath"
	"redEnv_v1/tools"
	"time"
)

func main() {
	port := tools.GetPort(fmt.Sprintf("%v%v", filepath.ConfRoot, filepath.PortConf))
	rand.Seed(time.Now().UnixNano())

	//通过数据库中EID最大的红包初始化当前红包的eid
	var rec dbtools.Record
	dbtools.Db.Last(&rec)
	handler.CurrEid = rec.Id

	r := gin.Default()

	//路由
	initRouter(r)

	r.Run(fmt.Sprintf(":%v", port))
}