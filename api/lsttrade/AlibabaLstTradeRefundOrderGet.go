package lsttrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lsttrade"
)

// Alibabalsttraderefundorderget 零售通退款订单查询
// alibaba.lst.trade.refund.order.get
//
// 零售通退款订单查询
func Alibabalsttraderefundorderget(clt *core.SDKClient, req *lsttrade.AlibabalsttraderefundordergetAPIRequest, session string) (*lsttrade.AlibabalsttraderefundordergetAPIResponse, error) {
	var resp lsttrade.AlibabalsttraderefundordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
