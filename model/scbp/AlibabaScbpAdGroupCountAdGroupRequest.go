package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
统计adgroup数量 APIRequest
alibaba.scbp.ad.group.count.ad.group

统计adgroup数量
*/
type AlibabaScbpAdGroupCountAdGroupRequest struct {
    model.Params

    // 计划id
    campaignId   int64 

    // 查询条件
    adGroupQuery   *AdGroupQueryDto 

    // 用户信息
    topContext   *TopContextDto 

}

func NewAlibabaScbpAdGroupCountAdGroupRequest() *AlibabaScbpAdGroupCountAdGroupRequest{
    return &AlibabaScbpAdGroupCountAdGroupRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdGroupCountAdGroupRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.group.count.ad.group"
}

func (r AlibabaScbpAdGroupCountAdGroupRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdGroupCountAdGroupRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r AlibabaScbpAdGroupCountAdGroupRequest) GetCampaignId() int64 {
    return r.campaignId
}

func (r *AlibabaScbpAdGroupCountAdGroupRequest) SetAdGroupQuery(adGroupQuery *AdGroupQueryDto) error {
    r.adGroupQuery = adGroupQuery
    r.Set("ad_group_query", adGroupQuery)
    return nil
}

func (r AlibabaScbpAdGroupCountAdGroupRequest) GetAdGroupQuery() *AdGroupQueryDto {
    return r.adGroupQuery
}

func (r *AlibabaScbpAdGroupCountAdGroupRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

func (r AlibabaScbpAdGroupCountAdGroupRequest) GetTopContext() *TopContextDto {
    return r.topContext
}

