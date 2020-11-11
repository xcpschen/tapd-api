package bugs

import (
	"fmt"

	. "github.com/xcpschen/tapd-api/lib"
)

type BugReq struct {
	TapdReq
	WorkspacesID int64
	IsList       bool
	Page         int
	Limit        int
}

func (p *BugReq) Req() *ClientReq {

	return &ClientReq{
		Method: "GET",
		URL:    fmt.Sprintf("bugs?workspace_id=%d", p.WorkspacesID),
	}
}

type BusResponse []map[string]Bug
type Bug struct {
	WorkspaceID      int64  `json:"workspace_id,string"`
	ID               int64  `json:"id,string"`
	IterationID      int64  `json:"iteration_id,string"`
	RegressionNumber int    `json:"regression_number,string"`
	Title            string `json:"title"`
	Description      string `json:"description"`
	Priority         string `json:"priority"`
	Severity         string `json:"severity"`
	Module           string `json:"module"`
	Status           string `json:"status"`
	Reporter         string `json:"reporter"`
	Deadline         string `json:"deadline"`
	Created          string `json:"created"`
	Bugtype          string `json:"bugtype"`
	Resolved         string `json:"resolved"`
	Closed           string `json:"closed"`
	Modified         string `json:"modified"`
	Lastmodify       string `json:"lastmodify"`
	Auditer          string `json:"auditer"`
	De               string `json:"de"`
	Fixer            string `json:"fixer"`
	VersionTest      string `json:"version_test"`
	VersionReport    string `json:"version_report"`
	VersionClose     string `json:"version_close"`
	VersionFix       string `json:"version_fix"`
	BaselineFind     string `json:"baseline_find"`
	BaselineJoin     string `json:"baseline_join"`
	BaselineClose    string `json:"baselineClose"`
	BaselineTest     string `json:"baseline_test"`
	Sourcephase      string `json:"sourcephase"`
	Te               string `json:"te"`
	CurrentOwner     string `json:"current_owner"`
	Resolution       string `json:"resolution"`
	Source           string `json:"source"`
	Originphase      string `json:"originphase"`
	Confirmer        string `json:"confirmer"`
	Milestone        string `json:"milestone"`
	Participator     string `json:"participator"`
	Closer           string `json:"closer"`
	Platform         string `json:"platform"`
	OS               string `json:"os"`
	Testtype         string `json:"testtype"`
	Testphase        string `json:"testphase"`
	Frequency        string `json:"frequency"`
	CC               string `json:"cc"`
	Flows            string `json:"flows"`
	Feature          string `json:"feature"`
	Testmode         string `json:"testmode"`
	Estimate         string `json:"estimate"`
	IssueID          string `json:"issue_id"`
	CreatedFrom      string `json:"created_from"`
	InProgressTime   string `json:"in_progress_time"`
	VerifyTime       string `json:"verify_time"`
	RejectTime       string `json:"reject_time"`
	ReopenTime       string `json:"reopen_time"`
	AuditTime        string `json:"audit_time"`
	SuspendTime      string `json:"suspend_time"`
	Due              string `json:"due"`
	Begin            string `json:"begin"`

	CustomFields map[string]string
}

const bugsKey string = "Bug"

func (p *BugReq) GetReSponse() (data []Bug, err error) {
	tmp := BusResponse{}
	if err = p.SetReSponse(&tmp); err != nil {
		return
	}
	data = make([]Bug, len(tmp))
	for i, obj := range tmp {
		data[i] = obj[bugsKey]
	}
	return
}

type AddBug struct {
	TapdReq
}

func (p *AddBug) Req() *ClientReq {
	return &ClientReq{
		Method: "POST",
		URL:    "bugs",
		Param:  p.Param,
	}
}
func (p *AddBug) GetReSponse() (data Bug, err error) {
	tmp := map[string]Bug{}
	if err = p.SetReSponse(&tmp); err != nil {
		return
	}
	data = tmp[bugsKey]
	return
}
