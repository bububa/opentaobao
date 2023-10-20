package alitripcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripcar"
)

// Taobaoalitripcarorderagentcancel 司机或客服取消订单
// taobao.alitrip.car.order.agent.cancel
//
// 司机或客服取消订单后通知飞猪订单取消
func Taobaoalitripcarorderagentcancel(clt *core.SDKClient, req *alitripcar.TaobaoalitripcarorderagentcancelAPIRequest, session string) (*alitripcar.TaobaoalitripcarorderagentcancelAPIResponse, error) {
	var resp alitripcar.TaobaoalitripcarorderagentcancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
