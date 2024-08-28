package drug

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drug"
)

// AlibabaAlihealthNrTradeOrderGetorderdetail 根据订单id获取单条订单详情
// alibaba.alihealth.nr.trade.order.getorderdetail
//
// 阿里健康O2O，获取订单详情，修复组合商品价格精度问题
func AlibabaAlihealthNrTradeOrderGetorderdetail(ctx context.Context, clt *core.SDKClient, req *drug.AlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest, resp *drug.AlibabaAlihealthNrTradeOrderGetorderdetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
