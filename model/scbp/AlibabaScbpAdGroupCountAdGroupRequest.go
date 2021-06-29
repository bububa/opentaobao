package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
统计adgroup数量 API请求
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

// 初始化AlibabaScbpAdGroupCountAdGroupRequest对象
func NewAlibabaScbpAdGroupCountAdGroupRequest() *AlibabaScbpAdGroupCountAdGroupRequest{
    return &AlibabaScbpAdGroupCountAdGroupRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdGroupCountAdGroupRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.group.count.ad.group"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdGroupCountAdGroupRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *AlibabaScbpAdGroupCountAdGroupRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdGroupCountAdGroupRequest) GetCampaignId() int64 {
    return r.campaignId
}
// AdGroupQuery Setter
// 查询条件
func (r *AlibabaScbpAdGroupCountAdGroupRequest) SetAdGroupQuery(adGroupQuery *AdGroupQueryDto) error {
    r.adGroupQuery = adGroupQuery
    r.Set("ad_group_query", adGroupQuery)
    return nil
}

// AdGroupQuery Getter
func (r AlibabaScbpAdGroupCountAdGroupRequest) GetAdGroupQuery() *AdGroupQueryDto {
    return r.adGroupQuery
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdGroupCountAdGroupRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdGroupCountAdGroupRequest) GetTopContext() *TopContextDto {
    return r.topContext
}
