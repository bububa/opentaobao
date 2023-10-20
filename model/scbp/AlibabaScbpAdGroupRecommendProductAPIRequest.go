package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadgrouprecommendproductAPIRequest 推品 API请求
// alibaba.scbp.ad.group.recommend.product
//
// 推品
type AlibabascbpadgrouprecommendproductAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 推品查询条件
	_recommendQuery *ProductRecommendQueryDto
}

// NewAlibabascbpadgrouprecommendproductRequest 初始化AlibabascbpadgrouprecommendproductAPIRequest对象
func NewAlibabascbpadgrouprecommendproductRequest() *AlibabascbpadgrouprecommendproductAPIRequest {
	return &AlibabascbpadgrouprecommendproductAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadgrouprecommendproductAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.group.recommend.product"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadgrouprecommendproductAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadgrouprecommendproductAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadgrouprecommendproductAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadgrouprecommendproductAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetRecommendQuery is RecommendQuery Setter
// 推品查询条件
func (r *AlibabascbpadgrouprecommendproductAPIRequest) SetRecommendQuery(_recommendQuery *ProductRecommendQueryDto) error {
	r._recommendQuery = _recommendQuery
	r.Set("recommend_query", _recommendQuery)
	return nil
}

// GetRecommendQuery RecommendQuery Getter
func (r AlibabascbpadgrouprecommendproductAPIRequest) GetRecommendQuery() *ProductRecommendQueryDto {
	return r._recommendQuery
}
