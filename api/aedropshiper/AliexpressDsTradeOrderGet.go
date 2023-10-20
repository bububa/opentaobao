package aedropshiper

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// AliexpressDsTradeOrderGet 交易订单查询
// aliexpress.ds.trade.order.get
//
// 交易订单查询
func AliexpressDsTradeOrderGet(clt *core.SDKClient, req *aedropshiper.AliexpressDsTradeOrderGetAPIRequest, resp *aedropshiper.AliexpressDsTradeOrderGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
