package lsttrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lsttrade"
)

// Alibabalsttradesellerorderdetailquery 订单详情查看(卖家视角)
// alibaba.lst.trade.seller.order.detail.query
//
// 订单详情查看(卖家视角)
func Alibabalsttradesellerorderdetailquery(clt *core.SDKClient, req *lsttrade.AlibabalsttradesellerorderdetailqueryAPIRequest, session string) (*lsttrade.AlibabalsttradesellerorderdetailqueryAPIResponse, error) {
	var resp lsttrade.AlibabalsttradesellerorderdetailqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
