package lsttrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lsttrade"
)

// Alibabalsttradeshiporderquery 供应商数据开放--发货单接口
// alibaba.lst.trade.shiporder.query
//
// 供应商数据开放--发货单接口
func Alibabalsttradeshiporderquery(clt *core.SDKClient, req *lsttrade.AlibabalsttradeshiporderqueryAPIRequest, session string) (*lsttrade.AlibabalsttradeshiporderqueryAPIResponse, error) {
	var resp lsttrade.AlibabalsttradeshiporderqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
