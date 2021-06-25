package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川验证找回密码验证码 APIRequest
taobao.baichuan.openaccount.resetcode.check

百川验证找回密码验证码
*/
type TaobaoBaichuanOpenaccountResetcodeCheckRequest struct {
    model.Params

    // name
    name   string 

}

func NewTaobaoBaichuanOpenaccountResetcodeCheckRequest() *TaobaoBaichuanOpenaccountResetcodeCheckRequest{
    return &TaobaoBaichuanOpenaccountResetcodeCheckRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBaichuanOpenaccountResetcodeCheckRequest) GetApiMethodName() string {
    return "taobao.baichuan.openaccount.resetcode.check"
}

func (r TaobaoBaichuanOpenaccountResetcodeCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBaichuanOpenaccountResetcodeCheckRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoBaichuanOpenaccountResetcodeCheckRequest) GetName() string {
    return r.name
}

