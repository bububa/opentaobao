package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川账号注册 APIRequest
taobao.baichuan.openaccount.register

百川账号注册
*/
type TaobaoBaichuanOpenaccountRegisterRequest struct {
    model.Params

    // name
    name   string 

}

func NewTaobaoBaichuanOpenaccountRegisterRequest() *TaobaoBaichuanOpenaccountRegisterRequest{
    return &TaobaoBaichuanOpenaccountRegisterRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBaichuanOpenaccountRegisterRequest) GetApiMethodName() string {
    return "taobao.baichuan.openaccount.register"
}

func (r TaobaoBaichuanOpenaccountRegisterRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBaichuanOpenaccountRegisterRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoBaichuanOpenaccountRegisterRequest) GetName() string {
    return r.name
}

