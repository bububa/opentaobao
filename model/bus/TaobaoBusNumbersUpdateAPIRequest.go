package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobusnumbersupdateAPIRequest 汽车票车次更新服务 API请求
// taobao.bus.numbers.update
//
// 用于汽车票车次信息的新增、更新和逻辑删除
type TaobaobusnumbersupdateAPIRequest struct {
	model.Params
	// 请求参数
	_paramTopBusNumberUpdateRQ *TopBusNumberUpdateRq
}

// NewTaobaobusnumbersupdateRequest 初始化TaobaobusnumbersupdateAPIRequest对象
func NewTaobaobusnumbersupdateRequest() *TaobaobusnumbersupdateAPIRequest {
	return &TaobaobusnumbersupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobusnumbersupdateAPIRequest) GetApiMethodName() string {
	return "taobao.bus.numbers.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobusnumbersupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobusnumbersupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamTopBusNumberUpdateRQ is ParamTopBusNumberUpdateRQ Setter
// 请求参数
func (r *TaobaobusnumbersupdateAPIRequest) SetParamTopBusNumberUpdateRQ(_paramTopBusNumberUpdateRQ *TopBusNumberUpdateRq) error {
	r._paramTopBusNumberUpdateRQ = _paramTopBusNumberUpdateRQ
	r.Set("param_top_bus_number_update_r_q", _paramTopBusNumberUpdateRQ)
	return nil
}

// GetParamTopBusNumberUpdateRQ ParamTopBusNumberUpdateRQ Getter
func (r TaobaobusnumbersupdateAPIRequest) GetParamTopBusNumberUpdateRQ() *TopBusNumberUpdateRq {
	return r._paramTopBusNumberUpdateRQ
}
