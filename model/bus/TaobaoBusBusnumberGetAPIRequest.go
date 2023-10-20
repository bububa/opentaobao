package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobusbusnumbergetAPIRequest 汽车票车次查询 API请求
// taobao.bus.busnumber.get
//
// 提供汽车票车次查询服务
type TaobaobusbusnumbergetAPIRequest struct {
	model.Params
	// 车次查询入参
	_paramBusNumberSearchRQ *BusNumberSearchRq
}

// NewTaobaobusbusnumbergetRequest 初始化TaobaobusbusnumbergetAPIRequest对象
func NewTaobaobusbusnumbergetRequest() *TaobaobusbusnumbergetAPIRequest {
	return &TaobaobusbusnumbergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobusbusnumbergetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.busnumber.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobusbusnumbergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobusbusnumbergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBusNumberSearchRQ is ParamBusNumberSearchRQ Setter
// 车次查询入参
func (r *TaobaobusbusnumbergetAPIRequest) SetParamBusNumberSearchRQ(_paramBusNumberSearchRQ *BusNumberSearchRq) error {
	r._paramBusNumberSearchRQ = _paramBusNumberSearchRQ
	r.Set("param_bus_number_search_r_q", _paramBusNumberSearchRQ)
	return nil
}

// GetParamBusNumberSearchRQ ParamBusNumberSearchRQ Getter
func (r TaobaobusbusnumbergetAPIRequest) GetParamBusNumberSearchRQ() *BusNumberSearchRq {
	return r._paramBusNumberSearchRQ
}
