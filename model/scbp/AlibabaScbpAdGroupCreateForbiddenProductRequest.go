package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建屏蔽品 API请求
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

// 初始化AlibabaScbpAdGroupCreateForbiddenProductRequest对象
func NewAlibabaScbpAdGroupCreateForbiddenProductRequest() *AlibabaScbpAdGroupCreateForbiddenProductRequest{
    return &AlibabaScbpAdGroupCreateForbiddenProductRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdGroupCreateForbiddenProductRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.group.create.forbidden.product"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdGroupCreateForbiddenProductRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *AlibabaScbpAdGroupCreateForbiddenProductRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdGroupCreateForbiddenProductRequest) GetCampaignId() int64 {
    return r.campaignId
}
// ForbiddenProductBatchOperation Setter
// 查询条件
func (r *AlibabaScbpAdGroupCreateForbiddenProductRequest) SetForbiddenProductBatchOperation(forbiddenProductBatchOperation *ForbiddenProductBatchOperationDto) error {
    r.forbiddenProductBatchOperation = forbiddenProductBatchOperation
    r.Set("forbidden_product_batch_operation", forbiddenProductBatchOperation)
    return nil
}

// ForbiddenProductBatchOperation Getter
func (r AlibabaScbpAdGroupCreateForbiddenProductRequest) GetForbiddenProductBatchOperation() *ForbiddenProductBatchOperationDto {
    return r.forbiddenProductBatchOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdGroupCreateForbiddenProductRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdGroupCreateForbiddenProductRequest) GetTopContext() *TopContextDto {
    return r.topContext
}
