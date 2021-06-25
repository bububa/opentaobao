package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川登录二次验证 APIRequest
taobao.baichuan.openaccount.logindoublecheck

百川登录二次验证
*/
type TaobaoBaichuanOpenaccountLogindoublecheckRequest struct {
    model.Params

    // name
    name   string 

}

func NewTaobaoBaichuanOpenaccountLogindoublecheckRequest() *TaobaoBaichuanOpenaccountLogindoublecheckRequest{
    return &TaobaoBaichuanOpenaccountLogindoublecheckRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBaichuanOpenaccountLogindoublecheckRequest) GetApiMethodName() string {
    return "taobao.baichuan.openaccount.logindoublecheck"
}

func (r TaobaoBaichuanOpenaccountLogindoublecheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBaichuanOpenaccountLogindoublecheckRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoBaichuanOpenaccountLogindoublecheckRequest) GetName() string {
    return r.name
}

