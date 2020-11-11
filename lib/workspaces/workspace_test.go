package workspaces_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/csl/tapd-api/lib"
	"github.com/csl/tapd-api/lib/workspaces"
)

func TestProjectList(t *testing.T) {
	project := &workspaces.Project{CompanyID: 43}
	client := lib.NewClient("", "")
	if err := client.Do(project); err != nil {
		t.Fatalf(err.Error())
	} else {
		data, err := project.GetReSponse()
		if err != nil {
			t.Fatalf(err.Error())
		}
		b, _ := json.Marshal(data)
		fmt.Println(string(b))

	}
}

func TestProjectUser(t *testing.T) {
	users := &workspaces.ProjectUser{WorkspacesID: 000}
	client := lib.NewClient("", "")
	if err := client.Do(users); err != nil {
		t.Fatalf(err.Error())
	} else {
		data, err := users.GetReSponse()
		if err != nil {
			t.Fatalf(err.Error())
		}
		b, _ := json.Marshal(data)
		fmt.Println(string(b))

	}
}
