package alitripcar

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripcar"
)

// TaobaoAlitripCarOrderRefund 用户投诉达成一致后给用户退款
// taobao.alitrip.car.order.refund
//
// 用户投诉后，供应商客服与客户沟通达成一致后通知飞猪给客户退款。退款金额以接口回调金额为准。
func TaobaoAlitripCarOrderRefund(ctx context.Context, clt *core.SDKClient, req *alitripcar.TaobaoAlitripCarOrderRefundAPIRequest, resp *alitripcar.TaobaoAlitripCarOrderRefundAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
