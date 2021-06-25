package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取目的地数据 APIRequest
taobao.bus.lastplace.get

传入城市 获取对应的目的地
*/
type TaobaoBusLastplaceGetRequest struct {
    model.Params

    // 目的地查询参数
    paramLastPlaceSearchRQ   *ParamLastPlaceSearchRq 

}

func NewTaobaoBusLastplaceGetRequest() *TaobaoBusLastplaceGetRequest{
    return &TaobaoBusLastplaceGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBusLastplaceGetRequest) GetApiMethodName() string {
    return "taobao.bus.lastplace.get"
}

func (r TaobaoBusLastplaceGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBusLastplaceGetRequest) SetParamLastPlaceSearchRQ(paramLastPlaceSearchRQ *ParamLastPlaceSearchRq) error {
    r.paramLastPlaceSearchRQ = paramLastPlaceSearchRQ
    r.Set("param_last_place_search_r_q", paramLastPlaceSearchRQ)
    return nil
}

func (r TaobaoBusLastplaceGetRequest) GetParamLastPlaceSearchRQ() *ParamLastPlaceSearchRq {
    return r.paramLastPlaceSearchRQ
}

