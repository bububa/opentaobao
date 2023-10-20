package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIRequest 生成臻享卡德比分销二维码 API请求
// alitrip.merchant.galaxy.derby.member.generate.seller.qrcode
//
// 生成臻享卡德比分销二维码
type AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIRequest struct {
	model.Params
	// tenant_key
	_tenantKey string
	// token
	_token string
	// 用户code
	_code string
}

// NewAlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeRequest 初始化AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIRequest对象
func NewAlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeRequest() *AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIRequest {
	return &AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r._code = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.generate.seller.qrcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// tenant_key
func (r *AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// token
func (r *AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIRequest) GetToken() string {
	return r._token
}

// SetCode is Code Setter
// 用户code
func (r *AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIRequest) GetCode() string {
	return r._code
}

var poolAlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeRequest()
	},
}

// GetAlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeRequest 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIRequest
func GetAlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIRequest() *AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIRequest {
	return poolAlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIRequest.Get().(*AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIRequest)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIRequest 将 AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIRequest(v *AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIRequest.Put(v)
}
