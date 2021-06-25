package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-查询定向推广计划列表并返回计划基础信息 APIRequest
alibaba.scbp.target.ad.plan.list

定向推广-查询定向推广计划列表并返回计划基础信息
*/
type AlibabaScbpTargetAdPlanListRequest struct {
    model.Params

    // TopP4pQuickCampaignQuery
    topP4pQuickCampaignQuery   *TopP4pQuickCampaignQueryDto 

}

func NewAlibabaScbpTargetAdPlanListRequest() *AlibabaScbpTargetAdPlanListRequest{
    return &AlibabaScbpTargetAdPlanListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpTargetAdPlanListRequest) GetApiMethodName() string {
    return "alibaba.scbp.target.ad.plan.list"
}

func (r AlibabaScbpTargetAdPlanListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpTargetAdPlanListRequest) SetTopP4pQuickCampaignQuery(topP4pQuickCampaignQuery *TopP4pQuickCampaignQueryDto) error {
    r.topP4pQuickCampaignQuery = topP4pQuickCampaignQuery
    r.Set("top_p4p_quick_campaign_query", topP4pQuickCampaignQuery)
    return nil
}

func (r AlibabaScbpTargetAdPlanListRequest) GetTopP4pQuickCampaignQuery() *TopP4pQuickCampaignQueryDto {
    return r.topP4pQuickCampaignQuery
}

