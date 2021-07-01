package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoIdleRecycleRefundDetailAPIRequest
闲鱼回收退款详情V2 API请求
taobao.idle.recycle.refund.detail

回收订单退款详情，主要包括退款状态，超时时间，和同意退款的卖家退货地址信息 */
type TaobaoIdleRecycleRefundDetailAPIRequest struct {
	model.Params
	// 订单号
	_bizOrderId int64
}

// New
