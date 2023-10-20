package tbtrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// TaobaoTradeFullinfoGet 获取单笔交易的详细信息
// taobao.trade.fullinfo.get
//
// 获取单笔交易的详细信息
// &lt;br/&gt;1. 只有单笔订单的情况下Trade数据结构中才包含商品相关的信息
// &lt;br/&gt;2. 获取到的Order中的payment字段在单笔子订单时包含物流费用，多笔子订单时不包含物流费用
// &lt;br/&gt;3. 获取红包金额使用字段：tmall_coupon_fee（天猫商家订单使用的红包金额）
// &lt;br/&gt;4. 请按需获取字段，减少TOP系统的压力
// &lt;br/&gt;&lt;strong&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=tradeapi&#34; target=&#34;_blank&#34;&gt;点击查看更多交易API说明&lt;/a&gt;&lt;/strong&gt;
func TaobaoTradeFullinfoGet(clt *core.SDKClient, req *tbtrade.TaobaoTradeFullinfoGetAPIRequest, resp *tbtrade.TaobaoTradeFullinfoGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
