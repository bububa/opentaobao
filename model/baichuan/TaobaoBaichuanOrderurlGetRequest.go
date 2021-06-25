package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川订单详情 APIRequest
taobao.baichuan.orderurl.get

百川订单详情
*/
type TaobaoBaichuanOrderurlGetRequest struct {
    model.Params

    // name
    name   string 

}

func NewTaobaoBaichuanOrderurlGetRequest() *TaobaoBaichuanOrderurlGetRequest{
    return &TaobaoBaichuanOrderurlGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBaichuanOrderurlGetRequest) GetApiMethodName() string {
    return "taobao.baichuan.orderurl.get"
}

func (r TaobaoBaichuanOrderurlGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBaichuanOrderurlGetRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoBaichuanOrderurlGetRequest) GetName() string {
    return r.name
}

