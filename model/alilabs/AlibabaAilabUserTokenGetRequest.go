package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
三方账号获取 token API请求
alibaba.ailab.user.token.get

inside 设备的三方 app，通过 extId、schema 生成 token
*/
type AlibabaAilabUserTokenGetRequest struct {
    model.Params
    // 三方用户的唯一ID
    _merchantUserId   string
    // 开放平台申请的schema
    _schemaKey   string
    // 用户点击同意授权，则会有授权结果：success/fail，此结果通过 callBackUrl 回调给三方 如果授权账号重复授权给已授权的淘宝账号，幂等返回成功 url 的调用是 表单 post 的方式， request body success example: merchantUserId=xxx&result=success request body fail example: merchantUserId=xxx&result=fail
    _callBackUrl   string
}

// 初始化AlibabaAilabUserTokenGetRequest对象
func NewAlibabaAilabUserTokenGetRequest() *AlibabaAilabUserTokenGetRequest{
    return &AlibabaAilabUserTokenGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabUserTokenGetRequest) GetApiMethodName() string {
    return "alibaba.ailab.user.token.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabUserTokenGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MerchantUserId Setter
// 三方用户的唯一ID
func (r *AlibabaAilabUserTokenGetRequest) SetMerchantUserId(_merchantUserId string) error {
    r._merchantUserId = _merchantUserId
    r.Set("merchant_user_id", _merchantUserId)
    return nil
}

// MerchantUserId Getter
func (r AlibabaAilabUserTokenGetRequest) GetMerchantUserId() string {
    return r._merchantUserId
}
// SchemaKey Setter
// 开放平台申请的schema
func (r *AlibabaAilabUserTokenGetRequest) SetSchemaKey(_schemaKey string) error {
    r._schemaKey = _schemaKey
    r.Set("schema_key", _schemaKey)
    return nil
}

// SchemaKey Getter
func (r AlibabaAilabUserTokenGetRequest) GetSchemaKey() string {
    return r._schemaKey
}
// CallBackUrl Setter
// 用户点击同意授权，则会有授权结果：success/fail，此结果通过 callBackUrl 回调给三方 如果授权账号重复授权给已授权的淘宝账号，幂等返回成功 url 的调用是 表单 post 的方式， request body success example: merchantUserId=xxx&result=success request body fail example: merchantUserId=xxx&result=fail
func (r *AlibabaAilabUserTokenGetRequest) SetCallBackUrl(_callBackUrl string) error {
    r._callBackUrl = _callBackUrl
    r.Set("call_back_url", _callBackUrl)
    return nil
}

// CallBackUrl Getter
func (r AlibabaAilabUserTokenGetRequest) GetCallBackUrl() string {
    return r._callBackUrl
}