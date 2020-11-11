package bugs_test

import (
	"fmt"
	"testing"
)

func TestBugs(t *testing.T) {
	// req := &bugs.BugReq{
	// 	WorkspacesID: int64(49974820),
	// }
	// client := lib.NewClient("g7!fK6fy", "CD37F684-F728-28B7-76E4-54C06BB4FC5C")
	// if err := client.Do(req); err != nil {
	// 	t.Fatalf(err.Error())
	// } else {
	// 	data, err := req.GetReSponse()
	// 	if err != nil {
	// 		t.Fatalf(err.Error())
	// 	}
	// 	b, _ := json.Marshal(data)
	// 	fmt.Println(string(b))

	// }
	b := new(B)
	fmt.Println(b.ShowName("I'm b"))
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
