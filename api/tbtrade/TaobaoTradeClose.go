package tbtrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// TaobaoTradeClose 卖家关闭一笔交易
// taobao.trade.close
//
// 关闭一笔订单，可以是主订单或子订单。当订单从创建到关闭时间小于10s的时候，会报“CLOSE_TRADE_TOO_FAST”错误。
func TaobaoTradeClose(clt *core.SDKClient, req *tbtrade.TaobaoTradeCloseAPIRequest, resp *tbtrade.TaobaoTradeCloseAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
