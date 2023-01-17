package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabUserAuthorizedQueryAPIRequest 查询授权状态接口 API请求
// alibaba.ailab.user.authorized.query
//
// 查询三方用户授权状态
type AlibabaAilabUserAuthorizedQueryAPIRequest struct {
	model.Params
	// 开放平台申请的schema
	_schemaKey string
	// 三方用户的唯一ID
	_merchantUserId string
}

// NewAlibabaAilabUserAuthorizedQueryRequest 初始化AlibabaAilabUserAuthorizedQueryAPIRequest对象
func NewAlibabaAilabUserAuthorizedQueryRequest() *AlibabaAilabUserAuthorizedQueryAPIRequest {
	return &AlibabaAilabUserAuthorizedQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabUserAuthorizedQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ailab.user.authorized.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabUserAuthorizedQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabUserAuthorizedQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchemaKey is SchemaKey Setter
// 开放平台申请的schema
func (r *AlibabaAilabUserAuthorizedQueryAPIRequest) SetSchemaKey(_schemaKey string) error {
	r._schemaKey = _schemaKey
	r.Set("schema_key", _schemaKey)
	return nil
}

// GetSchemaKey SchemaKey Getter
func (r AlibabaAilabUserAuthorizedQueryAPIRequest) GetSchemaKey() string {
	return r._schemaKey
}

// SetMerchantUserId is MerchantUserId Setter
// 三方用户的唯一ID
func (r *AlibabaAilabUserAuthorizedQueryAPIRequest) SetMerchantUserId(_merchantUserId string) error {
	r._merchantUserId = _merchantUserId
	r.Set("merchant_user_id", _merchantUserId)
	return nil
}

// GetMerchantUserId MerchantUserId Getter
func (r AlibabaAilabUserAuthorizedQueryAPIRequest) GetMerchantUserId() string {
	return r._merchantUserId
}
