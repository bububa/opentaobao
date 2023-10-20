package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest 删除关键词 API请求
// alibaba.scbp.ad.keyword.delete.keyword.batch
//
// 删除关键词
type AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 计划id
	_campaignId int64
	// 请求参数
	_keywordQuery *KeywordQuery
}

// NewAlibabaScbpAdKeywordDeleteKeywordBatchRequest 初始化AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest对象
func NewAlibabaScbpAdKeywordDeleteKeywordBatchRequest() *AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest {
	return &AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest) Reset() {
	r._topContext = nil
	r._campaignId = 0
	r._keywordQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.delete.keyword.batch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetKeywordQuery is KeywordQuery Setter
// 请求参数
func (r *AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest) SetKeywordQuery(_keywordQuery *KeywordQuery) error {
	r._keywordQuery = _keywordQuery
	r.Set("keyword_query", _keywordQuery)
	return nil
}

// GetKeywordQuery KeywordQuery Getter
func (r AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest) GetKeywordQuery() *KeywordQuery {
	return r._keywordQuery
}

var poolAlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdKeywordDeleteKeywordBatchRequest()
	},
}

// GetAlibabaScbpAdKeywordDeleteKeywordBatchRequest 从 sync.Pool 获取 AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest
func GetAlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest() *AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest {
	return poolAlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest.Get().(*AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest)
}

// ReleaseAlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest 将 AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest(v *AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest.Put(v)
}
