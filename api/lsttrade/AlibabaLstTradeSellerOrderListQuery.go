package lsttrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lsttrade"
)

// Alibabalsttradesellerorderlistquery 订单列表查看(卖家视角)
// alibaba.lst.trade.seller.order.list.query
//
// 卖家视角订单查询，查询授权经销商订单列表
func Alibabalsttradesellerorderlistquery(clt *core.SDKClient, req *lsttrade.AlibabalsttradesellerorderlistqueryAPIRequest, session string) (*lsttrade.AlibabalsttradesellerorderlistqueryAPIResponse, error) {
	var resp lsttrade.AlibabalsttradesellerorderlistqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
