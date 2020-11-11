package workspaces

import (
	"fmt"

	. "github.com/csl/tapd-api/lib"
)

type Project struct {
	TapdReq
	CompanyID int64
}
type Workspace struct {
	ID          int64  `json:"id,string"`
	CreatorID   int64  `json:"creator_id,string"`
	Secrecy     int    `json:"secrecy,string"`
	MemberCount int    `json:"member_count"`
	Name        string `json:"name"`
	PrettyName  string `json:"pretty_name"`
	Status      string `json:"status"`
	Created     string `json:"created"`
	Creator     string `json:"creator"`
}

func (p *Project) Req() *ClientReq {
	return &ClientReq{
		Method: "GET",
		URL:    fmt.Sprintf("workspaces/projects?company_id=%d", p.CompanyID),
	}
}

const workspaceKey string = "Workspace"

type ProjectRsp struct {
	TapdBaseRespon
	Data []map[string]Workspace `json:"data"`
}

func (p *Project) GetReSponse() (data []Workspace, err error) {
	tmp := []map[string]Workspace{}
	if err = p.SetReSponse(&tmp); err != nil {
		return
	}
	data = make([]Workspace, len(tmp))
	for i, obj := range tmp {
		data[i] = obj[workspaceKey]
	}
	return
}
