package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"redEnv_v1/app/redEnv/dbtools"
	"redEnv_v1/app/redEnv/statuscode"
)

func OpenHandler(c *gin.Context) {
	type Jsin struct {
		Uid        int `json:"uid"`
		EnvelopeId int `json:"envelope_id"`
	}
	var jsin Jsin
	if err := c.ShouldBindJSON(&jsin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	opened, val, err := dbtools.OpenGet(jsin.Uid, jsin.EnvelopeId)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": statuscode.NoThisEnv,
			"msg": "This user doesnt have this red envelope",
		})
		return
	}

	if opened == 1 {
		c.JSON(http.StatusOK, gin.H{
			"code": statuscode.AlreadyOpened,
			"msg": "this envelope has already been opened",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": statuscode.OK,
		"msg": "success",
		"data": gin.H{
			"value": val,
		},
	})

	go dbtools.OpenWrite(jsin.Uid, jsin.EnvelopeId, val)
}