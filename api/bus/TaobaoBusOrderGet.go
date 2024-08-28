package bus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusOrderGet 汽车票订单查询
// taobao.bus.order.get
//
// 商家汽车票订单查询
func TaobaoBusOrderGet(ctx context.Context, clt *core.SDKClient, req *bus.TaobaoBusOrderGetAPIRequest, resp *bus.TaobaoBusOrderGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
