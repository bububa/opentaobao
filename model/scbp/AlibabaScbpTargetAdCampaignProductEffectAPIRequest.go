package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbptargetadcampaignproducteffectAPIRequest 定向推广-获取计划中产品推广效果 API请求
// alibaba.scbp.target.ad.campaign.product.effect
//
// 定向推广-获取计划中产品推广效果
type AlibabascbptargetadcampaignproducteffectAPIRequest struct {
	model.Params
	// TopP4pQuickEffectQuery
	_topP4pQuickEffectQuery *TopP4pQuickEffectQuery
}

// NewAlibabascbptargetadcampaignproducteffectRequest 初始化AlibabascbptargetadcampaignproducteffectAPIRequest对象
func NewAlibabascbptargetadcampaignproducteffectRequest() *AlibabascbptargetadcampaignproducteffectAPIRequest {
	return &AlibabascbptargetadcampaignproducteffectAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbptargetadcampaignproducteffectAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.target.ad.campaign.product.effect"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbptargetadcampaignproducteffectAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbptargetadcampaignproducteffectAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopP4pQuickEffectQuery is TopP4pQuickEffectQuery Setter
// TopP4pQuickEffectQuery
func (r *AlibabascbptargetadcampaignproducteffectAPIRequest) SetTopP4pQuickEffectQuery(_topP4pQuickEffectQuery *TopP4pQuickEffectQuery) error {
	r._topP4pQuickEffectQuery = _topP4pQuickEffectQuery
	r.Set("top_p4p_quick_effect_query", _topP4pQuickEffectQuery)
	return nil
}

// GetTopP4pQuickEffectQuery TopP4pQuickEffectQuery Getter
func (r AlibabascbptargetadcampaignproducteffectAPIRequest) GetTopP4pQuickEffectQuery() *TopP4pQuickEffectQuery {
	return r._topP4pQuickEffectQuery
}
