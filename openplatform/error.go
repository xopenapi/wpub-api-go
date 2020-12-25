package openplatform

import (
	"encoding/json"
	"errors"
	"strconv"
)

type ApiError struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func CheckApiError(body []byte) error {
	apiErr := ApiError{}
	err := json.Unmarshal(body, &apiErr)
	if err != nil {
		return err
	}
	if apiErr.Errcode != 0 {
		return errors.New("[" + strconv.FormatInt(apiErr.Errcode, 10) + "][" + apiErr.Errmsg + "]")
	}
	return nil
}
