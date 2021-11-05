package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"redEnv_v1/app/redEnv/conftools"
	"redEnv_v1/app/redEnv/dbtools"
	"redEnv_v1/app/redEnv/statuscode"
	"redEnv_v1/filepath"
	"time"
)

var N int	//用户最多抢几次
var p float64	//获得红包的概率
var lower int	//红包最低金额
var upper int	//红包最高金额
var CurrEid int	//当前红包eid


func init() {
	N, p, lower, upper = conftools.GetEnvConfig(fmt.Sprintf("%v%v", filepath.ConfRoot, filepath.EnvConf))
	fmt.Println(p)
}

func SnatchHandler(c *gin.Context) {
	type JsIn struct {
		Uid int `json:"uid"`
	}
	var jsin JsIn
	if err := c.ShouldBindJSON(&jsin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	cnt, err := dbtools.SnatchGet(jsin.Uid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": statuscode.NoSuchUser,
			"msg": "No such user",
		})
		return
	}

	if cnt >= N {
		c.JSON(http.StatusOK, gin.H{
			"code": statuscode.TooManyEnv,
			"msg": "The number of red envelopes reached the upper limit",
		})
		return
	}

	eid, val, flag := getEnv()

	if flag == false {
		c.JSON(http.StatusOK, gin.H{
			"code": statuscode.Thankyou,
			"msg": "只能说运气并不是恨好",
		})
		return
	}

	// 向数据库写入 后期改为向MQ发消息
	go dbtools.SnatchWrite(jsin.Uid, eid, val, time.Now().Unix(), N)

	c.JSON(http.StatusOK, gin.H{
		"code": statuscode.OK,
		"msg": "success",
		"data": gin.H{
			"envelope_id": eid,   // 红包id
			"max_count": N,     // 最多抢几次
			"cur_count": cnt,      // 当前为第几次抢
		},
	})
}

func getEnv() (int, int, bool) {
	rnum := rand.Float64()

	if rnum >= p {
		return 0, 0, false
	}
	CurrEid++
	return CurrEid, rand.Intn(upper - lower) + lower, true
}