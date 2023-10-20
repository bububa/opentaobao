package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbptargetadplanupdatetagsAPIRequest 定向推广 按照id操作推广计划的定向溢价标签，包括新增，删除和更新 API请求
// alibaba.scbp.target.ad.plan.update.tags
//
// 定向推广 按照id操作推广计划的定向溢价标签，包括新增，删除和更新
type AlibabascbptargetadplanupdatetagsAPIRequest struct {
	model.Params
	// 系统生成
	_paramTopP4pModifyQuickCampaignTagDTO *TopP4pModifyQuickCampaignTagDto
}

// NewAlibabascbptargetadplanupdatetagsRequest 初始化AlibabascbptargetadplanupdatetagsAPIRequest对象
func NewAlibabascbptargetadplanupdatetagsRequest() *AlibabascbptargetadplanupdatetagsAPIRequest {
	return &AlibabascbptargetadplanupdatetagsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbptargetadplanupdatetagsAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.target.ad.plan.update.tags"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbptargetadplanupdatetagsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbptargetadplanupdatetagsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamTopP4pModifyQuickCampaignTagDTO is ParamTopP4pModifyQuickCampaignTagDTO Setter
// 系统生成
func (r *AlibabascbptargetadplanupdatetagsAPIRequest) SetParamTopP4pModifyQuickCampaignTagDTO(_paramTopP4pModifyQuickCampaignTagDTO *TopP4pModifyQuickCampaignTagDto) error {
	r._paramTopP4pModifyQuickCampaignTagDTO = _paramTopP4pModifyQuickCampaignTagDTO
	r.Set("param_top_p4p_modify_quick_campaign_tag_d_t_o", _paramTopP4pModifyQuickCampaignTagDTO)
	return nil
}

// GetParamTopP4pModifyQuickCampaignTagDTO ParamTopP4pModifyQuickCampaignTagDTO Getter
func (r AlibabascbptargetadplanupdatetagsAPIRequest) GetParamTopP4pModifyQuickCampaignTagDTO() *TopP4pModifyQuickCampaignTagDto {
	return r._paramTopP4pModifyQuickCampaignTagDTO
}
