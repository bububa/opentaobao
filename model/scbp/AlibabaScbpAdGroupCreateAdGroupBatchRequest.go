package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建推广单元 APIRequest
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

func NewAlibabaScbpAdGroupCreateAdGroupBatchRequest() *AlibabaScbpAdGroupCreateAdGroupBatchRequest{
    return &AlibabaScbpAdGroupCreateAdGroupBatchRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdGroupCreateAdGroupBatchRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.group.create.ad.group.batch"
}

func (r AlibabaScbpAdGroupCreateAdGroupBatchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdGroupCreateAdGroupBatchRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r AlibabaScbpAdGroupCreateAdGroupBatchRequest) GetCampaignId() int64 {
    return r.campaignId
}

func (r *AlibabaScbpAdGroupCreateAdGroupBatchRequest) SetAdGroupBatchOperation(adGroupBatchOperation *AdGroupBatchOperationDto) error {
    r.adGroupBatchOperation = adGroupBatchOperation
    r.Set("ad_group_batch_operation", adGroupBatchOperation)
    return nil
}

func (r AlibabaScbpAdGroupCreateAdGroupBatchRequest) GetAdGroupBatchOperation() *AdGroupBatchOperationDto {
    return r.adGroupBatchOperation
}

func (r *AlibabaScbpAdGroupCreateAdGroupBatchRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

func (r AlibabaScbpAdGroupCreateAdGroupBatchRequest) GetTopContext() *TopContextDto {
    return r.topContext
}

