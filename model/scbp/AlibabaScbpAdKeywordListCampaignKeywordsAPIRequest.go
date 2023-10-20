package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadkeywordlistcampaignkeywordsAPIRequest 获取计划关键词 API请求
// alibaba.scbp.ad.keyword.list.campaign.keywords
//
// 获取计划关键词
type AlibabascbpadkeywordlistcampaignkeywordsAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 计划id
	_campaignId int64
	// 搜索条件
	_campaignKeywordQuery *CampaignKeywordQuery
}

// NewAlibabascbpadkeywordlistcampaignkeywordsRequest 初始化AlibabascbpadkeywordlistcampaignkeywordsAPIRequest对象
func NewAlibabascbpadkeywordlistcampaignkeywordsRequest() *AlibabascbpadkeywordlistcampaignkeywordsAPIRequest {
	return &AlibabascbpadkeywordlistcampaignkeywordsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadkeywordlistcampaignkeywordsAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.list.campaign.keywords"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadkeywordlistcampaignkeywordsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadkeywordlistcampaignkeywordsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadkeywordlistcampaignkeywordsAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadkeywordlistcampaignkeywordsAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabascbpadkeywordlistcampaignkeywordsAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabascbpadkeywordlistcampaignkeywordsAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetCampaignKeywordQuery is CampaignKeywordQuery Setter
// 搜索条件
func (r *AlibabascbpadkeywordlistcampaignkeywordsAPIRequest) SetCampaignKeywordQuery(_campaignKeywordQuery *CampaignKeywordQuery) error {
	r._campaignKeywordQuery = _campaignKeywordQuery
	r.Set("campaign_keyword_query", _campaignKeywordQuery)
	return nil
}

// GetCampaignKeywordQuery CampaignKeywordQuery Getter
func (r AlibabascbpadkeywordlistcampaignkeywordsAPIRequest) GetCampaignKeywordQuery() *CampaignKeywordQuery {
	return r._campaignKeywordQuery
}
