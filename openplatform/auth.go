package openplatform

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type CreatePreAuthCodeReq struct {
	ComponentAppId string `json:"component_appid"`
}

type CreatePreAuthCodeRes struct {
	PreAuthCode string `json:"pre_auth_code"`
	ExpiresIn   int64  `json:"expires_in"`
}

type QueryAuthReq struct {
	ComponentAppId    string `json:"component_appid"`
	AuthorizationCode string `json:"authorization_code"`
}

type QueryAuthRes struct {
	AuthorizationInfo AuthorizationInfo `json:"authorization_info"`
}

type AuthorizationInfo struct {
	AuthorizerAppId        string              `json:"authorizer_appid"`
	AuthorizerAccessToken  string              `json:"authorizer_access_token"`
	ExpiresIn              int64               `json:"expires_in"`
	AuthorizerRefreshToken string              `json:"authorizer_refresh_token"`
	FuncInfo               []FuncScopeCategory `json:"func_info"`
}

type FuncScopeCategory struct {
	Id int64 `json:"id"`
}

type GetAuthorizerTokenReq struct {
	ComponentAppId         string `json:"component_appid"`
	AuthorizerAppId        string `json:"authorizer_appid"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
}

type GetAuthorizerTokenRes struct {
	AuthorizerAccessToken  string `json:"authorizer_access_token"`
	ExpiresIn              int64  `json:"expires_in"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
}

type GetAuthorizerInfoReq struct {
	ComponentAppId       string `json:"component_appid"`
	AuthorizerAppId      string `json:"authorizer_appid"`
	ComponentAccessToken string `json:"component_access_token"`
}

type GetAuthorizerInfoRes struct {
	AuthorizerInfo    AuthorizerInfo    `json:"authorizer_info"`
	AuthorizationInfo AuthorizationInfo `json:"authorization_info"`
}

type AuthorizerInfo struct {
	NickName        string          `json:"nick_name"`         //string	昵称
	HeadImg         string          `json:"head_img"`          //string	头像
	ServiceTypeInfo ServiceTypeInfo `json:"service_type_info"` //object	公众号类型
	VerifyTypeInfo  VerifyTypeInfo  `json:"verify_type_info"`  //object	公众号认证类型
	UserName        string          `json:"user_name"`         //string	原始 ID
	PrincipalName   string          `json:"principal_name"`    //string	主体名称
	Alias           string          `json:"alias"`             //string	公众号所设置的微信号，可能为空
	BusinessInfo    BusinessInfo    //object	用以了解功能的开通状况（0代表未开通，1代表已开通），详见business_info 说明
	QrcodeURL       string          `json:"qrcode_url"` //string	二维码图片的 URL，开发者最好自行也进行保存
}

type ServiceTypeInfo struct {
	Id int64 `json:"id"`
}

type VerifyTypeInfo struct {
	Id int64 `json:"id"`
}

type BusinessInfo struct {
	OpenStore int64 `json:"open_store"`
	OpenScan  int64 `json:"open_scan"`
	OpenPay   int64 `json:"open_pay"`
	OpenCard  int64 `json:"open_card"`
	OpenShake int64 `json:"open_shake"`
}

func GetAuthorizerToken(componentAccessToken string, req *GetAuthorizerTokenReq) (*GetAuthorizerTokenRes, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/component/api_authorizer_token?component_access_token=%s",
		componentAccessToken),
		"application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	rspBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Printf("GetAuthorizerToken Rsp: %v \n", string(rspBody))
	err = CheckApiError(rspBody)
	if err != nil {
		return nil, err
	}
	var res = GetAuthorizerTokenRes{}
	err = json.Unmarshal(rspBody, &res)
	return &res, err
}

func GetAuthorizerInfo(componentAccessToken string, req *GetAuthorizerInfoReq) (*GetAuthorizerInfoRes, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_info?component_access_token=%s",
		componentAccessToken),
		"application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	rspBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Printf("GetAuthorizerInfo Rsp: %v \n", string(rspBody))
	err = CheckApiError(rspBody)
	if err != nil {
		return nil, err
	}
	var res = GetAuthorizerInfoRes{}
	err = json.Unmarshal(rspBody, &res)

	return &res, err
}

func CreatePreauthcode(componentAccessToken string, req *CreatePreAuthCodeReq) (*CreatePreAuthCodeRes, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/component/"+
		"api_create_preauthcode?component_access_token=%s", componentAccessToken),
		"application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	rspBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Printf("CreatePreAuthCode Rsp: %v \n", string(rspBody))
	err = CheckApiError(rspBody)
	if err != nil {
		return nil, err
	}
	var preAuthRes = CreatePreAuthCodeRes{}
	err = json.Unmarshal(rspBody, &preAuthRes)
	return &preAuthRes, err
}

func QueryAuth(componentAccessToken string, req *QueryAuthReq) (*QueryAuthRes, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/component/api_query_auth?component_access_token=%s",
		componentAccessToken),
		"application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	rspBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Printf("QueryAuth Rsp: %v \n", string(rspBody))
	err = CheckApiError(rspBody)
	if err != nil {
		return nil, err
	}
	var res = QueryAuthRes{}
	err = json.Unmarshal(rspBody, &res)
	return &res, err
}
