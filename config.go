package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type today struct {
	max_temp       int32
	min_temp       int32
	deff_temp      int32
	skycon_daytime string
	skycon_night   string
}
type appConfig struct {
	ID          string `yaml:"id"`
	Weather     `yaml:"weather"`
	Gemini      `yaml:"gemini"`
	Proxy       `yaml:"proxy"`
	Volce       `yaml:"volce"`
	Push        `yaml:"push"`
	ExecTime    `yaml:"exec_time"`
	AllowIp     []string `yaml:"allow_ip"`
	version     string
	config_file string
}
type Weather struct {
	Cor   string `yaml:"cor"`
	Token string `yaml:"token"`
}
type Gemini struct {
	ApiKey string `yaml:"api_key"`
	Prompt string `yaml:"prompt"`
}
type Volce struct {
	Token      string `yaml:"token"`
	Appid      string `yaml:"appid"`
	Voice_type string `yaml:"voice_type"`
}
type Push struct {
	DeviceToken     string `yaml:"device_token"`
	Topic           string `yaml:"topic"`
	ApnsHostName    string `yaml:"apns_host_name"`
	CreateTokenFile string `yaml:"create_token_file"`
}
type ExecTime struct {
	Hour   int `yaml:"hour"`
	Minute int `yaml:"minute"`
}
type Proxy struct {
	Socks5 string `yaml:"socks5"`
}

type flagStruct struct {
	config_file string
	version     bool
}

func init_flag() flagStruct {
	var f flagStruct
	flag.StringVar(&f.config_file, "c", "config.yaml", "打开配置文件")
	flag.BoolVar(&f.version, "v", false, "查看版本号")
	flag.Parse()
	return f
}
func init_config(flag flagStruct) {
	if flag.version {
		fmt.Println("v" + app.version)
		os.Exit(0)
	}

	yamlFile, err := os.ReadFile(flag.config_file)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &app)
	if err != nil {
		panic(err)
	}
	app.config_file = flag.config_file
	return

}
