package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消账号授权 APIRequest
alibaba.ailab.user.authorized.cancel

三方用户取消授权给天猫精灵用户
*/
type AlibabaAilabUserAuthorizedCancelRequest struct {
    model.Params

    // 三方用户的唯一ID
    merchantUserId   string 

    // 开放平台申请的schema
    schemaKey   string 

}

func NewAlibabaAilabUserAuthorizedCancelRequest() *AlibabaAilabUserAuthorizedCancelRequest{
    return &AlibabaAilabUserAuthorizedCancelRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabUserAuthorizedCancelRequest) GetApiMethodName() string {
    return "alibaba.ailab.user.authorized.cancel"
}

func (r AlibabaAilabUserAuthorizedCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabUserAuthorizedCancelRequest) SetMerchantUserId(merchantUserId string) error {
    r.merchantUserId = merchantUserId
    r.Set("merchant_user_id", merchantUserId)
    return nil
}

func (r AlibabaAilabUserAuthorizedCancelRequest) GetMerchantUserId() string {
    return r.merchantUserId
}

func (r *AlibabaAilabUserAuthorizedCancelRequest) SetSchemaKey(schemaKey string) error {
    r.schemaKey = schemaKey
    r.Set("schema_key", schemaKey)
    return nil
}

func (r AlibabaAilabUserAuthorizedCancelRequest) GetSchemaKey() string {
    return r.schemaKey
}

