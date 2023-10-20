package car

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/car"
)

// Taobaoalitripcarorderaccept 确认接单
// taobao.alitrip.car.order.accept
//
// 用来接收服务商确认接单信息
func Taobaoalitripcarorderaccept(clt *core.SDKClient, req *car.TaobaoalitripcarorderacceptAPIRequest, session string) (*car.TaobaoalitripcarorderacceptAPIResponse, error) {
	var resp car.TaobaoalitripcarorderacceptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
