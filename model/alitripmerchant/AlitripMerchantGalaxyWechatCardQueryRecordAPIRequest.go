package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyWechatCardQueryRecordAPIRequest 微信会员卡领取记录查询 API请求
// alitrip.merchant.galaxy.wechat.card.query.record
//
// 微信会员卡领取记录查询
type AlitripMerchantGalaxyWechatCardQueryRecordAPIRequest struct {
	model.Params
	// 租户Id
	_tenantKey string
	// 用户登录凭证
	_token string
}

// NewAlitripMerchantGalaxyWechatCardQueryRecordRequest 初始化AlitripMerchantGalaxyWechatCardQueryRecordAPIRequest对象
func NewAlitripMerchantGalaxyWechatCardQueryRecordRequest() *AlitripMerchantGalaxyWechatCardQueryRecordAPIRequest {
	return &AlitripMerchantGalaxyWechatCardQueryRecordAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyWechatCardQueryRecordAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyWechatCardQueryRecordAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.wechat.card.query.record"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyWechatCardQueryRecordAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyWechatCardQueryRecordAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户Id
func (r *AlitripMerchantGalaxyWechatCardQueryRecordAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyWechatCardQueryRecordAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户登录凭证
func (r *AlitripMerchantGalaxyWechatCardQueryRecordAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyWechatCardQueryRecordAPIRequest) GetToken() string {
	return r._token
}

var poolAlitripMerchantGalaxyWechatCardQueryRecordAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyWechatCardQueryRecordRequest()
	},
}

// GetAlitripMerchantGalaxyWechatCardQueryRecordRequest 从 sync.Pool 获取 AlitripMerchantGalaxyWechatCardQueryRecordAPIRequest
func GetAlitripMerchantGalaxyWechatCardQueryRecordAPIRequest() *AlitripMerchantGalaxyWechatCardQueryRecordAPIRequest {
	return poolAlitripMerchantGalaxyWechatCardQueryRecordAPIRequest.Get().(*AlitripMerchantGalaxyWechatCardQueryRecordAPIRequest)
}

// ReleaseAlitripMerchantGalaxyWechatCardQueryRecordAPIRequest 将 AlitripMerchantGalaxyWechatCardQueryRecordAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyWechatCardQueryRecordAPIRequest(v *AlitripMerchantGalaxyWechatCardQueryRecordAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyWechatCardQueryRecordAPIRequest.Put(v)
}
