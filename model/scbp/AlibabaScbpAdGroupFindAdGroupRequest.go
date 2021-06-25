package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询推广组 APIRequest
alibaba.scbp.ad.group.find.ad.group

查询推广组
*/
type AlibabaScbpAdGroupFindAdGroupRequest struct {
    model.Params

    // 计划id
    campaignId   int64 

    // 入参
    adGroupQuery   *AdGroupQueryDto 

    // 用户信息
    topContext   *TopContextDto 

}

func NewAlibabaScbpAdGroupFindAdGroupRequest() *AlibabaScbpAdGroupFindAdGroupRequest{
    return &AlibabaScbpAdGroupFindAdGroupRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdGroupFindAdGroupRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.group.find.ad.group"
}

func (r AlibabaScbpAdGroupFindAdGroupRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdGroupFindAdGroupRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r AlibabaScbpAdGroupFindAdGroupRequest) GetCampaignId() int64 {
    return r.campaignId
}

func (r *AlibabaScbpAdGroupFindAdGroupRequest) SetAdGroupQuery(adGroupQuery *AdGroupQueryDto) error {
    r.adGroupQuery = adGroupQuery
    r.Set("ad_group_query", adGroupQuery)
    return nil
}

func (r AlibabaScbpAdGroupFindAdGroupRequest) GetAdGroupQuery() *AdGroupQueryDto {
    return r.adGroupQuery
}

func (r *AlibabaScbpAdGroupFindAdGroupRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

func (r AlibabaScbpAdGroupFindAdGroupRequest) GetTopContext() *TopContextDto {
    return r.topContext
}

