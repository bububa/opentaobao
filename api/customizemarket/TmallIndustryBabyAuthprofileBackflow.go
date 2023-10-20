package customizemarket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/customizemarket"
)

// Tmallindustrybabyauthprofilebackflow 孕校云回流档案
// tmall.industry.baby.authprofile.backflow
//
// 孕校云回流档案
func Tmallindustrybabyauthprofilebackflow(clt *core.SDKClient, req *customizemarket.TmallindustrybabyauthprofilebackflowAPIRequest, session string) (*customizemarket.TmallindustrybabyauthprofilebackflowAPIResponse, error) {
	var resp customizemarket.TmallindustrybabyauthprofilebackflowAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
