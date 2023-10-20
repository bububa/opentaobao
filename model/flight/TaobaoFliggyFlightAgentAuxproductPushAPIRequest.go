package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFliggyFlightAgentAuxproductPushAPIRequest 飞猪机票辅营商品投放 API请求
// taobao.fliggy.flight.agent.auxproduct.push
//
// 廉航辅营产品投放接口
type TaobaoFliggyFlightAgentAuxproductPushAPIRequest struct {
	model.Params
	// 廉航辅营产品推送请求
	_pushAuxProductsRq *PushAuxProductsRq
}

// NewTaobaoFliggyFlightAgentAuxproductPushRequest 初始化TaobaoFliggyFlightAgentAuxproductPushAPIRequest对象
func NewTaobaoFliggyFlightAgentAuxproductPushRequest() *TaobaoFliggyFlightAgentAuxproductPushAPIRequest {
	return &TaobaoFliggyFlightAgentAuxproductPushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFliggyFlightAgentAuxproductPushAPIRequest) GetApiMethodName() string {
	return "taobao.fliggy.flight.agent.auxproduct.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFliggyFlightAgentAuxproductPushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFliggyFlightAgentAuxproductPushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPushAuxProductsRq is PushAuxProductsRq Setter
// 廉航辅营产品推送请求
func (r *TaobaoFliggyFlightAgentAuxproductPushAPIRequest) SetPushAuxProductsRq(_pushAuxProductsRq *PushAuxProductsRq) error {
	r._pushAuxProductsRq = _pushAuxProductsRq
	r.Set("push_aux_products_rq", _pushAuxProductsRq)
	return nil
}

// GetPushAuxProductsRq PushAuxProductsRq Getter
func (r TaobaoFliggyFlightAgentAuxproductPushAPIRequest) GetPushAuxProductsRq() *PushAuxProductsRq {
	return r._pushAuxProductsRq
}
