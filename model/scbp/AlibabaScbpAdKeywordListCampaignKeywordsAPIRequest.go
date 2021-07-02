package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest 获取计划关键词 API请求
// alibaba.scbp.ad.keyword.list.campaign.keywords
//
// 获取计划关键词
type AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
	// 搜索条件
	_campaignKeywordQuery *CampaignKeywordQuery
	// 用户信息
	_topContext *TopContextDto
}

// NewAlibabaScbpAdKeywordListCampaignKeywordsRequest 初始化AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest对象
func NewAlibabaScbpAdKeywordListCampaignKeywordsRequest() *AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest {
	return &AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.list.campaign.keywords"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CampaignId Setter
// 计划id
func (r *AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// Get CampaignId Getter
func (r AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// Set is CampaignKeywordQuery Setter
// 搜索条件
func (r *AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest) SetCampaignKeywordQuery(_campaignKeywordQuery *CampaignKeywordQuery) error {
	r._campaignKeywordQuery = _campaignKeywordQuery
	r.Set("campaign_keyword_query", _campaignKeywordQuery)
	return nil
}

// Get CampaignKeywordQuery Getter
func (r AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest) GetCampaignKeywordQuery() *CampaignKeywordQuery {
	return r._campaignKeywordQuery
}

// Set is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// Get TopContext Getter
func (r AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}
