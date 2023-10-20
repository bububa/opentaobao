package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadkeywordrecommendwordAPIRequest 推词 API请求
// alibaba.scbp.ad.keyword.recommend.word
//
// 推词
type AlibabascbpadkeywordrecommendwordAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 推词请求实体类
	_recommendQuery *RecommendKeywordQueryDto
}

// NewAlibabascbpadkeywordrecommendwordRequest 初始化AlibabascbpadkeywordrecommendwordAPIRequest对象
func NewAlibabascbpadkeywordrecommendwordRequest() *AlibabascbpadkeywordrecommendwordAPIRequest {
	return &AlibabascbpadkeywordrecommendwordAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadkeywordrecommendwordAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.recommend.word"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadkeywordrecommendwordAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadkeywordrecommendwordAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadkeywordrecommendwordAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadkeywordrecommendwordAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetRecommendQuery is RecommendQuery Setter
// 推词请求实体类
func (r *AlibabascbpadkeywordrecommendwordAPIRequest) SetRecommendQuery(_recommendQuery *RecommendKeywordQueryDto) error {
	r._recommendQuery = _recommendQuery
	r.Set("recommend_query", _recommendQuery)
	return nil
}

// GetRecommendQuery RecommendQuery Getter
func (r AlibabascbpadkeywordrecommendwordAPIRequest) GetRecommendQuery() *RecommendKeywordQueryDto {
	return r._recommendQuery
}
