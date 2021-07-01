package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

/* TaobaoTradeGet
获取单笔交易的部分信息(性能高)
taobao.trade.get

获取单笔交易的部分信息
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=tradeapi" target="_blank">点击查看更多交易API说明</a></strong> */
func TaobaoTradeGet(clt *core.SDKClient, req *trade.TaobaoTradeGetAPIRequest, session string) (*trade.TaobaoTradeGetAPIResponse, error) {
	var resp trade.TaobaoTradeGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
