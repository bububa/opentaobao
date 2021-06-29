package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-获取计划中产品推广效果 API请求
alibaba.scbp.target.ad.campaign.product.effect

定向推广-获取计划中产品推广效果
*/
type AlibabaScbpTargetAdCampaignProductEffectRequest struct {
    model.Params
    // TopP4pQuickEffectQuery
    _topP4pQuickEffectQuery   *TopP4pQuickEffectQuery
}

// 初始化AlibabaScbpTargetAdCampaignProductEffectRequest对象
func NewAlibabaScbpTargetAdCampaignProductEffectRequest() *AlibabaScbpTargetAdCampaignProductEffectRequest{
    return &AlibabaScbpTargetAdCampaignProductEffectRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpTargetAdCampaignProductEffectRequest) GetApiMethodName() string {
    return "alibaba.scbp.target.ad.campaign.product.effect"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpTargetAdCampaignProductEffectRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TopP4pQuickEffectQuery Setter
// TopP4pQuickEffectQuery
func (r *AlibabaScbpTargetAdCampaignProductEffectRequest) SetTopP4pQuickEffectQuery(_topP4pQuickEffectQuery *TopP4pQuickEffectQuery) error {
    r._topP4pQuickEffectQuery = _topP4pQuickEffectQuery
    r.Set("top_p4p_quick_effect_query", _topP4pQuickEffectQuery)
    return nil
}

// TopP4pQuickEffectQuery Getter
func (r AlibabaScbpTargetAdCampaignProductEffectRequest) GetTopP4pQuickEffectQuery() *TopP4pQuickEffectQuery {
    return r._topP4pQuickEffectQuery
}
