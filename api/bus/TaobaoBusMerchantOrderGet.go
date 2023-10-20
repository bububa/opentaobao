package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// Taobaobusmerchantorderget 商家侧查询订单详情
// taobao.bus.merchant.order.get
//
// 商家侧查询订单详情
func Taobaobusmerchantorderget(clt *core.SDKClient, req *bus.TaobaobusmerchantordergetAPIRequest, session string) (*bus.TaobaobusmerchantordergetAPIResponse, error) {
	var resp bus.TaobaobusmerchantordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
