package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyofferqueryAPIRequest 星河-offer查询 API请求
// alitrip.merchant.galaxy.offer.query
//
// 根据offer的ID查询offer信息
type AlitripmerchantgalaxyofferqueryAPIRequest struct {
	model.Params
	// 租户身份信息
	_tenantKey string
	// offer活动ID
	_offerIds string
	// 渠道来源
	_offerChannel string
}

// NewAlitripmerchantgalaxyofferqueryRequest 初始化AlitripmerchantgalaxyofferqueryAPIRequest对象
func NewAlitripmerchantgalaxyofferqueryRequest() *AlitripmerchantgalaxyofferqueryAPIRequest {
	return &AlitripmerchantgalaxyofferqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyofferqueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.offer.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyofferqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyofferqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户身份信息
func (r *AlitripmerchantgalaxyofferqueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyofferqueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetOfferIds is OfferIds Setter
// offer活动ID
func (r *AlitripmerchantgalaxyofferqueryAPIRequest) SetOfferIds(_offerIds string) error {
	r._offerIds = _offerIds
	r.Set("offer_ids", _offerIds)
	return nil
}

// GetOfferIds OfferIds Getter
func (r AlitripmerchantgalaxyofferqueryAPIRequest) GetOfferIds() string {
	return r._offerIds
}

// SetOfferChannel is OfferChannel Setter
// 渠道来源
func (r *AlitripmerchantgalaxyofferqueryAPIRequest) SetOfferChannel(_offerChannel string) error {
	r._offerChannel = _offerChannel
	r.Set("offer_channel", _offerChannel)
	return nil
}

// GetOfferChannel OfferChannel Getter
func (r AlitripmerchantgalaxyofferqueryAPIRequest) GetOfferChannel() string {
	return r._offerChannel
}
