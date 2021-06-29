package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改推广单元 API请求
alibaba.scbp.ad.group.update.ad.group.batch

修改推广单元
*/
type AlibabaScbpAdGroupUpdateAdGroupBatchRequest struct {
    model.Params
    // 计划id
    campaignId   int64
    // 入参
    adGroupBatchOperation   *AdGroupBatchOperationDto
    // 用户信息
    topContext   *TopContextDto
}

// 初始化AlibabaScbpAdGroupUpdateAdGroupBatchRequest对象
func NewAlibabaScbpAdGroupUpdateAdGroupBatchRequest() *AlibabaScbpAdGroupUpdateAdGroupBatchRequest{
    return &AlibabaScbpAdGroupUpdateAdGroupBatchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdGroupUpdateAdGroupBatchRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.group.update.ad.group.batch"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdGroupUpdateAdGroupBatchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *AlibabaScbpAdGroupUpdateAdGroupBatchRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdGroupUpdateAdGroupBatchRequest) GetCampaignId() int64 {
    return r.campaignId
}
// AdGroupBatchOperation Setter
// 入参
func (r *AlibabaScbpAdGroupUpdateAdGroupBatchRequest) SetAdGroupBatchOperation(adGroupBatchOperation *AdGroupBatchOperationDto) error {
    r.adGroupBatchOperation = adGroupBatchOperation
    r.Set("ad_group_batch_operation", adGroupBatchOperation)
    return nil
}

// AdGroupBatchOperation Getter
func (r AlibabaScbpAdGroupUpdateAdGroupBatchRequest) GetAdGroupBatchOperation() *AdGroupBatchOperationDto {
    return r.adGroupBatchOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdGroupUpdateAdGroupBatchRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdGroupUpdateAdGroupBatchRequest) GetTopContext() *TopContextDto {
    return r.topContext
}
