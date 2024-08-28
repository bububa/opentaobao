package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCostCenterEntitySet 设置外部成本中心人员信息
// alitrip.btrip.cost.center.entity.set
//
// 设置外部成本中心人员信息
func AlitripBtripCostCenterEntitySet(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripCostCenterEntitySetAPIRequest, resp *btrip.AlitripBtripCostCenterEntitySetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
