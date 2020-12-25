package openplatform

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"io"
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

func AesDecrypt(cipherData []byte, aesKey []byte) ([]byte, error) {
	k := len(aesKey) //PKCS#7
	if len(cipherData)%k != 0 {
		return nil, errors.New("crypto/cipher: ciphertext size is not multiple of aes key length")
	}
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	plainData := make([]byte, len(cipherData))
	blockMode.CryptBlocks(plainData, cipherData)
	return plainData, nil
}

func EncodingAESKey2AESKey(encodingKey string) []byte {
	data, _ := base64.StdEncoding.DecodeString(encodingKey + "=")
	return data
}
