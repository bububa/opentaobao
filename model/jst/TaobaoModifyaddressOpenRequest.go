package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝自助修改地址服务开通 API请求
taobao.modifyaddress.open

商家自助修改地址服务开通
*/
type TaobaoModifyaddressOpenRequest struct {
    model.Params
}

// 初始化TaobaoModifyaddressOpenRequest对象
func NewTaobaoModifyaddressOpenRequest() *TaobaoModifyaddressOpenRequest{
    return &TaobaoModifyaddressOpenRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoModifyaddressOpenRequest) GetApiMethodName() string {
    return "taobao.modifyaddress.open"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoModifyaddressOpenRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
