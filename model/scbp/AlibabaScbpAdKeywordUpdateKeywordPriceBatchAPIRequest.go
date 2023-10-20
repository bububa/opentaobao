package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest 修改关键词价格 API请求
// alibaba.scbp.ad.keyword.update.keyword.price.batch
//
// 修改关键词价格
type AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 计划id
	_campaignId int64
	// 请求参数
	_keywordUpdateQuery *KeywordUpdateQuery
}

// NewAlibabaScbpAdKeywordUpdateKeywordPriceBatchRequest 初始化AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest对象
func NewAlibabaScbpAdKeywordUpdateKeywordPriceBatchRequest() *AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest {
	return &AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest) Reset() {
	r._topContext = nil
	r._campaignId = 0
	r._keywordUpdateQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.update.keyword.price.batch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetKeywordUpdateQuery is KeywordUpdateQuery Setter
// 请求参数
func (r *AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest) SetKeywordUpdateQuery(_keywordUpdateQuery *KeywordUpdateQuery) error {
	r._keywordUpdateQuery = _keywordUpdateQuery
	r.Set("keyword_update_query", _keywordUpdateQuery)
	return nil
}

// GetKeywordUpdateQuery KeywordUpdateQuery Getter
func (r AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest) GetKeywordUpdateQuery() *KeywordUpdateQuery {
	return r._keywordUpdateQuery
}

var poolAlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdKeywordUpdateKeywordPriceBatchRequest()
	},
}

// GetAlibabaScbpAdKeywordUpdateKeywordPriceBatchRequest 从 sync.Pool 获取 AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest
func GetAlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest() *AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest {
	return poolAlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest.Get().(*AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest)
}

// ReleaseAlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest 将 AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest(v *AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest.Put(v)
}
