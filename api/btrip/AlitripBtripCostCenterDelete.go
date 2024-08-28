package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCostCenterDelete 删除外部成本中心
// alitrip.btrip.cost.center.delete
//
// 删除外部成本中心
func AlitripBtripCostCenterDelete(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripCostCenterDeleteAPIRequest, resp *btrip.AlitripBtripCostCenterDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
