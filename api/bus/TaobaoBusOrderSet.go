package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusOrderSet 汽车票下单接口
// taobao.bus.order.set
//
// 提供给汽车票商家进行下单
func TaobaoBusOrderSet(clt *core.SDKClient, req *bus.TaobaoBusOrderSetAPIRequest, resp *bus.TaobaoBusOrderSetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
