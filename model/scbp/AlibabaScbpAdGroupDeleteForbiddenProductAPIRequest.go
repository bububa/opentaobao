package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest 删除屏蔽品 API请求
// alibaba.scbp.ad.group.delete.forbidden.product
//
// 删除屏蔽品
type AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 计划id
	_campaignId int64
	// 请求参数
	_forbiddenProductBatchOperation *ForbiddenProductBatchOperationDto
}

// NewAlibabaScbpAdGroupDeleteForbiddenProductRequest 初始化AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest对象
func NewAlibabaScbpAdGroupDeleteForbiddenProductRequest() *AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest {
	return &AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest) Reset() {
	r._topContext = nil
	r._campaignId = 0
	r._forbiddenProductBatchOperation = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.group.delete.forbidden.product"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetForbiddenProductBatchOperation is ForbiddenProductBatchOperation Setter
// 请求参数
func (r *AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest) SetForbiddenProductBatchOperation(_forbiddenProductBatchOperation *ForbiddenProductBatchOperationDto) error {
	r._forbiddenProductBatchOperation = _forbiddenProductBatchOperation
	r.Set("forbidden_product_batch_operation", _forbiddenProductBatchOperation)
	return nil
}

// GetForbiddenProductBatchOperation ForbiddenProductBatchOperation Getter
func (r AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest) GetForbiddenProductBatchOperation() *ForbiddenProductBatchOperationDto {
	return r._forbiddenProductBatchOperation
}

var poolAlibabaScbpAdGroupDeleteForbiddenProductAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdGroupDeleteForbiddenProductRequest()
	},
}

// GetAlibabaScbpAdGroupDeleteForbiddenProductRequest 从 sync.Pool 获取 AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest
func GetAlibabaScbpAdGroupDeleteForbiddenProductAPIRequest() *AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest {
	return poolAlibabaScbpAdGroupDeleteForbiddenProductAPIRequest.Get().(*AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest)
}

// ReleaseAlibabaScbpAdGroupDeleteForbiddenProductAPIRequest 将 AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdGroupDeleteForbiddenProductAPIRequest(v *AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdGroupDeleteForbiddenProductAPIRequest.Put(v)
}
