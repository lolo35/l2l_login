package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lolo35/l2l_login/env"
)

type l2lAreasResp struct {
	Success bool  `json:"success"`
	Offset  int   `json:"offset"`
	Limit   int   `json:"limit"`
	Data    areas `json:"data"`
}

type areas []struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

func GetAreas(context *gin.Context) {
	l2l_endpoint := fmt.Sprintf("%sareas/?auth=%s&site=%s&fields=id,code,description&active=true", env.Env("L2L_URL"), env.Env("L2L_AUTH"), env.Env("L2L_SITE"))

	resp, err := http.Get(l2l_endpoint)

	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	var l2l_resp l2lAreasResp
	json.Unmarshal([]byte(body), &l2l_resp)

	context.JSON(http.StatusOK, l2l_resp)
}
