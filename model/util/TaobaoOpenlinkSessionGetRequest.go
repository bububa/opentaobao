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
type TaobaoOpenlinkSessionGetRequest struct {
    model.Params
    // 授权码
    code   string
}

// 初始化TaobaoOpenlinkSessionGetRequest对象
func NewTaobaoOpenlinkSessionGetRequest() *TaobaoOpenlinkSessionGetRequest{
    return &TaobaoOpenlinkSessionGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenlinkSessionGetRequest) GetApiMethodName() string {
    return "taobao.openlink.session.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenlinkSessionGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Code Setter
// 授权码
func (r *TaobaoOpenlinkSessionGetRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

// Code Getter
func (r TaobaoOpenlinkSessionGetRequest) GetCode() string {
    return r.code
}
