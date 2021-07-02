package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordCreateKeywordBatchAPIRequest 关键词添加 API请求
// alibaba.scbp.ad.keyword.create.keyword.batch
//
// 关键词添加
type AlibabaScbpAdKeywordCreateKeywordBatchAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
	// 请求参数
	_keywordQuery *KeywordQuery
	// 用户信息
	_topContext *TopContextDto
}

// NewAlibabaScbpAdKeywordCreateKeywordBatchRequest 初始化AlibabaScbpAdKeywordCreateKeywordBatchAPIRequest对象
func NewAlibabaScbpAdKeywordCreateKeywordBatchRequest() *AlibabaScbpAdKeywordCreateKeywordBatchAPIRequest {
	return &AlibabaScbpAdKeywordCreateKeywordBatchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordCreateKeywordBatchAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.create.keyword.batch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordCreateKeywordBatchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CampaignId Setter
// 计划id
func (r *AlibabaScbpAdKeywordCreateKeywordBatchAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// Get CampaignId Getter
func (r AlibabaScbpAdKeywordCreateKeywordBatchAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// Set is KeywordQuery Setter
// 请求参数
func (r *AlibabaScbpAdKeywordCreateKeywordBatchAPIRequest) SetKeywordQuery(_keywordQuery *KeywordQuery) error {
	r._keywordQuery = _keywordQuery
	r.Set("keyword_query", _keywordQuery)
	return nil
}

// Get KeywordQuery Getter
func (r AlibabaScbpAdKeywordCreateKeywordBatchAPIRequest) GetKeywordQuery() *KeywordQuery {
	return r._keywordQuery
}

// Set is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdKeywordCreateKeywordBatchAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// Get TopContext Getter
func (r AlibabaScbpAdKeywordCreateKeywordBatchAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}
