package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取当前授权用户手机号码 APIRequest
taobao.miniapp.user.phone.get

在商家应用中，获取当前授权用户手机号码
*/
type TaobaoMiniappUserPhoneGetRequest struct {
    model.Params

}

func NewTaobaoMiniappUserPhoneGetRequest() *TaobaoMiniappUserPhoneGetRequest{
    return &TaobaoMiniappUserPhoneGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMiniappUserPhoneGetRequest) GetApiMethodName() string {
    return "taobao.miniapp.user.phone.get"
}

func (r TaobaoMiniappUserPhoneGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


