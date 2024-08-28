package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripOpenCostCenterEntitySet 设置成本中心人员信息
// alitrip.btrip.open.cost.center.entity.set
//
// 设置成本中心人员信息
func AlitripBtripOpenCostCenterEntitySet(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripOpenCostCenterEntitySetAPIRequest, resp *btrip.AlitripBtripOpenCostCenterEntitySetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
