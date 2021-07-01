package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查看功能权限 API请求
taobao.simba.customers.sid.get

查询用户是否拥有某个功能权限
*/
type TaobaoSimbaCustomersSidGetAPIRequest struct {
    model.Params
}

// 初始化TaobaoSimbaCustomersSidGetAPIRequest对象
func NewTaobaoSimbaCustomersSidGetRequest() *TaobaoSimbaCustomersSidGetAPIRequest{
    return &TaobaoSimbaCustomersSidGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCustomersSidGetAPIRequest) GetApiMethodName() string {
    return "taobao.simba.customers.sid.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCustomersSidGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
