package tbtrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// TaobaoTradesSoldIncrementvGet 查询卖家已卖出的增量交易数据（根据入库时间）
// taobao.trades.sold.incrementv.get
//
// 搜索当前会话用户作为卖家已卖出的增量交易数据（只能获取到三个月以内的交易信息）
// &lt;br/&gt;1. 一次请求只能查询时间跨度为一天的增量交易记录，即end_create - start_create &lt;= 1天。
// &lt;br/&gt;2. 返回的数据结果是以订单入库时间的倒序排列的(该时间和订单修改时间不同)，通过从后往前翻页的方式可以避免漏单问题。
// &lt;br/&gt;3. 返回的数据结果只包含了订单的部分数据，可通过taobao.trade.fullinfo.get获取订单详情。
// &lt;br/&gt;&lt;strong&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=tradeapi&#34; target=&#34;_blank&#34;&gt;点击查看更多交易API说明&lt;/a&gt;&lt;/strong&gt;
func TaobaoTradesSoldIncrementvGet(ctx context.Context, clt *core.SDKClient, req *tbtrade.TaobaoTradesSoldIncrementvGetAPIRequest, resp *tbtrade.TaobaoTradesSoldIncrementvGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
