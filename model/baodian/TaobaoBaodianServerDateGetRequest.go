package baodian

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务器时间获取 APIRequest
taobao.baodian.server.date.get

获取服务器时间
*/
type TaobaoBaodianServerDateGetRequest struct {
    model.Params

}

func NewTaobaoBaodianServerDateGetRequest() *TaobaoBaodianServerDateGetRequest{
    return &TaobaoBaodianServerDateGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBaodianServerDateGetRequest) GetApiMethodName() string {
    return "taobao.baodian.server.date.get"
}

func (r TaobaoBaodianServerDateGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


