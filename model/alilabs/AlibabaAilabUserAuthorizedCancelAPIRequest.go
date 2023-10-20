package alilabs

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabUserAuthorizedCancelAPIRequest 取消账号授权 API请求
// alibaba.ailab.user.authorized.cancel
//
// 三方用户取消授权给天猫精灵用户
type AlibabaAilabUserAuthorizedCancelAPIRequest struct {
	model.Params
	// 三方用户的唯一ID
	_merchantUserId string
	// 开放平台申请的schema
	_schemaKey string
}

// NewAlibabaAilabUserAuthorizedCancelRequest 初始化AlibabaAilabUserAuthorizedCancelAPIRequest对象
func NewAlibabaAilabUserAuthorizedCancelRequest() *AlibabaAilabUserAuthorizedCancelAPIRequest {
	return &AlibabaAilabUserAuthorizedCancelAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAilabUserAuthorizedCancelAPIRequest) Reset() {
	r._merchantUserId = ""
	r._schemaKey = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabUserAuthorizedCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.ailab.user.authorized.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabUserAuthorizedCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabUserAuthorizedCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMerchantUserId is MerchantUserId Setter
// 三方用户的唯一ID
func (r *AlibabaAilabUserAuthorizedCancelAPIRequest) SetMerchantUserId(_merchantUserId string) error {
	r._merchantUserId = _merchantUserId
	r.Set("merchant_user_id", _merchantUserId)
	return nil
}

// GetMerchantUserId MerchantUserId Getter
func (r AlibabaAilabUserAuthorizedCancelAPIRequest) GetMerchantUserId() string {
	return r._merchantUserId
}

// SetSchemaKey is SchemaKey Setter
// 开放平台申请的schema
func (r *AlibabaAilabUserAuthorizedCancelAPIRequest) SetSchemaKey(_schemaKey string) error {
	r._schemaKey = _schemaKey
	r.Set("schema_key", _schemaKey)
	return nil
}

// GetSchemaKey SchemaKey Getter
func (r AlibabaAilabUserAuthorizedCancelAPIRequest) GetSchemaKey() string {
	return r._schemaKey
}

var poolAlibabaAilabUserAuthorizedCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAilabUserAuthorizedCancelRequest()
	},
}

// GetAlibabaAilabUserAuthorizedCancelRequest 从 sync.Pool 获取 AlibabaAilabUserAuthorizedCancelAPIRequest
func GetAlibabaAilabUserAuthorizedCancelAPIRequest() *AlibabaAilabUserAuthorizedCancelAPIRequest {
	return poolAlibabaAilabUserAuthorizedCancelAPIRequest.Get().(*AlibabaAilabUserAuthorizedCancelAPIRequest)
}

// ReleaseAlibabaAilabUserAuthorizedCancelAPIRequest 将 AlibabaAilabUserAuthorizedCancelAPIRequest 放入 sync.Pool
func ReleaseAlibabaAilabUserAuthorizedCancelAPIRequest(v *AlibabaAilabUserAuthorizedCancelAPIRequest) {
	v.Reset()
	poolAlibabaAilabUserAuthorizedCancelAPIRequest.Put(v)
}
