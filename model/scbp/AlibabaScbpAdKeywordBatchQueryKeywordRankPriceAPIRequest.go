package scbp

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIRequest) Reset() {
	r._keywordList = r._keywordList[:0]
	r._topContext = nil
	r._campaignId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.batch.query.keyword.rank.price"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdKeywordBatchQueryKeywordRankPriceRequest()
	},
}

// GetAlibabaScbpAdKeywordBatchQueryKeywordRankPriceRequest 从 sync.Pool 获取 AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIRequest
func GetAlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIRequest() *AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIRequest {
	return poolAlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIRequest.Get().(*AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIRequest)
}

// ReleaseAlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIRequest 将 AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIRequest(v *AlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdKeywordBatchQueryKeywordRankPriceAPIRequest.Put(v)
}
