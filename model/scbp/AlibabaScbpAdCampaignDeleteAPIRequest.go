package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadcampaigndeleteAPIRequest 删除计划 API请求
// alibaba.scbp.ad.campaign.delete
//
// 删除计划
type AlibabascbpadcampaigndeleteAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 操作对象
	_batchOperation *CampaignBatchOperationDto
}

// NewAlibabascbpadcampaigndeleteRequest 初始化AlibabascbpadcampaigndeleteAPIRequest对象
func NewAlibabascbpadcampaigndeleteRequest() *AlibabascbpadcampaigndeleteAPIRequest {
	return &AlibabascbpadcampaigndeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadcampaigndeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.campaign.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadcampaigndeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadcampaigndeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadcampaigndeleteAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadcampaigndeleteAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetBatchOperation is BatchOperation Setter
// 操作对象
func (r *AlibabascbpadcampaigndeleteAPIRequest) SetBatchOperation(_batchOperation *CampaignBatchOperationDto) error {
	r._batchOperation = _batchOperation
	r.Set("batch_operation", _batchOperation)
	return nil
}

// GetBatchOperation BatchOperation Getter
func (r AlibabascbpadcampaigndeleteAPIRequest) GetBatchOperation() *CampaignBatchOperationDto {
	return r._batchOperation
}
