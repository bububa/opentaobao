package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票车次查询 APIRequest
taobao.bus.busnumber.get

提供汽车票车次查询服务
*/
type TaobaoBusBusnumberGetRequest struct {
    model.Params

    // 车次查询入参
    paramBusNumberSearchRQ   *BusNumberSearchRq 

}

func NewTaobaoBusBusnumberGetRequest() *TaobaoBusBusnumberGetRequest{
    return &TaobaoBusBusnumberGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBusBusnumberGetRequest) GetApiMethodName() string {
    return "taobao.bus.busnumber.get"
}

func (r TaobaoBusBusnumberGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBusBusnumberGetRequest) SetParamBusNumberSearchRQ(paramBusNumberSearchRQ *BusNumberSearchRq) error {
    r.paramBusNumberSearchRQ = paramBusNumberSearchRQ
    r.Set("param_bus_number_search_r_q", paramBusNumberSearchRQ)
    return nil
}

func (r TaobaoBusBusnumberGetRequest) GetParamBusNumberSearchRQ() *BusNumberSearchRq {
    return r.paramBusNumberSearchRQ
}

