package tblogistics

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// AlibabaAscpLogisticsSellerSend 商家配送发货
// alibaba.ascp.logistics.seller.send
//
// 该API提供商家配送发货能力，使用该接口发货，交易订单状态会直接变成卖家已发货
func AlibabaAscpLogisticsSellerSend(ctx context.Context, clt *core.SDKClient, req *tblogistics.AlibabaAscpLogisticsSellerSendAPIRequest, resp *tblogistics.AlibabaAscpLogisticsSellerSendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
