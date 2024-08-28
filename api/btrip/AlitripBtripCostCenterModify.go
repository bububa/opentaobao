package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCostCenterModify 修改外部成本中心
// alitrip.btrip.cost.center.modify
//
// 修改外部成本中心，设置成员，设置支付宝账号，设置名称，编号等
func AlitripBtripCostCenterModify(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripCostCenterModifyAPIRequest, resp *btrip.AlitripBtripCostCenterModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
