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
    _campaignId   int64
    // 查询条件
    _forbiddenProductBatchOperation   *ForbiddenProductBatchOperationDto
    // 用户信息
    _topContext   *TopContextDto
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
func (r *AlibabaScbpAdGroupCreateForbiddenProductRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdGroupCreateForbiddenProductRequest) GetCampaignId() int64 {
    return r._campaignId
}
// ForbiddenProductBatchOperation Setter
// 查询条件
func (r *AlibabaScbpAdGroupCreateForbiddenProductRequest) SetForbiddenProductBatchOperation(_forbiddenProductBatchOperation *ForbiddenProductBatchOperationDto) error {
    r._forbiddenProductBatchOperation = _forbiddenProductBatchOperation
    r.Set("forbidden_product_batch_operation", _forbiddenProductBatchOperation)
    return nil
}

// ForbiddenProductBatchOperation Getter
func (r AlibabaScbpAdGroupCreateForbiddenProductRequest) GetForbiddenProductBatchOperation() *ForbiddenProductBatchOperationDto {
    return r._forbiddenProductBatchOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdGroupCreateForbiddenProductRequest) SetTopContext(_topContext *TopContextDto) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdGroupCreateForbiddenProductRequest) GetTopContext() *TopContextDto {
    return r._topContext
}
