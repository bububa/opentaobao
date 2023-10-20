package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadgroupdeleteadgroupbatchAPIRequest 删除推广单元 API请求
// alibaba.scbp.ad.group.delete.ad.group.batch
//
// 删除推广单元
type AlibabascbpadgroupdeleteadgroupbatchAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 计划id
	_campaignId int64
	// 请求参数
	_adGroupBatchOperation *AdGroupBatchOperationDto
}

// NewAlibabascbpadgroupdeleteadgroupbatchRequest 初始化AlibabascbpadgroupdeleteadgroupbatchAPIRequest对象
func NewAlibabascbpadgroupdeleteadgroupbatchRequest() *AlibabascbpadgroupdeleteadgroupbatchAPIRequest {
	return &AlibabascbpadgroupdeleteadgroupbatchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadgroupdeleteadgroupbatchAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.group.delete.ad.group.batch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadgroupdeleteadgroupbatchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadgroupdeleteadgroupbatchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadgroupdeleteadgroupbatchAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadgroupdeleteadgroupbatchAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabascbpadgroupdeleteadgroupbatchAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabascbpadgroupdeleteadgroupbatchAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetAdGroupBatchOperation is AdGroupBatchOperation Setter
// 请求参数
func (r *AlibabascbpadgroupdeleteadgroupbatchAPIRequest) SetAdGroupBatchOperation(_adGroupBatchOperation *AdGroupBatchOperationDto) error {
	r._adGroupBatchOperation = _adGroupBatchOperation
	r.Set("ad_group_batch_operation", _adGroupBatchOperation)
	return nil
}

// GetAdGroupBatchOperation AdGroupBatchOperation Getter
func (r AlibabascbpadgroupdeleteadgroupbatchAPIRequest) GetAdGroupBatchOperation() *AdGroupBatchOperationDto {
	return r._adGroupBatchOperation
}
