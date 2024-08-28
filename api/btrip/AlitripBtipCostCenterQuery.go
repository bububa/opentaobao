package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtipCostCenterQuery 查询外部成本中心
// alitrip.btip.cost.center.query
//
// 查询外部成本中心
func AlitripBtipCostCenterQuery(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtipCostCenterQueryAPIRequest, resp *btrip.AlitripBtipCostCenterQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
