package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJdsTradesStatisticsGet 获取订单数量统计结果
// taobao.jds.trades.statistics.get
//
// 获取订单数量统计结果
func TaobaoJdsTradesStatisticsGet(clt *core.SDKClient, req *jst.TaobaoJdsTradesStatisticsGetAPIRequest, resp *jst.TaobaoJdsTradesStatisticsGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
