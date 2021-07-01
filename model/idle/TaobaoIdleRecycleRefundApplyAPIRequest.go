package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoIdleRecycleRefundApplyAPIRequest
闲鱼回收交易退款申请V2 API请求
taobao.idle.recycle.refund.apply

回收商买家申请退款 */
type TaobaoIdleRecycleRefundApplyAPIRequest struct {
	model.Params
	// 退款申请
	_refundApply *RecycleRefundTopRequest
}

// New
