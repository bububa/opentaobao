package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建推广单元 API请求
alibaba.scbp.ad.group.create.ad.group.batch

创建推广单元
*/
type AlibabaScbpAdGroupCreateAdGroupBatchRequest struct {
    model.Params
    // 计划id
    campaignId   int64
    // 入参
    adGroupBatchOperation   *AdGroupBatchOperationDto
    // 用户信息
    topContext   *TopContextDto
}

// 初始化AlibabaScbpAdGroupCreateAdGroupBatchRequest对象
func NewAlibabaScbpAdGroupCreateAdGroupBatchRequest() *AlibabaScbpAdGroupCreateAdGroupBatchRequest{
    return &AlibabaScbpAdGroupCreateAdGroupBatchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdGroupCreateAdGroupBatchRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.group.create.ad.group.batch"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdGroupCreateAdGroupBatchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *AlibabaScbpAdGroupCreateAdGroupBatchRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdGroupCreateAdGroupBatchRequest) GetCampaignId() int64 {
    return r.campaignId
}
// AdGroupBatchOperation Setter
// 入参
func (r *AlibabaScbpAdGroupCreateAdGroupBatchRequest) SetAdGroupBatchOperation(adGroupBatchOperation *AdGroupBatchOperationDto) error {
    r.adGroupBatchOperation = adGroupBatchOperation
    r.Set("ad_group_batch_operation", adGroupBatchOperation)
    return nil
}

// AdGroupBatchOperation Getter
func (r AlibabaScbpAdGroupCreateAdGroupBatchRequest) GetAdGroupBatchOperation() *AdGroupBatchOperationDto {
    return r.adGroupBatchOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdGroupCreateAdGroupBatchRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdGroupCreateAdGroupBatchRequest) GetTopContext() *TopContextDto {
    return r.topContext
}
