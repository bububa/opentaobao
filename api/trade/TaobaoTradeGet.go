package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// TaobaoTradeGet 获取单笔交易的部分信息(性能高)
// taobao.trade.get
//
// 获取单笔交易的部分信息
// &lt;br/&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=tradeapi&#34; target=&#34;_blank&#34;&gt;点击查看更多交易API说明&lt;/a&gt;
func TaobaoTradeGet(clt *core.SDKClient, req *trade.TaobaoTradeGetAPIRequest, session string) (*trade.TaobaoTradeGetAPIResponse, error) {
	var resp trade.TaobaoTradeGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
