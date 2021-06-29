package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据获取压测用户的sessionKey API请求
taobao.streetest.session.get

根据正常用户sessionKey获取对应压测用户的sessionKey，该sessionKey只能用户服务商全链路压测
*/
type TaobaoStreetestSessionGetRequest struct {
    model.Params
}

// 初始化TaobaoStreetestSessionGetRequest对象
func NewTaobaoStreetestSessionGetRequest() *TaobaoStreetestSessionGetRequest{
    return &TaobaoStreetestSessionGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoStreetestSessionGetRequest) GetApiMethodName() string {
    return "taobao.streetest.session.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoStreetestSessionGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
