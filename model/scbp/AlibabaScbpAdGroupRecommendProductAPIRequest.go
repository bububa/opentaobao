package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdGroupRecommendProductAPIRequest 推品 API请求
// alibaba.scbp.ad.group.recommend.product
//
// 推品
type AlibabaScbpAdGroupRecommendProductAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 推品查询条件
	_recommendQuery *ProductRecommendQueryDto
}

// NewAlibabaScbpAdGroupRecommendProductRequest 初始化AlibabaScbpAdGroupRecommendProductAPIRequest对象
func NewAlibabaScbpAdGroupRecommendProductRequest() *AlibabaScbpAdGroupRecommendProductAPIRequest {
	return &AlibabaScbpAdGroupRecommendProductAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdGroupRecommendProductAPIRequest) Reset() {
	r._topContext = nil
	r._recommendQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdGroupRecommendProductAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.group.recommend.product"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdGroupRecommendProductAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdGroupRecommendProductAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdGroupRecommendProductAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdGroupRecommendProductAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetRecommendQuery is RecommendQuery Setter
// 推品查询条件
func (r *AlibabaScbpAdGroupRecommendProductAPIRequest) SetRecommendQuery(_recommendQuery *ProductRecommendQueryDto) error {
	r._recommendQuery = _recommendQuery
	r.Set("recommend_query", _recommendQuery)
	return nil
}

// GetRecommendQuery RecommendQuery Getter
func (r AlibabaScbpAdGroupRecommendProductAPIRequest) GetRecommendQuery() *ProductRecommendQueryDto {
	return r._recommendQuery
}

var poolAlibabaScbpAdGroupRecommendProductAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdGroupRecommendProductRequest()
	},
}

// GetAlibabaScbpAdGroupRecommendProductRequest 从 sync.Pool 获取 AlibabaScbpAdGroupRecommendProductAPIRequest
func GetAlibabaScbpAdGroupRecommendProductAPIRequest() *AlibabaScbpAdGroupRecommendProductAPIRequest {
	return poolAlibabaScbpAdGroupRecommendProductAPIRequest.Get().(*AlibabaScbpAdGroupRecommendProductAPIRequest)
}

// ReleaseAlibabaScbpAdGroupRecommendProductAPIRequest 将 AlibabaScbpAdGroupRecommendProductAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdGroupRecommendProductAPIRequest(v *AlibabaScbpAdGroupRecommendProductAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdGroupRecommendProductAPIRequest.Put(v)
}
