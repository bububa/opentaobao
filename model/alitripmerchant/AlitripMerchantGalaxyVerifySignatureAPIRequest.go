package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyVerifySignatureAPIRequest 微信账号生物认证 API请求
// alitrip.merchant.galaxy.verify.signature
//
// 在用户注册等场景下，如果账号风险等级过高，需要进行生物认证
type AlitripMerchantGalaxyVerifySignatureAPIRequest struct {
	model.Params
	// 租户id
	_tenantKey string
	// 用户token
	_token string
	// 微信认证参数
	_jsonString string
	// 微信认证参数
	_jsonSignature string
}

// NewAlitripMerchantGalaxyVerifySignatureRequest 初始化AlitripMerchantGalaxyVerifySignatureAPIRequest对象
func NewAlitripMerchantGalaxyVerifySignatureRequest() *AlitripMerchantGalaxyVerifySignatureAPIRequest {
	return &AlitripMerchantGalaxyVerifySignatureAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyVerifySignatureAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r._jsonString = ""
	r._jsonSignature = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyVerifySignatureAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.verify.signature"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyVerifySignatureAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyVerifySignatureAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户id
func (r *AlitripMerchantGalaxyVerifySignatureAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyVerifySignatureAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户token
func (r *AlitripMerchantGalaxyVerifySignatureAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyVerifySignatureAPIRequest) GetToken() string {
	return r._token
}

// SetJsonString is JsonString Setter
// 微信认证参数
func (r *AlitripMerchantGalaxyVerifySignatureAPIRequest) SetJsonString(_jsonString string) error {
	r._jsonString = _jsonString
	r.Set("json_string", _jsonString)
	return nil
}

// GetJsonString JsonString Getter
func (r AlitripMerchantGalaxyVerifySignatureAPIRequest) GetJsonString() string {
	return r._jsonString
}

// SetJsonSignature is JsonSignature Setter
// 微信认证参数
func (r *AlitripMerchantGalaxyVerifySignatureAPIRequest) SetJsonSignature(_jsonSignature string) error {
	r._jsonSignature = _jsonSignature
	r.Set("json_signature", _jsonSignature)
	return nil
}

// GetJsonSignature JsonSignature Getter
func (r AlitripMerchantGalaxyVerifySignatureAPIRequest) GetJsonSignature() string {
	return r._jsonSignature
}

var poolAlitripMerchantGalaxyVerifySignatureAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyVerifySignatureRequest()
	},
}

// GetAlitripMerchantGalaxyVerifySignatureRequest 从 sync.Pool 获取 AlitripMerchantGalaxyVerifySignatureAPIRequest
func GetAlitripMerchantGalaxyVerifySignatureAPIRequest() *AlitripMerchantGalaxyVerifySignatureAPIRequest {
	return poolAlitripMerchantGalaxyVerifySignatureAPIRequest.Get().(*AlitripMerchantGalaxyVerifySignatureAPIRequest)
}

// ReleaseAlitripMerchantGalaxyVerifySignatureAPIRequest 将 AlitripMerchantGalaxyVerifySignatureAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyVerifySignatureAPIRequest(v *AlitripMerchantGalaxyVerifySignatureAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyVerifySignatureAPIRequest.Put(v)
}
