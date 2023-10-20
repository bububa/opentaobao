package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadgroupupdateadgroupbatchAPIRequest 修改推广单元 API请求
// alibaba.scbp.ad.group.update.ad.group.batch
//
// 修改推广单元
type AlibabascbpadgroupupdateadgroupbatchAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 计划id
	_campaignId int64
	// 入参
	_adGroupBatchOperation *AdGroupBatchOperationDto
}

// NewAlibabascbpadgroupupdateadgroupbatchRequest 初始化AlibabascbpadgroupupdateadgroupbatchAPIRequest对象
func NewAlibabascbpadgroupupdateadgroupbatchRequest() *AlibabascbpadgroupupdateadgroupbatchAPIRequest {
	return &AlibabascbpadgroupupdateadgroupbatchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadgroupupdateadgroupbatchAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.group.update.ad.group.batch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadgroupupdateadgroupbatchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadgroupupdateadgroupbatchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadgroupupdateadgroupbatchAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadgroupupdateadgroupbatchAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabascbpadgroupupdateadgroupbatchAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabascbpadgroupupdateadgroupbatchAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetAdGroupBatchOperation is AdGroupBatchOperation Setter
// 入参
func (r *AlibabascbpadgroupupdateadgroupbatchAPIRequest) SetAdGroupBatchOperation(_adGroupBatchOperation *AdGroupBatchOperationDto) error {
	r._adGroupBatchOperation = _adGroupBatchOperation
	r.Set("ad_group_batch_operation", _adGroupBatchOperation)
	return nil
}

// GetAdGroupBatchOperation AdGroupBatchOperation Getter
func (r AlibabascbpadgroupupdateadgroupbatchAPIRequest) GetAdGroupBatchOperation() *AdGroupBatchOperationDto {
	return r._adGroupBatchOperation
}
