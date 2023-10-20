package car

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/car"
)

// Taobaoalitripcarordercomplete 服务完成API
// taobao.alitrip.car.order.complete
//
// 用来接收服务商订单流程完成信息
func Taobaoalitripcarordercomplete(clt *core.SDKClient, req *car.TaobaoalitripcarordercompleteAPIRequest, session string) (*car.TaobaoalitripcarordercompleteAPIResponse, error) {
	var resp car.TaobaoalitripcarordercompleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
