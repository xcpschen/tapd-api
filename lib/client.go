package lib

import (
	"encoding/base64"
	"net/http"
	"time"
)

//Client tapd Api 客户端
type Client struct {
	apiUser       string
	apiPassWord   string
	authorization string
}
type IRequest interface {
	Req() *ClientReq
	SaveRespone(*http.Response)
}

const (
	//TapdAPIBaseURL tapd api url
	TapdAPIBaseURL string = "https://api.tapd.cn/"
)

var httpClient = &http.Client{
	Timeout: time.Minute * 1,
	Transport: &http.Transport{
		MaxIdleConns:        10,
		MaxIdleConnsPerHost: 1,
		MaxConnsPerHost:     2,
		IdleConnTimeout:     time.Minute * 3,
	},
}

func NewClient(user, pwd string) *Client {
	return &Client{apiUser: user, apiPassWord: pwd}
}
func NewClientWithConfInter(conf Iconfig) *Client {
	user, paw := conf.GetApiUserAndPaw()
	return &Client{apiUser: user, apiPassWord: paw}
}
func (c *Client) setAuthorization() {
	str := c.apiUser + ":" + c.apiPassWord
	base64.StdEncoding.EncodeToString([]byte(str))
	c.authorization = "Basic " + base64.StdEncoding.EncodeToString([]byte(str))
}

func (c *Client) setHTTPReq(tr *ClientReq) *http.Request {
	if c.authorization == "" {
		c.setAuthorization()
	}
	r, _ := http.NewRequest(tr.Method, TapdAPIBaseURL+tr.URL, tr)
	r.Header.Add("Authorization", c.authorization)
	return r
}
func (c *Client) Do(r IRequest) error {
	req := c.setHTTPReq(r.Req())
	respone, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	r.SaveRespone(respone)
	return nil
}
