package idleisv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商ISV闲鱼用户信息查询 API请求
alibaba.idle.isv.user.query

服务商ISV闲鱼用户信息查询
*/
type AlibabaIdleIsvUserQueryAPIRequest struct {
    model.Params
}

// 初始化AlibabaIdleIsvUserQueryAPIRequest对象
func NewAlibabaIdleIsvUserQueryRequest() *AlibabaIdleIsvUserQueryAPIRequest{
    return &AlibabaIdleIsvUserQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvUserQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.idle.isv.user.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvUserQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
