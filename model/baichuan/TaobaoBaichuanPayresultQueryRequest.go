package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川支付完成回调 APIRequest
taobao.baichuan.payresult.query

百川支付完成回调
*/
type TaobaoBaichuanPayresultQueryRequest struct {
    model.Params

    // name
    name   string 

}

func NewTaobaoBaichuanPayresultQueryRequest() *TaobaoBaichuanPayresultQueryRequest{
    return &TaobaoBaichuanPayresultQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBaichuanPayresultQueryRequest) GetApiMethodName() string {
    return "taobao.baichuan.payresult.query"
}

func (r TaobaoBaichuanPayresultQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBaichuanPayresultQueryRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoBaichuanPayresultQueryRequest) GetName() string {
    return r.name
}

