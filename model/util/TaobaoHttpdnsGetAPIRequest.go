package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
TOPDNS配置 API请求
taobao.httpdns.get

获取TOP DNS配置
*/
type TaobaoHttpdnsGetAPIRequest struct {
    model.Params
}

// 初始化TaobaoHttpdnsGetAPIRequest对象
func NewTaobaoHttpdnsGetRequest() *TaobaoHttpdnsGetAPIRequest{
    return &TaobaoHttpdnsGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoHttpdnsGetAPIRequest) GetApiMethodName() string {
    return "taobao.httpdns.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoHttpdnsGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
