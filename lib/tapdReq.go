package lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
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

func (cr *ClientReq) GetReqBody() *strings.Reader {
	if reflect.ValueOf(cr.Param).IsNil() {
		return strings.NewReader("")
	}
	return strings.NewReader(cr.Param.Encode())
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

type IArr []interface{}

func (t *TapdReq) SetReSponse(obj interface{}) error {
	buf := bytes.NewBuffer([]byte{})
	_, err := io.Copy(buf, t.response.Body)
	if err != nil {
		return err
	}
	tmp := new(TapdBaseRespon)
	err = json.Unmarshal(buf.Bytes(), tmp)
	if err != nil {
		return errors.New("Get a error when parse response body " + err.Error())
	}
	defer t.response.Body.Close()

	if tmp.Status != SucessStatusCode {
		return errors.New(tmp.Info)
	}
	var b []byte
	objType := reflect.ValueOf(obj).Elem().Kind()
	tmpType := reflect.TypeOf(tmp.Data).Kind()
	if objType != tmpType {
		if tmpType == reflect.Map && objType == reflect.Slice {
			b, _ = json.Marshal(IArr{tmp.Data})
		} else if tmpType == reflect.Slice && objType == reflect.Map {
			b, _ = json.Marshal(tmp.Data)
			tr := IArr{}
			tr = append(tr, obj)
			if err := json.Unmarshal(b, &tr); err != nil {
				return errors.New("Get a error when parse data " + err.Error())
			}
			return nil
		} else {
			return errors.New("Get a error when parse data ,source typeof " + tmpType.String() + ",target type of " + objType.String())
		}
	} else {
		b, _ = json.Marshal(tmp.Data)
	}
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
