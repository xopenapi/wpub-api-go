package openplatform

import (
	"encoding/xml"
)

type InfoType string

const (
	InfoTypeComponentVerifyTicket InfoType = "component_verify_ticket" //返回ticket
	InfoTypeAuthorized            InfoType = "authorized"              //授权
	InfoTypeUnauthorized          InfoType = "unauthorized"            //取消授权
	InfoTypeUpdateAuthorized      InfoType = "updateauthorized"        //更新授权
)

type EncryptMessage struct {
	XMLName xml.Name `xml:"xml"`
	AppId   string   `xml:"AppId"`
	Encrypt string   `xml:"Encrypt"`
}

type Event struct {
	XMLName                      xml.Name `xml:"xml"`
	AppId                        string   `xml:"AppId"`
	CreateTime                   int64    `xml:"CreateTime"`
	InfoType                     InfoType `xml:"InfoType"`
	ComponentVerifyTicket        string   `xml:"ComponentVerifyTicket"`
	AuthorizerAppId              string   `xml:"AuthorizerAppid"`
	AuthorizationCode            string   `xml:"AuthorizationCode"`
	AuthorizationCodeExpiredTime int64    `xml:"AuthorizationCodeExpiredTime"`
	PreAuthCode                  string   `xml:"PreAuthCode"`
}
