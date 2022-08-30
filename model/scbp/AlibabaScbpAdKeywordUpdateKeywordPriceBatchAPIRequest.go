package scbp

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.update.keyword.price.batch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
