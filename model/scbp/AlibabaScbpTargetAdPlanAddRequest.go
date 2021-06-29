package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-新建计划 API请求
alibaba.scbp.target.ad.plan.add

定向推广-新建单条计划
*/
type AlibabaScbpTargetAdPlanAddRequest struct {
    model.Params
    // 定向推广基础信息
    _topP4pBasicQuickCampaign   *BasicQuickCampaign
}

// 初始化AlibabaScbpTargetAdPlanAddRequest对象
func NewAlibabaScbpTargetAdPlanAddRequest() *AlibabaScbpTargetAdPlanAddRequest{
    return &AlibabaScbpTargetAdPlanAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpTargetAdPlanAddRequest) GetApiMethodName() string {
    return "alibaba.scbp.target.ad.plan.add"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpTargetAdPlanAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TopP4pBasicQuickCampaign Setter
// 定向推广基础信息
func (r *AlibabaScbpTargetAdPlanAddRequest) SetTopP4pBasicQuickCampaign(_topP4pBasicQuickCampaign *BasicQuickCampaign) error {
    r._topP4pBasicQuickCampaign = _topP4pBasicQuickCampaign
    r.Set("top_p4p_basic_quick_campaign", _topP4pBasicQuickCampaign)
    return nil
}

// TopP4pBasicQuickCampaign Getter
func (r AlibabaScbpTargetAdPlanAddRequest) GetTopP4pBasicQuickCampaign() *BasicQuickCampaign {
    return r._topP4pBasicQuickCampaign
}
