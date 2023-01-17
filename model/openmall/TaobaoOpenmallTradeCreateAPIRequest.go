package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallTradeCreateAPIRequest 创建订单 API请求
// taobao.openmall.trade.create
//
// 创建Openmall订单
type TaobaoOpenmallTradeCreateAPIRequest struct {
	model.Params
	// 请求入参
	_paramTopTradeCreateDO *TopTradeCreateDo
}

// NewTaobaoOpenmallTradeCreateRequest 初始化TaobaoOpenmallTradeCreateAPIRequest对象
func NewTaobaoOpenmallTradeCreateRequest() *TaobaoOpenmallTradeCreateAPIRequest {
	return &TaobaoOpenmallTradeCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallTradeCreateAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.trade.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallTradeCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenmallTradeCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamTopTradeCreateDO is ParamTopTradeCreateDO Setter
// 请求入参
func (r *TaobaoOpenmallTradeCreateAPIRequest) SetParamTopTradeCreateDO(_paramTopTradeCreateDO *TopTradeCreateDo) error {
	r._paramTopTradeCreateDO = _paramTopTradeCreateDO
	r.Set("param_top_trade_create_d_o", _paramTopTradeCreateDO)
	return nil
}

// GetParamTopTradeCreateDO ParamTopTradeCreateDO Getter
func (r TaobaoOpenmallTradeCreateAPIRequest) GetParamTopTradeCreateDO() *TopTradeCreateDo {
	return r._paramTopTradeCreateDO
}
