package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// Taobaobusorderget 汽车票订单查询
// taobao.bus.order.get
//
// 商家汽车票订单查询
func Taobaobusorderget(clt *core.SDKClient, req *bus.TaobaobusordergetAPIRequest, session string) (*bus.TaobaobusordergetAPIResponse, error) {
	var resp bus.TaobaobusordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
