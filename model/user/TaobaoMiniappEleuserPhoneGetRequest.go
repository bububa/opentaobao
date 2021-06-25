package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取饿了么用户信息 APIRequest
taobao.miniapp.eleuser.phone.get

获取饿了么用户信息
*/
type TaobaoMiniappEleuserPhoneGetRequest struct {
    model.Params

}

func NewTaobaoMiniappEleuserPhoneGetRequest() *TaobaoMiniappEleuserPhoneGetRequest{
    return &TaobaoMiniappEleuserPhoneGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMiniappEleuserPhoneGetRequest) GetApiMethodName() string {
    return "taobao.miniapp.eleuser.phone.get"
}

func (r TaobaoMiniappEleuserPhoneGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


