package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// TaobaoIdleRecycleRefundCancleapply 闲鱼回收取消退款申请V2
// taobao.idle.recycle.refund.cancleapply
//
// 回收商的回收订单取消退款申请
func TaobaoIdleRecycleRefundCancleapply(clt *core.SDKClient, req *idle.TaobaoIdleRecycleRefundCancleapplyAPIRequest, resp *idle.TaobaoIdleRecycleRefundCancleapplyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
