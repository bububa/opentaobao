package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripApprovalUpdate 更新审批单
// alitrip.btrip.approval.update
//
// 更新审批单
func AlitripBtripApprovalUpdate(clt *core.SDKClient, req *btrip.AlitripBtripApprovalUpdateAPIRequest, resp *btrip.AlitripBtripApprovalUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
