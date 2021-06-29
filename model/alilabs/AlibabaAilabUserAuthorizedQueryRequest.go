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
type AlibabaAilabUserAuthorizedQueryRequest struct {
    model.Params
    // 开放平台申请的schema
    _schemaKey   string
    // 三方用户的唯一ID
    _merchantUserId   string
}

// 初始化AlibabaAilabUserAuthorizedQueryRequest对象
func NewAlibabaAilabUserAuthorizedQueryRequest() *AlibabaAilabUserAuthorizedQueryRequest{
    return &AlibabaAilabUserAuthorizedQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabUserAuthorizedQueryRequest) GetApiMethodName() string {
    return "alibaba.ailab.user.authorized.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabUserAuthorizedQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SchemaKey Setter
// 开放平台申请的schema
func (r *AlibabaAilabUserAuthorizedQueryRequest) SetSchemaKey(_schemaKey string) error {
    r._schemaKey = _schemaKey
    r.Set("schema_key", _schemaKey)
    return nil
}

// SchemaKey Getter
func (r AlibabaAilabUserAuthorizedQueryRequest) GetSchemaKey() string {
    return r._schemaKey
}
// MerchantUserId Setter
// 三方用户的唯一ID
func (r *AlibabaAilabUserAuthorizedQueryRequest) SetMerchantUserId(_merchantUserId string) error {
    r._merchantUserId = _merchantUserId
    r.Set("merchant_user_id", _merchantUserId)
    return nil
}

// MerchantUserId Getter
func (r AlibabaAilabUserAuthorizedQueryRequest) GetMerchantUserId() string {
    return r._merchantUserId
}
