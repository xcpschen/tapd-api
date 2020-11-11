package workspaces

import (
	"fmt"

	. "github.com/csl/tapd-api/lib"
)

type ProjectUser struct {
	TapdReq
	WorkspacesID int64
	fields       string
}

func (p *ProjectUser) Req() *ClientReq {
	return &ClientReq{
		Method: "GET",
		URL:    fmt.Sprintf("workspaces/users?workspace_id=%d&fields=user,role_id,email", p.WorkspacesID),
	}
}

type ProjectUserRsp struct {
	TapdBaseRespon
	Data []map[string]PUser `json:"data"`
}
type PUser struct {
	User    string   `json:"user"`
	Email   string   `json:"email"`
	Name    string   `json:"name"`
	RoleIDS []string `json:"role_id"`
}

const UserWorkspace string = "UserWorkspace"

func (p *ProjectUser) GetReSponse() (data []PUser, err error) {
	tmp := []map[string]PUser{}
	if err = p.SetReSponse(&tmp); err != nil {
		return
	}
	data = make([]PUser, len(tmp))
	for i, obj := range tmp {
		data[i] = obj[UserWorkspace]
	}
	return
}
