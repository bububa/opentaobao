package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadgroupcreateadgroupbatchAPIRequest 创建推广单元 API请求
// alibaba.scbp.ad.group.create.ad.group.batch
//
// 创建推广单元
type AlibabascbpadgroupcreateadgroupbatchAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 计划id
	_campaignId int64
	// 入参
	_adGroupBatchOperation *AdGroupBatchOperationDto
}

// NewAlibabascbpadgroupcreateadgroupbatchRequest 初始化AlibabascbpadgroupcreateadgroupbatchAPIRequest对象
func NewAlibabascbpadgroupcreateadgroupbatchRequest() *AlibabascbpadgroupcreateadgroupbatchAPIRequest {
	return &AlibabascbpadgroupcreateadgroupbatchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadgroupcreateadgroupbatchAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.group.create.ad.group.batch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadgroupcreateadgroupbatchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadgroupcreateadgroupbatchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadgroupcreateadgroupbatchAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadgroupcreateadgroupbatchAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabascbpadgroupcreateadgroupbatchAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabascbpadgroupcreateadgroupbatchAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetAdGroupBatchOperation is AdGroupBatchOperation Setter
// 入参
func (r *AlibabascbpadgroupcreateadgroupbatchAPIRequest) SetAdGroupBatchOperation(_adGroupBatchOperation *AdGroupBatchOperationDto) error {
	r._adGroupBatchOperation = _adGroupBatchOperation
	r.Set("ad_group_batch_operation", _adGroupBatchOperation)
	return nil
}

// GetAdGroupBatchOperation AdGroupBatchOperation Getter
func (r AlibabascbpadgroupcreateadgroupbatchAPIRequest) GetAdGroupBatchOperation() *AdGroupBatchOperationDto {
	return r._adGroupBatchOperation
}
