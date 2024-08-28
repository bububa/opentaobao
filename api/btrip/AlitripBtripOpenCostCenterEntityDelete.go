package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripOpenCostCenterEntityDelete 删除成本中心人员信息
// alitrip.btrip.open.cost.center.entity.delete
//
// 删除成本中心人员信息
func AlitripBtripOpenCostCenterEntityDelete(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripOpenCostCenterEntityDeleteAPIRequest, resp *btrip.AlitripBtripOpenCostCenterEntityDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
