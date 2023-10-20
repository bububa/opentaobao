package lsttrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lsttrade"
)

// Alibabalsttradeorderquerychange 订单id批量查询（品牌商视角）
// alibaba.lst.trade.order.querychange
//
// 根据品牌和时间段查询有变更记录的订单id
func Alibabalsttradeorderquerychange(clt *core.SDKClient, req *lsttrade.AlibabalsttradeorderquerychangeAPIRequest, session string) (*lsttrade.AlibabalsttradeorderquerychangeAPIResponse, error) {
	var resp lsttrade.AlibabalsttradeorderquerychangeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
