package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川用户名密码登录 APIRequest
taobao.baichuan.openaccount.login

百川用户名密码登录
*/
type TaobaoBaichuanOpenaccountLoginRequest struct {
    model.Params

    // name
    name   string 

}

func NewTaobaoBaichuanOpenaccountLoginRequest() *TaobaoBaichuanOpenaccountLoginRequest{
    return &TaobaoBaichuanOpenaccountLoginRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBaichuanOpenaccountLoginRequest) GetApiMethodName() string {
    return "taobao.baichuan.openaccount.login"
}

func (r TaobaoBaichuanOpenaccountLoginRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBaichuanOpenaccountLoginRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoBaichuanOpenaccountLoginRequest) GetName() string {
    return r.name
}

