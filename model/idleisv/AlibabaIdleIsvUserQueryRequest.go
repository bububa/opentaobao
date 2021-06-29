package idleisv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商ISV闲鱼用户信息查询 APIRequest
alibaba.idle.isv.user.query

服务商ISV闲鱼用户信息查询
*/
type AlibabaIdleIsvUserQueryRequest struct {
    model.Params

}

func NewAlibabaIdleIsvUserQueryRequest() *AlibabaIdleIsvUserQueryRequest{
    return &AlibabaIdleIsvUserQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleIsvUserQueryRequest) GetApiMethodName() string {
    return "alibaba.idle.isv.user.query"
}

func (r AlibabaIdleIsvUserQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


