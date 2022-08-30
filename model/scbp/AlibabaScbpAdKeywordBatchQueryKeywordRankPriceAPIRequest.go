package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIRequest 批量查询关键词前五名排价 API请求
// alibaba.scbp.ad.keyword.batch.query.keyword.rank.price
//
// 批量查询关键词前五名排价
type AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIRequest struct {
	model.Params
	// 关键词信息集合
	_keywordList []KeywordInfo
	// 用户信息
	_topContext *TopContextDto
	// 计划id
	_campaignId int64
}

// NewAlibabaScbpAdKeywordBatchQueryKeywordRankPriceRequest 初始化AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIRequest对象
func NewAlibabaScbpAdKeywordBatchQueryKeywordRankPriceRequest() *AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIRequest {
	return &AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.batch.query.keyword.rank.price"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetKeywordList is KeywordList Setter
// 关键词信息集合
func (r *AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIRequest) SetKeywordList(_keywordList []KeywordInfo) error {
	r._keywordList = _keywordList
	r.Set("keyword_list", _keywordList)
	return nil
}

// GetKeywordList KeywordList Getter
func (r AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIRequest) GetKeywordList() []KeywordInfo {
	return r._keywordList
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}
