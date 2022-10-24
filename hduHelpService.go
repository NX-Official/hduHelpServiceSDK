package hduHelpServiceSDK

import (
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"net/url"
	"time"
)

type HduHelpService struct {
	ClientId     string
	ClientSecret string
	TokenInfo    struct {
		AccessToken        string
		AccessTokenExpire  int64
		RefreshToken       string
		RefreshTokenExpire int64
		StaffId            string
	}
}

func New(clientId, clientSecret string, token ...string) *HduHelpService {
	var hduHelpService = &HduHelpService{
		ClientId:     clientId,
		ClientSecret: clientSecret,
	}
	if len(token) != 0 {
		hduHelpService.TokenInfo.AccessToken = token[0]
	}
	return hduHelpService
}

// GetAndSaveToken 从 api.hduhelp.com 获取 token
func (s *HduHelpService) GetAndSaveToken(code, state string) error {
	query := make(url.Values)
	query.Add("client_id", s.ClientId)
	query.Add("client_secret", s.ClientSecret)
	query.Add("grant_type", "authorization_code")
	query.Add("code", code)
	query.Add("state", state)

	reqUrl := url.URL{
		Scheme:   "https",
		Host:     "api.hduhelp.com",
		Path:     "/oauth/token",
		RawQuery: query.Encode(),
	}

	//fmt.Println(reqUrl.String())

	res := GetTokenResponse{}
	_, _, err := gorequest.New().
		Get(reqUrl.String()).
		Retry(3, time.Second, http.StatusBadRequest, http.StatusInternalServerError).
		EndStruct(&res)

	s.TokenInfo = struct {
		AccessToken        string
		AccessTokenExpire  int64
		RefreshToken       string
		RefreshTokenExpire int64
		StaffId            string
	}(res.Data)

	if err != nil {
		return errors.New(fmt.Sprintf("%v", err))
	} else if res.Error != 0 {
		return errors.New(fmt.Sprintf("%v", res.Msg))
	} else {
		return nil
	}
}

// GetStudentInfo 从 SalmonBase 获取学生信息
func (s *HduHelpService) GetStudentInfo() (GetStudentInfoResponse, error) {
	reqUrl := url.URL{
		Scheme: "https",
		Host:   "api.hduhelp.com",
		Path:   "/salmon_base/student/info",
	}
	res := GetStudentInfoResponse{}
	_, _, err := gorequest.New().
		Get(reqUrl.String()).
		Retry(3, time.Second, http.StatusBadRequest, http.StatusInternalServerError).
		AppendHeader("Authorization", "token "+s.TokenInfo.AccessToken).EndStruct(&res)
	if err != nil {
		return res, errors.New(fmt.Sprintf("%v", err))
	} else {
		return res, nil
	}
}
