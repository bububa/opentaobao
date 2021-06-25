package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川找回密码 APIRequest
taobao.baichuan.openaccount.password.reset

百川找回密码
*/
type TaobaoBaichuanOpenaccountPasswordResetRequest struct {
    model.Params

    // name
    name   string 

}

func NewTaobaoBaichuanOpenaccountPasswordResetRequest() *TaobaoBaichuanOpenaccountPasswordResetRequest{
    return &TaobaoBaichuanOpenaccountPasswordResetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBaichuanOpenaccountPasswordResetRequest) GetApiMethodName() string {
    return "taobao.baichuan.openaccount.password.reset"
}

func (r TaobaoBaichuanOpenaccountPasswordResetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBaichuanOpenaccountPasswordResetRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoBaichuanOpenaccountPasswordResetRequest) GetName() string {
    return r.name
}

