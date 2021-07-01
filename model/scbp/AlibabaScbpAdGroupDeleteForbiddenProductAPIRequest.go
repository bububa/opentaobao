package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest
删除屏蔽品 API请求
alibaba.scbp.ad.group.delete.forbidden.product

删除屏蔽品 */
type AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
	// 请求参数
	_forbiddenProductBatchOperation *ForbiddenProductBatchOperationDto
	// 用户信息
	_topContext *TopContextDto
}

// NewAlibabaScbpAdGroupDeleteForbiddenProductRequest 初始化AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest对象
func NewAlibabaScbpAdGroupDeleteForbiddenProductRequest() *AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest {
	return &AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.group.delete.forbidden.product"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CampaignId Setter
// 计划id
func (r *AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// Get CampaignId Getter
func (r AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// Set is ForbiddenProductBatchOperation Setter
// 请求参数
func (r *AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest) SetForbiddenProductBatchOperation(_forbiddenProductBatchOperation *ForbiddenProductBatchOperationDto) error {
	r._forbiddenProductBatchOperation = _forbiddenProductBatchOperation
	r.Set("forbidden_product_batch_operation", _forbiddenProductBatchOperation)
	return nil
}

// Get ForbiddenProductBatchOperation Getter
func (r AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest) GetForbiddenProductBatchOperation() *ForbiddenProductBatchOperationDto {
	return r._forbiddenProductBatchOperation
}

// Set is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// Get TopContext Getter
func (r AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}
