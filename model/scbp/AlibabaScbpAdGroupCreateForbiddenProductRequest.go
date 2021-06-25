package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建屏蔽品 APIRequest
alibaba.scbp.ad.group.create.forbidden.product

创建屏蔽品
*/
type AlibabaScbpAdGroupCreateForbiddenProductRequest struct {
    model.Params

    // 计划id
    campaignId   int64 

    // 查询条件
    forbiddenProductBatchOperation   *ForbiddenProductBatchOperationDto 

    // 用户信息
    topContext   *TopContextDto 

}

func NewAlibabaScbpAdGroupCreateForbiddenProductRequest() *AlibabaScbpAdGroupCreateForbiddenProductRequest{
    return &AlibabaScbpAdGroupCreateForbiddenProductRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdGroupCreateForbiddenProductRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.group.create.forbidden.product"
}

func (r AlibabaScbpAdGroupCreateForbiddenProductRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdGroupCreateForbiddenProductRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r AlibabaScbpAdGroupCreateForbiddenProductRequest) GetCampaignId() int64 {
    return r.campaignId
}

func (r *AlibabaScbpAdGroupCreateForbiddenProductRequest) SetForbiddenProductBatchOperation(forbiddenProductBatchOperation *ForbiddenProductBatchOperationDto) error {
    r.forbiddenProductBatchOperation = forbiddenProductBatchOperation
    r.Set("forbidden_product_batch_operation", forbiddenProductBatchOperation)
    return nil
}

func (r AlibabaScbpAdGroupCreateForbiddenProductRequest) GetForbiddenProductBatchOperation() *ForbiddenProductBatchOperationDto {
    return r.forbiddenProductBatchOperation
}

func (r *AlibabaScbpAdGroupCreateForbiddenProductRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

func (r AlibabaScbpAdGroupCreateForbiddenProductRequest) GetTopContext() *TopContextDto {
    return r.topContext
}

