package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbptargetadplanlistAPIRequest 定向推广-查询定向推广计划列表并返回计划基础信息 API请求
// alibaba.scbp.target.ad.plan.list
//
// 定向推广-查询定向推广计划列表并返回计划基础信息
type AlibabascbptargetadplanlistAPIRequest struct {
	model.Params
	// TopP4pQuickCampaignQuery
	_topP4pQuickCampaignQuery *TopP4pQuickCampaignQueryDto
}

// NewAlibabascbptargetadplanlistRequest 初始化AlibabascbptargetadplanlistAPIRequest对象
func NewAlibabascbptargetadplanlistRequest() *AlibabascbptargetadplanlistAPIRequest {
	return &AlibabascbptargetadplanlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbptargetadplanlistAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.target.ad.plan.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbptargetadplanlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbptargetadplanlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopP4pQuickCampaignQuery is TopP4pQuickCampaignQuery Setter
// TopP4pQuickCampaignQuery
func (r *AlibabascbptargetadplanlistAPIRequest) SetTopP4pQuickCampaignQuery(_topP4pQuickCampaignQuery *TopP4pQuickCampaignQueryDto) error {
	r._topP4pQuickCampaignQuery = _topP4pQuickCampaignQuery
	r.Set("top_p4p_quick_campaign_query", _topP4pQuickCampaignQuery)
	return nil
}

// GetTopP4pQuickCampaignQuery TopP4pQuickCampaignQuery Getter
func (r AlibabascbptargetadplanlistAPIRequest) GetTopP4pQuickCampaignQuery() *TopP4pQuickCampaignQueryDto {
	return r._topP4pQuickCampaignQuery
}
