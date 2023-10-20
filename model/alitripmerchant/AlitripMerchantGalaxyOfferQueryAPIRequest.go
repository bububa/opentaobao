package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyOfferQueryAPIRequest 星河-offer查询 API请求
// alitrip.merchant.galaxy.offer.query
//
// 根据offer的ID查询offer信息
type AlitripMerchantGalaxyOfferQueryAPIRequest struct {
	model.Params
	// 租户身份信息
	_tenantKey string
	// offer活动ID
	_offerIds string
	// 渠道来源
	_offerChannel string
}

// NewAlitripMerchantGalaxyOfferQueryRequest 初始化AlitripMerchantGalaxyOfferQueryAPIRequest对象
func NewAlitripMerchantGalaxyOfferQueryRequest() *AlitripMerchantGalaxyOfferQueryAPIRequest {
	return &AlitripMerchantGalaxyOfferQueryAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyOfferQueryAPIRequest) Reset() {
	r._tenantKey = ""
	r._offerIds = ""
	r._offerChannel = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyOfferQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.offer.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyOfferQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyOfferQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户身份信息
func (r *AlitripMerchantGalaxyOfferQueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyOfferQueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetOfferIds is OfferIds Setter
// offer活动ID
func (r *AlitripMerchantGalaxyOfferQueryAPIRequest) SetOfferIds(_offerIds string) error {
	r._offerIds = _offerIds
	r.Set("offer_ids", _offerIds)
	return nil
}

// GetOfferIds OfferIds Getter
func (r AlitripMerchantGalaxyOfferQueryAPIRequest) GetOfferIds() string {
	return r._offerIds
}

// SetOfferChannel is OfferChannel Setter
// 渠道来源
func (r *AlitripMerchantGalaxyOfferQueryAPIRequest) SetOfferChannel(_offerChannel string) error {
	r._offerChannel = _offerChannel
	r.Set("offer_channel", _offerChannel)
	return nil
}

// GetOfferChannel OfferChannel Getter
func (r AlitripMerchantGalaxyOfferQueryAPIRequest) GetOfferChannel() string {
	return r._offerChannel
}

var poolAlitripMerchantGalaxyOfferQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyOfferQueryRequest()
	},
}

// GetAlitripMerchantGalaxyOfferQueryRequest 从 sync.Pool 获取 AlitripMerchantGalaxyOfferQueryAPIRequest
func GetAlitripMerchantGalaxyOfferQueryAPIRequest() *AlitripMerchantGalaxyOfferQueryAPIRequest {
	return poolAlitripMerchantGalaxyOfferQueryAPIRequest.Get().(*AlitripMerchantGalaxyOfferQueryAPIRequest)
}

// ReleaseAlitripMerchantGalaxyOfferQueryAPIRequest 将 AlitripMerchantGalaxyOfferQueryAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyOfferQueryAPIRequest(v *AlitripMerchantGalaxyOfferQueryAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyOfferQueryAPIRequest.Put(v)
}
