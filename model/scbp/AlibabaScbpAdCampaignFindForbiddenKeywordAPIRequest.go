package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest 查询屏蔽词 API请求
// alibaba.scbp.ad.campaign.find.forbidden.keyword
//
// 查询屏蔽词
type AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
	// 用户信息
	_topContext *TopContextDto
}

// NewAlibabaScbpAdCampaignFindForbiddenKeywordRequest 初始化AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest对象
func NewAlibabaScbpAdCampaignFindForbiddenKeywordRequest() *AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest {
	return &AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.campaign.find.forbidden.keyword"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CampaignId Setter
// 计划id
func (r *AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// Get CampaignId Getter
func (r AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// Set is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// Get TopContext Getter
func (r AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}
