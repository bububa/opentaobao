package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest 计划关键词数目 API请求
// alibaba.scbp.ad.keyword.get.keyword.count.by.query
//
// 计划关键词数目
type AlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 请求参数
	_campaignKeywordQuery *CampaignKeywordQuery
}

// NewAlibabaScbpAdKeywordGetKeywordCountByQueryRequest 初始化AlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest对象
func NewAlibabaScbpAdKeywordGetKeywordCountByQueryRequest() *AlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest {
	return &AlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest) Reset() {
	r._topContext = nil
	r._campaignKeywordQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.get.keyword.count.by.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignKeywordQuery is CampaignKeywordQuery Setter
// 请求参数
func (r *AlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest) SetCampaignKeywordQuery(_campaignKeywordQuery *CampaignKeywordQuery) error {
	r._campaignKeywordQuery = _campaignKeywordQuery
	r.Set("campaign_keyword_query", _campaignKeywordQuery)
	return nil
}

// GetCampaignKeywordQuery CampaignKeywordQuery Getter
func (r AlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest) GetCampaignKeywordQuery() *CampaignKeywordQuery {
	return r._campaignKeywordQuery
}

var poolAlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdKeywordGetKeywordCountByQueryRequest()
	},
}

// GetAlibabaScbpAdKeywordGetKeywordCountByQueryRequest 从 sync.Pool 获取 AlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest
func GetAlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest() *AlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest {
	return poolAlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest.Get().(*AlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest)
}

// ReleaseAlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest 将 AlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest(v *AlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest.Put(v)
}
