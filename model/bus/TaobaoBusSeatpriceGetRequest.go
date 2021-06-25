package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票余票接口 APIRequest
taobao.bus.seatprice.get

提供给商家，查询汽车票班次余票
*/
type TaobaoBusSeatpriceGetRequest struct {
    model.Params

    // 余票请求参数
    paramBusSeatPriceRQ   *BusSeatPriceRq 

}

func NewTaobaoBusSeatpriceGetRequest() *TaobaoBusSeatpriceGetRequest{
    return &TaobaoBusSeatpriceGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBusSeatpriceGetRequest) GetApiMethodName() string {
    return "taobao.bus.seatprice.get"
}

func (r TaobaoBusSeatpriceGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBusSeatpriceGetRequest) SetParamBusSeatPriceRQ(paramBusSeatPriceRQ *BusSeatPriceRq) error {
    r.paramBusSeatPriceRQ = paramBusSeatPriceRQ
    r.Set("param_bus_seat_price_r_q", paramBusSeatPriceRQ)
    return nil
}

func (r TaobaoBusSeatpriceGetRequest) GetParamBusSeatPriceRQ() *BusSeatPriceRq {
    return r.paramBusSeatPriceRQ
}

