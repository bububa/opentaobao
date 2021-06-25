package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川检查注册验证码 APIRequest
taobao.baichuan.openaccount.registercode.check

百川检查注册验证码
*/
type TaobaoBaichuanOpenaccountRegistercodeCheckRequest struct {
    model.Params

    // name
    name   string 

}

func NewTaobaoBaichuanOpenaccountRegistercodeCheckRequest() *TaobaoBaichuanOpenaccountRegistercodeCheckRequest{
    return &TaobaoBaichuanOpenaccountRegistercodeCheckRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBaichuanOpenaccountRegistercodeCheckRequest) GetApiMethodName() string {
    return "taobao.baichuan.openaccount.registercode.check"
}

func (r TaobaoBaichuanOpenaccountRegistercodeCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBaichuanOpenaccountRegistercodeCheckRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoBaichuanOpenaccountRegistercodeCheckRequest) GetName() string {
    return r.name
}

