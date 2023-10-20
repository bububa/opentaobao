package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbptargetadplanoperationAPIRequest 定向推广-计划开启/暂停/删除 API请求
// alibaba.scbp.target.ad.plan.operation
//
// 定向推广-计划开启/暂停/删除
type AlibabascbptargetadplanoperationAPIRequest struct {
	model.Params
	// TopP4pModifyQuickCampaignDTO
	_topP4pModifyQuickCampaignDTO *TopP4pModifyQuickCampaignDto
}

// NewAlibabascbptargetadplanoperationRequest 初始化AlibabascbptargetadplanoperationAPIRequest对象
func NewAlibabascbptargetadplanoperationRequest() *AlibabascbptargetadplanoperationAPIRequest {
	return &AlibabascbptargetadplanoperationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbptargetadplanoperationAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.target.ad.plan.operation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbptargetadplanoperationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbptargetadplanoperationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopP4pModifyQuickCampaignDTO is TopP4pModifyQuickCampaignDTO Setter
// TopP4pModifyQuickCampaignDTO
func (r *AlibabascbptargetadplanoperationAPIRequest) SetTopP4pModifyQuickCampaignDTO(_topP4pModifyQuickCampaignDTO *TopP4pModifyQuickCampaignDto) error {
	r._topP4pModifyQuickCampaignDTO = _topP4pModifyQuickCampaignDTO
	r.Set("top_p4p_modify_quick_campaign_d_t_o", _topP4pModifyQuickCampaignDTO)
	return nil
}

// GetTopP4pModifyQuickCampaignDTO TopP4pModifyQuickCampaignDTO Getter
func (r AlibabascbptargetadplanoperationAPIRequest) GetTopP4pModifyQuickCampaignDTO() *TopP4pModifyQuickCampaignDto {
	return r._topP4pModifyQuickCampaignDTO
}
