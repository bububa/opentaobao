package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除推广单元 API请求
alibaba.scbp.ad.group.delete.ad.group.batch

删除推广单元
*/
type AlibabaScbpAdGroupDeleteAdGroupBatchRequest struct {
    model.Params
    // 计划id
    _campaignId   int64
    // 请求参数
    _adGroupBatchOperation   *AdGroupBatchOperationDto
    // 用户信息
    _topContext   *TopContextDto
}

// 初始化AlibabaScbpAdGroupDeleteAdGroupBatchRequest对象
func NewAlibabaScbpAdGroupDeleteAdGroupBatchRequest() *AlibabaScbpAdGroupDeleteAdGroupBatchRequest{
    return &AlibabaScbpAdGroupDeleteAdGroupBatchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdGroupDeleteAdGroupBatchRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.group.delete.ad.group.batch"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdGroupDeleteAdGroupBatchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *AlibabaScbpAdGroupDeleteAdGroupBatchRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdGroupDeleteAdGroupBatchRequest) GetCampaignId() int64 {
    return r._campaignId
}
// AdGroupBatchOperation Setter
// 请求参数
func (r *AlibabaScbpAdGroupDeleteAdGroupBatchRequest) SetAdGroupBatchOperation(_adGroupBatchOperation *AdGroupBatchOperationDto) error {
    r._adGroupBatchOperation = _adGroupBatchOperation
    r.Set("ad_group_batch_operation", _adGroupBatchOperation)
    return nil
}

// AdGroupBatchOperation Getter
func (r AlibabaScbpAdGroupDeleteAdGroupBatchRequest) GetAdGroupBatchOperation() *AdGroupBatchOperationDto {
    return r._adGroupBatchOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdGroupDeleteAdGroupBatchRequest) SetTopContext(_topContext *TopContextDto) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdGroupDeleteAdGroupBatchRequest) GetTopContext() *TopContextDto {
    return r._topContext
}
