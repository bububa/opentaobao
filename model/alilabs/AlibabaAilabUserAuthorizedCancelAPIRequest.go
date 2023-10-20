package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabuserauthorizedcancelAPIRequest 取消账号授权 API请求
// alibaba.ailab.user.authorized.cancel
//
// 三方用户取消授权给天猫精灵用户
type AlibabaailabuserauthorizedcancelAPIRequest struct {
	model.Params
	// 三方用户的唯一ID
	_merchantUserId string
	// 开放平台申请的schema
	_schemaKey string
}

// NewAlibabaailabuserauthorizedcancelRequest 初始化AlibabaailabuserauthorizedcancelAPIRequest对象
func NewAlibabaailabuserauthorizedcancelRequest() *AlibabaailabuserauthorizedcancelAPIRequest {
	return &AlibabaailabuserauthorizedcancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabuserauthorizedcancelAPIRequest) GetApiMethodName() string {
	return "alibaba.ailab.user.authorized.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabuserauthorizedcancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabuserauthorizedcancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMerchantUserId is MerchantUserId Setter
// 三方用户的唯一ID
func (r *AlibabaailabuserauthorizedcancelAPIRequest) SetMerchantUserId(_merchantUserId string) error {
	r._merchantUserId = _merchantUserId
	r.Set("merchant_user_id", _merchantUserId)
	return nil
}

// GetMerchantUserId MerchantUserId Getter
func (r AlibabaailabuserauthorizedcancelAPIRequest) GetMerchantUserId() string {
	return r._merchantUserId
}

// SetSchemaKey is SchemaKey Setter
// 开放平台申请的schema
func (r *AlibabaailabuserauthorizedcancelAPIRequest) SetSchemaKey(_schemaKey string) error {
	r._schemaKey = _schemaKey
	r.Set("schema_key", _schemaKey)
	return nil
}

// GetSchemaKey SchemaKey Getter
func (r AlibabaailabuserauthorizedcancelAPIRequest) GetSchemaKey() string {
	return r._schemaKey
}
