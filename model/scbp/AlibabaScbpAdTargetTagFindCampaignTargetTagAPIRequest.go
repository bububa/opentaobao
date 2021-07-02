package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdTargetTagFindCampaignTargetTagAPIRequest 查询标签数据 API请求
// alibaba.scbp.ad.target.tag.find.campaign.target.tag
//
// 查询标签数据
type AlibabaScbpAdTargetTagFindCampaignTargetTagAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
	// 请求参数
	_targetTagOperation *TargetTagOperationDto
	// 用户信息
	_topContext *TopContextDto
}

// NewAlibabaScbpAdTargetTagFindCampaignTargetTagRequest 初始化AlibabaScbpAdTargetTagFindCampaignTargetTagAPIRequest对象
func NewAlibabaScbpAdTargetTagFindCampaignTargetTagRequest() *AlibabaScbpAdTargetTagFindCampaignTargetTagAPIRequest {
	return &AlibabaScbpAdTargetTagFindCampaignTargetTagAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdTargetTagFindCampaignTargetTagAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.target.tag.find.campaign.target.tag"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdTargetTagFindCampaignTargetTagAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabaScbpAdTargetTagFindCampaignTargetTagAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabaScbpAdTargetTagFindCampaignTargetTagAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetTargetTagOperation is TargetTagOperation Setter
// 请求参数
func (r *AlibabaScbpAdTargetTagFindCampaignTargetTagAPIRequest) SetTargetTagOperation(_targetTagOperation *TargetTagOperationDto) error {
	r._targetTagOperation = _targetTagOperation
	r.Set("target_tag_operation", _targetTagOperation)
	return nil
}

// GetTargetTagOperation TargetTagOperation Getter
func (r AlibabaScbpAdTargetTagFindCampaignTargetTagAPIRequest) GetTargetTagOperation() *TargetTagOperationDto {
	return r._targetTagOperation
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdTargetTagFindCampaignTargetTagAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdTargetTagFindCampaignTargetTagAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}
