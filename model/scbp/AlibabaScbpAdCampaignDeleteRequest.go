package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除计划 API请求
alibaba.scbp.ad.campaign.delete

删除计划
*/
type AlibabaScbpAdCampaignDeleteRequest struct {
    model.Params
    // 操作对象
    batchOperation   *CampaignBatchOperationDto
    // 用户信息
    topContext   *TopContextDto
}

// 初始化AlibabaScbpAdCampaignDeleteRequest对象
func NewAlibabaScbpAdCampaignDeleteRequest() *AlibabaScbpAdCampaignDeleteRequest{
    return &AlibabaScbpAdCampaignDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdCampaignDeleteRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.campaign.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdCampaignDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BatchOperation Setter
// 操作对象
func (r *AlibabaScbpAdCampaignDeleteRequest) SetBatchOperation(batchOperation *CampaignBatchOperationDto) error {
    r.batchOperation = batchOperation
    r.Set("batch_operation", batchOperation)
    return nil
}

// BatchOperation Getter
func (r AlibabaScbpAdCampaignDeleteRequest) GetBatchOperation() *CampaignBatchOperationDto {
    return r.batchOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdCampaignDeleteRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdCampaignDeleteRequest) GetTopContext() *TopContextDto {
    return r.topContext
}
