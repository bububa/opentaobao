package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川手淘信任登录 APIRequest
taobao.baichuan.user.loginbytoken

百川手淘信任登录
*/
type TaobaoBaichuanUserLoginbytokenRequest struct {
    model.Params

    // name
    name   string 

}

func NewTaobaoBaichuanUserLoginbytokenRequest() *TaobaoBaichuanUserLoginbytokenRequest{
    return &TaobaoBaichuanUserLoginbytokenRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBaichuanUserLoginbytokenRequest) GetApiMethodName() string {
    return "taobao.baichuan.user.loginbytoken"
}

func (r TaobaoBaichuanUserLoginbytokenRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBaichuanUserLoginbytokenRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoBaichuanUserLoginbytokenRequest) GetName() string {
    return r.name
}

