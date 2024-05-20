package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lolo35/l2l_login/env"
)

type l2lMachinesResp struct {
	Success bool     `json:"success"`
	Offset  int      `json:"offset"`
	Limit   int      `json:"limit"`
	Data    machines `json:"data"`
}

type machines []struct {
	Id          int    `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Line        int    `json:"line"`
	Linecode    string `json:"linecode"`
}

func FetchMachinesForLines(context *gin.Context) {
	//fmt.Println(context.PostForm("lines"))
	type request struct {
		Lines string `form:"lines" json:"lines" binding:"required"`
	}

	var validate request
	if err := context.ShouldBind(&validate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	line_ids := context.PostForm("lines")
	lines := strings.SplitAfter(line_ids, ",")
	var data []machines

	for _, value := range lines {
		lineID := strings.Replace(value, ",", "", -1)
		url := fmt.Sprintf("%smachines/?auth=%s&site=%s&fields=id,code,description,line,linecode&line=%s&active=true", env.Env("L2L_URL"), env.Env("L2L_AUTH"), env.Env("L2L_SITE"), lineID)
		resp, err := http.Get(url)

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
		var response l2lMachinesResp
		json.Unmarshal([]byte(body), &response)
		if response.Success {
			data = append(data, response.Data)
		}
	}

	//fmt.Printf("len=%d cap=%d %v", len(data), cap(data), data)

	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}
