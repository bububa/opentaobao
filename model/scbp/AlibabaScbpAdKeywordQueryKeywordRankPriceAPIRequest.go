package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadkeywordquerykeywordrankpriceAPIRequest 查询关键词前五名排价 API请求
// alibaba.scbp.ad.keyword.query.keyword.rank.price
//
// 查询关键词前五名排价
type AlibabascbpadkeywordquerykeywordrankpriceAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
	// 请求参数
	_keywordQuery *KeywordQuery
	// 用户信息
	_topContext *TopContextDto
}

// NewAlibabascbpadkeywordquerykeywordrankpriceRequest 初始化AlibabascbpadkeywordquerykeywordrankpriceAPIRequest对象
func NewAlibabascbpadkeywordquerykeywordrankpriceRequest() *AlibabascbpadkeywordquerykeywordrankpriceAPIRequest {
	return &AlibabascbpadkeywordquerykeywordrankpriceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadkeywordquerykeywordrankpriceAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.query.keyword.rank.price"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadkeywordquerykeywordrankpriceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadkeywordquerykeywordrankpriceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabascbpadkeywordquerykeywordrankpriceAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabascbpadkeywordquerykeywordrankpriceAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetKeywordQuery is KeywordQuery Setter
// 请求参数
func (r *AlibabascbpadkeywordquerykeywordrankpriceAPIRequest) SetKeywordQuery(_keywordQuery *KeywordQuery) error {
	r._keywordQuery = _keywordQuery
	r.Set("keyword_query", _keywordQuery)
	return nil
}

// GetKeywordQuery KeywordQuery Getter
func (r AlibabascbpadkeywordquerykeywordrankpriceAPIRequest) GetKeywordQuery() *KeywordQuery {
	return r._keywordQuery
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadkeywordquerykeywordrankpriceAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadkeywordquerykeywordrankpriceAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}
