package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobusseatpricegetAPIRequest 汽车票余票接口 API请求
// taobao.bus.seatprice.get
//
// 提供给商家，查询汽车票班次余票
type TaobaobusseatpricegetAPIRequest struct {
	model.Params
	// 余票请求参数
	_paramBusSeatPriceRQ *BusSeatPriceRq
}

// NewTaobaobusseatpricegetRequest 初始化TaobaobusseatpricegetAPIRequest对象
func NewTaobaobusseatpricegetRequest() *TaobaobusseatpricegetAPIRequest {
	return &TaobaobusseatpricegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobusseatpricegetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.seatprice.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobusseatpricegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobusseatpricegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBusSeatPriceRQ is ParamBusSeatPriceRQ Setter
// 余票请求参数
func (r *TaobaobusseatpricegetAPIRequest) SetParamBusSeatPriceRQ(_paramBusSeatPriceRQ *BusSeatPriceRq) error {
	r._paramBusSeatPriceRQ = _paramBusSeatPriceRQ
	r.Set("param_bus_seat_price_r_q", _paramBusSeatPriceRQ)
	return nil
}

// GetParamBusSeatPriceRQ ParamBusSeatPriceRQ Getter
func (r TaobaobusseatpricegetAPIRequest) GetParamBusSeatPriceRQ() *BusSeatPriceRq {
	return r._paramBusSeatPriceRQ
}
