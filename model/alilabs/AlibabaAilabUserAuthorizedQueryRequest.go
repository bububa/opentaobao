package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询授权状态接口 APIRequest
alibaba.ailab.user.authorized.query

查询三方用户授权状态
*/
type AlibabaAilabUserAuthorizedQueryRequest struct {
    model.Params

    // 开放平台申请的schema
    schemaKey   string 

    // 三方用户的唯一ID
    merchantUserId   string 

}

func NewAlibabaAilabUserAuthorizedQueryRequest() *AlibabaAilabUserAuthorizedQueryRequest{
    return &AlibabaAilabUserAuthorizedQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabUserAuthorizedQueryRequest) GetApiMethodName() string {
    return "alibaba.ailab.user.authorized.query"
}

func (r AlibabaAilabUserAuthorizedQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabUserAuthorizedQueryRequest) SetSchemaKey(schemaKey string) error {
    r.schemaKey = schemaKey
    r.Set("schema_key", schemaKey)
    return nil
}

func (r AlibabaAilabUserAuthorizedQueryRequest) GetSchemaKey() string {
    return r.schemaKey
}

func (r *AlibabaAilabUserAuthorizedQueryRequest) SetMerchantUserId(merchantUserId string) error {
    r.merchantUserId = merchantUserId
    r.Set("merchant_user_id", merchantUserId)
    return nil
}

func (r AlibabaAilabUserAuthorizedQueryRequest) GetMerchantUserId() string {
    return r.merchantUserId
}

