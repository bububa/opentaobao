package drug

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drug"
)

// AlibabaAlihealthNrTradeOrderGet 获取订单详情
// alibaba.alihealth.nr.trade.order.get
//
// 阿里健康O2O，获取订单详情
func AlibabaAlihealthNrTradeOrderGet(ctx context.Context, clt *core.SDKClient, req *drug.AlibabaAlihealthNrTradeOrderGetAPIRequest, resp *drug.AlibabaAlihealthNrTradeOrderGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
