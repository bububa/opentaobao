package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川发送找回密码验证码 APIRequest
taobao.baichuan.openaccount.resetcode.send

百川发送找回密码验证码
*/
type TaobaoBaichuanOpenaccountResetcodeSendRequest struct {
    model.Params

    // name
    name   string 

}

func NewTaobaoBaichuanOpenaccountResetcodeSendRequest() *TaobaoBaichuanOpenaccountResetcodeSendRequest{
    return &TaobaoBaichuanOpenaccountResetcodeSendRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBaichuanOpenaccountResetcodeSendRequest) GetApiMethodName() string {
    return "taobao.baichuan.openaccount.resetcode.send"
}

func (r TaobaoBaichuanOpenaccountResetcodeSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBaichuanOpenaccountResetcodeSendRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoBaichuanOpenaccountResetcodeSendRequest) GetName() string {
    return r.name
}

