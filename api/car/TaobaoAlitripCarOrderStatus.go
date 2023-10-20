package car

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/car"
)

// Taobaoalitripcarorderstatus 商家订单状态改变通知接口（神州专车接口）
// taobao.alitrip.car.order.status
//
// 商家订单状态改变通知接口，神州专车专用接口！
func Taobaoalitripcarorderstatus(clt *core.SDKClient, req *car.TaobaoalitripcarorderstatusAPIRequest, session string) (*car.TaobaoalitripcarorderstatusAPIResponse, error) {
	var resp car.TaobaoalitripcarorderstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
