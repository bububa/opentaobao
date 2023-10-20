package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadtargettagfindcampaigntargettagAPIRequest 查询标签数据 API请求
// alibaba.scbp.ad.target.tag.find.campaign.target.tag
//
// 查询标签数据
type AlibabascbpadtargettagfindcampaigntargettagAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 计划id
	_campaignId int64
	// 请求参数
	_targetTagOperation *TargetTagOperationDto
}

// NewAlibabascbpadtargettagfindcampaigntargettagRequest 初始化AlibabascbpadtargettagfindcampaigntargettagAPIRequest对象
func NewAlibabascbpadtargettagfindcampaigntargettagRequest() *AlibabascbpadtargettagfindcampaigntargettagAPIRequest {
	return &AlibabascbpadtargettagfindcampaigntargettagAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadtargettagfindcampaigntargettagAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.target.tag.find.campaign.target.tag"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadtargettagfindcampaigntargettagAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadtargettagfindcampaigntargettagAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadtargettagfindcampaigntargettagAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadtargettagfindcampaigntargettagAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabascbpadtargettagfindcampaigntargettagAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabascbpadtargettagfindcampaigntargettagAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetTargetTagOperation is TargetTagOperation Setter
// 请求参数
func (r *AlibabascbpadtargettagfindcampaigntargettagAPIRequest) SetTargetTagOperation(_targetTagOperation *TargetTagOperationDto) error {
	r._targetTagOperation = _targetTagOperation
	r.Set("target_tag_operation", _targetTagOperation)
	return nil
}

// GetTargetTagOperation TargetTagOperation Getter
func (r AlibabascbpadtargettagfindcampaigntargettagAPIRequest) GetTargetTagOperation() *TargetTagOperationDto {
	return r._targetTagOperation
}
