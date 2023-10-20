package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest 获取计划关键词 API请求
// alibaba.scbp.ad.keyword.list.campaign.keywords
//
// 获取计划关键词
type AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 计划id
	_campaignId int64
	// 搜索条件
	_campaignKeywordQuery *CampaignKeywordQuery
}

// NewAlibabaScbpAdKeywordListCampaignKeywordsRequest 初始化AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest对象
func NewAlibabaScbpAdKeywordListCampaignKeywordsRequest() *AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest {
	return &AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest) Reset() {
	r._topContext = nil
	r._campaignId = 0
	r._campaignKeywordQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.list.campaign.keywords"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetCampaignKeywordQuery is CampaignKeywordQuery Setter
// 搜索条件
func (r *AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest) SetCampaignKeywordQuery(_campaignKeywordQuery *CampaignKeywordQuery) error {
	r._campaignKeywordQuery = _campaignKeywordQuery
	r.Set("campaign_keyword_query", _campaignKeywordQuery)
	return nil
}

// GetCampaignKeywordQuery CampaignKeywordQuery Getter
func (r AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest) GetCampaignKeywordQuery() *CampaignKeywordQuery {
	return r._campaignKeywordQuery
}

var poolAlibabaScbpAdKeywordListCampaignKeywordsAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdKeywordListCampaignKeywordsRequest()
	},
}

// GetAlibabaScbpAdKeywordListCampaignKeywordsRequest 从 sync.Pool 获取 AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest
func GetAlibabaScbpAdKeywordListCampaignKeywordsAPIRequest() *AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest {
	return poolAlibabaScbpAdKeywordListCampaignKeywordsAPIRequest.Get().(*AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest)
}

// ReleaseAlibabaScbpAdKeywordListCampaignKeywordsAPIRequest 将 AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdKeywordListCampaignKeywordsAPIRequest(v *AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdKeywordListCampaignKeywordsAPIRequest.Put(v)
}
