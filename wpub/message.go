package wpub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type MessageSendReq struct {
	ToUser   string          `json:"touser"`
	MsgType  string          `json:"msgtype"`
	Text     TextMessage     `json:"text"`
	Image    ImageMessage    `json:"image"`
	Voice    VoiceMessage    `json:"voice"`
	Video    VideoMessage    `json:"video"`
	Music    MusicMessage    `json:"music"`
	Articles ArticlesMessage `json:articles`
}

type TextMessage struct {
	Content string `json:"content"`
}

type ImageMessage struct {
	MediaId string `json:"media_id"`
}

type VoiceMessage struct {
	MediaId string `json:"media_id"`
}

type VideoMessage struct {
	MediaId      string `json:"media_id"`
	ThumbMediaId string `json:"thumb_media_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
}

type MusicMessage struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	MusicURL     string `json:"musicurl"`
	HQMusicURL   string `json:"hqmusicurl"`
	ThumbMediaId string `json:"thumb_media_id"`
}

type ArticlesMessage struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	PicURL      string `json:"picurl"`
}

func MessageCustomSend(message *MessageSendReq, accessToken string) error {

	body, err := json.Marshal(message)
	if err != nil {
		return err
	}
	resp, err := http.Post(fmt.Sprintf(
		"https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s",
		accessToken), "application/json", bytes.NewBuffer(body))

	if err != nil {
		return err
	}
	defer resp.Body.Close()
	rspBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Printf("SendMessage Rsp: %v \n", string(rspBody))

	err = CheckApiError(rspBody)
	return err
}
