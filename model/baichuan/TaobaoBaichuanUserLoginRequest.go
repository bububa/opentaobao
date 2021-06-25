package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川H5登录 APIRequest
taobao.baichuan.user.login

百川H5登录
*/
type TaobaoBaichuanUserLoginRequest struct {
    model.Params

    // name
    name   string 

}

func NewTaobaoBaichuanUserLoginRequest() *TaobaoBaichuanUserLoginRequest{
    return &TaobaoBaichuanUserLoginRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBaichuanUserLoginRequest) GetApiMethodName() string {
    return "taobao.baichuan.user.login"
}

func (r TaobaoBaichuanUserLoginRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBaichuanUserLoginRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoBaichuanUserLoginRequest) GetName() string {
    return r.name
}

