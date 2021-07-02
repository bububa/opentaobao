package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest 查询关键词前五名排价 API请求
// alibaba.scbp.ad.keyword.query.keyword.rank.price
//
// 查询关键词前五名排价
type AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
	// 请求参数
	_keywordQuery *KeywordQuery
	// 用户信息
	_topContext *TopContextDto
}

// NewAlibabaScbpAdKeywordQueryKeywordRankPriceRequest 初始化AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest对象
func NewAlibabaScbpAdKeywordQueryKeywordRankPriceRequest() *AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest {
	return &AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.query.keyword.rank.price"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CampaignId Setter
// 计划id
func (r *AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// Get CampaignId Getter
func (r AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// Set is KeywordQuery Setter
// 请求参数
func (r *AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest) SetKeywordQuery(_keywordQuery *KeywordQuery) error {
	r._keywordQuery = _keywordQuery
	r.Set("keyword_query", _keywordQuery)
	return nil
}

// Get KeywordQuery Getter
func (r AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest) GetKeywordQuery() *KeywordQuery {
	return r._keywordQuery
}

// Set is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// Get TopContext Getter
func (r AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}
