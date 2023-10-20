package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusMerchantOrderGet 商家侧查询订单详情
// taobao.bus.merchant.order.get
//
// 商家侧查询订单详情
func TaobaoBusMerchantOrderGet(clt *core.SDKClient, req *bus.TaobaoBusMerchantOrderGetAPIRequest, resp *bus.TaobaoBusMerchantOrderGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
