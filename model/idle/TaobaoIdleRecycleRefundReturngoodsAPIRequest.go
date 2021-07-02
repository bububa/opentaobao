package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoIdleRecycleRefundReturngoodsAPIRequest 闲鱼回收退货V2 API请求
// taobao.idle.recycle.refund.returngoods
//
// 回收商买家退货，填写退货运单号
type TaobaoIdleRecycleRefundReturngoodsAPIRequest struct {
	model.Params
	// 退货
	_param0 *RecycleReturnGoodsRequest
}

// NewTaobaoIdleRecycleRefundReturngoodsRequest 初始化TaobaoIdleRecycleRefundReturngoodsAPIRequest对象
func NewTaobaoIdleRecycleRefundReturngoodsRequest() *TaobaoIdleRecycleRefundReturngoodsAPIRequest {
	return &TaobaoIdleRecycleRefundReturngoodsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoIdleRecycleRefundReturngoodsAPIRequest) GetApiMethodName() string {
	return "taobao.idle.recycle.refund.returngoods"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoIdleRecycleRefundReturngoodsAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 退货
func (r *TaobaoIdleRecycleRefundReturngoodsAPIRequest) SetParam0(_param0 *RecycleReturnGoodsRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r TaobaoIdleRecycleRefundReturngoodsAPIRequest) GetParam0() *RecycleReturnGoodsRequest {
	return r._param0
}
