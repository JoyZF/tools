package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type LunarCalendarData struct {
	Code           int    `json:"code"`
	Msg            string `json:"msg"`
	Time           string `json:"time"`
	Zodiac         string `json:"zodiac"`
	GanZhiYear     string `json:"GanZhiYear"`
	GanZhiMonth    string `json:"GanZhiMonth"`
	GanZhiDay      string `json:"GanZhiDay"`
	GanZhiHour     string `json:"GanZhiHour"`
	NewYear        string `json:"NewYear"`
	NewMonth       string `json:"NewMonth"`
	NewDay         string `json:"NewDay"`
	Week           string `json:"Week"`
	Worktime       int    `json:"worktime"`
	Term           string `json:"term"`
	LunarYear      int    `json:"lunarYear"`
	LunarMonth     int    `json:"lunarMonth"`
	LunarDay       int    `json:"lunarDay"`
	LunarHour      string `json:"lunarHour"`
	LunarMonthName string `json:"lunarMonthName"`
	LunarDayName   string `json:"lunarDayName"`
	LunarLeapMonth int    `json:"lunarLeapMonth"`
	SolarFestival  string `json:"solarFestival"`
	IsBigMonth     bool   `json:"isBigMonth"`
	LiuRenMonth    string `json:"LiuRenMonth"`
	LiuRenDay      string `json:"LiuRenDay"`
	LiuRenHour     string `json:"LiuRenHour"`
}

func getLunarCalendar() (*LunarCalendarData, error) {
	resp, err := http.Get("https://api.t1qq.com/api/tool/day/time?key=7l0bDGCkKqJ64JlH9D87ISsxgs")
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	content, _ := io.ReadAll(resp.Body)
	var data LunarCalendarData
	_ = json.Unmarshal(content, &data)
	if data.Code != 200 {
		return nil, errors.New(data.Msg)
	}
	return &data, nil
}

func LunarCalendarTable() {
	lunarCalendar, err := getLunarCalendar()
	if err != nil {
		return
	}

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	table := widgets.NewTable()
	table.Title = time.Now().Format(time.DateOnly) + "老黄历"
	table.BorderStyle = ui.NewStyle(ui.ColorRed)
	table.Rows = [][]string{
		{"列名", "值"},
	}
	m := struct2Map(lunarCalendar)
	for k, v := range m {

		table.Rows = append(table.Rows, []string{k, fmt.Sprintf("%+v", v)})
	}
	table.TextStyle = ui.NewStyle(ui.ColorYellow)
	table.TitleStyle = ui.NewStyle(ui.ColorYellow)
	table.SetRect(0, 0, 50, 500)
	ui.Render(table)
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "c":

		}
	}
}

func struct2Map(data *LunarCalendarData) map[string]interface{} {
	m := make(map[string]interface{})
	bytes, _ := json.Marshal(data)
	_ = json.Unmarshal(bytes, &m)
	return m
}
