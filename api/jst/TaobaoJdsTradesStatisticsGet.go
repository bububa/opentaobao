package jst

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJdsTradesStatisticsGet 获取订单数量统计结果
// taobao.jds.trades.statistics.get
//
// 获取订单数量统计结果
func TaobaoJdsTradesStatisticsGet(ctx context.Context, clt *core.SDKClient, req *jst.TaobaoJdsTradesStatisticsGetAPIRequest, resp *jst.TaobaoJdsTradesStatisticsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
