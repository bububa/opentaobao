package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询推广组 API请求
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

// 初始化AlibabaScbpAdGroupFindAdGroupRequest对象
func NewAlibabaScbpAdGroupFindAdGroupRequest() *AlibabaScbpAdGroupFindAdGroupRequest{
    return &AlibabaScbpAdGroupFindAdGroupRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdGroupFindAdGroupRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.group.find.ad.group"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdGroupFindAdGroupRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *AlibabaScbpAdGroupFindAdGroupRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdGroupFindAdGroupRequest) GetCampaignId() int64 {
    return r.campaignId
}
// AdGroupQuery Setter
// 入参
func (r *AlibabaScbpAdGroupFindAdGroupRequest) SetAdGroupQuery(adGroupQuery *AdGroupQueryDto) error {
    r.adGroupQuery = adGroupQuery
    r.Set("ad_group_query", adGroupQuery)
    return nil
}

// AdGroupQuery Getter
func (r AlibabaScbpAdGroupFindAdGroupRequest) GetAdGroupQuery() *AdGroupQueryDto {
    return r.adGroupQuery
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdGroupFindAdGroupRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdGroupFindAdGroupRequest) GetTopContext() *TopContextDto {
    return r.topContext
}
