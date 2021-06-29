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
    _batchOperation   *CampaignBatchOperationDTO
    // 用户信息
    _topContext   *TopContextDTO
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
func (r *AlibabaScbpAdCampaignDeleteRequest) SetBatchOperation(_batchOperation *CampaignBatchOperationDTO) error {
    r._batchOperation = _batchOperation
    r.Set("batch_operation", _batchOperation)
    return nil
}

// BatchOperation Getter
func (r AlibabaScbpAdCampaignDeleteRequest) GetBatchOperation() *CampaignBatchOperationDTO {
    return r._batchOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdCampaignDeleteRequest) SetTopContext(_topContext *TopContextDTO) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdCampaignDeleteRequest) GetTopContext() *TopContextDTO {
    return r._topContext
}
