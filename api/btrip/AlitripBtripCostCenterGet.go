package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCostCenterGet 获取用户费用归属
// alitrip.btrip.cost.center.get
//
// 获取差旅申请用户的费用归属列表
func AlitripBtripCostCenterGet(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripCostCenterGetAPIRequest, resp *btrip.AlitripBtripCostCenterGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
