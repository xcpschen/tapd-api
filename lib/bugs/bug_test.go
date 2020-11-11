package bugs_test

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"

	"github.com/xcpschen/tapd-api/lib"
	"github.com/xcpschen/tapd-api/lib/bugs"
)

func TestBugs(t *testing.T) {
	req := &bugs.AddBug{}
	req.Param = url.Values{}
	req.Param.Add("workspace_id", "")
	req.Param.Add("title", "测试api接口")
	req.Param.Add("reporter", "tapd-api")
	fmt.Println(req.Param.Encode())
	client := lib.NewClient("", "")
	if err := client.Do(req); err != nil {
		t.Fatalf(err.Error())
	} else {
		data, err := req.GetReSponse()
		if err != nil {
			t.Fatalf(err.Error())
		}
		b, _ := json.Marshal(data)
		fmt.Println(string(b))
	}
}

type A struct {
	Name string
}

type B struct {
	A
}

func (a *A) ShowName() string {
	return "a"
}

func (b *B) ShowName(a string) string {
	fmt.Println(b.A.ShowName())
	return a
}
