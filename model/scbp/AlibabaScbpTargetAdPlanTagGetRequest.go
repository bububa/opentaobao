package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-获取计划的定向溢价数据 APIRequest
alibaba.scbp.target.ad.plan.tag.get

定向推广-获取计划的定向溢价数据
*/
type AlibabaScbpTargetAdPlanTagGetRequest struct {
    model.Params

    // 推广计划Id
    campaignId   int64 

}

func NewAlibabaScbpTargetAdPlanTagGetRequest() *AlibabaScbpTargetAdPlanTagGetRequest{
    return &AlibabaScbpTargetAdPlanTagGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpTargetAdPlanTagGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.target.ad.plan.tag.get"
}

func (r AlibabaScbpTargetAdPlanTagGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpTargetAdPlanTagGetRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r AlibabaScbpTargetAdPlanTagGetRequest) GetCampaignId() int64 {
    return r.campaignId
}

