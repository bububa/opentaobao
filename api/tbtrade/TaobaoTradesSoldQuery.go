package tbtrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// TaobaoTradesSoldQuery 根据收件人信息查询交易单号
// taobao.trades.sold.query
//
// 根据收件人信息查询交易单号。
func TaobaoTradesSoldQuery(clt *core.SDKClient, req *tbtrade.TaobaoTradesSoldQueryAPIRequest, resp *tbtrade.TaobaoTradesSoldQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
