package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpTargetAdCampaignProductEffectAPIRequest
定向推广-获取计划中产品推广效果 API请求
alibaba.scbp.target.ad.campaign.product.effect

定向推广-获取计划中产品推广效果 */
type AlibabaScbpTargetAdCampaignProductEffectAPIRequest struct {
	model.Params
	// TopP4pQuickEffectQuery
	_topP4pQuickEffectQuery *TopP4pQuickEffectQuery
}

// NewAlibabaScbpTargetAdCampaignProductEffectRequest 初始化AlibabaScbpTargetAdCampaignProductEffectAPIRequest对象
func NewAlibabaScbpTargetAdCampaignProductEffectRequest() *AlibabaScbpTargetAdCampaignProductEffectAPIRequest {
	return &AlibabaScbpTargetAdCampaignProductEffectAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpTargetAdCampaignProductEffectAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.target.ad.campaign.product.effect"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpTargetAdCampaignProductEffectAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TopP4pQuickEffectQuery Setter
// TopP4pQuickEffectQuery
func (r *AlibabaScbpTargetAdCampaignProductEffectAPIRequest) SetTopP4pQuickEffectQuery(_topP4pQuickEffectQuery *TopP4pQuickEffectQuery) error {
	r._topP4pQuickEffectQuery = _topP4pQuickEffectQuery
	r.Set("top_p4p_quick_effect_query", _topP4pQuickEffectQuery)
	return nil
}

// Get TopP4pQuickEffectQuery Getter
func (r AlibabaScbpTargetAdCampaignProductEffectAPIRequest) GetTopP4pQuickEffectQuery() *TopP4pQuickEffectQuery {
	return r._topP4pQuickEffectQuery
}
