package wpub

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GetNewAccessTokenReq struct {
	AppId     string
	AppSecret string
}

type GetNewAccessTokenRes struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

func GetAccessToken(req *GetNewAccessTokenReq) (*GetNewAccessTokenRes, error) {
	if req.AppId == "" || req.AppSecret == "" {
		return nil, errors.New("appid and secret required")
	}
	resp, err := http.Get(fmt.Sprintf(
		"https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", req.AppId, req.AppSecret))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Printf("GetNewAccessToken %v \n", string(body))
	err = CheckApiError(body)
	if err != nil {
		return nil, err
	}
	var newAccessToke GetNewAccessTokenRes
	err = json.Unmarshal(body, &newAccessToke)
	return &newAccessToke, err
}
