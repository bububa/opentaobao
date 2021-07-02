package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdGroupCreateForbiddenProductAPIRequest 创建屏蔽品 API请求
// alibaba.scbp.ad.group.create.forbidden.product
//
// 创建屏蔽品
type AlibabaScbpAdGroupCreateForbiddenProductAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
	// 查询条件
	_forbiddenProductBatchOperation *ForbiddenProductBatchOperationDto
	// 用户信息
	_topContext *TopContextDto
}

// NewAlibabaScbpAdGroupCreateForbiddenProductRequest 初始化AlibabaScbpAdGroupCreateForbiddenProductAPIRequest对象
func NewAlibabaScbpAdGroupCreateForbiddenProductRequest() *AlibabaScbpAdGroupCreateForbiddenProductAPIRequest {
	return &AlibabaScbpAdGroupCreateForbiddenProductAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdGroupCreateForbiddenProductAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.group.create.forbidden.product"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdGroupCreateForbiddenProductAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabaScbpAdGroupCreateForbiddenProductAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabaScbpAdGroupCreateForbiddenProductAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetForbiddenProductBatchOperation is ForbiddenProductBatchOperation Setter
// 查询条件
func (r *AlibabaScbpAdGroupCreateForbiddenProductAPIRequest) SetForbiddenProductBatchOperation(_forbiddenProductBatchOperation *ForbiddenProductBatchOperationDto) error {
	r._forbiddenProductBatchOperation = _forbiddenProductBatchOperation
	r.Set("forbidden_product_batch_operation", _forbiddenProductBatchOperation)
	return nil
}

// GetForbiddenProductBatchOperation ForbiddenProductBatchOperation Getter
func (r AlibabaScbpAdGroupCreateForbiddenProductAPIRequest) GetForbiddenProductBatchOperation() *ForbiddenProductBatchOperationDto {
	return r._forbiddenProductBatchOperation
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdGroupCreateForbiddenProductAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdGroupCreateForbiddenProductAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}
