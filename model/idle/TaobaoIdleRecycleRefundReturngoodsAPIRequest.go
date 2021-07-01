package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoIdleRecycleRefundReturngoodsAPIRequest
闲鱼回收退货V2 API请求
taobao.idle.recycle.refund.returngoods

回收商买家退货，填写退货运单号 */
type TaobaoIdleRecycleRefundReturngoodsAPIRequest struct {
	model.Params
	// 退货
	_param0 *RecycleReturnGoodsRequest
}

// New
