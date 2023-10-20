package aedropshiper

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// AliexpressDsTradeOrderGet 交易订单查询
// aliexpress.ds.trade.order.get
//
// 交易订单查询
func AliexpressDsTradeOrderGet(clt *core.SDKClient, req *aedropshiper.AliexpressDsTradeOrderGetAPIRequest, session string) (*aedropshiper.AliexpressDsTradeOrderGetAPIResponse, error) {
	var resp aedropshiper.AliexpressDsTradeOrderGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
