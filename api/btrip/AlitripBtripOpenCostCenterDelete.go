package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripOpenCostCenterDelete 删除成本中心
// alitrip.btrip.open.cost.center.delete
//
// 删除成本中心
func AlitripBtripOpenCostCenterDelete(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripOpenCostCenterDeleteAPIRequest, resp *btrip.AlitripBtripOpenCostCenterDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
