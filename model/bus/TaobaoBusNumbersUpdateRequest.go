package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票车次更新服务 APIRequest
taobao.bus.numbers.update

用于汽车票车次信息的新增、更新和逻辑删除
*/
type TaobaoBusNumbersUpdateRequest struct {
    model.Params

    // 请求参数
    paramTopBusNumberUpdateRQ   *TopBusNumberUpdateRq 

}

func NewTaobaoBusNumbersUpdateRequest() *TaobaoBusNumbersUpdateRequest{
    return &TaobaoBusNumbersUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBusNumbersUpdateRequest) GetApiMethodName() string {
    return "taobao.bus.numbers.update"
}

func (r TaobaoBusNumbersUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBusNumbersUpdateRequest) SetParamTopBusNumberUpdateRQ(paramTopBusNumberUpdateRQ *TopBusNumberUpdateRq) error {
    r.paramTopBusNumberUpdateRQ = paramTopBusNumberUpdateRQ
    r.Set("param_top_bus_number_update_r_q", paramTopBusNumberUpdateRQ)
    return nil
}

func (r TaobaoBusNumbersUpdateRequest) GetParamTopBusNumberUpdateRQ() *TopBusNumberUpdateRq {
    return r.paramTopBusNumberUpdateRQ
}

