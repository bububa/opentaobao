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
    _campaignId   int64
    // 入参
    _adGroupBatchOperation   *AdGroupBatchOperationDTO
    // 用户信息
    _topContext   *TopContextDTO
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
func (r *AlibabaScbpAdGroupCreateAdGroupBatchRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdGroupCreateAdGroupBatchRequest) GetCampaignId() int64 {
    return r._campaignId
}
// AdGroupBatchOperation Setter
// 入参
func (r *AlibabaScbpAdGroupCreateAdGroupBatchRequest) SetAdGroupBatchOperation(_adGroupBatchOperation *AdGroupBatchOperationDTO) error {
    r._adGroupBatchOperation = _adGroupBatchOperation
    r.Set("ad_group_batch_operation", _adGroupBatchOperation)
    return nil
}

// AdGroupBatchOperation Getter
func (r AlibabaScbpAdGroupCreateAdGroupBatchRequest) GetAdGroupBatchOperation() *AdGroupBatchOperationDTO {
    return r._adGroupBatchOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdGroupCreateAdGroupBatchRequest) SetTopContext(_topContext *TopContextDTO) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdGroupCreateAdGroupBatchRequest) GetTopContext() *TopContextDTO {
    return r._topContext
}
