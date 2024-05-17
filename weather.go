package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	log "github.com/tengfei-xy/go-log"
)

var CAIYUN_CODE = map[string]string{
	"CLEAR_DAY":           "晴",
	"CLEAR_NIGHT":         "晴",
	"PARTLY_CLOUDY_DAY":   "多云",
	"PARTLY_CLOUDY_NIGHT": "多云",
	"CLOUDY":              "阴",
	"LIGHT_HAZE":          "轻度雾霾",
	"MODERATE_HAZE":       "中度雾霾",
	"HEAVY_HAZE":          "重度雾霾",
	"LIGHT_RAIN":          "小雨",
	"MODERATE_RAIN":       "中雨",
	"HEAVY_RAIN":          "大雨",
	"STORM_RAIN":          "暴雨",
	"FOG":                 "雾",
	"LIGHT_SNOW":          "小雪",
	"MODERATE_SNOW":       "中雪",
	"HEAVY_SNOW":          "大雪",
	"STORM_SNOW":          "暴雪",
	"DUST":                "浮尘",
	"SAND":                "沙尘",
	"WIND":                "大风",
}

func get_weather_daily() (today, error) {
	var dr DailyReq
	f, err := get_weather_daily_req()
	if err != nil {
		log.Error(f)
		return today{}, err

	}
	if err := json.Unmarshal(f, &dr); err != nil {
		log.Error(f)
		log.Error(dr)
		return today{}, err
	}
	if dr.Status != "ok" {
		log.Error(f)
		log.Error(dr)
		return today{}, err
	}
	var today today
	today.max_temp = int32(dr.Result.Daily.Temperature[0].Max)
	today.min_temp = int32(dr.Result.Daily.Temperature[0].Min)
	today.skycon_daytime = CAIYUN_CODE[dr.Result.Daily.Skycon08H20H[0].Value]
	today.skycon_night = CAIYUN_CODE[dr.Result.Daily.Skycon20H32H[0].Value]

	return today, nil
}
func get_weather_daily_req() ([]byte, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.caiyunapp.com/v2.6/%s/%s/daily?dailysteps=1", app.Weather.Token, app.Cor), nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	// req.Header.Set("Content-Type", `application/json`)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	resp_data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("状态码:%d 文本:%s", resp.StatusCode, string(resp_data))
	}

	return resp_data, nil
}

