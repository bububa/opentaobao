package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripapprovalmodify 修改审批单
// alitrip.btrip.approval.modify
//
// 修改审批单
func Alitripbtripapprovalmodify(clt *core.SDKClient, req *btrip.AlitripbtripapprovalmodifyAPIRequest, session string) (*btrip.AlitripbtripapprovalmodifyAPIResponse, error) {
	var resp btrip.AlitripbtripapprovalmodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
