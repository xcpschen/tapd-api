package lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type TapdReq struct {
	Param    url.Values
	response *http.Response
}

type ClientReq struct {
	Header      map[string]string
	lastReadInt int
	Method      string
	URL         string
	DataType    string
	Param       url.Values
	Data        interface{}
}

//请求内容格式
const (
	ClientReqJSONDataType     string = "json"
	ClientReqFormDataDataType string = "formdata"
)

func (cr *ClientReq) Read(b []byte) (n int, err error) {

	var (
		mb []byte
		ml int
	)
	l := len(b)
	mb = []byte(cr.Param.Encode())
	ml = len(mb)
	if ml <= 0 {
		return 0, io.EOF
	}
	if cr.lastReadInt == 0 {
		b = mb[:l]
		cr.lastReadInt = l
		n = l
	} else {
		if ml-cr.lastReadInt >= l {
			tmp := cr.lastReadInt + l
			b = mb[cr.lastReadInt:l]
			cr.lastReadInt = tmp
			n = l
		} else {
			n = ml - cr.lastReadInt
			b = mb[cr.lastReadInt:]
			cr.lastReadInt = ml
			err = io.EOF
		}
	}
	return
}

type TapdBaseRespon struct {
	Status int         `json:"status"`
	Info   string      `json:"info"`
	Data   interface{} `json:"data"`
}

const (
	//SucessStatusCode  成功
	SucessStatusCode int = 1
)

func (t *TapdReq) SaveRespone(r *http.Response) {
	t.response = r
}

func (t *TapdReq) SetReSponse(obj interface{}) error {
	buf := bytes.NewBuffer([]byte{})
	_, err := io.Copy(buf, t.response.Body)
	if err != nil {
		return err
	}
	tmp := new(TapdBaseRespon)
	fmt.Println(buf.String())
	err = json.Unmarshal(buf.Bytes(), tmp)
	if err != nil {
		return errors.New("Get a error when parse response body " + err.Error())
	}
	defer t.response.Body.Close()

	if tmp.Status != SucessStatusCode {
		return errors.New(tmp.Info)
	}
	b, _ := json.Marshal(tmp.Data)
	if err := json.Unmarshal(b, obj); err != nil {
		return errors.New("Get a error when parse data " + err.Error())
	}
	return nil
}

func (t *TapdReq) GetReSponseStatus() int {
	if t.response == nil {
		return 0
	}
	return t.response.StatusCode
}
