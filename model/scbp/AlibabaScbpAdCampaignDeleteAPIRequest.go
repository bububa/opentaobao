package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdCampaignDeleteAPIRequest 删除计划 API请求
// alibaba.scbp.ad.campaign.delete
//
// 删除计划
type AlibabaScbpAdCampaignDeleteAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 操作对象
	_batchOperation *CampaignBatchOperationDto
}

// NewAlibabaScbpAdCampaignDeleteRequest 初始化AlibabaScbpAdCampaignDeleteAPIRequest对象
func NewAlibabaScbpAdCampaignDeleteRequest() *AlibabaScbpAdCampaignDeleteAPIRequest {
	return &AlibabaScbpAdCampaignDeleteAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdCampaignDeleteAPIRequest) Reset() {
	r._topContext = nil
	r._batchOperation = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdCampaignDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.campaign.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdCampaignDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdCampaignDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdCampaignDeleteAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdCampaignDeleteAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetBatchOperation is BatchOperation Setter
// 操作对象
func (r *AlibabaScbpAdCampaignDeleteAPIRequest) SetBatchOperation(_batchOperation *CampaignBatchOperationDto) error {
	r._batchOperation = _batchOperation
	r.Set("batch_operation", _batchOperation)
	return nil
}

// GetBatchOperation BatchOperation Getter
func (r AlibabaScbpAdCampaignDeleteAPIRequest) GetBatchOperation() *CampaignBatchOperationDto {
	return r._batchOperation
}

var poolAlibabaScbpAdCampaignDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdCampaignDeleteRequest()
	},
}

// GetAlibabaScbpAdCampaignDeleteRequest 从 sync.Pool 获取 AlibabaScbpAdCampaignDeleteAPIRequest
func GetAlibabaScbpAdCampaignDeleteAPIRequest() *AlibabaScbpAdCampaignDeleteAPIRequest {
	return poolAlibabaScbpAdCampaignDeleteAPIRequest.Get().(*AlibabaScbpAdCampaignDeleteAPIRequest)
}

// ReleaseAlibabaScbpAdCampaignDeleteAPIRequest 将 AlibabaScbpAdCampaignDeleteAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdCampaignDeleteAPIRequest(v *AlibabaScbpAdCampaignDeleteAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdCampaignDeleteAPIRequest.Put(v)
}
