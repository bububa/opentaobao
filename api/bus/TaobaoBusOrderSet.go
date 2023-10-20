package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// Taobaobusorderset 汽车票下单接口
// taobao.bus.order.set
//
// 提供给汽车票商家进行下单
func Taobaobusorderset(clt *core.SDKClient, req *bus.TaobaobusordersetAPIRequest, session string) (*bus.TaobaobusordersetAPIResponse, error) {
	var resp bus.TaobaobusordersetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
