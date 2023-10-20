package scbp

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest) Reset() {
	r._campaignId = 0
	r._keywordQuery = nil
	r._topContext = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.query.keyword.rank.price"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetKeywordQuery is KeywordQuery Setter
// 请求参数
func (r *AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest) SetKeywordQuery(_keywordQuery *KeywordQuery) error {
	r._keywordQuery = _keywordQuery
	r.Set("keyword_query", _keywordQuery)
	return nil
}

// GetKeywordQuery KeywordQuery Getter
func (r AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest) GetKeywordQuery() *KeywordQuery {
	return r._keywordQuery
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

var poolAlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdKeywordQueryKeywordRankPriceRequest()
	},
}

// GetAlibabaScbpAdKeywordQueryKeywordRankPriceRequest 从 sync.Pool 获取 AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest
func GetAlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest() *AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest {
	return poolAlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest.Get().(*AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest)
}

// ReleaseAlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest 将 AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest(v *AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest.Put(v)
}
