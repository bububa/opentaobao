package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest 创建推广单元 API请求
// alibaba.scbp.ad.group.create.ad.group.batch
//
// 创建推广单元
type AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 计划id
	_campaignId int64
	// 入参
	_adGroupBatchOperation *AdGroupBatchOperationDto
}

// NewAlibabaScbpAdGroupCreateAdGroupBatchRequest 初始化AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest对象
func NewAlibabaScbpAdGroupCreateAdGroupBatchRequest() *AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest {
	return &AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest) Reset() {
	r._topContext = nil
	r._campaignId = 0
	r._adGroupBatchOperation = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.group.create.ad.group.batch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetAdGroupBatchOperation is AdGroupBatchOperation Setter
// 入参
func (r *AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest) SetAdGroupBatchOperation(_adGroupBatchOperation *AdGroupBatchOperationDto) error {
	r._adGroupBatchOperation = _adGroupBatchOperation
	r.Set("ad_group_batch_operation", _adGroupBatchOperation)
	return nil
}

// GetAdGroupBatchOperation AdGroupBatchOperation Getter
func (r AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest) GetAdGroupBatchOperation() *AdGroupBatchOperationDto {
	return r._adGroupBatchOperation
}

var poolAlibabaScbpAdGroupCreateAdGroupBatchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdGroupCreateAdGroupBatchRequest()
	},
}

// GetAlibabaScbpAdGroupCreateAdGroupBatchRequest 从 sync.Pool 获取 AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest
func GetAlibabaScbpAdGroupCreateAdGroupBatchAPIRequest() *AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest {
	return poolAlibabaScbpAdGroupCreateAdGroupBatchAPIRequest.Get().(*AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest)
}

// ReleaseAlibabaScbpAdGroupCreateAdGroupBatchAPIRequest 将 AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdGroupCreateAdGroupBatchAPIRequest(v *AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdGroupCreateAdGroupBatchAPIRequest.Put(v)
}
