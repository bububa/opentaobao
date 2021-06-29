package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取授权账号对应的OpenUid API请求
taobao.openuid.get

获取授权账号对应的OpenUid
*/
type TaobaoOpenuidGetRequest struct {
    model.Params
}

// 初始化TaobaoOpenuidGetRequest对象
func NewTaobaoOpenuidGetRequest() *TaobaoOpenuidGetRequest{
    return &TaobaoOpenuidGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenuidGetRequest) GetApiMethodName() string {
    return "taobao.openuid.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenuidGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
