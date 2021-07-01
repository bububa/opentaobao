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
type AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest struct {
    model.Params
    // 计划id
    _campaignId   int64
    // 入参
    _adGroupBatchOperation   *AdGroupBatchOperationDto
    // 用户信息
    _topContext   *TopContextDto
}

// 初始化AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest对象
func NewAlibabaScbpAdGroupCreateAdGroupBatchRequest() *AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest{
    return &AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.group.create.ad.group.batch"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest) GetCampaignId() int64 {
    return r._campaignId
}
// AdGroupBatchOperation Setter
// 入参
func (r *AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest) SetAdGroupBatchOperation(_adGroupBatchOperation *AdGroupBatchOperationDto) error {
    r._adGroupBatchOperation = _adGroupBatchOperation
    r.Set("ad_group_batch_operation", _adGroupBatchOperation)
    return nil
}

// AdGroupBatchOperation Getter
func (r AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest) GetAdGroupBatchOperation() *AdGroupBatchOperationDto {
    return r._adGroupBatchOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest) SetTopContext(_topContext *TopContextDto) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest) GetTopContext() *TopContextDto {
    return r._topContext
}
