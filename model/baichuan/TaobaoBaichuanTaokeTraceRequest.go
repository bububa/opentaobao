package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川淘客打点 APIRequest
taobao.baichuan.taoke.trace

百川淘客打点
*/
type TaobaoBaichuanTaokeTraceRequest struct {
    model.Params

    // name
    name   string 

}

func NewTaobaoBaichuanTaokeTraceRequest() *TaobaoBaichuanTaokeTraceRequest{
    return &TaobaoBaichuanTaokeTraceRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBaichuanTaokeTraceRequest) GetApiMethodName() string {
    return "taobao.baichuan.taoke.trace"
}

func (r TaobaoBaichuanTaokeTraceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBaichuanTaokeTraceRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoBaichuanTaokeTraceRequest) GetName() string {
    return r.name
}

