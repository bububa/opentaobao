package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
城市接口 APIRequest
taobao.bus.city.get

汽车票出发城市获取接口，获取所有出发城市
*/
type TaobaoBusCityGetRequest struct {
    model.Params

}

func NewTaobaoBusCityGetRequest() *TaobaoBusCityGetRequest{
    return &TaobaoBusCityGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBusCityGetRequest) GetApiMethodName() string {
    return "taobao.bus.city.get"
}

func (r TaobaoBusCityGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


