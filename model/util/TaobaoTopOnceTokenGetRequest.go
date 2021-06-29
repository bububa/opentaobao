package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
网关一次性token获取 API请求
taobao.top.once.token.get

网关一次性token获取
*/
type TaobaoTopOnceTokenGetRequest struct {
    model.Params
    // sec_token
    secToken   string
}

// 初始化TaobaoTopOnceTokenGetRequest对象
func NewTaobaoTopOnceTokenGetRequest() *TaobaoTopOnceTokenGetRequest{
    return &TaobaoTopOnceTokenGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTopOnceTokenGetRequest) GetApiMethodName() string {
    return "taobao.top.once.token.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTopOnceTokenGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SecToken Setter
// sec_token
func (r *TaobaoTopOnceTokenGetRequest) SetSecToken(secToken string) error {
    r.secToken = secToken
    r.Set("sec_token", secToken)
    return nil
}

// SecToken Getter
func (r TaobaoTopOnceTokenGetRequest) GetSecToken() string {
    return r.secToken
}
