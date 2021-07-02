package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTargetAdCampaignTagEffectAPIRequest 定向推广-获取推广计划定向效果数据 API请求
// alibaba.scbp.target.ad.campaign.tag.effect
//
// 定向推广-获取推广计划定向效果数据
type AlibabaScbpTargetAdCampaignTagEffectAPIRequest struct {
	model.Params
	// 效果数据
	_topP4pQuickEffectQuery *TopP4pQuickEffectQuery
}

// NewAlibabaScbpTargetAdCampaignTagEffectRequest 初始化AlibabaScbpTargetAdCampaignTagEffectAPIRequest对象
func NewAlibabaScbpTargetAdCampaignTagEffectRequest() *AlibabaScbpTargetAdCampaignTagEffectAPIRequest {
	return &AlibabaScbpTargetAdCampaignTagEffectAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpTargetAdCampaignTagEffectAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.target.ad.campaign.tag.effect"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpTargetAdCampaignTagEffectAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTopP4pQuickEffectQuery is TopP4pQuickEffectQuery Setter
// 效果数据
func (r *AlibabaScbpTargetAdCampaignTagEffectAPIRequest) SetTopP4pQuickEffectQuery(_topP4pQuickEffectQuery *TopP4pQuickEffectQuery) error {
	r._topP4pQuickEffectQuery = _topP4pQuickEffectQuery
	r.Set("top_p4p_quick_effect_query", _topP4pQuickEffectQuery)
	return nil
}

// GetTopP4pQuickEffectQuery TopP4pQuickEffectQuery Getter
func (r AlibabaScbpTargetAdCampaignTagEffectAPIRequest) GetTopP4pQuickEffectQuery() *TopP4pQuickEffectQuery {
	return r._topP4pQuickEffectQuery
}
