package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取得当前登录用户的授权账户列表 API请求
taobao.simba.customers.authorized.get

取得当前登录用户的授权账户列表
*/
type TaobaoSimbaCustomersAuthorizedGetRequest struct {
    model.Params
}

// 初始化TaobaoSimbaCustomersAuthorizedGetRequest对象
func NewTaobaoSimbaCustomersAuthorizedGetRequest() *TaobaoSimbaCustomersAuthorizedGetRequest{
    return &TaobaoSimbaCustomersAuthorizedGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCustomersAuthorizedGetRequest) GetApiMethodName() string {
    return "taobao.simba.customers.authorized.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCustomersAuthorizedGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
