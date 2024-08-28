package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripApprovalModify 修改审批单
// alitrip.btrip.approval.modify
//
// 修改审批单
func AlitripBtripApprovalModify(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripApprovalModifyAPIRequest, resp *btrip.AlitripBtripApprovalModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
