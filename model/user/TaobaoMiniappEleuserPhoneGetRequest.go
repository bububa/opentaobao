package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取饿了么用户信息 API请求
taobao.miniapp.eleuser.phone.get

获取饿了么用户信息
*/
type TaobaoMiniappEleuserPhoneGetRequest struct {
    model.Params
}

// 初始化TaobaoMiniappEleuserPhoneGetRequest对象
func NewTaobaoMiniappEleuserPhoneGetRequest() *TaobaoMiniappEleuserPhoneGetRequest{
    return &TaobaoMiniappEleuserPhoneGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMiniappEleuserPhoneGetRequest) GetApiMethodName() string {
    return "taobao.miniapp.eleuser.phone.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMiniappEleuserPhoneGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
