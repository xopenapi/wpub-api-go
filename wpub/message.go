package wpub

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type EncryptResponse struct {
	XMLName      xml.Name `xml:"xml"`
	Encrypt      Encrypt  `xml:"Encrypt"`
	MsgSignature string   `xml:"MsgSignature"`
	TimeStamp    string   `xml:"TimeStamp"`
	Nonce        string   `xml:"Nonce"`
}

type Encrypt struct {
	Text string `xml:",cdata"`
}

type ResponseMessage struct {
	XMLName      xml.Name     `xml:"xml"`
	ToUserName   string       `xml:"ToUserName"`   //开发者微信号
	FromUserName string       `xml:"FromUserName"` //发送方微信号，若为普通用户，则是一个OpenID
	CreateTime   int64        `xml:"CreateTime"`   //消息创建时间 （整型）
	MsgType      MsgType      `xml:"MsgType"`      //消息类型,文本:text 图片:image 语音: voice 视频:video 小视频:shortvideo 地理位置:location 链接:link
	Content      string       `xml:"Content"`      //文本消息内容
	Image        ImageMessage `xml:"Image"`
	Voice        VoiceMessage `xml:"Voice"`
	Video        VideoMessage `xml:"Video"`
	Music        MusicMessage `xml:"Music"`
}

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
	MediaId string `json:"media_id" xml:"MediaId"`
}

type VoiceMessage struct {
	MediaId string `json:"media_id" xml:"MediaId"`
}

type VideoMessage struct {
	MediaId      string `json:"media_id" xml:"MediaId"`
	ThumbMediaId string `json:"thumb_media_id" xml:"ThumbMediaId"`
	Title        string `json:"title" xml:"Title"`
	Description  string `json:"description" xml:"Description"`
}

type MusicMessage struct {
	Title        string `json:"title" xml:"Title"`
	Description  string `json:"description" xml:"Description"`
	MusicURL     string `json:"musicurl" xml:"MusicUrl"`
	HQMusicURL   string `json:"hqmusicurl" xml:"HQMusicUrl"`
	ThumbMediaId string `json:"thumb_media_id" xml:"ThumbMediaId"`
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
