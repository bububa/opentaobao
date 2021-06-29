package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除屏蔽品 API请求
alibaba.scbp.ad.group.delete.forbidden.product

删除屏蔽品
*/
type AlibabaScbpAdGroupDeleteForbiddenProductRequest struct {
    model.Params
    // 计划id
    campaignId   int64
    // 请求参数
    forbiddenProductBatchOperation   *ForbiddenProductBatchOperationDto
    // 用户信息
    topContext   *TopContextDto
}

// 初始化AlibabaScbpAdGroupDeleteForbiddenProductRequest对象
func NewAlibabaScbpAdGroupDeleteForbiddenProductRequest() *AlibabaScbpAdGroupDeleteForbiddenProductRequest{
    return &AlibabaScbpAdGroupDeleteForbiddenProductRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdGroupDeleteForbiddenProductRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.group.delete.forbidden.product"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdGroupDeleteForbiddenProductRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *AlibabaScbpAdGroupDeleteForbiddenProductRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdGroupDeleteForbiddenProductRequest) GetCampaignId() int64 {
    return r.campaignId
}
// ForbiddenProductBatchOperation Setter
// 请求参数
func (r *AlibabaScbpAdGroupDeleteForbiddenProductRequest) SetForbiddenProductBatchOperation(forbiddenProductBatchOperation *ForbiddenProductBatchOperationDto) error {
    r.forbiddenProductBatchOperation = forbiddenProductBatchOperation
    r.Set("forbidden_product_batch_operation", forbiddenProductBatchOperation)
    return nil
}

// ForbiddenProductBatchOperation Getter
func (r AlibabaScbpAdGroupDeleteForbiddenProductRequest) GetForbiddenProductBatchOperation() *ForbiddenProductBatchOperationDto {
    return r.forbiddenProductBatchOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdGroupDeleteForbiddenProductRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdGroupDeleteForbiddenProductRequest) GetTopContext() *TopContextDto {
    return r.topContext
}
