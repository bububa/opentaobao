package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCostCenterEntityDelete 删除外部成本中心人员信息
// alitrip.btrip.cost.center.entity.delete
//
// 删除外部成本中心人员信息
func AlitripBtripCostCenterEntityDelete(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripCostCenterEntityDeleteAPIRequest, resp *btrip.AlitripBtripCostCenterEntityDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
