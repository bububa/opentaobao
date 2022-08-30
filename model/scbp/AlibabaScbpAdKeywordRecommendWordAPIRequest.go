package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordRecommendWordAPIRequest 推词 API请求
// alibaba.scbp.ad.keyword.recommend.word
//
// 推词
type AlibabaScbpAdKeywordRecommendWordAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 推词请求实体类
	_recommendQuery *RecommendKeywordQueryDto
}

// NewAlibabaScbpAdKeywordRecommendWordRequest 初始化AlibabaScbpAdKeywordRecommendWordAPIRequest对象
func NewAlibabaScbpAdKeywordRecommendWordRequest() *AlibabaScbpAdKeywordRecommendWordAPIRequest {
	return &AlibabaScbpAdKeywordRecommendWordAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordRecommendWordAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.recommend.word"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordRecommendWordAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdKeywordRecommendWordAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdKeywordRecommendWordAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetRecommendQuery is RecommendQuery Setter
// 推词请求实体类
func (r *AlibabaScbpAdKeywordRecommendWordAPIRequest) SetRecommendQuery(_recommendQuery *RecommendKeywordQueryDto) error {
	r._recommendQuery = _recommendQuery
	r.Set("recommend_query", _recommendQuery)
	return nil
}

// GetRecommendQuery RecommendQuery Getter
func (r AlibabaScbpAdKeywordRecommendWordAPIRequest) GetRecommendQuery() *RecommendKeywordQueryDto {
	return r._recommendQuery
}
