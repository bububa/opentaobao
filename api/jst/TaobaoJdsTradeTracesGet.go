package jst

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJdsTradeTracesGet 获取单条订单跟踪详情
// taobao.jds.trade.traces.get
//
// 获取聚石塔数据共享的交易全链路信息
func TaobaoJdsTradeTracesGet(ctx context.Context, clt *core.SDKClient, req *jst.TaobaoJdsTradeTracesGetAPIRequest, resp *jst.TaobaoJdsTradeTracesGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
