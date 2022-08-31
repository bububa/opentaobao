package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabUserTokenGetAPIRequest 三方账号获取 token API请求
// alibaba.ailab.user.token.get
//
// inside 设备的三方 app，通过 extId、schema 生成 token
type AlibabaAilabUserTokenGetAPIRequest struct {
	model.Params
	// 三方用户的唯一ID
	_merchantUserId string
	// 开放平台申请的schema
	_schemaKey string
	// 用户点击同意授权，则会有授权结果：success/fail，此结果通过 callBackUrl 回调给三方 如果授权账号重复授权给已授权的淘宝账号，幂等返回成功 url 的调用是 表单 post 的方式， request body success example: merchantUserId=xxx&result=success request body fail example: merchantUserId=xxx&result=fail
	_callBackUrl string
}

// NewAlibabaAilabUserTokenGetRequest 初始化AlibabaAilabUserTokenGetAPIRequest对象
func NewAlibabaAilabUserTokenGetRequest() *AlibabaAilabUserTokenGetAPIRequest {
	return &AlibabaAilabUserTokenGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabUserTokenGetAPIRequest) GetApiMethodName() string {
	return "alibaba.ailab.user.token.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabUserTokenGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetMerchantUserId is MerchantUserId Setter
// 三方用户的唯一ID
func (r *AlibabaAilabUserTokenGetAPIRequest) SetMerchantUserId(_merchantUserId string) error {
	r._merchantUserId = _merchantUserId
	r.Set("merchant_user_id", _merchantUserId)
	return nil
}

// GetMerchantUserId MerchantUserId Getter
func (r AlibabaAilabUserTokenGetAPIRequest) GetMerchantUserId() string {
	return r._merchantUserId
}

// SetSchemaKey is SchemaKey Setter
// 开放平台申请的schema
func (r *AlibabaAilabUserTokenGetAPIRequest) SetSchemaKey(_schemaKey string) error {
	r._schemaKey = _schemaKey
	r.Set("schema_key", _schemaKey)
	return nil
}

// GetSchemaKey SchemaKey Getter
func (r AlibabaAilabUserTokenGetAPIRequest) GetSchemaKey() string {
	return r._schemaKey
}

// SetCallBackUrl is CallBackUrl Setter
// 用户点击同意授权，则会有授权结果：success/fail，此结果通过 callBackUrl 回调给三方 如果授权账号重复授权给已授权的淘宝账号，幂等返回成功 url 的调用是 表单 post 的方式， request body success example: merchantUserId=xxx&amp;result=success request body fail example: merchantUserId=xxx&amp;result=fail
func (r *AlibabaAilabUserTokenGetAPIRequest) SetCallBackUrl(_callBackUrl string) error {
	r._callBackUrl = _callBackUrl
	r.Set("call_back_url", _callBackUrl)
	return nil
}

// GetCallBackUrl CallBackUrl Getter
func (r AlibabaAilabUserTokenGetAPIRequest) GetCallBackUrl() string {
	return r._callBackUrl
}
