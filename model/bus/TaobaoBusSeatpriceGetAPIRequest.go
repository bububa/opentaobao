package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusSeatpriceGetAPIRequest 汽车票余票接口 API请求
// taobao.bus.seatprice.get
//
// 提供给商家，查询汽车票班次余票
type TaobaoBusSeatpriceGetAPIRequest struct {
	model.Params
	// 余票请求参数
	_paramBusSeatPriceRQ *BusSeatPriceRq
}

// NewTaobaoBusSeatpriceGetRequest 初始化TaobaoBusSeatpriceGetAPIRequest对象
func NewTaobaoBusSeatpriceGetRequest() *TaobaoBusSeatpriceGetAPIRequest {
	return &TaobaoBusSeatpriceGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusSeatpriceGetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.seatprice.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusSeatpriceGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBusSeatpriceGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBusSeatPriceRQ is ParamBusSeatPriceRQ Setter
// 余票请求参数
func (r *TaobaoBusSeatpriceGetAPIRequest) SetParamBusSeatPriceRQ(_paramBusSeatPriceRQ *BusSeatPriceRq) error {
	r._paramBusSeatPriceRQ = _paramBusSeatPriceRQ
	r.Set("param_bus_seat_price_r_q", _paramBusSeatPriceRQ)
	return nil
}

// GetParamBusSeatPriceRQ ParamBusSeatPriceRQ Getter
func (r TaobaoBusSeatpriceGetAPIRequest) GetParamBusSeatPriceRQ() *BusSeatPriceRq {
	return r._paramBusSeatPriceRQ
}
