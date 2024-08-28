package lsttrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lsttrade"
)

// AlibabaLstTradeOrderQuerychange 订单id批量查询（品牌商视角）
// alibaba.lst.trade.order.querychange
//
// 根据品牌和时间段查询有变更记录的订单id
func AlibabaLstTradeOrderQuerychange(ctx context.Context, clt *core.SDKClient, req *lsttrade.AlibabaLstTradeOrderQuerychangeAPIRequest, resp *lsttrade.AlibabaLstTradeOrderQuerychangeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
