package tbtrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// Taobaotradeconfirmfeeget 获取交易确认收货费用
// taobao.trade.confirmfee.get
//
// 获取交易确认收货费用，可以获取主订单或子订单的确认收货费用
func Taobaotradeconfirmfeeget(clt *core.SDKClient, req *tbtrade.TaobaotradeconfirmfeegetAPIRequest, session string) (*tbtrade.TaobaotradeconfirmfeegetAPIResponse, error) {
	var resp tbtrade.TaobaotradeconfirmfeegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
