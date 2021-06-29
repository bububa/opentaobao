package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取当前授权用户手机号码 API请求
taobao.miniapp.user.phone.get

在商家应用中，获取当前授权用户手机号码
*/
type TaobaoMiniappUserPhoneGetRequest struct {
    model.Params
}

// 初始化TaobaoMiniappUserPhoneGetRequest对象
func NewTaobaoMiniappUserPhoneGetRequest() *TaobaoMiniappUserPhoneGetRequest{
    return &TaobaoMiniappUserPhoneGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMiniappUserPhoneGetRequest) GetApiMethodName() string {
    return "taobao.miniapp.user.phone.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMiniappUserPhoneGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
