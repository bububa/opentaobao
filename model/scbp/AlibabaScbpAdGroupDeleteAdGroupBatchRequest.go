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
type AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest struct {
    model.Params
    // 计划id
    _campaignId   int64
    // 请求参数
    _adGroupBatchOperation   *AdGroupBatchOperationDTO
    // 用户信息
    _topContext   *TopContextDTO
}

// 初始化AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest对象
func NewAlibabaScbpAdGroupDeleteAdGroupBatchRequest() *AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest{
    return &AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.group.delete.ad.group.batch"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest) GetCampaignId() int64 {
    return r._campaignId
}
// AdGroupBatchOperation Setter
// 请求参数
func (r *AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest) SetAdGroupBatchOperation(_adGroupBatchOperation *AdGroupBatchOperationDTO) error {
    r._adGroupBatchOperation = _adGroupBatchOperation
    r.Set("ad_group_batch_operation", _adGroupBatchOperation)
    return nil
}

// AdGroupBatchOperation Getter
func (r AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest) GetAdGroupBatchOperation() *AdGroupBatchOperationDTO {
    return r._adGroupBatchOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest) SetTopContext(_topContext *TopContextDTO) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest) GetTopContext() *TopContextDTO {
    return r._topContext
}
