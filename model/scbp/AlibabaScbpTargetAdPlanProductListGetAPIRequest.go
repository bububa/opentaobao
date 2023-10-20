package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbptargetadplanproductlistgetAPIRequest 定向推广-获取推广计划产品列表 API请求
// alibaba.scbp.target.ad.plan.product.list.get
//
// 定向推广-获取推广计划产品列表
type AlibabascbptargetadplanproductlistgetAPIRequest struct {
	model.Params
	// TopP4pQuickProductQuery
	_topP4pQuickProductQuery *TopP4pQuickProductQuery
}

// NewAlibabascbptargetadplanproductlistgetRequest 初始化AlibabascbptargetadplanproductlistgetAPIRequest对象
func NewAlibabascbptargetadplanproductlistgetRequest() *AlibabascbptargetadplanproductlistgetAPIRequest {
	return &AlibabascbptargetadplanproductlistgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbptargetadplanproductlistgetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.target.ad.plan.product.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbptargetadplanproductlistgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbptargetadplanproductlistgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopP4pQuickProductQuery is TopP4pQuickProductQuery Setter
// TopP4pQuickProductQuery
func (r *AlibabascbptargetadplanproductlistgetAPIRequest) SetTopP4pQuickProductQuery(_topP4pQuickProductQuery *TopP4pQuickProductQuery) error {
	r._topP4pQuickProductQuery = _topP4pQuickProductQuery
	r.Set("top_p4p_quick_product_query", _topP4pQuickProductQuery)
	return nil
}

// GetTopP4pQuickProductQuery TopP4pQuickProductQuery Getter
func (r AlibabascbptargetadplanproductlistgetAPIRequest) GetTopP4pQuickProductQuery() *TopP4pQuickProductQuery {
	return r._topP4pQuickProductQuery
}
