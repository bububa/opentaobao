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
type AlitripMerchantGalaxyOfferQueryRequest struct {
    model.Params
    // 租户身份信息
    tenantKey   string
    // offer活动ID
    offerIds   string
    // 渠道来源
    offerChannel   string
}

// 初始化AlitripMerchantGalaxyOfferQueryRequest对象
func NewAlitripMerchantGalaxyOfferQueryRequest() *AlitripMerchantGalaxyOfferQueryRequest{
    return &AlitripMerchantGalaxyOfferQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyOfferQueryRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.offer.query"
}

// IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyOfferQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TenantKey Setter
// 租户身份信息
func (r *AlitripMerchantGalaxyOfferQueryRequest) SetTenantKey(tenantKey string) error {
    r.tenantKey = tenantKey
    r.Set("tenant_key", tenantKey)
    return nil
}

// TenantKey Getter
func (r AlitripMerchantGalaxyOfferQueryRequest) GetTenantKey() string {
    return r.tenantKey
}
// OfferIds Setter
// offer活动ID
func (r *AlitripMerchantGalaxyOfferQueryRequest) SetOfferIds(offerIds string) error {
    r.offerIds = offerIds
    r.Set("offer_ids", offerIds)
    return nil
}

// OfferIds Getter
func (r AlitripMerchantGalaxyOfferQueryRequest) GetOfferIds() string {
    return r.offerIds
}
// OfferChannel Setter
// 渠道来源
func (r *AlitripMerchantGalaxyOfferQueryRequest) SetOfferChannel(offerChannel string) error {
    r.offerChannel = offerChannel
    r.Set("offer_channel", offerChannel)
    return nil
}

// OfferChannel Getter
func (r AlitripMerchantGalaxyOfferQueryRequest) GetOfferChannel() string {
    return r.offerChannel
}
