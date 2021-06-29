package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-获取推广计划定向效果数据 API请求
alibaba.scbp.target.ad.campaign.tag.effect

定向推广-获取推广计划定向效果数据
*/
type AlibabaScbpTargetAdCampaignTagEffectRequest struct {
    model.Params
    // 效果数据
    _topP4pQuickEffectQuery   *TopP4pQuickEffectQuery
}

// 初始化AlibabaScbpTargetAdCampaignTagEffectRequest对象
func NewAlibabaScbpTargetAdCampaignTagEffectRequest() *AlibabaScbpTargetAdCampaignTagEffectRequest{
    return &AlibabaScbpTargetAdCampaignTagEffectRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpTargetAdCampaignTagEffectRequest) GetApiMethodName() string {
    return "alibaba.scbp.target.ad.campaign.tag.effect"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpTargetAdCampaignTagEffectRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TopP4pQuickEffectQuery Setter
// 效果数据
func (r *AlibabaScbpTargetAdCampaignTagEffectRequest) SetTopP4pQuickEffectQuery(_topP4pQuickEffectQuery *TopP4pQuickEffectQuery) error {
    r._topP4pQuickEffectQuery = _topP4pQuickEffectQuery
    r.Set("top_p4p_quick_effect_query", _topP4pQuickEffectQuery)
    return nil
}

// TopP4pQuickEffectQuery Getter
func (r AlibabaScbpTargetAdCampaignTagEffectRequest) GetTopP4pQuickEffectQuery() *TopP4pQuickEffectQuery {
    return r._topP4pQuickEffectQuery
}
