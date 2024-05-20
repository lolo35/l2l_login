package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lolo35/l2l_login/env"
)

type LoginResponse struct {
	Success bool            `json:"success"`
	Lines   []LinesResponse `json:"lines"`
}

type LinesResponse struct {
	Id         int    `json:"id"`
	Code       string `json:"code"`
	Costcenter string `json:"costcenter"`
	Externalid string `json:"externalid"`
}

func Login(context *gin.Context) {
	type request struct {
		Lines    string `json:"lines" form:"lines" binding:"required"`
		Machines string `json:"machines" form:"machines" binding:"required"`
		Operator string `json:"operator" form:"operator" binding:"required"`
	}

	var validate request
	if err := context.Bind(&validate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	lineIds := context.PostForm("lines")
	machineids := context.PostForm("machines")
	operator := context.PostForm("operator")

	var operator_id string

	switch len(operator) {
	case 4:
		operator_id = fmt.Sprintf("0%s", operator)
	case 3:
		operator_id = fmt.Sprintf("00%s", operator)
	default:
		operator_id = operator
	}

	l2l_endpoint := fmt.Sprintf("%susers/clock_in_by_externalid/%s/", env.Env("L2L_URL"), operator_id)
	method := "POST"

	trimmedLines := strings.TrimRight(lineIds, ",")
	trimmedMachines := strings.TrimRight(machineids, ",")

	lines := strings.Split(trimmedLines, ",")
	machines := strings.Split(trimmedMachines, ",")

	auth := fmt.Sprintf("auth=%s&site=%s", env.Env("L2L_AUTH"), env.Env("L2L_SITE"))

	startS, endS := concatDates()

	start := fmt.Sprintf("&start=%s", startS)
	end := fmt.Sprintf("&end=%s", endS)

	fmt.Println(start, " ", end)

	var urlEncodedLines strings.Builder
	for _, value := range lines {
		urlEncodedLines.WriteString(fmt.Sprintf("&line_id=%s", value))
	}

	var urlEncodedMachines strings.Builder
	for _, value := range machines {
		urlEncodedMachines.WriteString(fmt.Sprintf("&machine_id=%s", value))
	}

	payload := strings.NewReader(fmt.Sprintf("%s%s%s%s%s", auth, urlEncodedLines.String(), urlEncodedMachines.String(), start, end))

	client := &http.Client{}

	req, err := http.NewRequest(method, l2l_endpoint, payload)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var response LoginResponse
	json.Unmarshal([]byte(body), &response)

	context.JSON(http.StatusOK, response)
}

func concatDates() (string, string) {
	const layoutISO = "15:04"

	now := time.Now()

	nowf := now.Format(layoutISO)
	shift := establishShift(nowf, layoutISO)

	startTime, endTime := establishStartAndEndTime(nowf)

	startDate, endDate := establishDate(shift)

	start := fmt.Sprintf("%s %s", startDate, startTime)
	end := fmt.Sprintf("%s %s", endDate, endTime)

	return start, end
}

func establishShift(now string, layout string) int {
	shiftStruct := []struct {
		start string
		end   string
		shift int
	}{
		{"06:00", "14:30", 1},
		{"14:30", "23:00", 2},
		{"23:00", "24:00", 3},
		{"00:00", "06:00", 4},
	}
	var shift int
	for _, t := range shiftStruct {
		start, _ := time.Parse(layout, t.start)
		end, _ := time.Parse(layout, t.end)
		now, _ := time.Parse(layout, now)

		if now.After(start) && now.Before(end) {
			shift = t.shift
			break
		}
	}
	return shift
}

func establishStartAndEndTime(now string) (string, string) {
	layout := "15:04"
	// dateLayout := "2006-01-02"
	// dateNow := time.Now()
	shiftStruct := []struct {
		start string
		check string
		end   string
		limit string
	}{
		{"06:00", "14:30", "14:30", "06:30"},
		{"14:30", "23:00", "23:00", "15:00"},
		{"23:00", "24:00", "06:00", "23:30"},
		{"00:00", "06:00", "06:00", "00:00"},
	}
	var startF string
	var endF string
	for _, t := range shiftStruct {
		start, _ := time.Parse(layout, t.start)
		check, _ := time.Parse(layout, t.check)
		end, _ := time.Parse(layout, t.end)
		limit, _ := time.Parse(layout, t.limit)
		now, _ := time.Parse(layout, now)

		if now.After(start) && now.Before(limit) {
			startF = start.Format(layout)
			endF = end.Format(layout)
			break
		}
		if now.After(start) && now.Before(check) {
			startF = now.Format(layout)
			endF = end.Format(layout)
		}
	}
	return startF, endF
}

func establishDate(shift int) (string, string) {
	layoutDate := "2006-01-02"
	layoutTime := "15:04"
	now, _ := time.Parse(layoutTime, "23:01")
	newNow := time.Now()

	var startDate string
	var endDate string

	startDate = newNow.Format(layoutDate)
	endDate = newNow.Format(layoutDate)
	if shift == 3 {
		start, _ := time.Parse(layoutTime, "23:00")
		end, _ := time.Parse(layoutTime, "24:00")
		if now.After(start) && now.Before(end) {
			d := newNow.AddDate(0, 0, 1)
			endDate = d.Format(layoutDate)
			fmt.Println("Adding 1 day")
		}
	}

	return startDate, endDate
}
