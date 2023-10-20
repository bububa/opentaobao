package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadkeywordbatchquerykeywordrankpriceAPIRequest 批量查询关键词前五名排价 API请求
// alibaba.scbp.ad.keyword.batch.query.keyword.rank.price
//
// 批量查询关键词前五名排价
type AlibabascbpadkeywordbatchquerykeywordrankpriceAPIRequest struct {
	model.Params
	// 关键词信息集合
	_keywordList []KeywordInfo
	// 用户信息
	_topContext *TopContextDto
	// 计划id
	_campaignId int64
}

// NewAlibabascbpadkeywordbatchquerykeywordrankpriceRequest 初始化AlibabascbpadkeywordbatchquerykeywordrankpriceAPIRequest对象
func NewAlibabascbpadkeywordbatchquerykeywordrankpriceRequest() *AlibabascbpadkeywordbatchquerykeywordrankpriceAPIRequest {
	return &AlibabascbpadkeywordbatchquerykeywordrankpriceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadkeywordbatchquerykeywordrankpriceAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.batch.query.keyword.rank.price"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadkeywordbatchquerykeywordrankpriceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadkeywordbatchquerykeywordrankpriceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeywordList is KeywordList Setter
// 关键词信息集合
func (r *AlibabascbpadkeywordbatchquerykeywordrankpriceAPIRequest) SetKeywordList(_keywordList []KeywordInfo) error {
	r._keywordList = _keywordList
	r.Set("keyword_list", _keywordList)
	return nil
}

// GetKeywordList KeywordList Getter
func (r AlibabascbpadkeywordbatchquerykeywordrankpriceAPIRequest) GetKeywordList() []KeywordInfo {
	return r._keywordList
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadkeywordbatchquerykeywordrankpriceAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadkeywordbatchquerykeywordrankpriceAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabascbpadkeywordbatchquerykeywordrankpriceAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabascbpadkeywordbatchquerykeywordrankpriceAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}
