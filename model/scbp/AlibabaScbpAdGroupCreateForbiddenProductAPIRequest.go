package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadgroupcreateforbiddenproductAPIRequest 创建屏蔽品 API请求
// alibaba.scbp.ad.group.create.forbidden.product
//
// 创建屏蔽品
type AlibabascbpadgroupcreateforbiddenproductAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 计划id
	_campaignId int64
	// 查询条件
	_forbiddenProductBatchOperation *ForbiddenProductBatchOperationDto
}

// NewAlibabascbpadgroupcreateforbiddenproductRequest 初始化AlibabascbpadgroupcreateforbiddenproductAPIRequest对象
func NewAlibabascbpadgroupcreateforbiddenproductRequest() *AlibabascbpadgroupcreateforbiddenproductAPIRequest {
	return &AlibabascbpadgroupcreateforbiddenproductAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadgroupcreateforbiddenproductAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.group.create.forbidden.product"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadgroupcreateforbiddenproductAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadgroupcreateforbiddenproductAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadgroupcreateforbiddenproductAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadgroupcreateforbiddenproductAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabascbpadgroupcreateforbiddenproductAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabascbpadgroupcreateforbiddenproductAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetForbiddenProductBatchOperation is ForbiddenProductBatchOperation Setter
// 查询条件
func (r *AlibabascbpadgroupcreateforbiddenproductAPIRequest) SetForbiddenProductBatchOperation(_forbiddenProductBatchOperation *ForbiddenProductBatchOperationDto) error {
	r._forbiddenProductBatchOperation = _forbiddenProductBatchOperation
	r.Set("forbidden_product_batch_operation", _forbiddenProductBatchOperation)
	return nil
}

// GetForbiddenProductBatchOperation ForbiddenProductBatchOperation Getter
func (r AlibabascbpadgroupcreateforbiddenproductAPIRequest) GetForbiddenProductBatchOperation() *ForbiddenProductBatchOperationDto {
	return r._forbiddenProductBatchOperation
}
