package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-offer查询 APIRequest
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

func NewAlitripMerchantGalaxyOfferQueryRequest() *AlitripMerchantGalaxyOfferQueryRequest{
    return &AlitripMerchantGalaxyOfferQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripMerchantGalaxyOfferQueryRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.offer.query"
}

func (r AlitripMerchantGalaxyOfferQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripMerchantGalaxyOfferQueryRequest) SetTenantKey(tenantKey string) error {
    r.tenantKey = tenantKey
    r.Set("tenant_key", tenantKey)
    return nil
}

func (r AlitripMerchantGalaxyOfferQueryRequest) GetTenantKey() string {
    return r.tenantKey
}

func (r *AlitripMerchantGalaxyOfferQueryRequest) SetOfferIds(offerIds string) error {
    r.offerIds = offerIds
    r.Set("offer_ids", offerIds)
    return nil
}

func (r AlitripMerchantGalaxyOfferQueryRequest) GetOfferIds() string {
    return r.offerIds
}

func (r *AlitripMerchantGalaxyOfferQueryRequest) SetOfferChannel(offerChannel string) error {
    r.offerChannel = offerChannel
    r.Set("offer_channel", offerChannel)
    return nil
}

func (r AlitripMerchantGalaxyOfferQueryRequest) GetOfferChannel() string {
    return r.offerChannel
}

