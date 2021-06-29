package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-查询定向推广计划列表并返回计划基础信息 API请求
alibaba.scbp.target.ad.plan.list

定向推广-查询定向推广计划列表并返回计划基础信息
*/
type AlibabaScbpTargetAdPlanListRequest struct {
    model.Params
    // TopP4pQuickCampaignQuery
    _topP4pQuickCampaignQuery   *TopP4pQuickCampaignQueryDTO
}

// 初始化AlibabaScbpTargetAdPlanListRequest对象
func NewAlibabaScbpTargetAdPlanListRequest() *AlibabaScbpTargetAdPlanListRequest{
    return &AlibabaScbpTargetAdPlanListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpTargetAdPlanListRequest) GetApiMethodName() string {
    return "alibaba.scbp.target.ad.plan.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpTargetAdPlanListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TopP4pQuickCampaignQuery Setter
// TopP4pQuickCampaignQuery
func (r *AlibabaScbpTargetAdPlanListRequest) SetTopP4pQuickCampaignQuery(_topP4pQuickCampaignQuery *TopP4pQuickCampaignQueryDTO) error {
    r._topP4pQuickCampaignQuery = _topP4pQuickCampaignQuery
    r.Set("top_p4p_quick_campaign_query", _topP4pQuickCampaignQuery)
    return nil
}

// TopP4pQuickCampaignQuery Getter
func (r AlibabaScbpTargetAdPlanListRequest) GetTopP4pQuickCampaignQuery() *TopP4pQuickCampaignQueryDTO {
    return r._topP4pQuickCampaignQuery
}
