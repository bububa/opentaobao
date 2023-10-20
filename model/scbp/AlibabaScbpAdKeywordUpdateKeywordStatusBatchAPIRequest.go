package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadkeywordupdatekeywordstatusbatchAPIRequest 修改关键词状态 API请求
// alibaba.scbp.ad.keyword.update.keyword.status.batch
//
// 修改关键词状态
type AlibabascbpadkeywordupdatekeywordstatusbatchAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 计划id
	_campaignId int64
	// 更新数据
	_keywordUpdateQuery *KeywordUpdateQuery
}

// NewAlibabascbpadkeywordupdatekeywordstatusbatchRequest 初始化AlibabascbpadkeywordupdatekeywordstatusbatchAPIRequest对象
func NewAlibabascbpadkeywordupdatekeywordstatusbatchRequest() *AlibabascbpadkeywordupdatekeywordstatusbatchAPIRequest {
	return &AlibabascbpadkeywordupdatekeywordstatusbatchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadkeywordupdatekeywordstatusbatchAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.update.keyword.status.batch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadkeywordupdatekeywordstatusbatchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadkeywordupdatekeywordstatusbatchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadkeywordupdatekeywordstatusbatchAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadkeywordupdatekeywordstatusbatchAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabascbpadkeywordupdatekeywordstatusbatchAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabascbpadkeywordupdatekeywordstatusbatchAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetKeywordUpdateQuery is KeywordUpdateQuery Setter
// 更新数据
func (r *AlibabascbpadkeywordupdatekeywordstatusbatchAPIRequest) SetKeywordUpdateQuery(_keywordUpdateQuery *KeywordUpdateQuery) error {
	r._keywordUpdateQuery = _keywordUpdateQuery
	r.Set("keyword_update_query", _keywordUpdateQuery)
	return nil
}

// GetKeywordUpdateQuery KeywordUpdateQuery Getter
func (r AlibabascbpadkeywordupdatekeywordstatusbatchAPIRequest) GetKeywordUpdateQuery() *KeywordUpdateQuery {
	return r._keywordUpdateQuery
}
