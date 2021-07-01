package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取授权session信息 API请求
taobao.openlink.session.get

帮助第三方isv生成三方session
*/
type TaobaoOpenlinkSessionGetAPIRequest struct {
    model.Params
    // 授权码
    _code   string
}

// 初始化TaobaoOpenlinkSessionGetAPIRequest对象
func NewTaobaoOpenlinkSessionGetRequest() *TaobaoOpenlinkSessionGetAPIRequest{
    return &TaobaoOpenlinkSessionGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenlinkSessionGetAPIRequest) GetApiMethodName() string {
    return "taobao.openlink.session.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenlinkSessionGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Code Setter
// 授权码
func (r *TaobaoOpenlinkSessionGetAPIRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r TaobaoOpenlinkSessionGetAPIRequest) GetCode() string {
    return r._code
}
