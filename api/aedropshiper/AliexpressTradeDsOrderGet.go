package aedropshiper

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// AliexpressTradeDsOrderGet 买家查询订单详情
// aliexpress.trade.ds.order.get
//
// 买家查询订单详情，用于dropshipper
func AliexpressTradeDsOrderGet(clt *core.SDKClient, req *aedropshiper.AliexpressTradeDsOrderGetAPIRequest, resp *aedropshiper.AliexpressTradeDsOrderGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
