package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川H5登录二次验证 APIRequest
taobao.baichuan.user.logindoublecheck

百川H5登录二次验证
*/
type TaobaoBaichuanUserLogindoublecheckRequest struct {
    model.Params

    // name
    name   string 

}

func NewTaobaoBaichuanUserLogindoublecheckRequest() *TaobaoBaichuanUserLogindoublecheckRequest{
    return &TaobaoBaichuanUserLogindoublecheckRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBaichuanUserLogindoublecheckRequest) GetApiMethodName() string {
    return "taobao.baichuan.user.logindoublecheck"
}

func (r TaobaoBaichuanUserLogindoublecheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBaichuanUserLogindoublecheckRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoBaichuanUserLogindoublecheckRequest) GetName() string {
    return r.name
}

