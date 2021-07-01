package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票余票接口 API请求
taobao.bus.seatprice.get

提供给商家，查询汽车票班次余票
*/
type TaobaoBusSeatpriceGetAPIRequest struct {
    model.Params
    // 余票请求参数
    _paramBusSeatPriceRQ   *BusSeatPriceRq
}

// 初始化TaobaoBusSeatpriceGetAPIRequest对象
func NewTaobaoBusSeatpriceGetRequest() *TaobaoBusSeatpriceGetAPIRequest{
    return &TaobaoBusSeatpriceGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBusSeatpriceGetAPIRequest) GetApiMethodName() string {
    return "taobao.bus.seatprice.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBusSeatpriceGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamBusSeatPriceRQ Setter
// 余票请求参数
func (r *TaobaoBusSeatpriceGetAPIRequest) SetParamBusSeatPriceRQ(_paramBusSeatPriceRQ *BusSeatPriceRq) error {
    r._paramBusSeatPriceRQ = _paramBusSeatPriceRQ
    r.Set("param_bus_seat_price_r_q", _paramBusSeatPriceRQ)
    return nil
}

// ParamBusSeatPriceRQ Getter
func (r TaobaoBusSeatpriceGetAPIRequest) GetParamBusSeatPriceRQ() *BusSeatPriceRq {
    return r._paramBusSeatPriceRQ
}
