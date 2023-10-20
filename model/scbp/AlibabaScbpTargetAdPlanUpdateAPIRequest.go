package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbptargetadplanupdateAPIRequest 定向推广-更新推广计划的基础信息 API请求
// alibaba.scbp.target.ad.plan.update
//
// 定向推广-更新推广计划的基础信息
type AlibabascbptargetadplanupdateAPIRequest struct {
	model.Params
	// TopP4pBasicQuickCampaign
	_topP4pBasicQuickCampaign *TopP4pBasicQuickCampaign
}

// NewAlibabascbptargetadplanupdateRequest 初始化AlibabascbptargetadplanupdateAPIRequest对象
func NewAlibabascbptargetadplanupdateRequest() *AlibabascbptargetadplanupdateAPIRequest {
	return &AlibabascbptargetadplanupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbptargetadplanupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.target.ad.plan.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbptargetadplanupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbptargetadplanupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopP4pBasicQuickCampaign is TopP4pBasicQuickCampaign Setter
// TopP4pBasicQuickCampaign
func (r *AlibabascbptargetadplanupdateAPIRequest) SetTopP4pBasicQuickCampaign(_topP4pBasicQuickCampaign *TopP4pBasicQuickCampaign) error {
	r._topP4pBasicQuickCampaign = _topP4pBasicQuickCampaign
	r.Set("top_p4p_basic_quick_campaign", _topP4pBasicQuickCampaign)
	return nil
}

// GetTopP4pBasicQuickCampaign TopP4pBasicQuickCampaign Getter
func (r AlibabascbptargetadplanupdateAPIRequest) GetTopP4pBasicQuickCampaign() *TopP4pBasicQuickCampaign {
	return r._topP4pBasicQuickCampaign
}
