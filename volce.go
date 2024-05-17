package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"

	log "github.com/tengfei-xy/go-log"

	"github.com/google/uuid"
)

func player(filename string) {
	exec.Command("mpv", filename).Run()

}
func get_volce(text string) (string, error) {
	d := get_volce_data(text)
	j, err := json.Marshal(d)
	if err != nil {
		log.Error(err)
		return "", err
	}
	b, err := get_volce_req(j)
	if err != nil {
		log.Error(err)
		return "", err
	}
	var res VolceRes
	err = json.Unmarshal(b, &res)
	if err != nil {
		log.Error(err)
		return "", err
	}
	if err := check_res(res); err != nil {
		log.Error(err)
		return "", err
	}

	return base64_to_file(res.Data)
}
func get_volce_req(data []byte) ([]byte, error) {
	const url = "https://openspeech.bytedance.com/api/v1/tts"
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", `Bearer;`+app.Volce.Token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf("内部错误:%v", err)
		return nil, err

	}
	resp_data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Errorf("状态码:%d", resp.StatusCode)
		log.Errorf(string(resp_data))
		return nil, err
	}
	return resp_data, nil
}
func get_volce_data(text string) VolceReq {
	return VolceReq{
		App{
			Appid:   app.Volce.Appid,
			Token:   app.Volce.Token,
			Cluster: "volcano_tts",
		},
		User{
			UID: createID(),
		},
		Audio{
			VoiceType:       app.Voice_type,
			Encoding:        "mp3",
			CompressionRate: 1,
			Rate:            24000,
			SpeedRatio:      1.0,
			VolumeRatio:     1.0,
			PitchRatio:      1.0,
			Emotion:         "happy",
			Language:        "cn",
		},
		Request{
			Reqid:           uuid.New().String(),
			Text:            text,
			TextType:        "plain",
			Operation:       "query",
			SilenceDuration: "125",
			WithFrontend:    "1",
			FrontendType:    "unitTson",
			PureEnglishOpt:  "1",
		},
	}
}
func check_res(d VolceRes) error {
	if d.Code != 3000 {
		return fmt.Errorf("code:%s ,message:%s", d.Code, d.Message)
	}
	return nil
}
func base64_to_file(base64Str string) (string, error) {
	// 2. 解码 Base64 字符串
	decodedData, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return "", err
	}
	time.Now().Date()
	// 3. 指定要保存的文件名
	filename := fmt.Sprintf("%s.mp3", GetStrDate())

	// 4. 将解码后的数据写入文件
	err = os.WriteFile(filename, decodedData, 0644) // 0644 是文件权限
	if err != nil {
		return "", err
	}
	return filename, nil

}
func createID() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	const pool = "qazwsxedcrfvtgbyhnujmikolpQAZWSXEDCRFVTGBYHNUJMIKOLP1234567890"
	bytes := make([]byte, 15)
	for i := 0; i < 15; i++ {
		bytes[i] = pool[r.Intn(len(pool))]
	}
	return string(bytes)
}
func GetStrDate() string {
	y, m, d := time.Now().Date()
	sy := strconv.Itoa(int(y))
	sm := strconv.Itoa(int(m))
	sd := strconv.Itoa(d)
	if len(sm) == 1 {
		sm = "0" + sm
	}
	if len(sd) == 1 {
		sd = "0" + sd
	}
	return sy + sm + sd
}

type VolceReq struct {
	App     App     `json:"app"`
	User    User    `json:"user"`
	Audio   Audio   `json:"audio"`
	Request Request `json:"request"`
}
type App struct {
	Appid   string `json:"appid"`
	Token   string `json:"token"`
	Cluster string `json:"cluster"`
}
type User struct {
	UID string `json:"uid"`
}
type Audio struct {
	VoiceType       string  `json:"voice_type"`
	Encoding        string  `json:"encoding"`
	CompressionRate int     `json:"compression_rate"`
	Rate            int     `json:"rate"`
	SpeedRatio      float64 `json:"speed_ratio"`
	VolumeRatio     float64 `json:"volume_ratio"`
	PitchRatio      float64 `json:"pitch_ratio"`
	Emotion         string  `json:"emotion"`
	Language        string  `json:"language"`
}
type Request struct {
	Reqid           string `json:"reqid"`
	Text            string `json:"text"`
	TextType        string `json:"text_type"`
	Operation       string `json:"operation"`
	SilenceDuration string `json:"silence_duration"`
	WithFrontend    string `json:"with_frontend"`
	FrontendType    string `json:"frontend_type"`
	PureEnglishOpt  string `json:"pure_english_opt"`
}
type VolceRes struct {
	Reqid     string   `json:"reqid"`
	Code      int      `json:"code"`
	Operation string   `json:"operation"`
	Message   string   `json:"message"`
	Sequence  int      `json:"sequence"`
	Data      string   `json:"data"`
	Addition  Addition `json:"addition"`
}
type Addition struct {
	Description string `json:"description"`
	Duration    string `json:"duration"`
	Frontend    string `json:"frontend"`
}
