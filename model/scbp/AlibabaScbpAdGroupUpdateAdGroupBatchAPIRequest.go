package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest 修改推广单元 API请求
// alibaba.scbp.ad.group.update.ad.group.batch
//
// 修改推广单元
type AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 计划id
	_campaignId int64
	// 入参
	_adGroupBatchOperation *AdGroupBatchOperationDto
}

// NewAlibabaScbpAdGroupUpdateAdGroupBatchRequest 初始化AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest对象
func NewAlibabaScbpAdGroupUpdateAdGroupBatchRequest() *AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest {
	return &AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest) Reset() {
	r._topContext = nil
	r._campaignId = 0
	r._adGroupBatchOperation = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.group.update.ad.group.batch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetAdGroupBatchOperation is AdGroupBatchOperation Setter
// 入参
func (r *AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest) SetAdGroupBatchOperation(_adGroupBatchOperation *AdGroupBatchOperationDto) error {
	r._adGroupBatchOperation = _adGroupBatchOperation
	r.Set("ad_group_batch_operation", _adGroupBatchOperation)
	return nil
}

// GetAdGroupBatchOperation AdGroupBatchOperation Getter
func (r AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest) GetAdGroupBatchOperation() *AdGroupBatchOperationDto {
	return r._adGroupBatchOperation
}

var poolAlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdGroupUpdateAdGroupBatchRequest()
	},
}

// GetAlibabaScbpAdGroupUpdateAdGroupBatchRequest 从 sync.Pool 获取 AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest
func GetAlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest() *AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest {
	return poolAlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest.Get().(*AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest)
}

// ReleaseAlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest 将 AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest(v *AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest.Put(v)
}
