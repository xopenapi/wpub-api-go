package wpub

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type UploadMediaRes struct {
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt int64  `json:"created_at"`
}

func MediaUpload(filepath, mediaType, accessToken string) (string, error) {
	var resp UploadMediaRes
	err := uploadFile(fmt.Sprintf(
		"https://api.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s",
		accessToken, mediaType), filepath, &resp)
	if err != nil {
		return "", err
	}
	return resp.MediaId, err
}

func MediaGet(mediaId string, filepath string, accessToken string) error {
	err := DownloadFile(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/media/get?access_token=%s&media_id=%s",
		accessToken, mediaId), filepath)
	if err != nil {
		return err
	}
	return err
}

func DownloadFile(url string, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	rspBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	_, err = io.Copy(file, bytes.NewReader(rspBody))
	return err
}

func uploadFile(url string, filePath string, resp *UploadMediaRes) error {
	info, err := os.Stat(filePath)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return errors.New(fmt.Sprintf("path %s is a directory not a file", filePath))
	}
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)
	fileWriter1, err := bodyWriter.CreateFormFile("image", filepath.Base(info.Name()))
	if err != nil {
		return nil
	}
	file1, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file1.Close()
	_, err = io.Copy(fileWriter1, file1)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	req, err := http.NewRequest("POST", url, bodyBuffer)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", contentType)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	buf, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(buf))

	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		if len(buf) > 200 {
			return errors.New(fmt.Sprintf("recv status %d, body \n%s", res.StatusCode, buf[:199]))
		} else {
			return errors.New(fmt.Sprintf("recv status %d, body \n%s", res.StatusCode, buf))
		}
	}

	if err := json.Unmarshal(buf, &resp); err != nil {
		if len(buf) > 200 {
			return errors.New(fmt.Sprintf("marshal error '%v', body '%s'", err, buf[:199]))
		}
		return errors.New(fmt.Sprintf("marshal error '%v', body '%s'", err, buf))
	}
	return nil
}
