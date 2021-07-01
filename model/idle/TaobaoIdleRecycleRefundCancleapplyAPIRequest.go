package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoIdleRecycleRefundCancleapplyAPIRequest
闲鱼回收取消退款申请V2 API请求
taobao.idle.recycle.refund.cancleapply

回收商的回收订单取消退款申请 */
type TaobaoIdleRecycleRefundCancleapplyAPIRequest struct {
	model.Params
	// 订单号
	_bizOrderId int64
}

// New
