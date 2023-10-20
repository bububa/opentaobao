package idle

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoIdleRecycleRefundReturngoodsAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoIdleRecycleRefundReturngoodsAPIRequest) GetApiMethodName() string {
	return "taobao.idle.recycle.refund.returngoods"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoIdleRecycleRefundReturngoodsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoIdleRecycleRefundReturngoodsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 退货
func (r *TaobaoIdleRecycleRefundReturngoodsAPIRequest) SetParam0(_param0 *RecycleReturnGoodsRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoIdleRecycleRefundReturngoodsAPIRequest) GetParam0() *RecycleReturnGoodsRequest {
	return r._param0
}

var poolTaobaoIdleRecycleRefundReturngoodsAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoIdleRecycleRefundReturngoodsRequest()
	},
}

// GetTaobaoIdleRecycleRefundReturngoodsRequest 从 sync.Pool 获取 TaobaoIdleRecycleRefundReturngoodsAPIRequest
func GetTaobaoIdleRecycleRefundReturngoodsAPIRequest() *TaobaoIdleRecycleRefundReturngoodsAPIRequest {
	return poolTaobaoIdleRecycleRefundReturngoodsAPIRequest.Get().(*TaobaoIdleRecycleRefundReturngoodsAPIRequest)
}

// ReleaseTaobaoIdleRecycleRefundReturngoodsAPIRequest 将 TaobaoIdleRecycleRefundReturngoodsAPIRequest 放入 sync.Pool
func ReleaseTaobaoIdleRecycleRefundReturngoodsAPIRequest(v *TaobaoIdleRecycleRefundReturngoodsAPIRequest) {
	v.Reset()
	poolTaobaoIdleRecycleRefundReturngoodsAPIRequest.Put(v)
}
