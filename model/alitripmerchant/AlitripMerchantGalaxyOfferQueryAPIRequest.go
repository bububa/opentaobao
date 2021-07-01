package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-offer查询 API请求
alitrip.merchant.galaxy.offer.query

根据offer的ID查询offer信息
*/
type AlitripMerchantGalaxyOfferQueryAPIRequest struct {
    model.Params
    // 租户身份信息
    _tenantKey   string
    // offer活动ID
    _offerIds   string
    // 渠道来源
    _offerChannel   string
}

// 初始化AlitripMerchantGalaxyOfferQueryAPIRequest对象
func NewAlitripMerchantGalaxyOfferQueryRequest() *AlitripMerchantGalaxyOfferQueryAPIRequest{
    return &AlitripMerchantGalaxyOfferQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyOfferQueryAPIRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.offer.query"
}

// IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyOfferQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TenantKey Setter
// 租户身份信息
func (r *AlitripMerchantGalaxyOfferQueryAPIRequest) SetTenantKey(_tenantKey string) error {
    r._tenantKey = _tenantKey
    r.Set("tenant_key", _tenantKey)
    return nil
}

// TenantKey Getter
func (r AlitripMerchantGalaxyOfferQueryAPIRequest) GetTenantKey() string {
    return r._tenantKey
}
// OfferIds Setter
// offer活动ID
func (r *AlitripMerchantGalaxyOfferQueryAPIRequest) SetOfferIds(_offerIds string) error {
    r._offerIds = _offerIds
    r.Set("offer_ids", _offerIds)
    return nil
}

// OfferIds Getter
func (r AlitripMerchantGalaxyOfferQueryAPIRequest) GetOfferIds() string {
    return r._offerIds
}
// OfferChannel Setter
// 渠道来源
func (r *AlitripMerchantGalaxyOfferQueryAPIRequest) SetOfferChannel(_offerChannel string) error {
    r._offerChannel = _offerChannel
    r.Set("offer_channel", _offerChannel)
    return nil
}

// OfferChannel Getter
func (r AlitripMerchantGalaxyOfferQueryAPIRequest) GetOfferChannel() string {
    return r._offerChannel
}
