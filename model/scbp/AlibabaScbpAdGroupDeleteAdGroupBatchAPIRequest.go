package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest 删除推广单元 API请求
// alibaba.scbp.ad.group.delete.ad.group.batch
//
// 删除推广单元
type AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 计划id
	_campaignId int64
	// 请求参数
	_adGroupBatchOperation *AdGroupBatchOperationDto
}

// NewAlibabaScbpAdGroupDeleteAdGroupBatchRequest 初始化AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest对象
func NewAlibabaScbpAdGroupDeleteAdGroupBatchRequest() *AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest {
	return &AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest) Reset() {
	r._topContext = nil
	r._campaignId = 0
	r._adGroupBatchOperation = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.group.delete.ad.group.batch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetAdGroupBatchOperation is AdGroupBatchOperation Setter
// 请求参数
func (r *AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest) SetAdGroupBatchOperation(_adGroupBatchOperation *AdGroupBatchOperationDto) error {
	r._adGroupBatchOperation = _adGroupBatchOperation
	r.Set("ad_group_batch_operation", _adGroupBatchOperation)
	return nil
}

// GetAdGroupBatchOperation AdGroupBatchOperation Getter
func (r AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest) GetAdGroupBatchOperation() *AdGroupBatchOperationDto {
	return r._adGroupBatchOperation
}

var poolAlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdGroupDeleteAdGroupBatchRequest()
	},
}

// GetAlibabaScbpAdGroupDeleteAdGroupBatchRequest 从 sync.Pool 获取 AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest
func GetAlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest() *AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest {
	return poolAlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest.Get().(*AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest)
}

// ReleaseAlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest 将 AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest(v *AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest.Put(v)
}
