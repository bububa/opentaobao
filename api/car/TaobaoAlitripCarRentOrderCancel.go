package car

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/car"
)

// Taobaoalitripcarrentordercancel 租车-取消订单
// taobao.alitrip.car.rent.order.cancel
//
// 服务商主动取消用户订单或者拒绝取消订单.
func Taobaoalitripcarrentordercancel(clt *core.SDKClient, req *car.TaobaoalitripcarrentordercancelAPIRequest, session string) (*car.TaobaoalitripcarrentordercancelAPIResponse, error) {
	var resp car.TaobaoalitripcarrentordercancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
