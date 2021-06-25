package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
城市变更 APIRequest
taobao.bus.agent.city.change

代理商通知城市变更，比如可售变为不可售等
*/
type TaobaoBusAgentCityChangeRequest struct {
    model.Params

    // 城市变更请求对象
    paramCityChangeRQ   *CityChangeRq 

}

func NewTaobaoBusAgentCityChangeRequest() *TaobaoBusAgentCityChangeRequest{
    return &TaobaoBusAgentCityChangeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBusAgentCityChangeRequest) GetApiMethodName() string {
    return "taobao.bus.agent.city.change"
}

func (r TaobaoBusAgentCityChangeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBusAgentCityChangeRequest) SetParamCityChangeRQ(paramCityChangeRQ *CityChangeRq) error {
    r.paramCityChangeRQ = paramCityChangeRQ
    r.Set("param_city_change_r_q", paramCityChangeRQ)
    return nil
}

func (r TaobaoBusAgentCityChangeRequest) GetParamCityChangeRQ() *CityChangeRq {
    return r.paramCityChangeRQ
}

