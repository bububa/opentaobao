package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadkeywordgetkeywordcountbyqueryAPIRequest 计划关键词数目 API请求
// alibaba.scbp.ad.keyword.get.keyword.count.by.query
//
// 计划关键词数目
type AlibabascbpadkeywordgetkeywordcountbyqueryAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 请求参数
	_campaignKeywordQuery *CampaignKeywordQuery
}

// NewAlibabascbpadkeywordgetkeywordcountbyqueryRequest 初始化AlibabascbpadkeywordgetkeywordcountbyqueryAPIRequest对象
func NewAlibabascbpadkeywordgetkeywordcountbyqueryRequest() *AlibabascbpadkeywordgetkeywordcountbyqueryAPIRequest {
	return &AlibabascbpadkeywordgetkeywordcountbyqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadkeywordgetkeywordcountbyqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.get.keyword.count.by.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadkeywordgetkeywordcountbyqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadkeywordgetkeywordcountbyqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadkeywordgetkeywordcountbyqueryAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadkeywordgetkeywordcountbyqueryAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignKeywordQuery is CampaignKeywordQuery Setter
// 请求参数
func (r *AlibabascbpadkeywordgetkeywordcountbyqueryAPIRequest) SetCampaignKeywordQuery(_campaignKeywordQuery *CampaignKeywordQuery) error {
	r._campaignKeywordQuery = _campaignKeywordQuery
	r.Set("campaign_keyword_query", _campaignKeywordQuery)
	return nil
}

// GetCampaignKeywordQuery CampaignKeywordQuery Getter
func (r AlibabascbpadkeywordgetkeywordcountbyqueryAPIRequest) GetCampaignKeywordQuery() *CampaignKeywordQuery {
	return r._campaignKeywordQuery
}
