package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川发送注册验证码 APIRequest
taobao.baichuan.openaccount.registercode.send

百川发送注册验证码
*/
type TaobaoBaichuanOpenaccountRegistercodeSendRequest struct {
    model.Params

    // name
    name   string 

}

func NewTaobaoBaichuanOpenaccountRegistercodeSendRequest() *TaobaoBaichuanOpenaccountRegistercodeSendRequest{
    return &TaobaoBaichuanOpenaccountRegistercodeSendRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBaichuanOpenaccountRegistercodeSendRequest) GetApiMethodName() string {
    return "taobao.baichuan.openaccount.registercode.send"
}

func (r TaobaoBaichuanOpenaccountRegistercodeSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBaichuanOpenaccountRegistercodeSendRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoBaichuanOpenaccountRegistercodeSendRequest) GetName() string {
    return r.name
}

