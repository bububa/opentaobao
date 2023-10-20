package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadgroupdeleteforbiddenproductAPIRequest 删除屏蔽品 API请求
// alibaba.scbp.ad.group.delete.forbidden.product
//
// 删除屏蔽品
type AlibabascbpadgroupdeleteforbiddenproductAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 计划id
	_campaignId int64
	// 请求参数
	_forbiddenProductBatchOperation *ForbiddenProductBatchOperationDto
}

// NewAlibabascbpadgroupdeleteforbiddenproductRequest 初始化AlibabascbpadgroupdeleteforbiddenproductAPIRequest对象
func NewAlibabascbpadgroupdeleteforbiddenproductRequest() *AlibabascbpadgroupdeleteforbiddenproductAPIRequest {
	return &AlibabascbpadgroupdeleteforbiddenproductAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadgroupdeleteforbiddenproductAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.group.delete.forbidden.product"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadgroupdeleteforbiddenproductAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadgroupdeleteforbiddenproductAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadgroupdeleteforbiddenproductAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadgroupdeleteforbiddenproductAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabascbpadgroupdeleteforbiddenproductAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabascbpadgroupdeleteforbiddenproductAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetForbiddenProductBatchOperation is ForbiddenProductBatchOperation Setter
// 请求参数
func (r *AlibabascbpadgroupdeleteforbiddenproductAPIRequest) SetForbiddenProductBatchOperation(_forbiddenProductBatchOperation *ForbiddenProductBatchOperationDto) error {
	r._forbiddenProductBatchOperation = _forbiddenProductBatchOperation
	r.Set("forbidden_product_batch_operation", _forbiddenProductBatchOperation)
	return nil
}

// GetForbiddenProductBatchOperation ForbiddenProductBatchOperation Getter
func (r AlibabascbpadgroupdeleteforbiddenproductAPIRequest) GetForbiddenProductBatchOperation() *ForbiddenProductBatchOperationDto {
	return r._forbiddenProductBatchOperation
}
