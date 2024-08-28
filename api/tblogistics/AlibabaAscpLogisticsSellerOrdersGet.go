package tblogistics

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// AlibabaAscpLogisticsSellerOrdersGet 商家配送核销订单列表查询
// alibaba.ascp.logistics.seller.orders.get
//
// 商家配送核销订单列表查询
func AlibabaAscpLogisticsSellerOrdersGet(ctx context.Context, clt *core.SDKClient, req *tblogistics.AlibabaAscpLogisticsSellerOrdersGetAPIRequest, resp *tblogistics.AlibabaAscpLogisticsSellerOrdersGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
