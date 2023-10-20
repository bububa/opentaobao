package lsttrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lsttrade"
)

// Alibabalsttradeorderrefundlistquery 查询退款单列表(卖家视角)
// alibaba.lst.trade.order.refund.list.query
//
// 查询退款单列表(卖家视角)
func Alibabalsttradeorderrefundlistquery(clt *core.SDKClient, req *lsttrade.AlibabalsttradeorderrefundlistqueryAPIRequest, session string) (*lsttrade.AlibabalsttradeorderrefundlistqueryAPIResponse, error) {
	var resp lsttrade.AlibabalsttradeorderrefundlistqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
