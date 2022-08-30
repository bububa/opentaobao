package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// TaobaoTradesSimpleSoldGet 查询卖家已卖出的交易简易数据
// taobao.trades.simple.sold.get
//
// 搜索当前会话用户作为卖家已卖出的交易数据（只能获取到三个月以内的交易信息）
// <br/>1. 返回的数据结果是以订单的创建时间倒序排列的。
// <br/>2. 返回的数据结果只包含了订单的部分数据，可通过taobao.trade.simple.get获取订单信息。
// <br/>注意：type字段的说明，如果该字段不传，接口默认只查4种类型订单，非默认查询的订单是不返回。遇到订单查不到的情况的，解决办法就是type加上订单类型就可正常返回了。
// <br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=tradeapi" target="_blank">点击查看更多交易API说明</a></strong>
func TaobaoTradesSimpleSoldGet(clt *core.SDKClient, req *trade.TaobaoTradesSimpleSoldGetAPIRequest, session string) (*trade.TaobaoTradesSimpleSoldGetAPIResponse, error) {
	var resp trade.TaobaoTradesSimpleSoldGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
