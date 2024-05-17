package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	log "github.com/tengfei-xy/go-log"

	"net/http"
)

func get_ai_text(prompt string, text string) (string, error) {
	d, err := get_ai_req(fmt.Sprintf("%s%s", prompt, text))
	if err != nil {
		return "", err
	}
	var air AIReq
	err = json.Unmarshal(d, &air)
	switch air.Candidates[0].FinishReason {
	case "SAFETY":
		log.Errorf("FinishReason:%s", air.Candidates[0].FinishReason)
		return "", fmt.Errorf("failed")
	case "STOP":
		break
	default:
		log.Errorf("FinishReason:%s", air.Candidates[0].FinishReason)
		log.Errorf("%s", string(d))
	}

	if err == nil {
		text := air.Candidates[0].Content.Parts[0].Text
		return text, nil
	}

	return "", nil
}
func set_ai_text(text string) string {
	text = strings.ReplaceAll(text, "\n", "")
	return strings.ReplaceAll(text, "**", "")
}

func get_ai_req(text string) ([]byte, error) {
	data := `{ "contents":[{"parts":[{"text": "` + text + `"}]}]}`
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1/models/gemini-pro:generateContent?key=%s", app.ApiKey)
	client := get_client()
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", `application/json`)

	resp, err := client.Do(req)
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

type AIReq struct {
	Candidates    []Candidates  `json:"candidates"`
	UsageMetadata UsageMetadata `json:"usageMetadata"`
}
type Parts struct {
	Text string `json:"text"`
}
type Content struct {
	Parts []Parts `json:"parts"`
	Role  string  `json:"role"`
}
type SafetyRatings struct {
	Category    string `json:"category"`
	Probability string `json:"probability"`
}
type Candidates struct {
	Content       Content         `json:"content"`
	FinishReason  string          `json:"finishReason"`
	Index         int             `json:"index"`
	SafetyRatings []SafetyRatings `json:"safetyRatings"`
}
type UsageMetadata struct {
	PromptTokenCount     int `json:"promptTokenCount"`
	CandidatesTokenCount int `json:"candidatesTokenCount"`
	TotalTokenCount      int `json:"totalTokenCount"`
}
