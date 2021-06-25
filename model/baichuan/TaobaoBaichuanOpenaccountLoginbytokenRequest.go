package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川TOKEN 登录 APIRequest
taobao.baichuan.openaccount.loginbytoken

百川TOKEN 登录
*/
type TaobaoBaichuanOpenaccountLoginbytokenRequest struct {
    model.Params

    // name
    name   string 

}

func NewTaobaoBaichuanOpenaccountLoginbytokenRequest() *TaobaoBaichuanOpenaccountLoginbytokenRequest{
    return &TaobaoBaichuanOpenaccountLoginbytokenRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBaichuanOpenaccountLoginbytokenRequest) GetApiMethodName() string {
    return "taobao.baichuan.openaccount.loginbytoken"
}

func (r TaobaoBaichuanOpenaccountLoginbytokenRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBaichuanOpenaccountLoginbytokenRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoBaichuanOpenaccountLoginbytokenRequest) GetName() string {
    return r.name
}

