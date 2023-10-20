package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusOrderGet 汽车票订单查询
// taobao.bus.order.get
//
// 商家汽车票订单查询
func TaobaoBusOrderGet(clt *core.SDKClient, req *bus.TaobaoBusOrderGetAPIRequest, resp *bus.TaobaoBusOrderGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
