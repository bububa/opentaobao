package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCostCenterEntityAdd 增加外部成本中心人员信息
// alitrip.btrip.cost.center.entity.add
//
// 增加外部成本中心人员信息
func AlitripBtripCostCenterEntityAdd(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripCostCenterEntityAddAPIRequest, resp *btrip.AlitripBtripCostCenterEntityAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
