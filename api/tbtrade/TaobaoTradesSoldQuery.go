package tbtrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// TaobaoTradesSoldQuery 根据收件人信息查询交易单号
// taobao.trades.sold.query
//
// 根据收件人信息查询交易单号。
func TaobaoTradesSoldQuery(clt *core.SDKClient, req *tbtrade.TaobaoTradesSoldQueryAPIRequest, session string) (*tbtrade.TaobaoTradesSoldQueryAPIResponse, error) {
	var resp tbtrade.TaobaoTradesSoldQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
