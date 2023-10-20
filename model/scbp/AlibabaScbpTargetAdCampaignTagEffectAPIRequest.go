package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbptargetadcampaigntageffectAPIRequest 定向推广-获取推广计划定向效果数据 API请求
// alibaba.scbp.target.ad.campaign.tag.effect
//
// 定向推广-获取推广计划定向效果数据
type AlibabascbptargetadcampaigntageffectAPIRequest struct {
	model.Params
	// 效果数据
	_topP4pQuickEffectQuery *TopP4pQuickEffectQuery
}

// NewAlibabascbptargetadcampaigntageffectRequest 初始化AlibabascbptargetadcampaigntageffectAPIRequest对象
func NewAlibabascbptargetadcampaigntageffectRequest() *AlibabascbptargetadcampaigntageffectAPIRequest {
	return &AlibabascbptargetadcampaigntageffectAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbptargetadcampaigntageffectAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.target.ad.campaign.tag.effect"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbptargetadcampaigntageffectAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbptargetadcampaigntageffectAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopP4pQuickEffectQuery is TopP4pQuickEffectQuery Setter
// 效果数据
func (r *AlibabascbptargetadcampaigntageffectAPIRequest) SetTopP4pQuickEffectQuery(_topP4pQuickEffectQuery *TopP4pQuickEffectQuery) error {
	r._topP4pQuickEffectQuery = _topP4pQuickEffectQuery
	r.Set("top_p4p_quick_effect_query", _topP4pQuickEffectQuery)
	return nil
}

// GetTopP4pQuickEffectQuery TopP4pQuickEffectQuery Getter
func (r AlibabascbptargetadcampaigntageffectAPIRequest) GetTopP4pQuickEffectQuery() *TopP4pQuickEffectQuery {
	return r._topP4pQuickEffectQuery
}
