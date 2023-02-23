package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallTradeRenderAPIRequest 渲染订单价格 API请求
// taobao.openmall.trade.render
//
// 请求渲染订单价格
type TaobaoOpenmallTradeRenderAPIRequest struct {
	model.Params
	// 请求入参
	_paramTopTradeCreateDO *TopTradeCreateDo
}

// NewTaobaoOpenmallTradeRenderRequest 初始化TaobaoOpenmallTradeRenderAPIRequest对象
func NewTaobaoOpenmallTradeRenderRequest() *TaobaoOpenmallTradeRenderAPIRequest {
	return &TaobaoOpenmallTradeRenderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallTradeRenderAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.trade.render"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallTradeRenderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenmallTradeRenderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamTopTradeCreateDO is ParamTopTradeCreateDO Setter
// 请求入参
func (r *TaobaoOpenmallTradeRenderAPIRequest) SetParamTopTradeCreateDO(_paramTopTradeCreateDO *TopTradeCreateDo) error {
	r._paramTopTradeCreateDO = _paramTopTradeCreateDO
	r.Set("param_top_trade_create_d_o", _paramTopTradeCreateDO)
	return nil
}

// GetParamTopTradeCreateDO ParamTopTradeCreateDO Getter
func (r TaobaoOpenmallTradeRenderAPIRequest) GetParamTopTradeCreateDO() *TopTradeCreateDo {
	return r._paramTopTradeCreateDO
}
