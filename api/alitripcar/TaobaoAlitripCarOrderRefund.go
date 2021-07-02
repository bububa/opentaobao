package alitripcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripcar"
)

// TaobaoAlitripCarOrderRefund 用户投诉达成一致后给用户退款
// taobao.alitrip.car.order.refund
//
// 用户投诉后，供应商客服与客户沟通达成一致后通知飞猪给客户退款。退款金额以接口回调金额为准。
func TaobaoAlitripCarOrderRefund(clt *core.SDKClient, req *alitripcar.TaobaoAlitripCarOrderRefundAPIRequest, session string) (*alitripcar.TaobaoAlitripCarOrderRefundAPIResponse, error) {
	var resp alitripcar.TaobaoAlitripCarOrderRefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
