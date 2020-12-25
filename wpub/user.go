package wpub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ListUserOpenIdRes struct {
	Total      int32      `json:"subscribe"`   //	关注该公众账号的总用户数
	Count      int32      `json:"count"`       //拉取的OPENID个数，最大值为10000
	Data       OpenIdData `json:"data"`        //列表数据，OPENID的列表
	NextOpenid string     `json:"next_openid"` //拉取列表的最后一个用户的OPENID
}

type OpenIdData struct {
	OpenId []string `json:"openid"`
}

type UserInfo struct {
	Subscribe int32  `json:"subscribe"` //用户是否订阅该公众号标识，值为0时，代表此用户没有关注该公众号，拉取不到其余信息。
	Openid    string `json:"openid"`    //用户的标识，对当前公众号唯一
	Nickname  string `json:"nickname"`  //用户的昵称
	Sex       int32  `json:"sex"`       //用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	Language  string `json:"language"`  //用户的语言，简体中文为zh_CN
	City      string `json:"city"`      //用户所在城市
	Province  string `json:"province"`  //用户所在省份
	Country   string `json:"country"`   //用户所在国家
	/*用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），
	用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。*/
	HeadImgURL    string  `json:"headimgurl"`
	SubscribeTime int64   `json:"subscribe_time"` //用户关注时间，为时间戳。如果用户曾多次关注，则取最后关注时间
	UnionId       string  `json:"unionid"`        //只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
	Remark        string  `json:"remark"`         //公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
	GroupId       int64   `json:"groupid"`        //用户所在的分组ID（兼容旧的用户分组接口）
	TagIdList     []int64 `json:"tagid_list"`     //用户被打上的标签ID列表
	/*返回用户关注的渠道来源，ADD_SCENE_SEARCH 公众号搜索，ADD_SCENE_ACCOUNT_MIGRATION 公众号迁移，ADD_SCENE_PROFILE_CARD 名片分享，
	ADD_SCENE_QR_CODE 扫描二维码，ADD_SCENE_PROFILE_LINK 图文页内名称点击，ADD_SCENE_PROFILE_ITEM 图文页右上角菜单，
	ADD_SCENE_PAID 支付后关注，ADD_SCENE_WECHAT_ADVERTISEMENT 微信广告，ADD_SCENE_OTHERS 其他*/
	SubscribeScene string `json:"subscribe_scene"`
	QrScene        int64  `json:"qr_scene"`     // 二维码扫码场景（开发者自定义）
	QrSceneStr     string `json:"qr_scene_str"` //二维码扫码场景描述（开发者自定义）
}

type BatchGetUserReq struct {
	UserList []*UserList `json:"user_list"`
}

type BatchGetUserRes struct {
	UserInfoList []*UserInfo `json:"user_info_list"`
}

type UserList struct {
	OpenId string `json:"openid"`
	Lang   string `json:"lang"`
}

type UserTag struct {
	// 公众平台标签id
	Id int64 `json:"id"`
	// 标签名称
	Name string `json:"name"`
}

func ListUserOpenId(accessToken string, nextOpenid string) (*ListUserOpenIdRes, error) {
	resp, err := http.Get(fmt.Sprintf(
		"https://api.weixin.qq.com/cgi-bin/user/get?access_token=%s&next_openid=%s", accessToken, nextOpenid))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = CheckApiError(body)
	if err != nil {
		return nil, err
	}
	log.Printf("ListUserOpenId Rsp : %v \n", string(body))

	openid := ListUserOpenIdRes{}
	err = json.Unmarshal(body, &openid)
	return &openid, err
}

func GetUserInfo(accessToken string, openId string) (*UserInfo, error) {
	resp, err := http.Get(fmt.Sprintf(
		"https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=zh_CN", accessToken, openId))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Printf("GetUserDetail Rsp : %v \n", string(body))
	err = CheckApiError(body)
	if err != nil {
		return nil, err
	}
	userDetail := UserInfo{}
	json.Unmarshal(body, &userDetail)
	return &userDetail, err
}

func BatchGetUserInfo(accessToken string, openId []string) (*BatchGetUserRes, error) {
	var userList []*UserList
	for _, v := range openId {
		user := UserList{
			OpenId: v,
			Lang:   "zh_CN",
		}
		userList = append(userList, &user)
	}

	req := BatchGetUserReq{
		UserList: userList,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(fmt.Sprintf(
		"https://api.weixin.qq.com/cgi-bin/user/info/batchget?access_token=%s",
		accessToken), "application/json", bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	rspBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Printf("BatchGetUser Rsp : %v \n", string(rspBody))
	err = CheckApiError(rspBody)
	if err != nil {
		return nil, err
	}
	batchRes := BatchGetUserRes{}
	json.Unmarshal(rspBody, &batchRes)
	return &batchRes, err
}
