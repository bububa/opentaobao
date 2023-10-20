package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripapprovalupdate 更新审批单
// alitrip.btrip.approval.update
//
// 更新审批单
func Alitripbtripapprovalupdate(clt *core.SDKClient, req *btrip.AlitripbtripapprovalupdateAPIRequest, session string) (*btrip.AlitripbtripapprovalupdateAPIResponse, error) {
	var resp btrip.AlitripbtripapprovalupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
