package jst

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJdsTradesStatisticsDiff 订单全链路状态统计差异比较
// taobao.jds.trades.statistics.diff
//
// 订单全链路状态统计差异比较
func TaobaoJdsTradesStatisticsDiff(ctx context.Context, clt *core.SDKClient, req *jst.TaobaoJdsTradesStatisticsDiffAPIRequest, resp *jst.TaobaoJdsTradesStatisticsDiffAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
