package main

import (
	"fmt"
	"os"
	"time"

	"github.com/tengfei-xy/go-log"
)

var app appConfig

func main() {

	app.version = "0.1"

	init_config(init_flag())

	go func() {
		for {
			check_public_ip()
			time.Sleep((time.Duration(1) * time.Hour))
		}
	}()
	for {

		n := time.Now()
		if !(app.ExecTime.Hour == n.Hour() && app.ExecTime.Minute == n.Minute()) {
			time.Sleep(time.Duration(1) * time.Minute)
			continue
		}

		t, err := get_weather_daily()
		if err != nil {
			os.Exit(1)
		}
		text := fmt.Sprintf("天气输出: 最高%d摄氏度,最低%d摄氏度,白天%s，晚上%s。", t.max_temp, t.min_temp, t.skycon_daytime, t.skycon_night)
		log.Infof("%s", text)

		var ai_text string
		for {
			ai_text, err = get_ai_text(app.Prompt, text)
			if err == nil {
				break
			}
		}
		ai_text = set_ai_text(ai_text)
		log.Infof("AI输出: %s", ai_text)
		mp3, err := get_volce(ai_text)
		if err != nil {
			//
		}
		log.Infof("文件输出: %s", mp3)
		player(mp3)
	}

}
