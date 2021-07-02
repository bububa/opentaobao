package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusBusnumberGetAPIRequest 汽车票车次查询 API请求
// taobao.bus.busnumber.get
//
// 提供汽车票车次查询服务
type TaobaoBusBusnumberGetAPIRequest struct {
	model.Params
	// 车次查询入参
	_paramBusNumberSearchRQ *BusNumberSearchRq
}

// NewTaobaoBusBusnumberGetRequest 初始化TaobaoBusBusnumberGetAPIRequest对象
func NewTaobaoBusBusnumberGetRequest() *TaobaoBusBusnumberGetAPIRequest {
	return &TaobaoBusBusnumberGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusBusnumberGetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.busnumber.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusBusnumberGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamBusNumberSearchRQ Setter
// 车次查询入参
func (r *TaobaoBusBusnumberGetAPIRequest) SetParamBusNumberSearchRQ(_paramBusNumberSearchRQ *BusNumberSearchRq) error {
	r._paramBusNumberSearchRQ = _paramBusNumberSearchRQ
	r.Set("param_bus_number_search_r_q", _paramBusNumberSearchRQ)
	return nil
}

// Get ParamBusNumberSearchRQ Getter
func (r TaobaoBusBusnumberGetAPIRequest) GetParamBusNumberSearchRQ() *BusNumberSearchRq {
	return r._paramBusNumberSearchRQ
}
