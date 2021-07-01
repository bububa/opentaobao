package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest
修改推广单元 API请求
alibaba.scbp.ad.group.update.ad.group.batch

修改推广单元 */
type AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
	// 入参
	_adGroupBatchOperation *AdGroupBatchOperationDto
	// 用户信息
	_topContext *TopContextDto
}

// NewAlibabaScbpAdGroupUpdateAdGroupBatchRequest 初始化AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest对象
func NewAlibabaScbpAdGroupUpdateAdGroupBatchRequest() *AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest {
	return &AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.group.update.ad.group.batch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CampaignId Setter
// 计划id
func (r *AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// Get CampaignId Getter
func (r AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// Set is AdGroupBatchOperation Setter
// 入参
func (r *AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest) SetAdGroupBatchOperation(_adGroupBatchOperation *AdGroupBatchOperationDto) error {
	r._adGroupBatchOperation = _adGroupBatchOperation
	r.Set("ad_group_batch_operation", _adGroupBatchOperation)
	return nil
}

// Get AdGroupBatchOperation Getter
func (r AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest) GetAdGroupBatchOperation() *AdGroupBatchOperationDto {
	return r._adGroupBatchOperation
}

// Set is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// Get TopContext Getter
func (r AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}
