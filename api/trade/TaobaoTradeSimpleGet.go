package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// TaobaoTradeSimpleGet 获取交易订单的简易信息
// taobao.trade.simple.get
//
// 获取CRM单表交易的详情
// &lt;br/&gt;1. 只有单笔订单的情况下Trade数据结构中才包含商品相关的信息
// &lt;br/&gt;2. 获取到的Order中的payment字段在单笔子订单时包含物流费用，多笔子订单时不包含物流费用
// &lt;br/&gt;3. 获取红包优惠金额可以使用字段 coupon_fee
// &lt;br/&gt;4. 请按需获取字段，减少TOP系统的压力
// &lt;br/&gt;&lt;strong&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=tradeapi&#34; target=&#34;_blank&#34;&gt;点击查看更多交易API说明&lt;/a&gt;&lt;/strong&gt;
func TaobaoTradeSimpleGet(clt *core.SDKClient, req *trade.TaobaoTradeSimpleGetAPIRequest, session string) (*trade.TaobaoTradeSimpleGetAPIResponse, error) {
	var resp trade.TaobaoTradeSimpleGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
