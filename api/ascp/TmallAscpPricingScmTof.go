package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Tmallascppricingscmtof TOF&SCM营销域对接-成本录入设置
// tmall.ascp.pricing.scm.tof
//
// TOF&amp;SCM营销域对接-成本录入设置
func Tmallascppricingscmtof(clt *core.SDKClient, req *ascp.TmallascppricingscmtofAPIRequest, session string) (*ascp.TmallascppricingscmtofAPIResponse, error) {
	var resp ascp.TmallascppricingscmtofAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
