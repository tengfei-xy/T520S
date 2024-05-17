package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/tengfei-xy/go-log"
)

func push_message(text string) error {
	token, err := getTokenFile()
	if err != nil {
		log.Error(err)
		return err
	}
	data := `{"aps":{"alert":"` + text + `"}}`
	url := fmt.Sprintf("https://%s/3/device/%s", app.Push.ApnsHostName, app.Push.DeviceToken)
	client := http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return err
	}
	log.Infof(token)
	req.Header.Set("authorization", `bearer `+token)
	req.Header.Set("apns-push-type", `alert`)
	req.Header.Set("apns-topic", app.Push.Topic)

	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("内部错误:%v", err)
		return err

	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Errorf("状态码:%d", resp.StatusCode)
		return err
	}
	log.Info(text)
	return nil
}

func getTokenFile() (string, error) {
	cmd := exec.Command("bash", app.Push.CreateTokenFile)
	r, err := cmd.Output()
	return string(r), err
}
