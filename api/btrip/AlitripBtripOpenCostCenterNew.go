package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripOpenCostCenterNew 新增成本中心
// alitrip.btrip.open.cost.center.new
//
// 新增成本中心
func AlitripBtripOpenCostCenterNew(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripOpenCostCenterNewAPIRequest, resp *btrip.AlitripBtripOpenCostCenterNewAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