type DailyReq struct {
	Status     string    `json:"status"`
	APIVersion string    `json:"api_version"`
	APIStatus  string    `json:"api_status"`
	Lang       string    `json:"lang"`
	Unit       string    `json:"unit"`
	Tzshift    int       `json:"tzshift"`
	Timezone   string    `json:"timezone"`
	ServerTime int       `json:"server_time"`
	Location   []float64 `json:"location"`
	Result     Result    `json:"result"`
}
type Sunrise struct {
	Time string `json:"time"`
}
type Sunset struct {
	Time string `json:"time"`
}
type Astro struct {
	Date    string  `json:"date"`
	Sunrise Sunrise `json:"sunrise"`
	Sunset  Sunset  `json:"sunset"`
}
type Precipitation08H20H struct {
	Date        string  `json:"date"`
	Max         float64 `json:"max"`
	Min         float64 `json:"min"`
	Avg         float64 `json:"avg"`
	Probability int     `json:"probability"`
}
type Precipitation20H32H struct {
	Date        string  `json:"date"`
	Max         float64 `json:"max"`
	Min         float64 `json:"min"`
	Avg         float64 `json:"avg"`
	Probability int     `json:"probability"`
}
type Precipitation struct {
	Date        string  `json:"date"`
	Max         float64 `json:"max"`
	Min         float64 `json:"min"`
	Avg         float64 `json:"avg"`
	Probability int     `json:"probability"`
}
type Temperature struct {
	Date string  `json:"date"`
	Max  float64 `json:"max"`
	Min  float64 `json:"min"`
	Avg  float64 `json:"avg"`
}
type Temperature08H20H struct {
	Date string  `json:"date"`
	Max  float64 `json:"max"`
	Min  float64 `json:"min"`
	Avg  float64 `json:"avg"`
}
type Temperature20H32H struct {
	Date string  `json:"date"`
	Max  float64 `json:"max"`
	Min  float64 `json:"min"`
	Avg  float64 `json:"avg"`
}
type Wind struct {
	Date string `json:"date"`
	Max  Max    `json:"max"`
	Min  Min    `json:"min"`
	Avg  Avg    `json:"avg"`
}
type Wind08H20H struct {
	Date string `json:"date"`
	Max  Max    `json:"max"`
	Min  Min    `json:"min"`
	Avg  Avg    `json:"avg"`
}
type Wind20H32H struct {
	Date string `json:"date"`
	Max  Max    `json:"max"`
	Min  Min    `json:"min"`
	Avg  Avg    `json:"avg"`
}
type Humidity struct {
	Date string  `json:"date"`
	Max  float64 `json:"max"`
	Min  float64 `json:"min"`
	Avg  float64 `json:"avg"`
}
type Cloudrate struct {
	Date string  `json:"date"`
	Max  float64 `json:"max"`
	Min  float64 `json:"min"`
	Avg  float64 `json:"avg"`
}
type Pressure struct {
	Date string  `json:"date"`
	Max  float64 `json:"max"`
	Min  float64 `json:"min"`
	Avg  float64 `json:"avg"`
}
type Visibility struct {
	Date string  `json:"date"`
	Max  float64 `json:"max"`
	Min  float64 `json:"min"`
	Avg  float64 `json:"avg"`
}
type Dswrf struct {
	Date string  `json:"date"`
	Max  float64 `json:"max"`
	Min  float64 `json:"min"`
	Avg  float64 `json:"avg"`
}
type Max struct {
	Chn int `json:"chn"`
	Usa int `json:"usa"`
}
type Avg struct {
	Chn int `json:"chn"`
	Usa int `json:"usa"`
}
type Min struct {
	Chn int `json:"chn"`
	Usa int `json:"usa"`
}
type Aqi struct {
	Date string `json:"date"`
	Max  Max    `json:"max"`
	Avg  Avg    `json:"avg"`
	Min  Min    `json:"min"`
}
type Pm25 struct {
	Date string `json:"date"`
	Max  int    `json:"max"`
	Avg  int    `json:"avg"`
	Min  int    `json:"min"`
}
type AirQuality struct {
	Aqi  []Aqi  `json:"aqi"`
	Pm25 []Pm25 `json:"pm25"`
}
type Skycon struct {
	Date  string `json:"date"`
	Value string `json:"value"`
}
type Skycon08H20H struct {
	Date  string `json:"date"`
	Value string `json:"value"`
}
type Skycon20H32H struct {
	Date  string `json:"date"`
	Value string `json:"value"`
}
type Ultraviolet struct {
	Date  string `json:"date"`
	Index string `json:"index"`
	Desc  string `json:"desc"`
}
type CarWashing struct {
	Date  string `json:"date"`
	Index string `json:"index"`
	Desc  string `json:"desc"`
}
type Dressing struct {
	Date  string `json:"date"`
	Index string `json:"index"`
	Desc  string `json:"desc"`
}
type Comfort struct {
	Date  string `json:"date"`
	Index string `json:"index"`
	Desc  string `json:"desc"`
}
type ColdRisk struct {
	Date  string `json:"date"`
	Index string `json:"index"`
	Desc  string `json:"desc"`
}
type LifeIndex struct {
	Ultraviolet []Ultraviolet `json:"ultraviolet"`
	CarWashing  []CarWashing  `json:"carWashing"`
	Dressing    []Dressing    `json:"dressing"`
	Comfort     []Comfort     `json:"comfort"`
	ColdRisk    []ColdRisk    `json:"coldRisk"`
}
type Daily struct {
	Status              string                `json:"status"`
	Astro               []Astro               `json:"astro"`
	Precipitation08H20H []Precipitation08H20H `json:"precipitation_08h_20h"`
	Precipitation20H32H []Precipitation20H32H `json:"precipitation_20h_32h"`
	Precipitation       []Precipitation       `json:"precipitation"`
	Temperature         []Temperature         `json:"temperature"`
	Temperature08H20H   []Temperature08H20H   `json:"temperature_08h_20h"`
	Temperature20H32H   []Temperature20H32H   `json:"temperature_20h_32h"`
	Wind                []Wind                `json:"wind"`
	Wind08H20H          []Wind08H20H          `json:"wind_08h_20h"`
	Wind20H32H          []Wind20H32H          `json:"wind_20h_32h"`
	Humidity            []Humidity            `json:"humidity"`
	Cloudrate           []Cloudrate           `json:"cloudrate"`
	Pressure            []Pressure            `json:"pressure"`
	Visibility          []Visibility          `json:"visibility"`
	Dswrf               []Dswrf               `json:"dswrf"`
	AirQuality          AirQuality            `json:"air_quality"`
	Skycon              []Skycon              `json:"skycon"`
	Skycon08H20H        []Skycon08H20H        `json:"skycon_08h_20h"`
	Skycon20H32H        []Skycon20H32H        `json:"skycon_20h_32h"`
	LifeIndex           LifeIndex             `json:"life_index"`
}
type Result struct {
	Daily   Daily `json:"daily"`
	Primary int   `json:"primary"`
}
