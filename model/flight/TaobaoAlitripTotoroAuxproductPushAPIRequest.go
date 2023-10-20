package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptotoroauxproductpushAPIRequest 廉航辅营产品投放 API请求
// taobao.alitrip.totoro.auxproduct.push
//
// 廉航辅营产品投放接口
type TaobaoalitriptotoroauxproductpushAPIRequest struct {
	model.Params
	// 廉航辅营产品推送请求
	_pushAuxProductsRq *PushAuxProductsRq
}

// NewTaobaoalitriptotoroauxproductpushRequest 初始化TaobaoalitriptotoroauxproductpushAPIRequest对象
func NewTaobaoalitriptotoroauxproductpushRequest() *TaobaoalitriptotoroauxproductpushAPIRequest {
	return &TaobaoalitriptotoroauxproductpushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptotoroauxproductpushAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.totoro.auxproduct.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptotoroauxproductpushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptotoroauxproductpushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPushAuxProductsRq is PushAuxProductsRq Setter
// 廉航辅营产品推送请求
func (r *TaobaoalitriptotoroauxproductpushAPIRequest) SetPushAuxProductsRq(_pushAuxProductsRq *PushAuxProductsRq) error {
	r._pushAuxProductsRq = _pushAuxProductsRq
	r.Set("push_aux_products_rq", _pushAuxProductsRq)
	return nil
}

// GetPushAuxProductsRq PushAuxProductsRq Getter
func (r TaobaoalitriptotoroauxproductpushAPIRequest) GetPushAuxProductsRq() *PushAuxProductsRq {
	return r._pushAuxProductsRq
}
