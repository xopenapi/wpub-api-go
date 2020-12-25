package openplatform

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type GetComponentAccessTokenReq struct {
	ComponentAppId        string `json:"component_appid"`
	ComponentAppSecret    string `json:"component_appsecret"`
	ComponentVerifyTicket string `json:"component_verify_ticket"`
}

type GetComponentAccessTokenRes struct {
	ComponentAccessToken string `json:"component_access_token"`
	ExpiresIn            int64  `json:"expires_in"`
}

func GetComponentAccessToken(req *GetComponentAccessTokenReq) (*GetComponentAccessTokenRes, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post("https://api.weixin.qq.com/cgi-bin/component/api_component_token",
		"application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	rspBody, err := ioutil.ReadAll(resp.Body)
	log.Printf("GetComponentAccessToken Rsp: %v \n", string(rspBody))
	err = CheckApiError(rspBody)
	if err != nil {
		return nil, err
	}
	var tokenRes = GetComponentAccessTokenRes{}
	err = json.Unmarshal(rspBody, &tokenRes)
	return &tokenRes, err
}
