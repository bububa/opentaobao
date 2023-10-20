package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripApprovalNew 新建审批单
// alitrip.btrip.approval.new
//
// 用户新建审批单
func AlitripBtripApprovalNew(clt *core.SDKClient, req *btrip.AlitripBtripApprovalNewAPIRequest, resp *btrip.AlitripBtripApprovalNewAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
