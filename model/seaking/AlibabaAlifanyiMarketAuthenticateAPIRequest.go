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
type AlibabaAlifanyiMarketAuthenticateAPIRequest struct {
    model.Params
    // 有效时长
    _expireTime   int64
}

// 初始化AlibabaAlifanyiMarketAuthenticateAPIRequest对象
func NewAlibabaAlifanyiMarketAuthenticateRequest() *AlibabaAlifanyiMarketAuthenticateAPIRequest{
    return &AlibabaAlifanyiMarketAuthenticateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlifanyiMarketAuthenticateAPIRequest) GetApiMethodName() string {
    return "alibaba.alifanyi.market.authenticate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlifanyiMarketAuthenticateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ExpireTime Setter
// 有效时长
func (r *AlibabaAlifanyiMarketAuthenticateAPIRequest) SetExpireTime(_expireTime int64) error {
    r._expireTime = _expireTime
    r.Set("expire_time", _expireTime)
    return nil
}

// ExpireTime Getter
func (r AlibabaAlifanyiMarketAuthenticateAPIRequest) GetExpireTime() int64 {
    return r._expireTime
}
