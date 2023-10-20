package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadkeywordupdatekeywordpricebatchAPIRequest 修改关键词价格 API请求
// alibaba.scbp.ad.keyword.update.keyword.price.batch
//
// 修改关键词价格
type AlibabascbpadkeywordupdatekeywordpricebatchAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 计划id
	_campaignId int64
	// 请求参数
	_keywordUpdateQuery *KeywordUpdateQuery
}

// NewAlibabascbpadkeywordupdatekeywordpricebatchRequest 初始化AlibabascbpadkeywordupdatekeywordpricebatchAPIRequest对象
func NewAlibabascbpadkeywordupdatekeywordpricebatchRequest() *AlibabascbpadkeywordupdatekeywordpricebatchAPIRequest {
	return &AlibabascbpadkeywordupdatekeywordpricebatchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadkeywordupdatekeywordpricebatchAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.update.keyword.price.batch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadkeywordupdatekeywordpricebatchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadkeywordupdatekeywordpricebatchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadkeywordupdatekeywordpricebatchAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadkeywordupdatekeywordpricebatchAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabascbpadkeywordupdatekeywordpricebatchAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabascbpadkeywordupdatekeywordpricebatchAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetKeywordUpdateQuery is KeywordUpdateQuery Setter
// 请求参数
func (r *AlibabascbpadkeywordupdatekeywordpricebatchAPIRequest) SetKeywordUpdateQuery(_keywordUpdateQuery *KeywordUpdateQuery) error {
	r._keywordUpdateQuery = _keywordUpdateQuery
	r.Set("keyword_update_query", _keywordUpdateQuery)
	return nil
}

// GetKeywordUpdateQuery KeywordUpdateQuery Getter
func (r AlibabascbpadkeywordupdatekeywordpricebatchAPIRequest) GetKeywordUpdateQuery() *KeywordUpdateQuery {
	return r._keywordUpdateQuery
}
