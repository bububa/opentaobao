package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取开放平台出口IP段 API请求
taobao.top.ipout.get

获取开放平台出口IP段
*/
type TaobaoTopIpoutGetRequest struct {
    model.Params
}

// 初始化TaobaoTopIpoutGetRequest对象
func NewTaobaoTopIpoutGetRequest() *TaobaoTopIpoutGetRequest{
    return &TaobaoTopIpoutGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTopIpoutGetRequest) GetApiMethodName() string {
    return "taobao.top.ipout.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTopIpoutGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
