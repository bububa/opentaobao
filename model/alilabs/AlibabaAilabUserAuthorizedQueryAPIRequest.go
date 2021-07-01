package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询授权状态接口 API请求
alibaba.ailab.user.authorized.query

查询三方用户授权状态
*/
type AlibabaAilabUserAuthorizedQueryAPIRequest struct {
    model.Params
    // 开放平台申请的schema
    _schemaKey   string
    // 三方用户的唯一ID
    _merchantUserId   string
}

// 初始化AlibabaAilabUserAuthorizedQueryAPIRequest对象
func NewAlibabaAilabUserAuthorizedQueryRequest() *AlibabaAilabUserAuthorizedQueryAPIRequest{
    return &AlibabaAilabUserAuthorizedQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabUserAuthorizedQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.ailab.user.authorized.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabUserAuthorizedQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SchemaKey Setter
// 开放平台申请的schema
func (r *AlibabaAilabUserAuthorizedQueryAPIRequest) SetSchemaKey(_schemaKey string) error {
    r._schemaKey = _schemaKey
    r.Set("schema_key", _schemaKey)
    return nil
}

// SchemaKey Getter
func (r AlibabaAilabUserAuthorizedQueryAPIRequest) GetSchemaKey() string {
    return r._schemaKey
}
// MerchantUserId Setter
// 三方用户的唯一ID
func (r *AlibabaAilabUserAuthorizedQueryAPIRequest) SetMerchantUserId(_merchantUserId string) error {
    r._merchantUserId = _merchantUserId
    r.Set("merchant_user_id", _merchantUserId)
    return nil
}

// MerchantUserId Getter
func (r AlibabaAilabUserAuthorizedQueryAPIRequest) GetMerchantUserId() string {
    return r._merchantUserId
}
