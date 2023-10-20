package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTargetAdPlanAddAPIRequest 定向推广-新建计划 API请求
// alibaba.scbp.target.ad.plan.add
//
// 定向推广-新建单条计划
type AlibabaScbpTargetAdPlanAddAPIRequest struct {
	model.Params
	// 定向推广基础信息
	_topP4pBasicQuickCampaign *BasicQuickCampaign
}

// NewAlibabaScbpTargetAdPlanAddRequest 初始化AlibabaScbpTargetAdPlanAddAPIRequest对象
func NewAlibabaScbpTargetAdPlanAddRequest() *AlibabaScbpTargetAdPlanAddAPIRequest {
	return &AlibabaScbpTargetAdPlanAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpTargetAdPlanAddAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.target.ad.plan.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpTargetAdPlanAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpTargetAdPlanAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopP4pBasicQuickCampaign is TopP4pBasicQuickCampaign Setter
// 定向推广基础信息
func (r *AlibabaScbpTargetAdPlanAddAPIRequest) SetTopP4pBasicQuickCampaign(_topP4pBasicQuickCampaign *BasicQuickCampaign) error {
	r._topP4pBasicQuickCampaign = _topP4pBasicQuickCampaign
	r.Set("top_p4p_basic_quick_campaign", _topP4pBasicQuickCampaign)
	return nil
}

// GetTopP4pBasicQuickCampaign TopP4pBasicQuickCampaign Getter
func (r AlibabaScbpTargetAdPlanAddAPIRequest) GetTopP4pBasicQuickCampaign() *BasicQuickCampaign {
	return r._topP4pBasicQuickCampaign
}
