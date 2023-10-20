package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripApprovalModify 修改审批单
// alitrip.btrip.approval.modify
//
// 修改审批单
func AlitripBtripApprovalModify(clt *core.SDKClient, req *btrip.AlitripBtripApprovalModifyAPIRequest, resp *btrip.AlitripBtripApprovalModifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
