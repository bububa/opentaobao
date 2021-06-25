package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除计划 APIRequest
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

func NewAlibabaScbpAdCampaignDeleteRequest() *AlibabaScbpAdCampaignDeleteRequest{
    return &AlibabaScbpAdCampaignDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdCampaignDeleteRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.campaign.delete"
}

func (r AlibabaScbpAdCampaignDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdCampaignDeleteRequest) SetBatchOperation(batchOperation *CampaignBatchOperationDto) error {
    r.batchOperation = batchOperation
    r.Set("batch_operation", batchOperation)
    return nil
}

func (r AlibabaScbpAdCampaignDeleteRequest) GetBatchOperation() *CampaignBatchOperationDto {
    return r.batchOperation
}

func (r *AlibabaScbpAdCampaignDeleteRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

func (r AlibabaScbpAdCampaignDeleteRequest) GetTopContext() *TopContextDto {
    return r.topContext
}

