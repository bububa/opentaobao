package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadkeyworddeletekeywordbatchAPIRequest 删除关键词 API请求
// alibaba.scbp.ad.keyword.delete.keyword.batch
//
// 删除关键词
type AlibabascbpadkeyworddeletekeywordbatchAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 计划id
	_campaignId int64
	// 请求参数
	_keywordQuery *KeywordQuery
}

// NewAlibabascbpadkeyworddeletekeywordbatchRequest 初始化AlibabascbpadkeyworddeletekeywordbatchAPIRequest对象
func NewAlibabascbpadkeyworddeletekeywordbatchRequest() *AlibabascbpadkeyworddeletekeywordbatchAPIRequest {
	return &AlibabascbpadkeyworddeletekeywordbatchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadkeyworddeletekeywordbatchAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.delete.keyword.batch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadkeyworddeletekeywordbatchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadkeyworddeletekeywordbatchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadkeyworddeletekeywordbatchAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadkeyworddeletekeywordbatchAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabascbpadkeyworddeletekeywordbatchAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabascbpadkeyworddeletekeywordbatchAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetKeywordQuery is KeywordQuery Setter
// 请求参数
func (r *AlibabascbpadkeyworddeletekeywordbatchAPIRequest) SetKeywordQuery(_keywordQuery *KeywordQuery) error {
	r._keywordQuery = _keywordQuery
	r.Set("keyword_query", _keywordQuery)
	return nil
}

// GetKeywordQuery KeywordQuery Getter
func (r AlibabascbpadkeyworddeletekeywordbatchAPIRequest) GetKeywordQuery() *KeywordQuery {
	return r._keywordQuery
}
