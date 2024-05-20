package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lolo35/l2l_login/env"
)

type l2lResp struct {
	Success bool `json:"success"`
	Limit   int  `json:"limit"`
	Offset  int  `json:"offset"`
	Data    line `json:"data"`
}

type line []struct {
	Id          int    `json:"id"`
	Code        string `json:"code"`
	Area        int    `json:"area"`
	Description string `json:"description"`
}

func FetchLines(c *gin.Context) {
	l2l_endpoint := fmt.Sprintf("%slines/?auth=%s&site=%s&active=true&fields=id,code,area,description", env.Env("L2L_URL"), env.Env("L2L_AUTH"), env.Env("L2L_SITE"))

	resp, err := http.Get(l2l_endpoint)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var content l2lResp
	json.Unmarshal([]byte(body), &content)
	//fmt.Println(content)

	c.JSON(http.StatusOK, content)
}

func FetchLinesInArea(context *gin.Context) {
	type request struct {
		Area_id int `form:"area_id" json:"area_id" binding:"required"`
	}

	var validate request
	if err := context.ShouldBind(&validate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	area_id := context.Query("area_id")
	l2l_endpoint := fmt.Sprintf("%slines/?auth=%s&site=%s&active=true&fields=id,code,area,description&area=%s", env.Env("L2L_URL"), env.Env("L2L_AUTH"), env.Env("L2L_SITE"), area_id)

	resp, err := http.Get(l2l_endpoint)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var response l2lResp
	json.Unmarshal([]byte(body), &response)
	context.JSON(http.StatusOK, response)
}
