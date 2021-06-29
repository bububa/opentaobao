package seaking

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
第三方授权 API请求
alibaba.alifanyi.market.authenticate

第三方授权，获取授权码
*/
type AlibabaAlifanyiMarketAuthenticateRequest struct {
    model.Params
    // 有效时长
    expireTime   int64
}

// 初始化AlibabaAlifanyiMarketAuthenticateRequest对象
func NewAlibabaAlifanyiMarketAuthenticateRequest() *AlibabaAlifanyiMarketAuthenticateRequest{
    return &AlibabaAlifanyiMarketAuthenticateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlifanyiMarketAuthenticateRequest) GetApiMethodName() string {
    return "alibaba.alifanyi.market.authenticate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlifanyiMarketAuthenticateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ExpireTime Setter
// 有效时长
func (r *AlibabaAlifanyiMarketAuthenticateRequest) SetExpireTime(expireTime int64) error {
    r.expireTime = expireTime
    r.Set("expire_time", expireTime)
    return nil
}

// ExpireTime Getter
func (r AlibabaAlifanyiMarketAuthenticateRequest) GetExpireTime() int64 {
    return r.expireTime
}
