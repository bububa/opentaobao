package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripOpenCostCenterEntityAdd 增加成本中心人员信息
// alitrip.btrip.open.cost.center.entity.add
//
// 增加成本中心人员信息
func AlitripBtripOpenCostCenterEntityAdd(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripOpenCostCenterEntityAddAPIRequest, resp *btrip.AlitripBtripOpenCostCenterEntityAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
