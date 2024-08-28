package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangInventoryFutureplan 负卖计划
// alibaba.dchain.aoxiang.inventory.futureplan
//
// 负卖计划。底层有白名单控制，并非对所有商家开放。如果需要使用，请联系对应的小二增加白名单
func AlibabaDchainAoxiangInventoryFutureplan(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangInventoryFutureplanAPIRequest, resp *ascp.AlibabaDchainAoxiangInventoryFutureplanAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
