package seaking

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
第三方授权 APIRequest
alibaba.alifanyi.market.authenticate

第三方授权，获取授权码
*/
type AlibabaAlifanyiMarketAuthenticateRequest struct {
    model.Params

    // 有效时长
    expireTime   int64 

}

func NewAlibabaAlifanyiMarketAuthenticateRequest() *AlibabaAlifanyiMarketAuthenticateRequest{
    return &AlibabaAlifanyiMarketAuthenticateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlifanyiMarketAuthenticateRequest) GetApiMethodName() string {
    return "alibaba.alifanyi.market.authenticate"
}

func (r AlibabaAlifanyiMarketAuthenticateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlifanyiMarketAuthenticateRequest) SetExpireTime(expireTime int64) error {
    r.expireTime = expireTime
    r.Set("expire_time", expireTime)
    return nil
}

func (r AlibabaAlifanyiMarketAuthenticateRequest) GetExpireTime() int64 {
    return r.expireTime
}

