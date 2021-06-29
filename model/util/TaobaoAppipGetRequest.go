package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取ISV发起请求服务器IP API请求
taobao.appip.get

获取ISV发起请求服务器IP
*/
type TaobaoAppipGetRequest struct {
    model.Params
}

// 初始化TaobaoAppipGetRequest对象
func NewTaobaoAppipGetRequest() *TaobaoAppipGetRequest{
    return &TaobaoAppipGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAppipGetRequest) GetApiMethodName() string {
    return "taobao.appip.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAppipGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
