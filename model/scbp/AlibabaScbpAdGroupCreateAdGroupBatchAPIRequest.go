package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest 创建推广单元 API请求
// alibaba.scbp.ad.group.create.ad.group.batch
//
// 创建推广单元
type AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
	// 入参
	_adGroupBatchOperation *AdGroupBatchOperationDto
	// 用户信息
	_topContext *TopContextDto
}

// NewAlibabaScbpAdGroupCreateAdGroupBatchRequest 初始化AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest对象
func NewAlibabaScbpAdGroupCreateAdGroupBatchRequest() *AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest {
	return &AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.group.create.ad.group.batch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdGroupCreateAdGroupBatchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
