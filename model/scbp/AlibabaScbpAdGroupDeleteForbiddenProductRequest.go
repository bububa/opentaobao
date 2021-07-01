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
type AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest struct {
    model.Params
    // 计划id
    _campaignId   int64
    // 请求参数
    _forbiddenProductBatchOperation   *ForbiddenProductBatchOperationDTO
    // 用户信息
    _topContext   *TopContextDTO
}

// 初始化AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest对象
func NewAlibabaScbpAdGroupDeleteForbiddenProductRequest() *AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest{
    return &AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.group.delete.forbidden.product"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest) GetCampaignId() int64 {
    return r._campaignId
}
// ForbiddenProductBatchOperation Setter
// 请求参数
func (r *AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest) SetForbiddenProductBatchOperation(_forbiddenProductBatchOperation *ForbiddenProductBatchOperationDTO) error {
    r._forbiddenProductBatchOperation = _forbiddenProductBatchOperation
    r.Set("forbidden_product_batch_operation", _forbiddenProductBatchOperation)
    return nil
}

// ForbiddenProductBatchOperation Getter
func (r AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest) GetForbiddenProductBatchOperation() *ForbiddenProductBatchOperationDTO {
    return r._forbiddenProductBatchOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest) SetTopContext(_topContext *TopContextDTO) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest) GetTopContext() *TopContextDTO {
    return r._topContext
}
