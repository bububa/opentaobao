package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIRequest 德比无限次券核销通知接口 API请求
// alitrip.merchant.galaxy.derby.voucher.card.unlimited.change.callback
//
// 德比无限次券核销通知接口
type AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIRequest struct {
	model.Params
	// 1
	_tenantKey string
	// 方法标识
	_derbyMethod string
	// 德比回调json信息
	_derbyBody string
}

// NewAlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackRequest 初始化AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIRequest对象
func NewAlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackRequest() *AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIRequest {
	return &AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIRequest) Reset() {
	r._tenantKey = ""
	r._derbyMethod = ""
	r._derbyBody = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.voucher.card.unlimited.change.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 1
func (r *AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetDerbyMethod is DerbyMethod Setter
// 方法标识
func (r *AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIRequest) SetDerbyMethod(_derbyMethod string) error {
	r._derbyMethod = _derbyMethod
	r.Set("derby_method", _derbyMethod)
	return nil
}

// GetDerbyMethod DerbyMethod Getter
func (r AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIRequest) GetDerbyMethod() string {
	return r._derbyMethod
}

// SetDerbyBody is DerbyBody Setter
// 德比回调json信息
func (r *AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIRequest) SetDerbyBody(_derbyBody string) error {
	r._derbyBody = _derbyBody
	r.Set("derby_body", _derbyBody)
	return nil
}

// GetDerbyBody DerbyBody Getter
func (r AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIRequest) GetDerbyBody() string {
	return r._derbyBody
}

var poolAlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackRequest()
	},
}

// GetAlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackRequest 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIRequest
func GetAlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIRequest() *AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIRequest {
	return poolAlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIRequest.Get().(*AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIRequest)
}

// ReleaseAlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIRequest 将 AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIRequest(v *AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIRequest.Put(v)
}
