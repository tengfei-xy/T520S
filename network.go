package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	log "github.com/tengfei-xy/go-log"
	tools "github.com/tengfei-xy/go-tools"
	"gopkg.in/yaml.v3"
)

func check_public_ip() error {
	ip, err := get_public_ip()
	if err != nil {
		return err
	}
	if tools.ListHasString(app.AllowIp, ip) {
		return nil
	}
	msg := fmt.Sprintf("%s 新增%s", app.ID, ip)
	tools.ListAddString(&app.AllowIp, ip)
	if err := push_message(msg); err != nil {
		log.Errorf("发送消息:%s %v", msg, err)
		panic(err)
	}
	// 将配置数据转换为 YAML 格式
	yamlData, err := yaml.Marshal(&app)
	if err != nil {
		return err
	}

	// 写入 YAML 文件
	err = os.WriteFile(app.config_file, yamlData, 0644) // 0644 是文件权限
	if err != nil {
		return err
	}
	return nil

}
func get_public_ip() (string, error) {

	const url = "https://2024.ipchaxun.com"
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 Safari/605.1.15`)
	req.Header.Set("Host", "cip.cc")

	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("内部错误:%v", err)
		return "", err

	}
	defer resp.Body.Close()

	resp_data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		log.Errorf("状态码:%d", resp.StatusCode)
		return "", err
	}
	var publicip PublicIPReq
	if err := json.Unmarshal(resp_data, &publicip); err != nil {
		return "", err
	}

	return publicip.IP, nil
}

type PublicIPReq struct {
	Ret  string   `json:"ret"`
	IP   string   `json:"ip"`
	Data []string `json:"data"`
}
