package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoidlerecyclerefundreturngoodsAPIRequest 闲鱼回收退货V2 API请求
// taobao.idle.recycle.refund.returngoods
//
// 回收商买家退货，填写退货运单号
type TaobaoidlerecyclerefundreturngoodsAPIRequest struct {
	model.Params
	// 退货
	_param0 *RecycleReturnGoodsRequest
}

// NewTaobaoidlerecyclerefundreturngoodsRequest 初始化TaobaoidlerecyclerefundreturngoodsAPIRequest对象
func NewTaobaoidlerecyclerefundreturngoodsRequest() *TaobaoidlerecyclerefundreturngoodsAPIRequest {
	return &TaobaoidlerecyclerefundreturngoodsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoidlerecyclerefundreturngoodsAPIRequest) GetApiMethodName() string {
	return "taobao.idle.recycle.refund.returngoods"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoidlerecyclerefundreturngoodsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoidlerecyclerefundreturngoodsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 退货
func (r *TaobaoidlerecyclerefundreturngoodsAPIRequest) SetParam0(_param0 *RecycleReturnGoodsRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoidlerecyclerefundreturngoodsAPIRequest) GetParam0() *RecycleReturnGoodsRequest {
	return r._param0
}
