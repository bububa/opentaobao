package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-更新推广计划的基础信息 API请求
alibaba.scbp.target.ad.plan.update

定向推广-更新推广计划的基础信息
*/
type AlibabaScbpTargetAdPlanUpdateRequest struct {
    model.Params
    // TopP4pBasicQuickCampaign
    _topP4pBasicQuickCampaign   *TopP4pBasicQuickCampaign
}

// 初始化AlibabaScbpTargetAdPlanUpdateRequest对象
func NewAlibabaScbpTargetAdPlanUpdateRequest() *AlibabaScbpTargetAdPlanUpdateRequest{
    return &AlibabaScbpTargetAdPlanUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpTargetAdPlanUpdateRequest) GetApiMethodName() string {
    return "alibaba.scbp.target.ad.plan.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpTargetAdPlanUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TopP4pBasicQuickCampaign Setter
// TopP4pBasicQuickCampaign
func (r *AlibabaScbpTargetAdPlanUpdateRequest) SetTopP4pBasicQuickCampaign(_topP4pBasicQuickCampaign *TopP4pBasicQuickCampaign) error {
    r._topP4pBasicQuickCampaign = _topP4pBasicQuickCampaign
    r.Set("top_p4p_basic_quick_campaign", _topP4pBasicQuickCampaign)
    return nil
}

// TopP4pBasicQuickCampaign Getter
func (r AlibabaScbpTargetAdPlanUpdateRequest) GetTopP4pBasicQuickCampaign() *TopP4pBasicQuickCampaign {
    return r._topP4pBasicQuickCampaign
}
