package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTargetAdPlanUpdateAPIRequest 定向推广-更新推广计划的基础信息 API请求
// alibaba.scbp.target.ad.plan.update
//
// 定向推广-更新推广计划的基础信息
type AlibabaScbpTargetAdPlanUpdateAPIRequest struct {
	model.Params
	// TopP4pBasicQuickCampaign
	_topP4pBasicQuickCampaign *TopP4pBasicQuickCampaign
}

// NewAlibabaScbpTargetAdPlanUpdateRequest 初始化AlibabaScbpTargetAdPlanUpdateAPIRequest对象
func NewAlibabaScbpTargetAdPlanUpdateRequest() *AlibabaScbpTargetAdPlanUpdateAPIRequest {
	return &AlibabaScbpTargetAdPlanUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpTargetAdPlanUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.target.ad.plan.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpTargetAdPlanUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TopP4pBasicQuickCampaign Setter
// TopP4pBasicQuickCampaign
func (r *AlibabaScbpTargetAdPlanUpdateAPIRequest) SetTopP4pBasicQuickCampaign(_topP4pBasicQuickCampaign *TopP4pBasicQuickCampaign) error {
	r._topP4pBasicQuickCampaign = _topP4pBasicQuickCampaign
	r.Set("top_p4p_basic_quick_campaign", _topP4pBasicQuickCampaign)
	return nil
}

// Get TopP4pBasicQuickCampaign Getter
func (r AlibabaScbpTargetAdPlanUpdateAPIRequest) GetTopP4pBasicQuickCampaign() *TopP4pBasicQuickCampaign {
	return r._topP4pBasicQuickCampaign
}
