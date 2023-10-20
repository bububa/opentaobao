package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenmalltraderenderAPIRequest 渲染订单价格 API请求
// taobao.openmall.trade.render
//
// 请求渲染订单价格
type TaobaoopenmalltraderenderAPIRequest struct {
	model.Params
	// 请求入参
	_paramTopTradeCreateDO *TopTradeCreateDo
}

// NewTaobaoopenmalltraderenderRequest 初始化TaobaoopenmalltraderenderAPIRequest对象
func NewTaobaoopenmalltraderenderRequest() *TaobaoopenmalltraderenderAPIRequest {
	return &TaobaoopenmalltraderenderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenmalltraderenderAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.trade.render"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenmalltraderenderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenmalltraderenderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamTopTradeCreateDO is ParamTopTradeCreateDO Setter
// 请求入参
func (r *TaobaoopenmalltraderenderAPIRequest) SetParamTopTradeCreateDO(_paramTopTradeCreateDO *TopTradeCreateDo) error {
	r._paramTopTradeCreateDO = _paramTopTradeCreateDO
	r.Set("param_top_trade_create_d_o", _paramTopTradeCreateDO)
	return nil
}

// GetParamTopTradeCreateDO ParamTopTradeCreateDO Getter
func (r TaobaoopenmalltraderenderAPIRequest) GetParamTopTradeCreateDO() *TopTradeCreateDo {
	return r._paramTopTradeCreateDO
}
