package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCostCenterNew 新建外部成本中心
// alitrip.btrip.cost.center.new
//
// 新建外部成本中心
func AlitripBtripCostCenterNew(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripCostCenterNewAPIRequest, resp *btrip.AlitripBtripCostCenterNewAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
